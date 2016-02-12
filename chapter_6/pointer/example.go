package main

import "fmt"

func fakeZero(x int) {
	x = 0
}

func zero(xPtr *int) {
	*xPtr = 0
}

func main() {
	x := 5
	fakeZero(x)
	fmt.Println(x)
	zero(&x)
	fmt.Println(x)
}
