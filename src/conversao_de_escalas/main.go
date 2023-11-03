package main

import (
	"fmt"
)

func main() {
	var kelvin float64

	fmt.Print("Digite a temperatura em Kelvin: ")
	fmt.Scan(&kelvin)

	celsius := convertToCelsius(float64(kelvin))
	fahrenheit := convertToFahrenheit(float64(kelvin))

	fmt.Printf("%.2f Kelvin é equivalente a %.2f Celsius e %.2f Fahrenheit.\n", kelvin, celsius, fahrenheit)
}

func convertToCelsius( kelvin float64) float64{
	return  kelvin - 273.15
}

func convertToFahrenheit(kelvin float64)  float64{
	return (kelvin-273.15)*9/5 + 32
}