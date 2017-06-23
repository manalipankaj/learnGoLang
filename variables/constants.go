package main

import "fmt"

// Can be declared directly
const PI = 3.14

// Can be declared inside block
const (
	ONE = 1
	TWO = 2
	THREE = 3
	FOUR = 4
	FIVE = 5
	SIX = 6
)

func main() {
	// Can be declared inside function
	const MAIN = "this is main function"
	fmt.Println(ONE, TWO, THREE, FOUR, FIVE, SIX, PI, MAIN)
}