package main

import (
	"fmt"
	"math"
)

var a float64 = 0.0
var b float64 = 0.0
var c float64 = 0.0
var D float64 = 0.0

// TODO: implement the function computeQuadraticFormula
func computeQuadraticFormula(a float64, b float64, c float64) {
	D := math.Pow(b, 2) - 4*a*c

	if D > 0 {
		solution1 := (-b + math.Sqrt(D)) / (2 * a)
		solution2 := (-b - math.Sqrt(D)) / (2 * a)
		fmt.Println(solution1, solution2)
	} else if D == 0 {
		solution := -b / (2 * a)
		fmt.Println(solution)
	} else {
		fmt.Println("Keine Lösung")
	}
}
func main() {
	// TODO: call the function computeQuadraticFormula
	computeQuadraticFormula(2.0, 3.0, 4.0) //keine Lösung
}
