package mailer

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"

	"gitlab.com/gowtham-munukutla/vashisht-api/config"
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

	message += fmt.Sprintf("Subject: %s\r\n", mail.subject)
	message += "\r\n" + mail.body

	return message
}

func SendMail(body string, subject string, recipients []string) error {

	appConfig := config.GetAppConfig()
	mailConfig := appConfig.MailConfig

	mail := Mail{}
	mail.senderName = mailConfig.MailSenderConfig.Name
	mail.senderId = mailConfig.MailSenderConfig.Email
	mail.toIds = recipients
	mail.subject = subject
	mail.body = body

	messageBody := mail.BuildMessage()

	smtpServer := SmtpServer{host: mailConfig.SMTPConfig.Host, port: mailConfig.SMTPConfig.Port}

	//build an auth
	auth := smtp.PlainAuth("", mail.senderId, mailConfig.MailSenderConfig.Password, smtpServer.host)

	// Gmail will reject connection if it's not secure
	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpServer.host,
	}

	conn, err := tls.Dial("tcp", smtpServer.ServerName(), tlsconfig)
	if err != nil {
		return err
	}

	client, err := smtp.NewClient(conn, smtpServer.host)
	if err != nil {
		return err
	}

	// step 1: Use Auth
	if err = client.Auth(auth); err != nil {
		return err
	}

	// step 2: add all from and to
	if err = client.Mail(mail.senderId); err != nil {
		return err
	}
	for _, k := range mail.toIds {
		if err = client.Rcpt(k); err != nil {
			return err
		}
	}

	// Data
	w, err := client.Data()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(messageBody))
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	err = client.Quit()
	return err
}
