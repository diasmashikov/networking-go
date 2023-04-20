package main

import (
	"fmt"
	"net"
	"net/http"
)


func handleConnection(conn net.Conn) {
    defer conn.Close()

    buf := make([]byte, 1024)
    for {
        n, err := conn.Read(buf)
        if err != nil {
            fmt.Println("Error reading:", err.Error())
            return
        }

        fmt.Println("Received message:", string(buf[:n]))
        _, err = conn.Write([]byte("Message received\n"))
        if err != nil {
            fmt.Println("Error writing:", err.Error())
            return
        }
    }
}

func tcpServer() {
	listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        return
    }
    defer listener.Close()

    fmt.Println("Listening on port 8080")

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting:", err.Error())
            continue
        }

        go handleConnection(conn)
    }
}

func openTcpServer() {
	listener, err := net.Listen("tcp", "0.0.0.0:1234")
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        return
    }
    defer listener.Close()
    fmt.Println("Listening on", listener.Addr())

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err.Error())
            continue
        }
        defer conn.Close()
        fmt.Println("Accepted connection from", conn.RemoteAddr())

        // Handle incoming connection here...
    }
}

func tcpAndHttpServer() {
	// Create a TCP listener on port 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}

	fmt.Println("Listening on port 8080")

	// Create an HTTP server that listens on the TCP listener
	httpServer := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Handle the HTTP request and send a response
			fmt.Fprintf(w, "Hello, World!")
		}),
	}

	// Start the HTTP server and handle incoming requests
	err = httpServer.Serve(listener)
	if err != nil {
		fmt.Println("Error serving:", err.Error())
		return
	}
}


func main() {
	openTcpServer()
}

