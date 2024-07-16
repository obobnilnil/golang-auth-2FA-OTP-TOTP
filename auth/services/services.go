package services

import (
	"auth_git/auth/models"
	"auth_git/auth/repositories"
	"auth_git/config/proxyLine"
	"auth_git/utilts/generate"
	"auth_git/utilts/sendEmailFunctions"
	"auth_git/utilts/utility"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

type ServicePort interface {
	RequestEmailForValidateOTPChicCRMServices(email string) (string, error) // Email
	ValidateOTPFromRequestEmailChicCRMServices(receivedOTP, receivedReferenceID string) error
	QrTOTPChicCRMServices(accountName string, value int) (string, bool, error)
	ValidateQrTOTPChicCRMServices(validateQrTOTP models.ValidateQrTOTP) (bool, error)
	DeleteKeyQrTOTPChicCRMServices(deleteKeyQrTOTP string) error
	RequestEmailForValidateOTPLineChicCRMServices(email_line string) (string, error) // Line
}

type serviceAdapter struct {
	otpStore   map[string]int
	otpKeys    map[string]*otp.Key
	refIDStore map[string]string
	r          repositories.RepositoryPort
}

func NewServiceAdapter(r repositories.RepositoryPort) ServicePort {
	return &serviceAdapter{
		otpStore:   make(map[string]int), // int input for validate
		otpKeys:    make(map[string]*otp.Key),
		refIDStore: make(map[string]string),
		r:          r,
	}
}

func (s *serviceAdapter) RequestEmailForValidateOTPChicCRMServices(email string) (string, error) {
	otp := generate.GenerateOTP(6)
	referenceID := generate.GenerateReferenceID(6)
	s.otpStore[email] = otp
	s.refIDStore[email] = referenceID
	// subject := "Verify your identity with an OTP."
	// body := fmt.Sprintf("Your OTP code is: %d", otp) // %d for integer
	subject := "Verify your identity with an OTP."
	body := fmt.Sprintf("Please Use OTP provided below to verify your identity<br>Your OTP code is: %d<br>Your Reference No is: %s", otp, referenceID)

	if err := sendEmailFunctions.SendEmailOTP(email, subject, body); err != nil {
		log.Printf("Failed to send OTP via Email: %v\n", err)
		return "", err
	}
	return referenceID, nil
}

func (s *serviceAdapter) ValidateOTPFromRequestEmailChicCRMServices(receivedOTP, receivedReferenceID string) error {
	for email, expectedOTP := range s.otpStore {
		expectedReferenceID, existsRefID := s.refIDStore[email]

		if existsRefID && receivedOTP == strconv.Itoa(expectedOTP) && receivedReferenceID == expectedReferenceID {
			delete(s.otpStore, email)
			delete(s.refIDStore, email)
			return nil
		}
	}
	return errors.New("OTP or reference ID is invalid")
}
func (s *serviceAdapter) QrTOTPChicCRMServices(accountName string, value int) (string, bool, error) {
	if value != 1 {
		return "", false, nil
	}
	if _, found := s.otpKeys[accountName]; found {
		return "", false, errors.New("AccountName already exists")
	}
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "your_issuer",
		AccountName: accountName,
		Algorithm:   otp.AlgorithmSHA512,
		SecretSize:  32,
		Period:      30,
	})
	if err != nil {
		return "", false, err
	}

	s.otpKeys[accountName] = key
	qrCodeURL := utility.GenerateQRCodeURL(key)

	return qrCodeURL, true, nil
}

func (s *serviceAdapter) ValidateQrTOTPChicCRMServices(validateQrTOTP models.ValidateQrTOTP) (bool, error) {
	key, found := s.otpKeys[validateQrTOTP.AccountName]
	if !found {
		return false, errors.New("no OTP key available, account name does not match")
	}

	serverOTP, err := totp.GenerateCode(key.Secret(), time.Now())
	if err != nil {
		return false, err
	}

	return validateQrTOTP.OTP == serverOTP, nil
}

func (s *serviceAdapter) DeleteKeyQrTOTPChicCRMServices(deleteKeyQrTOTP string) error {
	if _, found := s.otpKeys[deleteKeyQrTOTP]; found {
		delete(s.otpKeys, deleteKeyQrTOTP)
		return nil
	}
	return errors.New("OTP key for " + deleteKeyQrTOTP + " not found")
}

func (s *serviceAdapter) RequestEmailForValidateOTPLineChicCRMServices(email_line string) (string, error) {
	otp := generate.GenerateOTP(6)
	referenceID := generate.GenerateReferenceID(6)
	s.otpStore[email_line] = otp
	s.refIDStore[email_line] = referenceID
	lineID, err := s.r.RequestEmailForValidateOTPLineChicCRMRepositories(email_line)
	if err != nil {
		log.Println(err)
		return "", err
	}
	message := map[string]interface{}{
		"to": lineID.LineID,
		"messages": []map[string]interface{}{
			{
				"type": "text",
				"text": "Your OTP is: " + strconv.Itoa(otp) + "\nReferenceID: " + referenceID,
			},
		},
	}
	fmt.Println(message)

	// Marshal the message to JSON
	body, err := json.Marshal(message)
	if err != nil {
		return "", err
	}

	respBody, statusCode, err := proxyLine.ProxyBackend(body)
	if err != nil {
		return "", err
	}

	// Check if the status code is a success
	if statusCode != http.StatusOK {
		return "", errors.New("received non-ok status code from LINE API: " + strconv.Itoa(statusCode))
	}

	// Optionally handle the response body...
	_ = respBody // Do something with the response body if needed

	return referenceID, nil
}
