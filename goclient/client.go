package main

import "os"
import "net"
import "fmt"
import "time"

func main() {
	// connect to this socket
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:8081", os.Args[1]))
	if err != nil {
		panic(err)
	}
	total := int64(0)
	for i := 0; i < 100; i++ {
		data := make([]byte, 8)
		st := time.Now()
		n, err := conn.Write(data)
		if n != 8 || err != nil {
			fmt.Printf("failed to write the whole message %v", err)
			return
		}
		n, err = conn.Read(data)
		if n != 8 || err != nil {
			fmt.Printf("failed to read the whole messagei %v", err)
			return
		}
		total += time.Now().Sub(st).Nanoseconds()
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("avg latency %d nanoseconds (%d microseconds)\n",
		total/100, total/100000)
}
