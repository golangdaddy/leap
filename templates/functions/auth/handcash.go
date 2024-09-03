package main

import (
	"fmt"
	"net/http"

	"github.com/golangdaddy/leap/sdk/cloudfunc"
)

// handcash-success
func (app *App) HandcashEntrypointSuccess(w http.ResponseWriter, r *http.Request) {

	if cloudfunc.HandleCORS(w, r, "*") {
		return
	}

	switch r.Method {

	case "GET":

		authToken, err := cloudfunc.QueryParam(r, "authToken")
		if err != nil {
			cloudfunc.HttpError(w, err, http.StatusBadRequest)
			return
		}

		// manage internal user
		profile, err := app.Handcash().GetProfile(app.Context(), authToken)
		if err != nil {
			cloudfunc.HttpError(w, err, http.StatusInternalServerError)
			return
		}

		username := profile.PublicProfile.Handle
		email := profile.PublicProfile.Paymail

		// find if email is conflicting
		user, err := app.GetUserByEmail(email)
		if err != nil {

			println("making handcash user")

			var status int
			user, status, err = app.CreateUser("handcash", email, username, "$")
			if err != nil {
				cloudfunc.HttpError(w, err, status)
				return
			}

		}

		secret := app.Token256()

		otp := NewOTP(email, user.Meta.ID)

		// hash the OTP secret to set the firestore record
		if _, err := app.Firestore().Collection("otp").Doc(
			app.SeedDigest(secret),
		).Set(app.Context(), otp); err != nil {
			cloudfunc.HttpError(w, err, http.StatusInternalServerError)
			return
		}

		http.Redirect(
			w,
			r,
			fmt.Sprintf("%shome?otp=%s&authToken=%s", "{{.Config.WebAPI}}", secret, authToken),
			http.StatusTemporaryRedirect,
		)

		return
	}
}
