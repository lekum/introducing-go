package main

import "fmt"

func main() {
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64 = 0
	for i, value := range x {
		fmt.Println("Position:", i, "Value:", value)
		total += value
	}
	fmt.Println("Average:", total/float64(len(x)))
}
