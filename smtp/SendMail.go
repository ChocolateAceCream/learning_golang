package main

import (
	gomail "gopkg.in/gomail.v2"
)

func main() {

	msg := gomail.NewMessage()
	msg.SetHeader("From", "nuodi@hotmail.com")
	msg.SetHeader("To", "344234485@qq.com")
	msg.SetHeader("Subject", "dfdfd")
	msg.SetBody("text/html", "<b>This is the body of the mail</b>")
	msg.Attach("timg.jpg")

	n := gomail.NewDialer("smtp.office365.com", 587, "nuodi@hotmail.com", "*Di996962648")

	// Send the email
	if err := n.DialAndSend(msg); err != nil {
		panic(err)
	}

}
