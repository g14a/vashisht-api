package mailer

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"log"
	"strings"

	"gitlab.com/gowtham-munukutla/vashisht-api/config"
	"gitlab.com/gowtham-munukutla/vashisht-api/models"
	"gopkg.in/gomail.v2"
)

type Mail struct {
	senderId   string
	senderName string
	toIds      []string
	subject    string
	body       string
}

type SmtpServer struct {
	host string
	port string
}

func (s *SmtpServer) ServerName() string {
	return s.host + ":" + s.port
}

func (mail *Mail) BuildMessage() string {
	message := ""
	message += fmt.Sprintf("From: %s<%s>\r\n", mail.senderName, mail.senderId)
	if len(mail.toIds) > 0 {
		message += fmt.Sprintf("To: %s\r\n", strings.Join(mail.toIds, ";"))
	}

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	message += fmt.Sprintf("Subject: %s\r\n", mail.subject)
	message += "\r\n" + mime
	message += "\r\n" + mail.body

	log.Println(message)

	return message
}

func SendRegistrationEmail(user *models.User) error {
	t := template.New("template.html")
	t, _ = t.ParseFiles("mailer/template.html")

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, user); err != nil {
		return err
	}

	result := tpl.String()

	return SendMail(result, "Vashisht 2018 Registration", user.EmailAddress)

}

func SendMail(body string, subject string, recipient string) error {

	appConfig := config.GetAppConfig()
	mailConfig := appConfig.MailConfig

	smtpHost := mailConfig.SMTPConfig.Host
	smtpPort := mailConfig.SMTPConfig.Port

	senderName := mailConfig.MailSenderConfig.Name
	senderMailID := mailConfig.MailSenderConfig.Email
	senderAuthPassword := mailConfig.MailSenderConfig.Password

	d := gomail.NewDialer(smtpHost, smtpPort, senderMailID, senderAuthPassword)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := gomail.NewMessage()
	m.SetHeader("From", fmt.Sprintf("%s <%s>", senderName, senderMailID))
	m.SetHeader("To", recipient)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	// Send the email to Bob
	return d.DialAndSend(m)
}
