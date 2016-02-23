package main

import (
	"fmt"
	"time"
)

func writer(c chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Println("Sending", i)
		c <- i
		fmt.Println("Sent", i)
	}
}

func reader(c <-chan int) {
	for {
		select {
		case res := <-c:
			fmt.Println("Received", res)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	bufchan := make(chan int, 3)
	go writer(bufchan)
	go reader(bufchan)
	var input string
	fmt.Scanln(&input)

}
