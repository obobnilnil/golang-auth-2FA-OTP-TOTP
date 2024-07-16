package configSMTP

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	SMTPServer   string
	SMTPPort     int
	SMTPUsername string
	SMTPPassword string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	SMTPServer = os.Getenv("SMTP_SERVER")
	SMTPUsername = os.Getenv("SMTP_USERNAME")
	SMTPPassword = os.Getenv("SMTP_PASSWORD")
	SMTPPort, err = strconv.Atoi(os.Getenv("SMTP_PORT"))

	if err != nil {
		log.Fatalf("Error converting SMTP_PORT to int: %v", err)
	}
}
