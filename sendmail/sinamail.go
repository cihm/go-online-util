package sendmail

import (
	"flag"
	"fmt"
	"log"
	"net/smtp"
)

var (
	subject  = flag.String("s", "", "subject of the mail")
	body     = flag.String("b", "", "body of themail")
	reciMail = flag.String("m", "", "recipient mail address")
)

func SendSina() {
	// Set up authentication information.
	flag.Parse()
	sub := fmt.Sprintf("subject: %s\r\n\r\n", *subject)
	content := *body
	//mailList := strings.Split(*reciMail, ",")
	receiver := []string{"g1007125goodman@yahoo.com.tw", "g1007125goodman@gmail.com"}

	auth := smtp.PlainAuth(
		"",
		"sina@sina.com",
		"qp!",
		"smtp.sina.com",
		//"smtp.gmail.com",
	)
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		"smtp.sina.com:587",
		auth,
		"sina@sina.com",
		receiver,
		[]byte(sub+content),
	)
	if err != nil {
		log.Fatal("!!!", err)
	}
}
