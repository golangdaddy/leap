package utils

import (
	"context"

	"github.com/golangdaddy/leap/sdk/common"
	"github.com/richardboase/npgpublic/models"
)

type Generic struct {
	Meta models.Internals
}

func GetMetadata(app *common.App, id string) (*models.Internals, error) {

	dst := &Generic{}

	i := models.Internal(id)
	path := i.DocPath()

	println("GET DOCUMENT", path)

	doc, err := app.Firestore().Doc(path).Get(context.Background())
	if err != nil {
		return nil, err
	}
	return &dst.Meta, doc.DataTo(dst)
}

func GetDocument(app *common.App, id string, dst interface{}) error {

	i := models.Internal(id)
	path := i.DocPath()

	println("GET DOCUMENT", path)

	doc, err := app.Firestore().Doc(path).Get(context.Background())
	if err != nil {
		return err
	}
	return doc.DataTo(dst)
}
