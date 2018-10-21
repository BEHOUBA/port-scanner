package main

import (
	"fmt"
	"net"
	"strconv"
)

func checkPort(port int) (isOpen bool) {
	_, err := net.Dial("tcp", "127.0.0.1:"+strconv.Itoa(port))
	if err != nil {
		isOpen = false
	} else {
		isOpen = true
	}
	return
}

func main() {
	fmt.Println("Started scanning all ports...")

	p := make(chan int)
	for i := 0; i <= 10000; i++ {
		port := i
		go func(int) {
			if isOpen := checkPort(port); isOpen == true {
				fmt.Printf("port %d is opend \n", port)
			}
			p <- port
		}(port)
	}
	for i := 0; i <= 10000; i++ {
		<-p
		if i == 10000 {
			close(p)
		}
	}
}
