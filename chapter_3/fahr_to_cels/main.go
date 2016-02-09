package main

import "fmt"

func main() {
	fmt.Printf("Enter a temperature (°F): ")
	var temperatureFahr float64
	fmt.Scanf("%f", &temperatureFahr)
	var temperatureCels float64 = (temperatureFahr - 32.0) * (5.0 / 9.0)
	fmt.Printf("%.2f °C\n", temperatureCels)
}
