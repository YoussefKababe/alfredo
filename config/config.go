package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	// AppPort is the port that the app will listen on.
	AppPort string
	// VerifyToken is needed to verify connection between the messenger platform and the app.
	VerifyToken string
	// PageToken is needed to communicate with the messenger platform.
	PageToken string
	// FirebaseProjectID is needed to communicate with Firebase.
	FirebaseProjectID string
	// FirebaseSecret is needed to communicate with Firebase.
	FirebaseSecret string
	// DropboxKey is needed to communicate with Dropbox.
	DropboxKey string
	// DropboxSecret is needed to communicate with Dropbox.
	DropboxSecret string
	// DropboxRedirect is the page users are sent to after authorizing the app on Dropbox
	DropboxRedirect string
)

// Initialize loads config environment variables
func Initialize() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	VerifyToken = os.Getenv("VERIFY_TOKEN")
	PageToken = os.Getenv("PAGE_TOKEN")
	FirebaseProjectID = os.Getenv("FIREBASE_PROJECT_ID")
	FirebaseSecret = os.Getenv("FIREBASE_SECRET")
	DropboxKey = os.Getenv("DROPBOX_KEY")
	DropboxSecret = os.Getenv("DROPBOX_SECRET")
	DropboxRedirect = os.Getenv("DROPBOX_REDIRECT")

	AppPort = os.Getenv("PORT")
	if AppPort == "" {
		AppPort = "8080"
	}
}
