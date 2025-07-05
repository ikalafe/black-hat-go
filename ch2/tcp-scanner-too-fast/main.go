package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 65535; i++ {
		wg.Add(i)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("127.0.0.1:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			connError := conn.Close()
			fmt.Println("Connection Error: ", connError)
			fmt.Printf("%d open\n", j)
		}(i)
	}

	wg.Wait()
}
