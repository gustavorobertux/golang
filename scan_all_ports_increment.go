package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	var IP string

	fmt.Print("TARGET IP: ")
	fmt.Scan(&IP)

	for PORT := 1; PORT < 65535; PORT++ {

		TARGET := fmt.Sprintf("%s:%d", IP, PORT)

		// Change the value to 10 if you want to use it on your local network.

		_, err := net.DialTimeout("tcp", TARGET, 600*time.Millisecond)

		if err != nil {
		// comment the line below if you want to see only the open ports
			fmt.Printf("Port %d is closed\n", PORT)
		} else {
			fmt.Printf("Port %d is open\n", PORT)
		}
	}
}
