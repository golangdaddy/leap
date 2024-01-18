package common

import (
	"context"
	"log"
	"sync"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/pubsub"
	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"
)

type GCPClients struct {
	projectID         string
	storage           *storage.Client
	firebase          *firebase.App
	firestore         *firestore.Client
	pubsub            *pubsub.Client
	firestoreDatabase string
	sync.RWMutex
}

// Firebase lazy loads the client when needed
func (self *GCPClients) Firebase() *firebase.App {

	client := self.firebase

	if client == nil {
		conf := &firebase.Config{ProjectID: self.projectID}
		var err error
		self.firebase, err = firebase.NewApp(context.Background(), conf)
		if err != nil {
			log.Fatalln(err)
		}
		return self.firebase
	}
	return client
}

// Firestore exposes and initalises the firestore db
func (self *GCPClients) Firestore() *firestore.Client {

	self.RLock()
	client := self.firestore
	self.RUnlock()

	if client == nil {
		self.Lock()
		defer self.Unlock()
		var err error
		ctx := context.Background()
		log.Println("connecting to firestore with database:", self.firestoreDatabase)
		self.firestore, err = firestore.NewClientWithDatabase(ctx, self.projectID, self.firestoreDatabase)
		if err != nil {
			log.Fatalf("Failed to create Firestore client with database: %s %v", self.firestoreDatabase, err)
		}
		return self.firestore
	}

	return client
}

func (self *GCPClients) GCS() *storage.Client {

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
		self.Unlock()
		return self.storage
	}

	return client
}

func (self *GCPClients) PubSub() *pubsub.Client {

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
		self.Unlock()
		return self.pubsub
	}

	return client
}
