package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Enter a search term:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	term := scanner.Text()

	// Perform a Google search for the search term
	resp, err := http.Get("https://www.google.com/search?q=" + term)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// Print the response status and header information
	fmt.Println("Response status:", resp.Status)
	for k, v := range resp.Header {
		fmt.Println(k, ":", v)
	}
}
