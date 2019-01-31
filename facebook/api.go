package facebook

//https://ithelp.ithome.com.tw/articles/10204840
//go get golang.org/x/oauth2
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

var (
	oauthConf = &oauth2.Config{
		ClientID:     "1774734012835603",
		ClientSecret: "48d0fb010742b5a78f68d31d85810186",
		RedirectURL:  "http://localhost:8080/oauth2callback",
		Scopes:       []string{"public_profile"},
		Endpoint:     facebook.Endpoint,
	}
	oauthStateString = "gamilms"
)

const htmlIndex = `<html><body>
Logged in with <a href="/login">facebook</a>
</body></html>
`

func handleMain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(htmlIndex))
}

func handleFacebookLogin(w http.ResponseWriter, r *http.Request) {
	URL, err := url.Parse(oauthConf.Endpoint.AuthURL)
	if err != nil {
		log.Fatal("Parse: ", err)
	}
	parameters := url.Values{}
	parameters.Add("client_id", oauthConf.ClientID)
	parameters.Add("scope", strings.Join(oauthConf.Scopes, " "))
	parameters.Add("redirect_uri", oauthConf.RedirectURL)
	parameters.Add("response_type", "code")
	parameters.Add("state", oauthStateString)
	URL.RawQuery = parameters.Encode()
	url := URL.String()
	fmt.Println(url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleFacebookCallback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")
	fmt.Println("code: %s\n", code)
	token, err := oauthConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Printf("oauthConf.Exchange() failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	fmt.Println("token: %s\n", token)
	fmt.Println("token.AccessToken: %s\n", token.AccessToken)
	fmt.Println("token.RefreshToken: %s\n", token.RefreshToken)
	fmt.Println("token.Expiry: %s\n", token.Expiry)

	resp, err := http.Get("https://graph.facebook.com/me?access_token=" +
		url.QueryEscape(token.AccessToken))
	if err != nil {
		fmt.Printf("Get: %s\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ReadAll: %s\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	var mapResult map[string]interface{}
	//使用 json.Unmarshal(data []byte, v interface{})进行转换,返回 error 信息
	if err := json.Unmarshal([]byte(response), &mapResult); err != nil {
		fmt.Println("error: %s\n", err)
	}

	if val, ok := mapResult["name"]; ok {
		fmt.Println("login success:" + val.(string))
	} else {
		fmt.Println("login fail")
	}

	//fmt.Printf("parseResponseBody: %s\n", string(response))
	fmt.Println("parseResponseBody: %s\n", string(response))

	//getUserAvator
	getUserAvator(token.AccessToken)

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func getUserAvator(accessToken string) {
	resp, err := http.Get("https://graph.facebook.com/v2.11/292218118227980/picture?redirect=false&access_token=" + url.QueryEscape(accessToken))
	if err != nil {
		fmt.Printf("Get: %s\n", err)
		return
	}
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ReadAll: %s\n", err)
		return
	}

	//fmt.Printf("parseResponseBody: %s\n", string(response))
	fmt.Println("parseResponseBody: %s\n", string(response))

	var mapResult map[string]map[string]interface{}
	//使用 json.Unmarshal(data []byte, v interface{})进行转换,返回 error 信息
	if err := json.Unmarshal([]byte(response), &mapResult); err != nil {
		fmt.Println("error: %s\n", err)
	}

	if val, ok := mapResult["data"]["url"]; ok {
		fmt.Println("avatar success:" + val.(string))
	} else {
		fmt.Println("avatar fail")
	}
}

func FacebookFlow() {
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/login", handleFacebookLogin)
	http.HandleFunc("/oauth2callback", handleFacebookCallback)
	fmt.Print("Started running on http://localhost:8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
