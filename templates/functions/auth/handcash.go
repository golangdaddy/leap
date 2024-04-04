package main

import (
	"fmt"
	"net/http"

	"github.com/golangdaddy/leap/sdk/cloudfunc"
	"github.com/kr/pretty"
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

		pretty.Println(profile)

		println("making handcash user")

		username := profile.PublicProfile.Handle
		email := profile.PublicProfile.Paymail

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
