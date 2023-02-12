package main

import (
    "fmt"
    "net"
    "os"
    "time"
)

func main() {
    host := os.Args[1]
    port := os.Args[2]

    // Connect to the target host and port
    conn, err := net.DialTimeout("tcp", host+":"+port, 5*time.Second)
    if err != nil {
        fmt.Println("Error:", err.Error())
        return
    }
    defer conn.Close()

    // Send the request to the server
    fmt.Fprintf(conn, "GET / HTTP/1.1\r\nHost: %s\r\n\r\n", host)

    // Read the response from the server
    buffer := make([]byte, 1024)
    n, _ := conn.Read(buffer)

    // Print the response
    response := string(buffer[:n])
    fmt.Println(response)
}
