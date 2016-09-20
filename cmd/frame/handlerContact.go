package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/asaskevich/govalidator"
	"encoding/json"
	"log"
)

func handlerContact(c *gin.Context) {

	var body struct {
		Name    string `json:"name" valid:"required"`
		Email   string `json:"email" valid:"email,required"`
		Message string `json:"message" valid:"required"`
	}


	err := json.NewDecoder(c.Request.Body).Decode(&body)
	if err != nil {
		EXCEPTION(err.Error())
	}
	valid, err := govalidator.ValidateStruct(body)
	if err != nil {
		EXCEPTION(err.Error())
	}

	println(valid)

	mailConf := MailConfig{}
	mailConf.Data = body
	mailConf.From = config.SMTP.From.Name + " <" + config.SMTP.From.Address + ">"
	//mailConf.To = config.SystemEmail
	mailConf.To = "im7mortal@gmail.com"
	mailConf.Subject = config.CompanyName + " contact form"
	//mailConf.ReplyTo = body.Email
	mailConf.ReplyTo = "im7mortal@gmail.com"
	mailConf.HtmlPath = "views/contact/email-html.html"

	if err := mailConf.SendMail(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message" : "Email wasn't send. Please try another time or later.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "Success.",
	})
}

func EXCEPTION(i interface{}) {
	log.Panicln(i)
}