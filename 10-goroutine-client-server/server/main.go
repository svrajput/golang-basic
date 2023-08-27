package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	//write server program to handle concurrent client connections.
	listner, err := net.Listen("tcp", "localhost:8000")

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listner.Accept()

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(" Server is started ... ")
		go handleConn(conn)
	}

}

// handleConn - utility function
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, "response from server\n")
		if err != nil {
			return
		}
		time.Sleep(time.Second)
	}
}
