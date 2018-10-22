package main

import "fmt"

func main() {

	host, firstPort, lastPort := getUserInput()

	fmt.Println("Started scanning all ports...")

	p := make(chan int)

	for i := firstPort; i <= lastPort; i++ {
		port := i
		go func(int) {
			if isOpen := checkPort(host, port); isOpen == true {
				fmt.Printf("port %d is opened \n", port)
			}
			p <- port
		}(port)
	}
	for i := firstPort; i <= lastPort; i++ {
		<-p
		if i == lastPort {
			close(p)
		}
	}

}
