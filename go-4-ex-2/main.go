package main

import (
	"fmt"
	"math"
)

func computeHypotenuse(a float64, b float64) {
	aSquare, bSquare := math.Pow(a, 2), math.Pow(b, 2)
	cHyp := math.Sqrt(aSquare + bSquare)
	fmt.Println(cHyp)
}

// TODO: implement the function computeHypotenuse using math.Pow and math.Sqrt

func main() {
	// TODO: call the function computeHypotenuse
	computeHypotenuse(2.4, 4.1) //4.750789408087881
}
