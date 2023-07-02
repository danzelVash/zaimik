package smtp

import (
	"bytes"
	"crypto/tls"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
	"html/template"
	"net"
	"net/mail"
	"os"
	"strings"
)

type EmailParams struct {
	TemplateName string
	TemplateVars interface{}
	Destination  string
	Subject      string
}

func SendEmail(params EmailParams) error {
	address, err := mail.ParseAddress(params.Destination)
	if err != nil {
		return BadEmail
	}

	domain := strings.Split(address.Address, "@")[1]

	mx, err := net.LookupMX(domain)
	if err != nil || len(mx) == 0 {
		return BadEmail
	}

	var body bytes.Buffer
	t, err := template.ParseFiles(params.TemplateName)
	if err != nil {
		return errors.Errorf("error while parsing email template %s: %s", params.TemplateName, err.Error())
	}

	err = t.Execute(&body, params.TemplateVars)
	if err != nil {
		return errors.Errorf("error while executing email template %s: %s", params.TemplateName, err.Error())
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", viper.GetString("email.address"))
	msg.SetHeader("To", params.Destination)
	msg.SetHeader("Subject", params.Subject)
	msg.SetBody("text/html", body.String())

	d := gomail.NewDialer("smtp.gmail.com", 587, viper.GetString("email.address"), os.Getenv("EMAIL_SUPPORT_SECRET"))
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err = d.DialAndSend(msg); err != nil {
		return err
	}

	return nil
}
