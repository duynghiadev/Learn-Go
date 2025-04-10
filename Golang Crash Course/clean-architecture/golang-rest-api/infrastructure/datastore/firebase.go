package datastore

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// InitFirestore initializes the Firestore client.
// Consider passing config values (projectID, credentialsPath) instead of hardcoding.
func InitFirestore(ctx context.Context, projectID string, credentialsPath string) (*firestore.Client, error) {
	sa := option.WithCredentialsFile(credentialsPath)
	conf := &firebase.Config{
		ProjectID: projectID,
	}

	app, err := firebase.NewApp(ctx, conf, sa)
	if err != nil {
		log.Printf("Datastore: Error initializing Firebase app: %v", err)
		return nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Printf("Datastore: Error initializing Firestore client: %v", err)
		return nil, err
	}

	log.Println("Datastore: Firestore client initialized successfully for project:", projectID)
	return client, nil
}
