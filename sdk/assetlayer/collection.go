package assetlayer

import (
	"encoding/json"

	"github.com/kr/pretty"
)

type Collection struct {
	SlotID       string `json:"slotId"`
	CollectionID string `json:"collectionId"`
	//
	CollectionName   string                 `json:"collectionName"`
	CollectionImage  string                 `json:"collectionImage"`
	CollectionBanner string                 `json:"collectionBanner"`
	Description      string                 `json:"description"`
	Type             string                 `json:"type"`
	Maximum          int                    `json:"maximum"`
	Tags             []string               `json:"tags"`
	RoyaltyRecipient map[string]interface{} `json:"royaltyRecipient"`
	Properties       map[string]interface{} `json:"properties"`
}

func (client *Client) DeactivateCollection(collectionID string) error {
	_, err := client.Try(
		"PUT",
		"/api/v1/collection/deactivate",
		map[string]string{
			"collectionId": collectionID,
		},
	)
	return err
}

func (client *Client) EnsureCollectionExists(slotID, collectionType, name, description, image string, maximum int, properties map[string]interface{}) (string, error) {
	collections, err := client.GetCollections(slotID)
	if err != nil {
		return "", err
	}
	exists := false
	var collection *Collection
	for _, collection = range collections {
		if collection.CollectionName == name {
			exists = true
			break
		}
	}
	if exists {
		println("returning existing collection: " + collection.CollectionID)
		return collection.CollectionID, nil
	}
	return client.NewCollection(slotID, collectionType, name, description, image, maximum, properties)
}

func (client *Client) NewCollection(slotID, collectionType, name, description, image string, maximum int, properties map[string]interface{}) (string, error) {

	collection := &Collection{
		SlotID: slotID,
		// Identical or Unique
		Type:            collectionType,
		CollectionName:  name,
		Description:     description,
		CollectionImage: image,
		Maximum:         maximum,
		Properties:      properties,
	}

	pretty.Println(collection)

	data, err := client.Try("POST", "/api/v1/collection/new", nil, collection)
	if err != nil {
		return "", err
	}
	m, err := assertMapStringInterface(data)
	if err != nil {
		return "", err
	}
	id, err := assertString(m["collectionId"])
	if err != nil {
		return "", err
	}
	return id, nil
}

func (client *Client) GetCollections(slotID string) ([]*Collection, error) {

	data, err := client.Try(
		"GET",
		"/api/v1/slot/collections",
		map[string]string{
			"slotId": slotID,
			"idOnly": "false",
		},
	)
	if err != nil {
		return nil, err
	}

	m, err := assertMapStringInterface(data)
	if err != nil {
		return nil, err
	}
	app, err := assertMapStringInterface(m["slot"])
	if err != nil {
		return nil, err
	}
	s, err := assertInterfaceArray(app["collections"])
	if err != nil {
		return nil, err
	}

	collections := []*Collection{}

	for _, item := range s {
		b, err := json.Marshal(item)
		if err != nil {
			return nil, err
		}
		collection := &Collection{}
		if err := json.Unmarshal(b, collection); err != nil {
			return nil, err
		}
		collections = append(collections, collection)
	}

	return collections, nil
}
