package proxyLine

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Proxy(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot read body"})
		return
	}
	req, err := http.NewRequest("POST", "https://api.line.me/v2/bot/message/push", bytes.NewBuffer(body))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot create request"})
		return
	}

	// add additional Header
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer your_authorization_token")

	// api request to LINE
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request to LINE API"})
		return
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response from LINE API"})
		return
	}
	c.Data(resp.StatusCode, "application/json", respBody)
}

func ProxyBackend(body []byte) ([]byte, int, error) { // validateOTP via LINE
	req, err := http.NewRequest("POST", "https://api.line.me/v2/bot/message/push", bytes.NewBuffer(body))
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	// Set headers
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer your_authorization_token")

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	// Read the response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return respBody, resp.StatusCode, nil
}
