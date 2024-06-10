package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"encoding/base64"
)

func main() {
	var host string
	var uri string
	var useSSL bool
	var username string
	var password string

	flag.StringVar(&host, "H", "localhost", "Elasticsearch host")
	flag.StringVar(&uri, "u", "/", "Elasticsearch URI")
	flag.BoolVar(&useSSL, "ssl", false, "Use SSL")
	flag.StringVar(&username, "username", "", "Elasticsearch username")
	flag.StringVar(&password, "password", "", "Elasticsearch password")
	flag.Parse()

	protocol := "http"
	if useSSL {
		protocol = "https"
	}

	url := fmt.Sprintf("%s://%s:%s", protocol, host, uri)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		os.Exit(2)
	}

	if username != "" && password != "" {
		auth := username + ":" + password
		b64 := base64.StdEncoding.EncodeToString([]byte(auth))
		req.Header.Set("Authorization", "Basic "+b64)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("CRITICAL - Unable to reach Elasticsearch:", err)
		os.Exit(2)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("OK - Elasticsearch is available")
		os.Exit(0)
	} else {
		fmt.Printf("CRITICAL - Elasticsearch returned status code %d\n", resp.StatusCode)
		os.Exit(2)
	}
}
