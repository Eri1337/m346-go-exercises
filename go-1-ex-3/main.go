package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var eyes = rand.Intn(5) + 1
	var when = time.Now()

	// TODO: use fmt.Fprintln instead!
	fmt.Fprintln(os.Stdout, "the dice shows", eyes, "eyes")

	// TODO: use fmt.Fprintln instead!
	fmt.Fprintln(os.Stdout, "the dice was rolled at", when)

	// TODO: how to write the output into eyes.txt and dice.log?
	eyesFile, _ := os.Create("eyes.txt") //ohne _ funktioniert es nicht, weiss nicht warum
	fmt.Fprintln(eyesFile, "the dice shows", eyes, "eyes")
	eyesFile.Close()

	logFile, _ := os.Create("dice.log") //ohne _ funktioniert es nicht, weiss nicht warum
	fmt.Fprintln(logFile, "the dice shows", eyes, "eyes")
	fmt.Fprintln(logFile, "the dice was rolled at", when)
	logFile.Close()

	// go run ex3/main.go TODO
}
