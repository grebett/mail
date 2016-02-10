// This package is a simple wrapper around gomail package
package mail

import (
	"gopkg.in/gomail.v2"
)

type Mailer struct {
	from   string
	dialer *gomail.Dialer
}

func (mailer *Mailer) Init(provider string, port int, from string, password string) error {
	mailer.from = from
	mailer.dialer = gomail.NewPlainDialer(provider, port, from, password)

	// try a dial to see if the infos are correct
	res, err := mailer.dialer.Dial()
	if err != nil {
		return err
	} else {
		res.Close()
		return nil
	}
}

func (mailer *Mailer) SendNewMessage(to string, subject string, body string) error { //attach?
	mail := gomail.NewMessage()
	mail.SetHeader("From", mailer.from)
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", body) // add a flag parameter to enable text/plain?

	if err := mailer.dialer.DialAndSend(mail); err != nil {
		return err
	} else {
		return nil
	}
}
