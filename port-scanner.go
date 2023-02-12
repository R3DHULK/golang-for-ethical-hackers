package main

import (
	"fmt"
	"net"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Enter a host to scan:")
	var host string
	fmt.Scanln(&host)

	fmt.Println("Enter the range of ports to scan (e.g. 1-1024):")
	var portRange string
	fmt.Scanln(&portRange)

	fmt.Println("Scanning...")
	ports := getPorts(portRange)
	for _, port := range ports {
		address := fmt.Sprintf("%s:%d", host, port)
		conn, err := net.DialTimeout("tcp", address, time.Second*10)
		if err == nil {
			fmt.Printf("%d open\n", port)
			conn.Close()
		}
	}
	fmt.Println("Scan complete.")
}

func getPorts(portRange string) []int {
	ports := []int{}
	parts := strings.Split(portRange, "-")
	start, err := strconv.Atoi(parts[0])
	if err != nil {
		return ports
	}
	end, err := strconv.Atoi(parts[1])
	if err != nil {
		return ports
	}
	for i := start; i <= end; i++ {
		ports = append(ports, i)
	}
	sort.Ints(ports)
	return ports
}
