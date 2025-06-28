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
		fmt.Println("❌ NewRequest", "|", name, "|", err)
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
		fmt.Println("✔️  Not Modified", "|", name)
	case http.StatusOK:
		fmt.Println("✔️  OK", "|", name)
	case http.StatusForbidden:
		fmt.Println("❌ Forbidden", "|", name)
	default:
		fmt.Println("❌ Error", "|", name, "|", resp.StatusCode, http.StatusText(resp.StatusCode))
	}
}

func handleError(name, url string, count int, err error) {
	msg := err.Error()
	switch {
	case strings.Contains(msg, "dial tcp") && strings.Contains(msg, "actively refused"):
		fmt.Println("❌ Blocked", "|", name)
	case strings.Contains(msg, "TLS handshake timeout"):
		fmt.Println("❌ Timeout", "|", name)
		if count < 4 {
			SendRequest(name, url, count)
		}
	case strings.Contains(msg, "forcibly closed"):
		fmt.Println("❌ Closed", "|", name)
	case strings.Contains(msg, "no such host"):
		fmt.Println("❌ NoHost", "|", name)
	default:
		fmt.Println("❌ SendRequest:", "|", name, "|", err)
	}
}
