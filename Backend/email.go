package backend

import (
	"fmt"
	"net/smtp"
)

//SendEmail function to send email
func SendEmail(email string, symbol string, target float64) {

	from := "balatamoghna.krypto@gmail.com"
	password := ">kWLC6--LEn':6G>"

	to := []string{email}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	msg := fmt.Sprintf("Your alert for %s has been triggered!\nThe current price of %s is %g!", symbol, symbol, target)

	message := `To: "Balatamoghna " <` + email + `>
From: "GoLang API Krypto" <` + from + `>
Subject: Crypto Alert Triggered!
	
` + msg + `
	`

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(message))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
