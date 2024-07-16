package sendEmailFunctions

import (
	"auth_git/config/configSMTP"
	"fmt"
	"net/smtp"

	"gopkg.in/gomail.v2"
)

func SendEmailRegister(to, subject, body string) error {

	smtpServer := configSMTP.SMTPServer
	smtpPort := configSMTP.SMTPPort
	senderEmail := configSMTP.SMTPUsername
	senderPassword := configSMTP.SMTPPassword

	from := senderEmail
	recipients := []string{to}

	// Create the email content in HTML format
	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-Version: 1.0;\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\";\r\n" +
		"\r\n" +
		body)
	// connect smtp server and sent email
	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpServer)
	err := smtp.SendMail(fmt.Sprintf("%s:%d", smtpServer, smtpPort), auth, from, recipients, msg)
	return err
}

func SendEmailOTP(to, subject, body string) error {

	from := configSMTP.SMTPUsername
	password := configSMTP.SMTPPassword

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer("your_smtp_server", 587, from, password)

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
