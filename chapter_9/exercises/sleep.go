package main

import (
	"fmt"
	"time"
)

func Sleep(delay int64) time.Time {
	timestamp := <-time.After(time.Duration(delay) * time.Second)
	return timestamp
}

func main() {
	fmt.Println("Calling Sleep(3)")
	fmt.Println("Start timestamp:", time.Now())
	res := Sleep(3)
	fmt.Println("End timestamp:  ", res)

}
