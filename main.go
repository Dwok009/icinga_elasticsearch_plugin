package main

import (
    "flag"
    "fmt"
    "net/http"
    "os"
)

func main() {
    var host string
    var port string
    var uri string
    var useSSL bool

    flag.StringVar(&host, "H", "localhost", "Elasticsearch host")
    flag.StringVar(&port, "P", "9200", "Elasticsearch port")
    flag.StringVar(&uri, "u", "/", "Elasticsearch URI")
    flag.BoolVar(&useSSL, "ssl", false, "Use SSL")
    flag.Parse()

    protocol := "http"
    if useSSL {
        protocol = "https"
    }

    url := fmt.Sprintf("%s://%s:%s%s", protocol, host, port, uri)
    resp, err := http.Get(url)
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