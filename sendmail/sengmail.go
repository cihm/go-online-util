package sendmail

import (
	"bytes"
	"fmt"
	"mime/quotedprintable"
	"net/smtp"
	"strings"
)

/**
	Modified from https://gist.github.com/jpillora/cb46d183eca0710d909a
	Thank you very much.


	if erro
	mtp error: 535 5.7.8 Username and Password not accepted. Learn more at
    5.7.8  https://support.google.com/mail/?p=BadCredentials q1sm16828610pfb.96

	check:
	https://serverfault.com/questions/635139/how-to-fix-send-mail-authorization-failed-534-5-7-14
**/
var (
	SMTPServer  string
	Receiver    []string
	Subject     string
	Message     string
	BodyMessage string
)

func init() {
	/**
		Gmail SMTP Server
	**/
	SMTPServer = "smtp.gmail.com"
	//The receiver needs to be in slice as the receive supports multiple receiver
	Receiver = []string{"g1007125goodman@gmail.com", "xyz@gmail.com", "larrypage@googlemail.com"}

	Subject = "Testing HTLML Email from golang"
	Message = `
		<!DOCTYPE HTML PULBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
		<html>
		<head>
		<meta http-equiv="content-type" content="text/html"; charset=ISO-8859-1">
		</head>
		<body>This is the body<br>
		<div class="moz-signature"><i><br>
		<br>
		Regards<br>
		Alex<br>
		<i></div>
		</body>
		</html>
		`
	//BodyMessage := WriteHTMLEmail(Receiver, Subject, Message)
}

type Sender struct {
	User     string
	Password string
}

func NewSender(Username, Password string) Sender {

	return Sender{Username, Password}
}

func (sender Sender) SendMail(Dest []string, Subject, bodyMessage string) {

	msg := "From: " + sender.User + "\n" +
		"To: " + strings.Join(Dest, ",") + "\n" +
		"Subject: " + Subject + "\n" + bodyMessage

	err := smtp.SendMail(SMTPServer+":587",
		smtp.PlainAuth("", sender.User, sender.Password, SMTPServer),
		sender.User, Dest, []byte(msg))

	if err != nil {

		fmt.Printf("smtp error: %s", err)
		return
	}

	fmt.Println("Mail sent successfully!")
}

func WriteEmail(dest []string, contentType, subject, bodyMessage string, sender Sender) string {

	header := make(map[string]string)
	header["From"] = sender.User

	receipient := ""

	for _, user := range dest {
		receipient = receipient + user
	}

	header["To"] = receipient
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = fmt.Sprintf("%s; charset=\"utf-8\"", contentType)
	header["Content-Transfer-Encoding"] = "quoted-printable"
	header["Content-Disposition"] = "inline"

	message := ""

	for key, value := range header {
		message += fmt.Sprintf("%s: %s\r\n", key, value)
	}

	var encodedMessage bytes.Buffer

	finalMessage := quotedprintable.NewWriter(&encodedMessage)
	finalMessage.Write([]byte(bodyMessage))
	finalMessage.Close()

	message += "\r\n" + encodedMessage.String()

	return message
}

func WriteHTMLEmail(dest []string, subject, bodyMessage string, sender Sender) string {

	return WriteEmail(dest, "text/html", subject, bodyMessage, sender)
}

func WritePlainEmail(dest []string, subject, bodyMessage string, sender Sender) string {

	return WriteEmail(dest, "text/plain", subject, bodyMessage, sender)
}
