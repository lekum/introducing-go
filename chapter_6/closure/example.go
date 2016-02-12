package main

import "fmt"

func makeEvenGenerator() func() uint {
	counter := uint(0)
	return func() uint {
		ret := counter
		counter += 2
		return ret
	}
}

func main() {
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
}
