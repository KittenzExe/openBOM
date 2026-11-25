package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
)

var jar http.CookieJar

func init() {
	var err error
	jar, err = cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
	}
}

func proxyBomRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Accept-Encoding", "")

	client := &http.Client{Jar: jar}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("BOM returned %d: %s", resp.StatusCode, string(body))
		return nil, fmt.Errorf("BOM API error: status %d", resp.StatusCode)
	}

	return body, nil
}
