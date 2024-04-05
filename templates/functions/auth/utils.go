package main

import (
	"fmt"
	"net/http"
)

func (app *App) CreateUser(mode, email, username string, prefix ...string) (user *User, status int, err error) {

	user = NewUser(mode, email, username)

	if !user.IsValid() {
		err = fmt.Errorf("username failed validation: %s", user.Username)
		status = http.StatusBadRequest
		return
	}

	// add any prefix to the username - stop collisions with normal users
	if len(prefix) > 0 {
		user.Username = fmt.Sprintf("%s%s", prefix[0], user.Username)
	}
	// find if email is conflicting
	if _, err = app.GetUserByEmail(user.Email); err == nil {
		err = fmt.Errorf("email already exists: %s", user.Email)
		status = http.StatusConflict
		return
	}

	// create new user
	if _, err = app.Firestore().Collection("users").Doc(user.Meta.ID).Set(app.Context(), user); err != nil {
		status = http.StatusInternalServerError
		return
	}

	// fail if a conflicting username exists
	if _, err = app.Firestore().Collection("usernames").Doc(user.Username).Get(app.Context()); err == nil {
		err = fmt.Errorf("username already exists: %s", user.Email)
		status = http.StatusConflict
		return
	}

	// create username association
	if _, err = app.Firestore().Collection("usernames").Doc(user.Username).Set(app.Context(), user.GetUsernameRef()); err != nil {
		status = http.StatusInternalServerError
		return
	}

	return
}
