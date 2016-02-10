package main

import "fmt"

func main() {
	x := [5]float64{89, 93, 77, 82, 83}
	y := [5]float64{
		89,
		93,
		77,
		82,
		83,
	}
	for _, i := range x {
		fmt.Println(i)
	}
	for _, i := range y {
		fmt.Println(i)
	}
	alf := x[0:3]
	fmt.Println(alf)
}
