package main

import "fmt"

func main() {
	var name string
	name = "Nada Hasni Muhammad"
	fmt.Println("My Name is", name)

	var age uint8
	age = 24
	fmt.Println("I am", age, "years old")

	var city = "Jember"
	fmt.Println("I live in", city)

	hobby := "Code ♥️"
	fmt.Println("I love to", hobby)

	var (
		movieName        = "The lord of the rings"
		movieDescription = "film about war between multiple races in middle earth"
	)

	fmt.Println("I love to watch", movieName)
	fmt.Println("It's about", movieDescription)
}
