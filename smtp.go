package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func Sender(recipient Recipient) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
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

	// handling the errors
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Successfully sent mail to all user in toList")

	//fmt.Println(from, password)

}
