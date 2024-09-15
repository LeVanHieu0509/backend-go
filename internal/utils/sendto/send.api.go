package sendto

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type MailRequest struct {
	ToEmail     string `json:"toEmail"`
	MessageBody string `json:"messageBody"`
	Subject     string `json:"subject"`
	Attactment  string `json:"attachment"`
}

func sendEmailToJavaByAPI(otp string, email string, purpose string) error {
	postUrl := "https://localhost:8001/email/send_text"

	mailRequest := MailRequest{
		ToEmail:     email,
		MessageBody: "OTP is " + otp,
		Subject:     "Verify OTP " + purpose,
		Attactment:  "path/to/email",
	}

	// convert struct to json
	requestBody, err := json.Marshal(mailRequest)

	if err != nil {
		return err
	}

	// Create Request
	req, err := http.NewRequest("POST", postUrl, bytes.NewBuffer(requestBody))

	if err != nil {
		return err
	}

	// PUT Header
	req.Header.Set("Content-Type", "application/json")

	// Execute request
	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Sprintln("Response status: ", resp.Status)
	return nil
}
