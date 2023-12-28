package functions

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/richardboase/npgpublic/models"
	"github.com/richardboase/npgpublic/sdk/cloudfunc"
	"github.com/richardboase/npgpublic/sdk/common"
	"github.com/richardboase/npgpublic/utils"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"google.golang.org/api/iterator"
)

// api-auth
func Entrypoint(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()

	app := common.NewApp()
	app.UseGCP({{.ProjectID}})

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
			email = strings.TrimSpace(email)
			iter := app.Firestore().Collection("users").Where("email", "==", email).Documents(ctx)
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

			SENDGRID_API_KEY, err := cloudfunc.GetSecretFromVolume("/sendgrid-key/sendgrid-key")
			if err != nil {
				cloudfunc.HttpError(w, err, http.StatusInternalServerError)
				return
			}
			println(SENDGRID_API_KEY)

			email, err := cloudfunc.QueryParam(r, "email")
			if err != nil {
				cloudfunc.HttpError(w, err, http.StatusBadRequest)
				return
			}

			email = strings.ToLower(strings.TrimSpace(email))

			user, err := utils.GetUserByEmail(app, email)
			if err != nil {
				cloudfunc.HttpError(w, err, http.StatusBadRequest)
				return
			}

			secret := app.Token256()

			from := mail.NewEmail("", "richard@ninjapunkgirls.com")
			to := mail.NewEmail(user.Username, email)
			subject := "Sending with Twilio SendGrid is Fun"
			plainTextContent := "and easy to do anywhere, even with Go follow this link: "

			htmlContent := fmt.Sprintf(
				`<h2>One-time-password link:</h2>
				<br/>
				<a href='http://localhost:3000/dashboard?otp=%s'>Debug</a>
				<br/>
				<a href='http://npgplatform.vercel.app/dashboard?otp=%s'>Login</a>
				`,
				secret,
				secret,
			)
			message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
			client := sendgrid.NewSendClient(SENDGRID_API_KEY)
			response, err := client.Send(message)
			if err != nil {
				log.Println(err)
			} else {
				fmt.Println(response.StatusCode)
				fmt.Println(response.Body)
				fmt.Println(response.Headers)
			}

			otp := models.NewOTP(email, user.Meta.ID)

			// hash the OTP secret to set the firestore record
			if _, err := app.Firestore().Collection("otp").Doc(
				app.SeedDigest(secret),
			).Set(ctx, otp); err != nil {
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

			user := models.NewUser(email, username)

			if !user.IsValid() {
				err := fmt.Errorf("username failed validation: %s", user.Username)
				cloudfunc.HttpError(w, err, http.StatusBadRequest)
				return
			}

			// find if email is conflicting
			if _, err := utils.GetUserByEmail(app, user.Email); err == nil {
				cloudfunc.HttpError(w, err, http.StatusConflict)
				return
			}

			// create new user
			if _, err := app.Firestore().Collection("users").Doc(user.Meta.ID).Set(ctx, user); err != nil {
				cloudfunc.HttpError(w, err, http.StatusInternalServerError)
				return
			}

			// fail if a conflicting username exists
			if _, err := app.Firestore().Collection("usernames").Doc(user.Username).Get(ctx); err == nil {
				cloudfunc.HttpError(w, err, http.StatusConflict)
				return
			}

			// create username association
			if _, err := app.Firestore().Collection("usernames").Doc(user.Username).Set(ctx, user.GetUsernameRef()); err != nil {
				cloudfunc.HttpError(w, err, http.StatusInternalServerError)
				return
			}
			return

		case "login":

			// get and delete(?) otp
			otp, err := utils.DebugGetOTP(app, r)
			if err != nil {
				cloudfunc.HttpError(w, err, http.StatusBadRequest)
				return
			}

			secret, expires, err := utils.CreateSessionSecret(app, otp)
			if err != nil {
				cloudfunc.HttpError(w, err, http.StatusInternalServerError)
				return
			}

			user, err := otp.GetUser(app)
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
