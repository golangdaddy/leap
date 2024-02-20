package main

import (
	"log"
	{{if .Object.Options.Order}}"google.golang.org/api/iterator"{{end}}
)

func (app *App) CreateDocument{{uppercase .Object.Name}}(parent *Internals, object *{{uppercase .Object.Name}}) error {
	
	log.Println("CREATING DOCUMENT", object.Meta.Class, object.Meta.ID)

	{{if eq false .Object.Options.Order}}/*{{end}}
	var order int
	iter := parent.Firestore(app.App).Collection(object.Meta.Class).Documents(app.Context())
	for {
		_, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Println(err)
			break
		}
		order++
	}
	object.Meta.Context.Order = order
	{{if eq false .Object.Options.Order}}*/{{end}}

	{{if ne nil .Object.Options.Assetlayer}}
	{{if eq false .Object.Options.Assetlayer.Wallet}}/*{{end}}
	// create app wallet
	{
		log.Println("CREATING WALLET")
		wallerUserID, err := app.Assetlayer().NewAppWallet(object.Meta.AssetlayerWalletID())
		if err != nil {
			return err
		}
		object.Meta.Wallet = wallerUserID
	}
	{{if eq false .Object.Options.Assetlayer.Wallet}}*/{{end}}

	{{if eq false .Object.Options.Assetlayer.Token}}/*{{end}}
	// create asset
	{
		log.Println("CREATING TOKEN")
		assetID, err := app.Assetlayer().MintAssetWithProperties(object.Meta.AssetlayerCollectionID(), object)
		if err != nil {
			return err
		}
		object.Meta.Asset = assetID
		{{if eq false .Object.Options.Assetlayer.Wallet}}
		if err := app.Assetlayer().SendAsset(assetID, "$"+object.Meta.AssetlayerWalletID()); err != nil {
			return err
		}
		{{end}}
	}
	{{if eq false .Object.Options.Assetlayer.Token}}*/{{end}}
	{{end}}
	
	// write new {{uppercase .Object.Name}} to the DB
	if err := object.Meta.SaveToFirestore(app.App, object); err != nil {
		return err
	}

	{{if eq nil .Object.Options.TopicCreate}}/*{{end}}
	b, err := app.MarshalJSON(object)
	if err != nil {
		return err
	}
	topicID := "{{.Object.Options.TopicCreate}}"
	result := app.PubSub().Topic(topicID).Publish(
		app.Context(),
		&pubsub.Message{Data: b},
	)
	msgID, err := result.Get(app.Context())
	if err != nil {
		return err
	}
	log.Println("PUBLISHED JOB TO TOPIC", topicID, msgID)
	{{if eq nil .Object.Options.TopicCreate}}*/{{end}}

	return nil
}