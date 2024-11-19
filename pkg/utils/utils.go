package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"

	"github.com/Rizabekus/doodocs-zipper-rest-api/internal/models"
	"gopkg.in/gomail.v2"
)

func GetCallerInfo() (string, int, string) {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		return "Unknown", 0, "Unknown"
	}

	fn := runtime.FuncForPC(pc)
	functionName := fn.Name()

	return file, line, functionName
}
func SendResponse(msg string, w http.ResponseWriter, statusCode int) {
	response := models.Response{Message: msg}

	responseJSON, err := json.Marshal(response)
	if err != nil {

		resp := models.Response{Message: "Internal Server Error"}
		internalErrorJSON, _ := json.Marshal(resp)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, string(internalErrorJSON), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(responseJSON)
}
func SendEmail(to []string, filename string, fileData []byte) error {

	fromEmail := os.Getenv("SMTP_EMAIL")
	fromPassword := os.Getenv("SMTP_PASSWORD")
	smtpHost := "smtp.gmail.com"
	smtpPort := 587

	if fromEmail == "" || fromPassword == "" {
		return errors.New("SMTP credentials are not set")
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", fromEmail)
	msg.SetHeader("To", to...)
	msg.SetHeader("Subject", "File delivery")
	msg.SetBody("text/plain", "Please find the attached file.")

	fileReader := bytes.NewReader(fileData)
	msg.Attach(filename, gomail.SetCopyFunc(func(w io.Writer) error {
		_, err := io.Copy(w, fileReader)
		return err
	}))

	dialer := gomail.NewDialer(smtpHost, smtpPort, fromEmail, fromPassword)
	if err := dialer.DialAndSend(msg); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
