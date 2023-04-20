package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
    // Connect to the remote server at address "remotehost:8000"
    conn, err := net.Dial("tcp", "golang.org:80")
    if err != nil {
        fmt.Println("Error connecting:", err)
        os.Exit(1)
    }
    defer conn.Close()

    // Send a request to the server
    fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")

    // Read the response from the server
    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
}