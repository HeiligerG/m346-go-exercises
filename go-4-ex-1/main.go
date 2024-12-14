package main

import (
	"fmt"
)

// float32			//float32
func computerGrade(gotPoints float64, maxPoints float64) {
	grade := ((gotPoints / maxPoints) * 5) + 1
	fmt.Println("Grade:", grade)
}

func main() {
	computerGrade(17.5, 28.0) //Grade: 4.125
	computerGrade(25.5, 28.0) //Grade: 5.553571428571429
	computerGrade(10.5, 28.0) //Grade: 2.875

	// TODO: call the function computeGrade
}
