package main

import "fmt"

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	max := x[0]
	for _, i := range x {
		fmt.Println("i:", i)
		if i > max {
			max = i
		}
	}
	fmt.Println(max)
}
