package main

import (
	"fmt"
	"net"
	"syscall"
)

func main() {
	fmt.Println("Launching server...")
	addr, err := net.ResolveTCPAddr("tcp", ":8081")
	if err != nil {
		panic(err)
	}
	ln, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.AcceptTCP()
		if err != nil {
			panic(err)
		}
		f, err := conn.File()
		if err != nil {
			panic(err)
		}
		fd := int(f.Fd())
		go func() {
			data := make([]byte, 8)
			for {
				n, err := syscall.Read(fd, data)
				if n != 8 || err != nil {
					fmt.Printf("failed to read the entire message %v", err)
					return
				}
				n, err = syscall.Write(fd, data)
				if n != 8 || err != nil {
					fmt.Printf("failed to send the message back %v", err)
					return
				}
			}
		}()
	}
}
