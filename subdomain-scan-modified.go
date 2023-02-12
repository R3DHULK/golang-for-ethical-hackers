package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	// Set the base domain and the list of subdomains to check
	fmt.Println("Enter a domain to scan (without https, http, www): ")
	var baseDomain string
	fmt.Scanln(&baseDomain)
	subdomains := []string{"www", "mail", "ftp","docs", "localhost", "webmail", "smtp","pop","ns1","webdisk","ns2","cpanel","whm","autodiscover","autoconfig","m","imap","test","ns","blog","pop3","dev","www2","admin","forum","news","vpn","ns3","mail2","new","mysql","old","lists","support","mobile","mx","static","docs","beta","shop","sql","secure","demo","cp","calendar","wiki","web","media","email","images","img","www1","intranet","portal","video","sip","dns2","api","cdn","stats","dns1","ns4","www3","dns","search","staging","server","mx1","chat","wap","my","svn","mail1","sites","proxy","ads","host","crm","cms","backup","mx2","lyncdiscover","info","apps","download","remote","db","forums","store","relay","files","newsletter","app","live","owa","en","start","sms","office","exchange","ipv4","Footer"}

	// Check each subdomain
	for _, subdomain := range subdomains {
		// Construct the full domain name
		domain := subdomain + "." + baseDomain

		// Check if the domain is reachable
		address, err := net.LookupHost(domain)
		if err != nil {
			fmt.Println(domain, "is unreachable")
			continue
		}

		// Print the domain and its IP addresses
		fmt.Println(domain, "has address(es)", strings.Join(address, ", "))
	}
}
