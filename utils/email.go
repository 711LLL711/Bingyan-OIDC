package utils

import (
	"OIDC/config"
	"OIDC/model/request"
	"bytes"
	"crypto/tls"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/k3a/html2text"
	"gopkg.in/gomail.v2"
)

type EmailData struct {
	URL       string
	FirstName string
	Subject   string
}

// ? Email template parser

func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	// Walk the directory to get all the files.
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)
}

func SendEmail(user *request.UserRegisterRequest, data *EmailData) {

	// Sender data.
	from := config.GlobalConfig.SMTP.From
	smtpPass := config.GlobalConfig.SMTP.Password
	smtpUser := config.GlobalConfig.SMTP.Username
	to := user.Email
	smtpHost := config.GlobalConfig.SMTP.Host
	smtpPort, _ := strconv.Atoi(config.GlobalConfig.SMTP.Port)

	var body bytes.Buffer

	template, err := ParseTemplateDir("templates")
	if err != nil {
		log.Fatal("Could not parse template", err)
	}

	template.ExecuteTemplate(&body, "verificationCode.html", &data)

	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", body.String())
	m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send Email
	//TODO: error:could not send mail:missing @
	if err := d.DialAndSend(m); err != nil {
		log.Fatal("Could not send email: ", err)
	}

}
