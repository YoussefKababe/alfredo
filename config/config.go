package config

import (
  "github.com/joho/godotenv"
  "log"
  "os"
)

var VerifyToken, PageToken string

func LoadEnvVars() {
  if err := godotenv.Load(); err != nil {
    log.Fatal(err)
  }
  VerifyToken = os.Getenv("VERIFY_TOKEN")
  PageToken = os.Getenv("PAGE_TOKEN")
}
