package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		var oddity string
		if i%2 == 0 {
			oddity = "even"
		} else {
			oddity = "odd"
		}
		fmt.Printf("%d - %s\n", i, oddity)
	}
}
