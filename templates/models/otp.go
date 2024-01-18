package main

import (
	"context"
	"time"

	"github.com/golangdaddy/leap/sdk/common"
)

type OTP struct {
	Email     string `json:"email" firestore:"email"`
	User      string `json:"user" firestore:"user"`
	Timestamp int64  `json:"timestamp" firestore:"timestamp"`
}

func NewOTP(email, userID string) *OTP {
	return &OTP{
		Email:     email,
		User:      userID,
		Timestamp: time.Now().UTC().Unix(),
	}
}

func (otp *OTP) GetUser(app *common.App) (*User, error) {
	doc, err := app.Firestore().Collection("users").Doc(otp.User).Get(context.Background())
	if err != nil {
		return nil, err
	}
	user := &User{}
	return user, doc.DataTo(user)
}
