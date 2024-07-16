package repositories

import (
	"auth_git/auth/models"
	"auth_git/utilts/encrypt"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type RepositoryPort interface {
	RequestEmailForValidateOTPLineChicCRMRepositories(email_line string) (models.EmailForLineIDResponse, error)
}

type repositoryAdapter struct {
	db *sql.DB
}

func NewRepositoryAdapter(db *sql.DB) RepositoryPort {
	return &repositoryAdapter{db: db}
}

func (r *repositoryAdapter) RequestEmailForValidateOTPLineChicCRMRepositories(email_line string) (models.EmailForLineIDResponse, error) { //
	var emailLineIDInfo models.EmailForLineIDResponse
	const (
		keyUsername = "your_credential"
		keyPassword = "your_credential"
	)
	if email_line == "" {
		return emailLineIDInfo, errors.New("email must not be empty")
	}
	fmt.Println(email_line)
	cipherUsername, err := encrypt.SendToFortanixSDKMSTokenizationEmailForMasking(email_line, keyUsername, keyPassword)
	if err != nil {
		log.Println(err)
		return emailLineIDInfo, err
	}
	err = r.db.QueryRow("SELECT orgmb_id, orgmb_line_id FROM organize_member WHERE orgmb_email = $1", cipherUsername).Scan(&emailLineIDInfo.ID, &emailLineIDInfo.LineID)
	if err != nil {
		log.Printf("ID or lineID does not match. Error: %v", err)
		return emailLineIDInfo, errors.New("id or lineID does not match")
	}
	return emailLineIDInfo, nil
}
