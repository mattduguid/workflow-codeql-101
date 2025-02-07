package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

func insecureTLS() {
	// 🚨 Vulnerability: Insecure TLS configuration
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, _ := client.Get("https://example.com")
	defer resp.Body.Close()

	fmt.Println("Insecure TLS request sent")
}
