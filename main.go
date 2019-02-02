package main

import (
	//"go-online-util/rsa"
	//"go-online-util/password"
	//"go-online-util/facebook"
	//"go-online-util/sendmail"
	//"go-online-util/concurrent"
	//"go-online-util/listener"
	//"go-online-util/cron"
	"go-online-util/timing"
)

func main() {

	/*
		Timing example

	*/
	timing.Flow1()

	/*
		Cron example

	*/
	//cron.CronFlow()

	/*
		Listener example

	*/
	//example 1
	//listener.ListenerFlow()
	//example 2
	//listener.EventFlow()

	/*
		Concurrent example

	*/
	//example 1
	//concurrent.TimeSample1()
	//example 2
	//concurrent.WorkPoolFlow()

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
	// sender := sendmail.NewSender("lewisli.acer@gmail.com", "your password")
	// bodyMessage := sendmail.WriteHTMLEmail(sendmail.Receiver, sendmail.Subject, sendmail.Message, sender)
	// sender.SendMail(sendmail.Receiver, sendmail.Subject, bodyMessage)
}
