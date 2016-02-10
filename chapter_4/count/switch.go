package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		switch i {
		case 1:
			fmt.Println("First one!")
		case 5:
			fmt.Println("Fifth!")
		case 10:
			fmt.Println("Last!")
		default:
			fmt.Println("Boring one:", i)
		}
	}
}
