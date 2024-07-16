package encrypt

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func SendToFortanixSDKMSTokenization(data string, username, password string) (string, error) {
	// เปลี่ยน URL นี้เป็น URL ใหม่ที่คุณต้องการใช้
	fortanixAPIURL := "your_urls"

	// แปลงข้อมูลเป็น Base64
	encodedData := base64.StdEncoding.EncodeToString([]byte(data))

	// ส่งข้อมูลไปยัง API พร้อมข้อมูลการรับรองความถูกต้อง (HTTP Basic Authentication)
	client := &http.Client{}
	reqBody := fmt.Sprintf(`{"alg": "AES", "mode": "FPE", "plain": "%s"}`, encodedData)
	req, err := http.NewRequest("POST", fortanixAPIURL, strings.NewReader(reqBody))
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(username, password) // กำหนด username และ password ในการรับรองความถูกต้อง
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// อ่านค่า "cipher" จากตอบรับ
	cipherData, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	// fmt.Println("Response from Fortanix SDK MS:", string(cipherData)) // เพิ่มบรรทัดนี้

	// สร้าง JSON object จากค่า "cipher" เท่านั้น
	var result struct {
		Cipher string `json:"cipher"`
	}
	err = json.Unmarshal(cipherData, &result)
	if err != nil {
		return "", err
	}

	return result.Cipher, nil
}

func SendToFortanixSDKMSTokenizationEmailForMasking(data string, username, password string) (string, error) { /// token only
	fortanixAPIURL := "your_url"
	encodedData := base64.StdEncoding.EncodeToString([]byte(data))

	// ส่งข้อมูลไปยัง API พร้อมข้อมูลการรับรองความถูกต้อง (HTTP Basic Authentication)
	client := &http.Client{}
	reqBody := fmt.Sprintf(`{"alg": "AES", "mode": "FPE", "plain": "%s"}`, encodedData)
	req, err := http.NewRequest("POST", fortanixAPIURL, strings.NewReader(reqBody))
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(username, password) // กำหนด username และ password ในการรับรองความถูกต้อง
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// อ่านค่า "cipher" จากตอบรับ
	cipherData, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// สร้าง JSON object จากค่า "cipher" เท่านั้น
	var result struct {
		Cipher string `json:"cipher"`
	}
	err = json.Unmarshal(cipherData, &result)
	if err != nil {
		return "", err
	}

	return result.Cipher, nil
}

func SendToFortanixSDKMSTokenizationPhoneForMasking(data string, username, password string) (string, error) {
	fortanixAPIURL := "your_url"

	// แปลงข้อมูลเป็น Base64
	encodedData := base64.StdEncoding.EncodeToString([]byte(data))

	// ส่งข้อมูลไปยัง API พร้อมข้อมูลการรับรองความถูกต้อง (HTTP Basic Authentication)
	client := &http.Client{}
	reqBody := fmt.Sprintf(`{"alg": "AES", "mode": "FPE", "plain": "%s"}`, encodedData)

	// เรียก API เส้นแรก เพื่อรับค่า "cipher"
	req1, err := http.NewRequest("POST", fortanixAPIURL, strings.NewReader(reqBody))
	if err != nil {
		return "", err
	}
	req1.SetBasicAuth(username, password)
	req1.Header.Add("Content-Type", "application/json")
	resp1, err := client.Do(req1)
	if err != nil {
		return "", err
	}
	defer resp1.Body.Close()

	// อ่านค่า "cipher" จากการเรียก API เส้นแรก
	cipherData, err := io.ReadAll(resp1.Body)
	if err != nil {
		return "", err
	}

	var result struct {
		Cipher string `json:"cipher"`
	}
	err = json.Unmarshal(cipherData, &result)
	if err != nil {
		return "", err
	}
	return result.Cipher, nil
}
