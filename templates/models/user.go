package main

import (
	"log"
	"regexp"
	"strings"
)

type Users []*User

type UserRef struct {
	Account  int
	ID       string
	Username string
}

func DemoUser() *User {
	return NewUser("john@doe.com", "john doe")
}

func NewUser(email, username string) *User {
	user := &User{
		Meta:     (Internals{}).NewInternals("users"),
		Email:    strings.ToLower(strings.TrimSpace(email)),
		Username: strings.ToLower(strings.TrimSpace(username)),
	}
	return user
}

type User struct {
	Meta Internals
	// user (0) or practitioner (1) or business (2)
	Account  int    `json:"account" firestore:"account"`
	Email    string `json:"email" firestore:"email"`
	Username string `json:"username" firestore:"username"`
}

func (user *User) Ref() UserRef {
	return UserRef{
		Account:  user.Account,
		ID:       user.Meta.ID,
		Username: user.Username,
	}
}

func (users Users) Refs() []UserRef {
	refs := []UserRef{}
	for _, user := range users {
		refs = append(refs, user.Ref())
	}
	return refs
}

func (user *User) IsValid() bool {
	log.Println(user.Username)

	if len(user.Username) < 6 {
		return false
	}
	if len(user.Username) > 24 {
		return false
	}
	if strings.Contains(user.Username, " ") {
		return false
	}
	if !isAlphanumeric(strings.Replace(user.Username, "_", "", -1)) {
		return false
	}
	return true
}

func isAlphanumeric(word string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(word)
}
