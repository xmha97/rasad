package main

import (
	"fmt"
	"net/http"
	"strings"
)

type site struct {
	name string
	url  string
}

func main() {
	sites := []site{
		{name: "Aparat", url: "https://www.aparat.com/"},
		{name: "Digikala", url: "https://www.digikala.com/"},
		{name: "Divaar", url: "https://divar.ir/"},
		{name: "Bale", url: "https://ble.ir/"},
		{name: "Eitaa", url: "https://eitaa.com/"},
		{name: "Zarebin", url: "https://zarebin.ir/"},
		{name: "Varzesh3", url: "https://www.varzesh3.com/"},
		{name: "Shaparak", url: "https://shaparak.ir/"},
		{name: "Blogfa", url: "https://www.blogfa.com/"},
		{name: "Telewebion", url: "https://telewebion.com/"},
		{name: "Google", url: "https://www.google.com/"},
		{name: "Google for Developers", url: "https://developers.google.com/"},
		{name: "Dart packages", url: "https://pub.dev/"},
		{name: "WhatsApp", url: "https://www.whatsapp.com/"},
		{name: "ChatGPT", url: "https://chatgpt.com/"},
		{name: "Instagram", url: "https://www.instagram.com/"},
		{name: "Facebook", url: "https://www.facebook.com/"},
		{name: "Discord", url: "https://discord.com/"},
		{name: "Telegram", url: "https://telegram.org/"},
		{name: "PornHub", url: "https://www.pornhub.com/"},
	}

	for _, s := range sites {
		sendRequest(s.name, s.url, 0)
	}
}

func sendRequest(name string, url string, count int) {
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
		if strings.Contains(err.Error(), "dial tcp 10.10.34.35:443: connectex: No connection could be made because the target machine actively refused it.") {
			fmt.Println("❌ Blocked", "|", name)
		} else if strings.Contains(err.Error(), "net/http: TLS handshake timeout") {
			fmt.Println("❌ Timeout", "|", name)
			if (count) < 4 {
				sendRequest(name, url, count+1)
			}
		} else if strings.Contains(err.Error(), "An existing connection was forcibly closed by the remote host.") {
			fmt.Println("❌ Closed", "|", name)
		} else if strings.Contains(err.Error(), "no such host") {
			fmt.Println("❌ NoHost", "|", name)
		} else {
			fmt.Println("❌ SendRequest:", "|", name, "|", err)
		}
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
