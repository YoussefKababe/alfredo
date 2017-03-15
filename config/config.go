package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// VerifyToken is needed to verify connection between the messenger platform and the app.
var VerifyToken string

// PageToken is needed to communicate with the messenger platform.
var PageToken string

// DropboxToken is needed to communicate with Dropbox.
var DropboxToken string

// FirebaseToken is needed to communicate with Firebase.
var FirebaseToken string

// LoadEnvVars loads config environment variables
func LoadEnvVars() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	VerifyToken = os.Getenv("VERIFY_TOKEN")
	PageToken = os.Getenv("PAGE_TOKEN")
	DropboxToken = os.Getenv("DROPBOX_TOKEN")
	FirebaseToken = os.Getenv("FIREBASE_TOKEN")
}
