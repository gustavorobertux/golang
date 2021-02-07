package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	var IP string
	var PORT string

	fmt.Print("TARGET IP: ")
	fmt.Scan(&IP)

	fmt.Print("TARGET PORT: ")
	fmt.Scan(&PORT)

	TARGET := fmt.Sprintf("%s:%s", IP, PORT)

	fmt.Println(TARGET)

// Change the value to 10 if you want to use it on your local network.

_, err := net.DialTimeout("tcp", TARGET, 600*time.Millisecond)

	if err != nil {
		fmt.Printf("Port %s is closed\n", PORT)
	} else {
		fmt.Printf("Port %s is open\n", PORT)
	}
}
