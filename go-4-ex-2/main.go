package main

import (
	"fmt"
	"math"
)

var c float64 = 0.0

// TODO: implement the function computeHypotenuse using math.Pow and math.Sqrt
func computeHypotenuse(a float64, b float64) {
	c := math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))

	fmt.Println(c)
}

func main() {
	// TODO: call the function computeHypotenuse

	computeHypotenuse(10, 5) // 11.18...
	computeHypotenuse(7, 4)  // 8.062...
	computeHypotenuse(9, 6)  // 10.816...

}
