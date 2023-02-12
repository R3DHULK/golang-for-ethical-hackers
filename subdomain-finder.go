package main

import (
    "fmt"
    "net"
    "os"
)

func main() {
    // Take domain name as input from the user
    fmt.Println("Enter the domain name: ")
    var domain string
    fmt.Scanln(&domain)

    subdomains := []string{}

    // Perform a DNS lookup for subdomains
    nsRecords, err := net.LookupNS(domain)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    // Append subdomains found to the subdomains slice
    for _, ns := range nsRecords {
        subdomains = append(subdomains, ns.Host)
    }

    // Print the subdomains found
    fmt.Println("Subdomains found:")
    fmt.Println(subdomains)
}
