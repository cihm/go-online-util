package gmailclient

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"net/mail"
	"os"
	"strings"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	fmt.Println(err)
	fmt.Println(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func Flow() {
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// Creates a oauth2.Config using the secret
	// The second parameter is the scope, in this case we only want to send email
	conf, err := google.ConfigFromJSON(b, gmail.GmailSendScope)
	if err != nil {
		log.Printf("Error: %v", err)
	}

	// // Creates a URL for the user to follow
	// url := conf.AuthCodeURL("CSRF", oauth2.AccessTypeOffline)
	// // Prints the URL to the terminal
	// fmt.Printf("Visit this URL: \n %v \n", url)

	// // Grabs the authorization code you paste into the terminal
	// var code string
	// _, err = fmt.Scan(&code)
	// if err != nil {
	// 	log.Printf("Error: %v", err)
	// }

	// // Exchange the auth code for an access token
	// tok, err := conf.Exchange(oauth2.NoContext, code)
	// if err != nil {
	// 	log.Printf("Error: %v", err)
	// }

	// Create the *http.Client using the access token
	//client := conf.Client(oauth2.NoContext, tok)
	client := getClient(conf)
	// Create a new gmail service using the client
	gmailService, err := gmail.New(client)
	if err != nil {
		log.Printf("Error: %v", err)
	}

	// New message for our gmail service to send
	var message gmail.Message

	// Compose the message
	//	messageStr := []byte(
	// "From: 集智台灣平台 <test@gmail.com>\r\n" +
	// 	"To: lewisli@gmail.com\r\n" +
	// 	"Subject: 您的驗證碼為\r\n\r\n" +
	// 	"1234567")

	header := make(map[string]string)
	//	header["From"] = encodeRFC2047("1234 ", "<sanpeople2020@gmail.com>")
	header["From"] = mime.QEncoding.Encode("utf-8", "台灣平台") + " <test@gmail.com>"
	header["To"] = "lewisli@gmail.com"
	header["Subject"] = mime.QEncoding.Encode("utf-8", "您的驗證碼為")
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = fmt.Sprintf("%s; charset=\"utf-8\"", "text/plain")
	//header["Content-Transfer-Encoding"] = "base64"
	header["Content-Transfer-Encoding"] = "quoted-printable"
	var msg string
	for k, v := range header {
		msg += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	msg += "\r\n" + "1234567"
	fmt.Println(msg)
	// Place messageStr into message.Raw in base64 encoded format
	//message.Raw = base64.URLEncoding.EncodeToString([]byte(msg))
	message.Raw = encodeWeb64String([]byte(msg))
	fmt.Println(message.Raw)

	// Send the message
	_, err = gmailService.Users.Messages.Send("me", &message).Do()
	if err != nil {
		fmt.Println(err)
		log.Printf("Error: %v", err)
	} else {
		fmt.Println("Message sent!")
	}
}
func encodeRFC2047(name, address string) string {
	// use mail's rfc2047 to encode any string
	addr := mail.Address{name, address}
	return strings.Trim(addr.String(), " <>")
}
func encodeWeb64String(b []byte) string {

	s := base64.URLEncoding.EncodeToString(b)

	var i = len(s) - 1
	for s[i] == '=' {
		i--
	}

	return s[0 : i+1]
}
