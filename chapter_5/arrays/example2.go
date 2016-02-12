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

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	slice1 = []int{1, 2, 3}
	slice2 = make([]int, 2)
	fmt.Println(slice2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)

}
