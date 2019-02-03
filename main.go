package main

import (
	//"go-online-util/rsa"
	//"go-online-util/password"
	//"go-online-util/facebook"
	//"go-online-util/sendmail"
	"go-online-util/concurrent"
	"os"
	"os/signal"
	"syscall"
	//"go-online-util/listener"
	//"go-online-util/cron"
	// "go-online-util/timing"
	//"go-online-util/nsq"
	//"os"
	///"os/signal"
	//"syscall"
	//"time"
)

func main() {

	/*
		NSQ example
		https://github.com/cihm/goNotes
	*/
	//go nsqexample.Consumerflow()
	//go nsqexample.Producerflow()
	//time.Sleep(time.Second * 2)

	/*
		Timing example

	*/
	//timing.Flow1()

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
	//example
	concurrent.AntFlow()

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

	// forever := make(chan bool)
	// fmt.Println(" [*] Waiting for logs. To exit press CTRL+C")
	// <-forever
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
}
