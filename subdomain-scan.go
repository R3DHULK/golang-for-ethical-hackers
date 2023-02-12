package main

import (
	"fmt"
	"net"
)

func main() {
	// Set the domain to scan for subdomains
	domain := "example.com"

	// Set the subdomains to check
	subdomains := []string{"www", "mail", "ftp", "docs", "localhost", "webmail", "smtp","pop","ns1","webdisk","ns2","cpanel","whm","autodiscover","autoconfig","m","imap","test","ns","blog","pop3","dev","www2","admin","forum","news","vpn","ns3","mail2","new","mysql","old","lists","support","mobile","mx","static","docs","beta","shop","sql","secure","demo","cp","calendar","wiki","web","media","email","images","img","www1","intranet","portal","video","sip","dns2","api","cdn","stats","dns1","ns4","www3","dns","search","staging","server","mx1","chat","wap","my","svn","mail1","sites","proxy","ads","host","crm","cms","backup","mx2","lyncdiscover","info","apps","download","remote","db","forums","store","relay","files","newsletter","app","live","owa","en","start","sms","office","exchange","ipv4","Footer"}

	// Check each subdomain
	for _, subdomain := range subdomains {
		// Construct the full domain name
		fullDomain := subdomain + "." + domain

		// Perform a DNS lookup for the domain
		_, err := net.LookupHost(fullDomain)
		if err == nil {
			// The domain was found, so it exists
			fmt.Println(fullDomain + " exists")
		} else {
			// The domain was not found
			fmt.Println(fullDomain + " does not exist")
		}
	}
}

