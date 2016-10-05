package main

import (
	"fmt"
	"math"
)

func add(args ...float64) float64 {
	total := 0.0
	for _, arg := range args {
		total += arg
	}
	return total
}

func main() {
	fmt.Printf("Result: %.2f\n", add(1.0, 2.7, -3.0, math.Sqrt(3)))
}
