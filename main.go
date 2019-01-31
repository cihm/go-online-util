package main

import (
	//"go-online-util/rsa"
	//"go-online-util/password"
	//"go-online-util/facebook"
	"go-online-util/sendmail"
)

func main() {

	/*
		Facebook example
	*/
	//facebook.FacebookFlow()

	/*
		Hash and salt string example
	*/
	//password.Runpassword()

	/*
		Ras flow example
	*/
	//rsa.RSAflow()

	/*
		Send mail example
	*/
	sender := sendmail.NewSender("lewisli.acer@gmail.com", "your password")
	bodyMessage := sendmail.WriteHTMLEmail(sendmail.Receiver, sendmail.Subject, sendmail.Message, sender)
	sender.SendMail(sendmail.Receiver, sendmail.Subject, bodyMessage)
}
