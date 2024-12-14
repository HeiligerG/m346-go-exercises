package main

import "fmt"

// TODO: implement the function convertFahrenheitToCelsius
func convertFahrenheitToCelsius(F float64) {
	C := (F - 32.0) * 5.0 / 9.0
	fmt.Println(C)
}

// TODO: implement the function convertCelsiusToFahrenheit
func convertCelsiusToFahrenheit(C float64) {
	F := (C * (9.0 / 5.0)) + 32.0
	fmt.Println(F)
}

func main() {
	// TODO: call the function convertCelsiusToFahrenheit
	convertCelsiusToFahrenheit(120) // 248
	convertCelsiusToFahrenheit(90)  // 194
	convertCelsiusToFahrenheit(20)  // 68
	// TODO: call the function convertFahrenheitToCelsius
	convertFahrenheitToCelsius(248) // 120
	convertFahrenheitToCelsius(194) // 90
	convertFahrenheitToCelsius(68)  // 20
}
