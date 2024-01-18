package cloudfunc

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

func VerifyUser(ctx context.Context, projectID string, r *http.Request) (*auth.Token, error) {

	errPrefix := "auth error"

	fire, err := firebase.NewApp(
		ctx,
		&firebase.Config{
			ProjectID: projectID,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %v\n", errPrefix, err)
	}
	client, err := fire.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %v\n", errPrefix, err)

	}

	header := strings.Split(r.Header.Get("Authorization"), " ")
	if len(header) != 2 || len(header[1]) < 10 {
		err := "invalid firebase id token"
		return nil, fmt.Errorf("%s: %v\n", errPrefix, err)
	}

	token, err := client.VerifyIDToken(ctx, header[1])
	if err != nil {
		return nil, fmt.Errorf("%s: %v\n", errPrefix, err)
	}

	return token, nil
}
