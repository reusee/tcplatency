package main

import (
	"fmt"
	"net"
	"os"
	"syscall"
	"time"
)

func main() {
	// connect to this socket
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:8081", os.Args[1]))
	if err != nil {
		panic(err)
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		panic(err)
	}
	n := 1000000
	data := make([]byte, 8)
	st := time.Now()
	f, err := conn.File()
	if err != nil {
		panic(err)
	}
	fd := int(f.Fd())
	for i := 0; i < n; i++ {
		n, err := syscall.Write(fd, data)
		if n != 8 || err != nil {
			fmt.Printf("failed to write the whole message %v", err)
			return
		}
		n, err = syscall.Read(fd, data)
		if n != 8 || err != nil {
			fmt.Printf("failed to read the whole messagei %v", err)
			return
		}
	}
	fmt.Printf("%v\n", time.Since(st)/time.Duration(n))
}
