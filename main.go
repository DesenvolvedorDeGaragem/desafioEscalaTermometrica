package main

import (
	"fmt"
)

func kelvinToCelsius(k float64) float64 {
	return k - 273.15
}

func main() {
	var kelvin float64
	fmt.Print("Digite a temperatura em Kelvin: ")
	fmt.Scanf("%f", &kelvin)

	celsius := kelvinToCelsius(kelvin)
	fmt.Printf("A temperatura em Celsius é: %.2f\n", celsius)
}
