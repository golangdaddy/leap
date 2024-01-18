package common

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MyCustomClaims struct {
	jwt.RegisteredClaims
	Custom map[string]interface{} `json:"custom"`
}

func (app *App) ValidateToken(tokenString string) (bool, string, error) {

	app.RLock()
	key := app.jwtSigningKey
	app.RUnlock()

	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return key, nil
		})
	if err != nil {
		return false, "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return false, "", nil
	}

	subject, err := claims.GetSubject()
	if err != nil {
		log.Println(err)
		return false, "", err
	}

	return true, subject, nil
}

func (app *App) NewAuthToken(subject string, customClaims map[string]interface{}) (string, error) {

	app.RLock()
	key := app.jwtSigningKey
	app.RUnlock()

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		MyCustomClaims{
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(time.Hour * 8)),
				Issuer:    "test",
				Subject:   subject,
			},
			customClaims,
		},
	)
	return token.SignedString(key)
}
