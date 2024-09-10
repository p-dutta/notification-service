package notificationSender

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io"
	"net/http"
	"notificationSubscriber/logger"
	"os"
)

// EmailPayload represents the structure of the email payload
type EmailPayload struct {
	From             From      `json:"from"`
	Subject          string    `json:"subject"`
	Content          []Content `json:"content"`
	Personalizations []To      `json:"personalizations"`
}

// From represents the sender's information
type From struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

// Content represents the email content type and value
type Content struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// To represents the recipient's information
type To struct {
	To []Recipient `json:"to"`
}

// Recipient represents the recipient's email and name
type Recipient struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

// SendEmail sends an email with customizable parameters
func SendEmail(apiKey, fromEmail, fromName, toEmail, toName, subject, content string) error {
	url := os.Getenv("NETCORE_EMAIL_API")

	payload := EmailPayload{
		From: From{
			Email: fromEmail,
			Name:  fromName,
		},
		Subject: subject,
		Content: []Content{
			{
				Type:  "html",
				Value: content,
			},
		},
		Personalizations: []To{
			{
				To: []Recipient{
					{
						Email: toEmail,
						Name:  toName,
					},
				},
			},
		},
	}

	// Marshal the payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		zapLog := logger.GetForFile("send-mail")
		zapLog.Error("Error marshaling payload:", zap.Error(err))
		return err
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		zapLog := logger.GetForFile("send-mail")
		zapLog.Error("Error creating request:", zap.Error(err))
		return err
	}

	req.Header.Set("api_key", apiKey)
	req.Header.Set("Content-Type", "application/json")

	// Send HTTP request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		zapLog := logger.GetForFile("send-mail")
		zapLog.Error("Error sending request:", zap.Error(err))
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			zapLog := logger.GetForFile("send-mail")
			zapLog.Error("Error closing response body:", zap.Error(err))
		}
	}(res.Body)

	// Read response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		zapLog := logger.GetForFile("send-mail")
		zapLog.Error("Error reading response body:", zap.Error(err))
		return err
	}

	fmt.Println(string(body))

	return nil
}
