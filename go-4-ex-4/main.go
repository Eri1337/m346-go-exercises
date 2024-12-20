package main

import (
	"fmt"
)

var C float64 = 0.0
var F float64 = 0.0

// TODO: implement the function convertCelsiusToFahrenheit
func convertCelsiusToFahrenheit(C float64) {
	F := (C * (9.0 / 5.0)) + 32.0 // hat vorher wegen fehlendem .0 falschen wert gegeben
	//Tip: bei rechnung immer gleichen Typ nutzen (hier float)
	fmt.Println(F)
}

func convertFahrenheitToCelsius(F float64) {
	C := (F - 32.0) * (5.0 / 9.0)

	fmt.Println(C)
}

// TODO: implement the function convertFahrenheitToCelsius

func main() {
	// TODO: call the function convertCelsiusToFahrenheit
	convertCelsiusToFahrenheit(10)  //50
	convertCelsiusToFahrenheit(100) //212
	convertCelsiusToFahrenheit(50)  //122

	// TODO: call the function convertFahrenheitToCelsius
	convertFahrenheitToCelsius(50)  //10
	convertFahrenheitToCelsius(212) //100
	convertFahrenheitToCelsius(122) //50

}
