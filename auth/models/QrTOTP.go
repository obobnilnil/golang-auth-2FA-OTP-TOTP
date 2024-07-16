package models

type QrTOTPRequest struct {
	Value       int    `json:"value"`
	AccountName string `json:"accountName"`
}

type ValidateQrTOTP struct {
	AccountName string `json:"accountName"`
	OTP         string `json:"otp"`
}

type DeleteKeyQrTOTP struct {
	AccountName string `json:"accountName"`
}
