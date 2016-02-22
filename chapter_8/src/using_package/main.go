package main

import "fmt"
import "alex/math"

func main() {
	xs := []float64{1, 2, 3, 4, 5}
	fmt.Println("Series:", xs)
	fmt.Println("Average:", math.Average(xs))
	fmt.Println("Min:", math.Min(xs))
	fmt.Println("Max:", math.Max(xs))
}
