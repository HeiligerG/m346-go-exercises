package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var eyes = rand.Intn(6) + 1
	var when = time.Now()

	eyesFile, err1 := os.Create("eyes.txt")
	if err1 != nil {
		fmt.Println("Error creating eyes.txt:", err1)
		return
	}
	defer eyesFile.Close()

	logFile, err2 := os.Create("dice.log")
	if err2 != nil {
		fmt.Println("Error creating dice.log:", err2)
		return
	}
	defer logFile.Close()
	// TODO: use fmt.Fprintln instead!

	fmt.Fprintln(os.Stdout, "the dice shows", eyes, "eyes")
	fmt.Fprintln(os.Stdout, "the dice was rolled at", when)

	fmt.Fprintln(eyesFile, eyes)
	fmt.Fprintln(logFile, "the dice shows", eyes, "eyes")
	fmt.Fprintln(logFile, "the dice was rolled at", when)

	fmt.Println("Output written to eyes.txt and dice.log")
}
