package main

import (
	//"go-online-util/rsa"
	//"go-online-util/password"
	//"go-online-util/facebook"
	//"go-online-util/sendmail"
	//"go-online-util/concurrent"
	//"go-online-util/redis"
	//go-online-util/ginapi"
	//"go-online-util/jaeger"
	//"go-online-util/oop"
	//"go-online-util/designpattern"
	//"go-online-util/gmail"
	"go-online-util/minio"
	"os"
	"os/signal"
	"syscall"
	//"go-online-util/redis"
	//"go-online-util/listener"
	//"go-online-util/cron"
	// "go-online-util/timing"
	//"go-online-util/nsq"
	//"go-online-util/elsearch"
	//"os"
	///"os/signal"
	//"syscall"
	//"time"
)

func main() {
	/*
		gmailclient
	*/
	serverUrl := "domainname:9000"
	accessKeyID := "432432432432"
	secretAccessKey := "zuf+34324345543dsgfreyt"
	go minio.Flow(serverUrl, accessKeyID, secretAccessKey)

	/*
		gmailclient
	*/
	//go gmailclient.Flow()

	/*
		designpattern
	*/
	//go designpattern.Proxyd()
	//go designpattern.Decorator()

	/*
		oop
	*/
	//go oop.Flow()

	/*
		https://github.com/yurishkuro/opentracing-tutorial/tree/master/go/lesson02
	*/
	// samplingServerURL := "http://speng8.3232.de:5778/sampling"
	// collectorEndpoint := "http://speng8.3232.de:14268/api/traces"
	// //go testjaeger.Mainflow(samplingServerURL, collectorEndpoint)
	// go testjaeger.MainFormatflow(samplingServerURL, collectorEndpoint)
	// go testjaeger.MainPublishflow(samplingServerURL, collectorEndpoint)
	// time.Sleep(time.Second * 10)

	// go testjaeger.MainCompleteflow(samplingServerURL, collectorEndpoint)

	/*
		https://studygolang.com/articles/11836
		https://studygolang.com/articles/11836
	*/
	//serverUrl := "rerererewr.de:6379"
	//go ginapi.Mainflow()

	/*
		https://github.com/go-redis/redis
	*/
	//serverUrl := "rerererewr.de:6379"
	//go redis.Testflow(serverUrl)

	/*
		Elastic search example
		https://olivere.github.io/elastic/
	*/
	//serverUrl := "http://speng3.rerwe.de:3231208"
	//go elsearch.Buildflow(serverUrl)
	//go elsearch.Updateflow(serverUrl)
	//go elsearch.Searchflow(serverUrl)

	/*
		NSQ example
		https://github.com/cihm/goNotes
		//topic > channel
		//topic can publish to multi channels
	*/
	//var aa, bb, dd int = 1, 2, 3
	// var a, b, d *int = &aa, &bb, &dd
	//go nsqexample.Consumerflow(a)
	// time.Sleep(time.Second * 2)
	// go nsqexample.Consumerflow(b)
	// time.Sleep(time.Second * 2)
	// go nsqexample.Consumerflow(d)
	// time.Sleep(time.Second * 2)
	// go nsqexample.Producerflow()
	// time.Sleep(time.Second * 2)

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
	//concurrent.AntFlow()

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
	//寄送不到yahoo
	// sender := sendmail.NewSender("tretretre@sina.com", "tretretre!", "smtp.sina.com")
	// bodyMessage := sendmail.WriteHTMLEmail(sendmail.Receiver, sendmail.Subject, sendmail.Message, sender)
	// sender.SendMail(sendmail.Receiver, sendmail.Subject, bodyMessage)
	//sendmail.SendSina()

	// forever := make(chan bool)
	// fmt.Println(" [*] Waiting for logs. To exit press CTRL+C")
	// <-forever
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
}
