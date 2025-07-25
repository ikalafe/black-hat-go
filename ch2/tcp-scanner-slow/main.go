package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 1; i <= 1024; i++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			continue
		}
		connectionErr := conn.Close()
		fmt.Println("Connection Error: ", connectionErr)
		fmt.Printf("%d open\n", i)
	}
}
