package main

import "net"
import "fmt"

func main() {
	fmt.Println("Launching server...")
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}
	conn, err := ln.Accept()
	if err != nil {
		panic(err)
	}
	data := make([]byte, 8)
	for {
		n, err := conn.Read(data)
		if n != 8 || err != nil {
			fmt.Printf("failed to read the entire message %v", err)
			return
		}
		n, err = conn.Write(data)
		if n != 8 || err != nil {
			fmt.Printf("failed to send the message back %v", err)
			return
		}
	}
}
