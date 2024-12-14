package main

import (
	"fmt"
	"math"
)

// TODO: implement the function computeQuadraticFormula

func computeDiskrimante(a, b, c float64) float64 {
	return math.Pow(b, 2) - 4*a*c
}

func computeQuadraticFormula(a, b, c float64) []float64 {
	D := computeDiskrimante(a, b, c)

	if D > 0 {
		x1 := (-b + math.Sqrt(D)) / (2 * a)
		x2 := (-b - math.Sqrt(D)) / (2 * a)
		return []float64{x1, x2}
	} else if D == 0 {
		x := (-b + math.Sqrt(D)) / (2 * a)
		return []float64{x}
	} else {
		return []float64{}
	}
}

func main() {

	// TODO: call the function computeQuadraticFormula
	testCases := []struct {
		a, b, c float64
	}{
		{3, 4, 1}, // Diskriminante > 0: zwei Lösungen
		{2, 4, 2}, // Diskriminante = 0: eine Lösung
		{3, 4, 2}, // Diskriminante < 0: keine Lösung
	}

	for _, tc := range testCases {
		D := computeDiskrimante(tc.a, tc.b, tc.c)
		solutions := computeQuadraticFormula(tc.a, tc.b, tc.c)

		fmt.Printf("Für a = %.2f, b = %.2f, c = %.2f:\n", tc.a, tc.b, tc.c)
		fmt.Printf("Diskriminante: %.2f\n", D)
		if len(solutions) == 2 {
			fmt.Printf("Zwei Lösungen: x1 = %.2f, x2 = %.2f\n\n", solutions[0], solutions[1])
		} else if len(solutions) == 1 {
			fmt.Printf("Eine Lösung: x = %.2f\n\n", solutions[0])
		} else {
			fmt.Printf("Keine Lösung\n\n")
		}
	}
}

//Für a = 3.00, b = 4.00, c = 1.00:
//Diskriminante: 4.00
//Zwei Lösungen: x1 = -0.33, x2 = -1.00
//
//Für a = 2.00, b = 4.00, c = 2.00:
//Diskriminante: 0.00
//Eine Lösung: x = -1.00
//
//Für a = 3.00, b = 4.00, c = 2.00:
//Diskriminante: -8.00
//Keine Lösung
