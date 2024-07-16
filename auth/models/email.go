package models

type ValidateBody struct {
	// Email       string `json:"email"` << in case use email, refID, OTP for validate
	OTP         string `json:"otp"`
	ReferenceID string `json:"referenceID"`
}

type EmailForLineIDResponse struct {
	ID     string
	LineID *string
}
