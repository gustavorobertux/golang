package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func main() {

	var IP string
	var PI string

	fmt.Print("TARGET IP: ")
	fmt.Scan(&IP)

	fmt.Print("TARGET PORT LIST, e.g 21,80,443,: ")
	fmt.Scan(&PI)

	PORTS := strings.Split(PI, ",")

	for _, PORT := range PORTS {

		TARGET := fmt.Sprintf("%s:%s", IP, PORT)

		_, err := net.DialTimeout("tcp", TARGET, 600*time.Millisecond)

		if err != nil {
			fmt.Printf("Port %s is closed\n", PORT)
		} else {
			fmt.Printf("Port %s is open\n", PORT)
		}
	}
}
