package firebase

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
	"gorm.io/gorm"
)

var CloudMessaging *FCM

type FCM struct {
	Clients map[int]*messaging.Client
}

func Initialize(clientCount int, notificationKey string, db *gorm.DB) *FCM {
	//! PROJECT DATABASE LOGGER
	FirebaseLogger = *NewLogger(db)

	var clients map[int]*messaging.Client = make(map[int]*messaging.Client)

	var i int = 1
	if clientCount < i {
		clientCount = i
	}

	for {
		if i > clientCount {
			break
		}

		clients[i] = initClient(notificationKey)
		i++
	}

	return &FCM{
		Clients: clients,
	}
}

func initClient(pkey string) *messaging.Client {
	c := context.Background()

	var opts []option.ClientOption
	if os.Getenv("SFTP_CLOUD") == "on" {
		//! CLOUD
		//! YOUR CLOUD PKEY KEY CODE HERE
	} else {
		//! LOCAL
		//! YOUR LOCAL PKEY CODE HERE
		opts = []option.ClientOption{option.WithCredentialsFile(pkey)}
	}

	app, err := firebase.NewApp(c, nil, opts...)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}

	client, err := app.Messaging(c)
	if err != nil {
		log.Fatalf("error initializing client %s:", err)
	}

	return client
}
