package main

import "github.com/xmha97/rasad/internal/checker"

func main() {
	sites := checker.DefaultSites()
	for _, s := range sites {
		checker.SendRequest(s.Name, s.URL, 0)
	}
}
