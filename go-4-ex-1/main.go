package main

import (
	"fmt"
)

var Grade float64 = 0.0

// TODO: implement the function computeGrade
func computeGrade(gotPoints float64, maxPoints float64) {
	Grade := ((float64(gotPoints) / float64(maxPoints)) * 5) + 1
	fmt.Println(Grade)
}

func main() {
	// TODO: call the function computeGrade

	computeGrade(17.5, 28.0) // 4.125
	computeGrade(20.0, 25.0) //5
	computeGrade(1.0, 10.0)  // 1.5
}
