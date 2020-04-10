package auth

import (
	"context"
	firebase "firebase.google.com/go"
	fbauth "firebase.google.com/go/auth"
	"google.golang.org/api/option"
	"log"
	secrets_firebase "server/etc/secrets/firebase"
	"server/etc/utils"
)

func NewClient(PROJECT_ID string) *fbauth.Client {
	cred, err := getCred(PROJECT_ID)
	if err != nil {
		log.Fatalln(err)
	}
	client, err := createClient(cred)
	if err != nil {
		log.Fatalln(err)
	}
	return client
}

func getCred(PROJECT_ID string) (string, error) {
	fb_cred := secrets_firebase.Sample
	absPath := utils.GetMainPath() + fb_cred
	return absPath, nil
}

func createClient(CRED_PATH string) (*fbauth.Client, error) {
	opt := option.WithCredentialsFile(CRED_PATH)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
		//log.Fatalf("error initializing app: %v\n", err)
	}
	client, err := app.Auth(context.Background())
	if err != nil {
		return nil, err
		//log.Fatalf("error getting Auth client: %v\n", err)
	}
	log.Println("\x1b[31m", "----- Firebase auth Client Spawn! -----", "\x1b[0m")
	return client, nil
}
