package main

import (
	"log"
	"os"
	"strconv"

	"github.com/go-mail/mail/v2"
	"github.com/joho/godotenv"
)

func main() {
	from := "test@lenslocked.com"
	to := "test@test.com"
	subject := "this is a test"
	plaintext := "this is the body"
	html := `<h1>Hello there buddy!</h1><p>This is the email</p><p>Hope you enjoy it</p>`

	msg := mail.NewMessage()
	msg.SetHeader("To", to)
	msg.SetHeader("From", from)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/plain", plaintext)
	msg.AddAlternative("text/html", html)
	msg.WriteTo(os.Stdout)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("STMP_HOST")
	portStr := os.Getenv("SMTP_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		panic(err)
	}
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")

	dialer := mail.NewDialer(host, port, username, password)
	sender, err := dialer.Dial()
	if err != nil {
		// TODO: Handle the error correctly
		panic(err)
	}
	defer sender.Close()
	err = mail.Send(sender, msg)
	if err != nil {
		// TODO: Handle the error correctly
		panic(err)
	}
}
