package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	var IP string
	var PORTS = []int{80, 443, 444, 65, 3389}

	fmt.Print("TARGET IP: ")
	fmt.Scan(&IP)

	for _, PORT := range PORTS {

		TARGET := fmt.Sprintf("%s:%d", IP, PORT)

		_, err := net.DialTimeout("tcp", TARGET, 600*time.Millisecond)

		if err != nil {
			fmt.Printf("Port %d is closed\n", PORT)
		} else {
			fmt.Printf("Port %d is open\n", PORT)
		}
	}
}
