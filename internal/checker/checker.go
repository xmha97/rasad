package checker

import (
	"fmt"
	"net/http"
	"strings"
)

type Site struct {
	Name string
	URL  string
}

func DefaultSites() []Site {
	return []Site{
		{Name: "Aparat", URL: "https://www.aparat.com/"},
		{Name: "Digikala", URL: "https://www.digikala.com/"},
		{Name: "Divaar", URL: "https://divar.ir/"},
		{Name: "Bale", URL: "https://ble.ir/"},
		{Name: "Eitaa", URL: "https://eitaa.com/"},
		{Name: "Zarebin", URL: "https://zarebin.ir/"},
		{Name: "Varzesh3", URL: "https://www.varzesh3.com/"},
		{Name: "Shaparak", URL: "https://shaparak.ir/"},
		{Name: "Blogfa", URL: "https://www.blogfa.com/"},
		{Name: "Telewebion", URL: "https://telewebion.com/"},
		{Name: "Google", URL: "https://www.google.com/"},
		{Name: "Google for Developers", URL: "https://developers.google.com/"},
		{Name: "Dart packages", URL: "https://pub.dev/"},
		{Name: "WhatsApp", URL: "https://www.whatsapp.com/"},
		{Name: "ChatGPT", URL: "https://chatgpt.com/"},
		{Name: "Instagram", URL: "https://www.instagram.com/"},
		{Name: "Facebook", URL: "https://www.facebook.com/"},
		{Name: "Discord", URL: "https://discord.com/"},
		{Name: "Telegram", URL: "https://telegram.org/"},
		{Name: "PornHub", URL: "https://www.pornhub.com/"},
	}
}

func SendRequest(name string, url string, count int) {
	count++
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("❌", name, "|", "NewRequest:", err)
		return
	}
	req.Header.Set("If-Modified-Since", "Wed, 21 Oct 2020 07:28:00 GMT")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		handleError(name, url, count, err)
		return
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusNotModified:
		fmt.Println("✔️", name, "|", "OK: NotModified")
	case http.StatusOK:
		fmt.Println("✔️", name, "|", "OK")
	case http.StatusForbidden:
		fmt.Println("❌", name, "|", "Forbidden")
	default:
		fmt.Println("❌", name, "|", "Error:", resp.StatusCode, http.StatusText(resp.StatusCode))
	}
}

func handleError(name, url string, count int, err error) {
	msg := err.Error()
	switch {
	case strings.Contains(msg, "dial tcp"):
		fmt.Println("❌", name, "|", "Blocked:", "Dial")
	case strings.Contains(msg, "read tcp"):
		fmt.Println("❌", name, "|", "Blocked:", "Read")
	case strings.Contains(msg, "TLS handshake timeout"):
		fmt.Println("❌", name, "|", "Timeout")
		if count < 4 {
			SendRequest(name, url, count)
		}
	case strings.Contains(msg, "forcibly closed"):
		fmt.Println("❌", name, "|", "Blocked:", "Closed")
	case strings.Contains(msg, "no such host"):
		fmt.Println("❌", name, "|", "Blocked:", "NoHost")
	default:
		fmt.Println("❌", name, "|", "SendRequest:", err)
	}
}
