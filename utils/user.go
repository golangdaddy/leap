package utils

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/golangdaddy/leap/sdk/common"
	"github.com/richardboase/npgpublic/models"
	"google.golang.org/api/iterator"
)

// UserCollection abstracts the handling of subdata to within the user object
func UserCollection(app *common.App, user *models.User, collectionID string) *firestore.CollectionRef {
	return UserRefCollection(app, user.Ref(), collectionID)
}

func UserRefCollection(app *common.App, userRef models.UserRef, collectionID string) *firestore.CollectionRef {
	return app.Firestore().Collection("users").Doc(userRef.ID).Collection(collectionID)
}

// RegionCollection abstracts the handling of subdata to within the country/region
func RegionCollection(app *common.App, user *models.User, collectionID string) *firestore.CollectionRef {
	return app.Firestore().Collection("countries").Doc(user.Meta.Context.Country).Collection("regions").Doc(user.Meta.Context.Region).Collection(collectionID)
}

func GetUserByUsername(app *common.App, username string) (*models.User, error) {
	doc, err := app.Firestore().Collection("usernames").Doc(username).Get(app.Context())
	if err != nil {
		return nil, err
	}
	record := &models.Username{}
	if err := doc.DataTo(record); err != nil {
		return nil, err
	}
	return GetUserByID(app, record.User.ID)
}

func GetUser(app *common.App, ref models.UserRef) (*models.User, error) {
	return GetUserByID(app, ref.ID)
}

func GetUserByID(app *common.App, id string) (*models.User, error) {
	doc, err := app.Firestore().Collection("users").Doc(id).Get(context.Background())
	if err != nil {
		return nil, err
	}
	user := &models.User{}
	return user, doc.DataTo(user)
}

func GetUserByEmail(app *common.App, email string) (*models.User, error) {

	iter := app.Firestore().Collection("users").Where("email", "==", email).Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		user := &models.User{}
		return user, doc.DataTo(user)
	}

	return nil, fmt.Errorf("no user forund via email: %s", email)
}
