package main

import (
	"fmt"
	"net"
)

func main() {
	// Perform a forward DNS lookup for the domain example.com
	addrs, err := net.LookupHost("example.com")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the results
	fmt.Println("Addresses:")
	for _, addr := range addrs {
		fmt.Println(addr)
	}
}
