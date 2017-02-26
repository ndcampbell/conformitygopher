package report

import (
	"log"
	"time"

	"github.com/ndcampbell/conformitygopher/configs"

	"gopkg.in/gomail.v2"
)

func SendEmail(config *configs.EmailConfig) {
	curDate := time.Now().Format("2006-01-02")
	subject := "ConformityGopher Report - " + curDate
	d := gomail.NewDialer(config.Host, config.Port, config.Username, config.Password)
	//Sends email to everyone in config list
	for _, to := range config.Recipients {
		m := buildMessage(config.Sender, to, subject)
		if err := d.DialAndSend(m); err != nil {
			log.Fatal(err)
		}
	}
	log.Println("Report sent")
}

func buildMessage(sender string, recipient string, subject string) *gomail.Message {
	m := gomail.NewMessage()
	m.SetHeader("From", sender)
	m.SetHeader("To", recipient)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", "Testing <b>Testing</b>")
	return m
}
