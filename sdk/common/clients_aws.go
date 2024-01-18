package common

import (
	"context"
	"log"
	"sync"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/pubsub"
	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"
	"github.com/ninjapunkgirls/sdk/graph"
)

type AWSClients struct {
	projectID string
	storage   *storage.Client
	firebase  *firebase.App
	firestore *firestore.Client
	pubsub    *pubsub.Client
	graph     map[string]*graph.GraphClient
	sync.RWMutex
}

// Firebase lazy loads the client when needed
func (self *GCPClients) AWSFirebase() *firebase.App {

	self.RLock()
	client := self.firebase
	self.RUnlock()

	if client == nil {
		self.Lock()
		conf := &firebase.Config{ProjectID: self.projectID}
		var err error
		self.firebase, err = firebase.NewApp(context.Background(), conf)
		if err != nil {
			log.Fatalln(err)
		}
		defer self.Unlock()
		return self.firebase
	}
	return client
}

// Firestore exposes and initalises the firestore db
func (self *GCPClients) AWSFirestore() *firestore.Client {

	self.RLock()
	client := self.firestore
	self.RUnlock()

	if client == nil {
		self.Lock()
		var err error
		self.firestore, err = self.firebase.Firestore(context.Background())
		if err != nil {
			log.Fatalln(err)
		}
		defer self.Unlock()
		return self.firestore
	}

	return client
}

func (self *GCPClients) S3() *storage.Client {

	self.RLock()
	client := self.storage
	self.RUnlock()

	if client == nil {
		self.Lock()
		var err error
		self.storage, err = storage.NewClient(context.Background())
		if err != nil {
			log.Fatalln(err)
		}
		defer self.Unlock()
		return self.storage
	}

	return client
}

func (self *GCPClients) SNS() *pubsub.Client {

	self.RLock()
	client := self.pubsub
	self.RUnlock()

	if client == nil {
		self.Lock()
		var err error
		self.pubsub, err = pubsub.NewClient(context.Background(), self.projectID)
		if err != nil {
			log.Fatal(err)
		}
		defer self.Unlock()
		return self.pubsub
	}

	return client
}
