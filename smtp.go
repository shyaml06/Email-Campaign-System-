// smtp.go
package main

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func Sender(recipient Recipient) error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	from := os.Getenv("MAIL")
	password := os.Getenv("PASSWD")
	toList := []string{recipient.Email}
	host := "smtp.gmail.com"
	port := "587"
	msg := "email from shyam's email campaign"
	body := []byte(msg)
	auth := smtp.PlainAuth("", from, password, host)
	err = smtp.SendMail(host+":"+port, auth, from, toList, body)

	if err != nil {
		return err
	}

	fmt.Println("Successfully sent mail to:", recipient.Email)
	return nil
}
