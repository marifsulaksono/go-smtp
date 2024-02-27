package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

const SMTP_HOST = "smtp.gmail.com"
const SMTP_PORT = 587

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error load env file : %v", err)
	}

	// set message
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", os.Getenv("SENDER_NAME"))
	mailer.SetHeader("To", "marifsulaksono@gmail.com")
	mailer.SetHeader("Subject", "Test Send New Email using Go")
	mailer.SetBody("text/html", "Hello, this is my testing result, enjoy :) !!")

	// set messsage dialer
	dialer := gomail.NewDialer(SMTP_HOST, SMTP_PORT, os.Getenv("AUTH_USER"), os.Getenv("AUTH_PASSWORD"))

	// dial and send message
	err = dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatalf("Error : %v", err)
	}

	log.Println("Success sent mail")
}
