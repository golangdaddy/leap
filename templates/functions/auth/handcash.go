package main

import (
	"context"
	"fmt"
	"log"
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

		// Get the current user's profile (Auth token was from oAuth callback)
		profile, err := app.Handcash().GetProfile(context.Background(), authToken)
		if err != nil {
			log.Fatalln("error: ", err)
		}

		username := profile.PublicProfile.Handle
		email := profile.PrivateProfile.Email

		user, status, err := app.CreateUser("handcash", email, username, "@")
		if err != nil {
			cloudfunc.HttpError(w, err, status)
			return
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
			fmt.Sprintf("%shome?otp=%s", "{{.WebAPI}}", secret),
			http.StatusTemporaryRedirect,
		)

		return
	}
}
