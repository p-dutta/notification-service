package notificationSender

import (
	"bytes"
	"fmt"
	"go.uber.org/zap"
	"io"
	"net/http"
	"net/url"
	"notificationSubscriber/logger"
)

func SendSms(smsAPI, username, password, msisdn, textBody, sid, datetime, blNumber string) error {
	// Create form data
	data := url.Values{}
	data.Set("user", username)
	data.Set("pass", password)
	data.Set("msisdn", msisdn)
	data.Set("text_body", textBody)
	data.Set("sid", sid)
	data.Set("datetime", datetime)
	data.Set("bl_number", blNumber)

	// Create a new request
	req, err := http.NewRequest("POST", smsAPI, bytes.NewBufferString(data.Encode()))
	if err != nil {
		zapLog := logger.GetForFile("send-sms")
		zapLog.Error("Error creating request:", zap.Error(err))
		return err
	}

	// Set the appropriate headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		zapLog := logger.GetForFile("send-sms")
		zapLog.Error("Error sending request:", zap.Error(err))
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			zapLog := logger.GetForFile("send-sms")
			zapLog.Error("Error closing response body:", zap.Error(err))
		}
	}(resp.Body)

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		zapLog := logger.GetForFile("send-sms")
		zapLog.Error("Error reading response body:", zap.Error(err))
		return err
	}
	fmt.Println("Response Body:", string(body))

	return nil

}
