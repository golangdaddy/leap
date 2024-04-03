package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golangdaddy/leap/sdk/cloudfunc"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"google.golang.org/api/iterator"
)

// api-auth
func (app *App) AuthEntrypoint(w http.ResponseWriter, r *http.Request) {

	if cloudfunc.HandleCORS(w, r, "*") {
		return
	}

	function, err := cloudfunc.QueryParam(r, "function")
	if err != nil {
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "GET":

		switch function {

		case "query":

			email, err := cloudfunc.QueryParam(r, "email")
			if err != nil {
				cloudfunc.HttpError(w, err, http.StatusInternalServerError)
				return
			}
			var result int = 0
			email = strings.TrimSpace(strings.ToLower(email))
			iter := app.Firestore().Collection("users").Where("email", "==", email).Documents(app.Context())
			for {
				_, err = iter.Next()
				if err == iterator.Done {
					break
				}
				if err != nil {
					break
				}
				result = 1
				break
			}
			reply := map[string]int{"result": result}
			if err := cloudfunc.ServeJSON(w, reply); err != nil {
				cloudfunc.HttpError(w, err, http.StatusInternalServerError)
				return
			}
			return

		case "otp":

			email, err := cloudfunc.QueryParam(r, "email")
			if err != nil {
				cloudfunc.HttpError(w, err, http.StatusBadRequest)
				return
			}

			email = strings.ToLower(strings.TrimSpace(email))

			user, err := app.GetUserByEmail(email)
			if err != nil {
				cloudfunc.HttpError(w, err, http.StatusBadRequest)
				return
			}

			secret := app.Token256()
			log.Println(secret)

			// send magic email link
			SENDGRID_API_KEY := os.Getenv("SENDGRID_API_KEY")
			from := mail.NewEmail("", "richard@ninjapunkgirls.com")
			to := mail.NewEmail(user.Username, email)
			subject := "MAGIC LINK for {{.SiteName}}"
			plainTextContent := fmt.Sprintf(
				"one time password link: %shome?otp=%s",
				"{{.WebAPI}}",
				secret,
			)

			htmlContent := fmt.Sprintf(
				`<h2>One-time-password link:</h2>
				<br/>
				<a href='http://localhost:3000/home?p={{.SiteName}}&otp=%s'>Debug</a>
				<br/>
				<br/>
				<a href='{{.WebAPI}}home?p={{.SiteName}}&otp=%s'>Click here to Login</a>
				`,
				secret,
				secret,
			)
			message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
			println(SENDGRID_API_KEY)
			client := sendgrid.NewSendClient(SENDGRID_API_KEY)
			response, err := client.Send(message)
			if err != nil {
				log.Println(err)
			} else {
				fmt.Println(response.StatusCode)
				fmt.Println(response.Body)
				fmt.Println(response.Headers)
			}

			otp := NewOTP(email, user.Meta.ID)

			// hash the OTP secret to set the firestore record
			if _, err := app.Firestore().Collection("otp").Doc(
				app.SeedDigest(secret),
			).Set(app.Context(), otp); err != nil {
				cloudfunc.HttpError(w, err, http.StatusInternalServerError)
				return
			}

			cloudfunc.ServeJSON(w, "please check your email...")
			return
		}

	case "POST":

		switch function {

		case "register":

			email, err := cloudfunc.QueryParam(r, "email")
			if err != nil {
				cloudfunc.HttpError(w, err, http.StatusBadRequest)
				return
			}

			username, err := cloudfunc.QueryParam(r, "username")
			if err != nil {
				cloudfunc.HttpError(w, err, http.StatusBadRequest)
				return
			}

			username = strings.Replace(strings.TrimSpace(strings.Replace(username, "_", " ", -1)), " ", "_", -1)

			_, status, err := app.CreateUser("email", email, username)
			if err != nil {
				cloudfunc.HttpError(w, err, status)
				return
			}

			return

		case "login":

			// get and delete(?) otp
			otp, err := app.DebugGetOTP(r)
			if err != nil {
				cloudfunc.HttpError(w, err, http.StatusBadRequest)
				return
			}

			secret, expires, err := app.CreateSessionSecret(otp)
			if err != nil {
				cloudfunc.HttpError(w, err, http.StatusInternalServerError)
				return
			}

			user, err := otp.GetUser(app.App)
			if err != nil {
				cloudfunc.HttpError(w, err, http.StatusInternalServerError)
				return
			}

			data := map[string]interface{}{
				"secret":  secret,
				"user":    user,
				"expires": expires,
			}

			if err := cloudfunc.ServeJSON(w, data); err != nil {
				cloudfunc.HttpError(w, err, http.StatusInternalServerError)
				return
			}
			return

		}

	}
}
