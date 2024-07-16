package decrypt

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	keyUsername = "your_credential"
	keyPassword = "your_credential"
)

func DetokenizationEmailForMasking(maskToken string) (string, error) {
	fortanixAPIURL := "your_url"

	// สร้าง JSON request โดยระบุ "cipher" ที่เป็นค่า "username_token"
	reqBody := fmt.Sprintf(`{"alg": "AES", "mode": "FPE", "cipher": "%s"}`, maskToken)

	client := &http.Client{}
	req, err := http.NewRequest("POST", fortanixAPIURL, strings.NewReader(reqBody))
	if err != nil {
		return "", err
	}

	// ตั้งค่าการรับรองความถูกต้อง (HTTP Basic Authentication)
	req.SetBasicAuth(keyUsername, keyPassword)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Log ตอบรับ
	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	// log.Printf("Response from Fortanix API: %s", string(responseBytes))

	// อ่านค่า "plain" จากการเรียก API
	var result struct {
		Plain string `json:"plain"`
	}
	err = json.Unmarshal(responseBytes, &result)
	if err != nil {
		return "", err
	}

	return result.Plain, nil
}

func Detokenize(usernameToken string) (string, error) {
	fortanixAPIURL := "your_url"

	// สร้าง JSON request โดยระบุ "cipher" ที่เป็นค่า "username_token"
	reqBody := fmt.Sprintf(`{"alg": "AES", "mode": "FPE", "cipher": "%s"}`, usernameToken)

	client := &http.Client{}
	req, err := http.NewRequest("POST", fortanixAPIURL, strings.NewReader(reqBody))
	if err != nil {
		return "", err
	}

	// ตั้งค่าการรับรองความถูกต้อง (HTTP Basic Authentication)
	req.SetBasicAuth(keyUsername, keyPassword)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Log ตอบรับ
	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	// log.Printf("Response from Fortanix API: %s", string(responseBytes))

	// อ่านค่า "plain" จากการเรียก API
	var result struct {
		Plain string `json:"plain"`
	}
	err = json.Unmarshal(responseBytes, &result)
	if err != nil {
		return "", err
	}

	return result.Plain, nil
}

func DetokenizationPhoneForMasking(mobilePhone string) (string, error) {
	fortanixAPIURL := "your_url"

	// สร้าง JSON request โดยระบุ "cipher" ที่เป็นค่า "username_token"
	reqBody := fmt.Sprintf(`{"alg": "AES", "mode": "FPE", "cipher": "%s"}`, mobilePhone)

	client := &http.Client{}
	req, err := http.NewRequest("POST", fortanixAPIURL, strings.NewReader(reqBody))
	if err != nil {
		return "", err
	}

	// ตั้งค่าการรับรองความถูกต้อง (HTTP Basic Authentication)
	req.SetBasicAuth(keyUsername, keyPassword)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Log ตอบรับ
	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	// log.Printf("Response from Fortanix API: %s", string(responseBytes))

	// อ่านค่า "plain" จากการเรียก API
	var result struct {
		Plain string `json:"plain"`
	}
	err = json.Unmarshal(responseBytes, &result)
	if err != nil {
		return "", err
	}

	return result.Plain, nil
}
