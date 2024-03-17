package main

import "fmt"

func main() {
	const phi = 3.14
	// This will error, const can not be reassigned
	// phi = 10.0
	fmt.Println("Phi constant =", phi)

	const (
		length uint = 100
		height      = 200
	)

	fmt.Println("Cube length =", length)
	fmt.Println("Cube height =", height)
}
