package report

import (
	"bytes"
	"html/template"
	"log"
	"time"

	"github.com/ndcampbell/conformitygopher/configs"

	"gopkg.in/gomail.v2"
)

type TemplateData struct {
	CurDate  string
	Resource string
}

func SendEmail(config *configs.EmailConfig) {
	curDate := time.Now().Format("2006-01-02")
	subject := "ConformityGopher Report - " + curDate
	d := gomail.NewDialer(config.Host, config.Port, config.Username, config.Password)
	//Sends email to everyone in config list
	data := TemplateData{CurDate: curDate, Resource: "ec2"}
	doc := buildTemplate(&data)
	for _, to := range config.Recipients {
		m := buildMessage(config.Sender, to, doc, subject)
		if err := d.DialAndSend(m); err != nil {
			log.Fatal(err)
		}
	}
	log.Println("Report sent")
}

func buildMessage(sender string, recipient string, doc string, subject string) *gomail.Message {
	m := gomail.NewMessage()
	m.SetHeader("From", sender)
	m.SetHeader("To", recipient)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", doc)
	return m
}

func buildTemplate(data *TemplateData) string {
	var doc bytes.Buffer
	t, err := template.ParseFiles("report/templates/emailtemplate.html")
	if err != nil {
		log.Fatal("error loading email template: ", err)
	}
	err = t.Execute(&doc, data)
	if err != nil {
		log.Fatal("error rendering template: ", err)
	}
	return doc.String()
}
