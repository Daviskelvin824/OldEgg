package controller

import (
	"net/smtp"

	"github.com/Daviskelvin824/OldEgg/config"
	"github.com/Daviskelvin824/OldEgg/models"
	"github.com/gin-gonic/gin"
)

func BlastEmail(c *gin.Context) {

	type RequestBody struct {
		MailSubject string `json:"mail_subject"`
		MailBody    string `json:"mail_body"`
	}

	var requestBody RequestBody
	c.ShouldBindJSON(&requestBody)

	auth := smtp.PlainAuth("", "oldegg2023@gmail.com", "Old.egg@2023oldegg", "smtp.gmail.com")

	var subscribedUsers []models.User

	config.DB.Model(models.User{}).Where("subscribed_to_email_offers_and_discounts = ?", "true").Find(&subscribedUsers)

	length := len(subscribedUsers)
	var to []string
	for i := 0; i < length; i++ {

		to = append(to, subscribedUsers[i].Email)

	}

	msg := "Subject: " + requestBody.MailSubject + "\n" + requestBody.MailBody
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"oldegg2023@gmail.com",
		to,
		[]byte(msg),
	)

	if err != nil {
		c.String(200, "Email Blast Error")
		return
	}

	c.String(200, "Email Sent Successfully")

}

func SendEmail(c *gin.Context) {

	type RequestBody struct {
		Recipient   string `json:"recipient"`
		MailSubject string `json:"mail_subject"`
		MailBody    string `json:"mail_body"`
	}

	var requestBody RequestBody
	c.ShouldBindJSON(&requestBody)

	auth := smtp.PlainAuth("", "oldegg2023@gmail.com", "Old.egg@2023oldegg", "smtp.gmail.com")

	msg := "Subject: " + requestBody.MailSubject + "\n" + requestBody.MailBody
	var to []string
	to = append(to, requestBody.Recipient)

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"oldegg2023@gmail.com",
		to,
		[]byte(msg),
	)

	if err != nil {
		c.String(200, "Send Error")
		return
	}

	c.String(200, "Email Sent Successfully")

}
