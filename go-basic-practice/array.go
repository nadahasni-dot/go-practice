package main

import "fmt"

func main() {
	var peoplesName [3]string

	peoplesName[0] = "Budi"
	peoplesName[1] = "Andi"
	peoplesName[2] = "Siti"
	// peoplesName[3] = "Siti" // result in error, out of bond [0,1,2]

	fmt.Println(peoplesName)

	var peoplesAge = [...]int{
		20,
		21,
		22,
	}

	fmt.Println(peoplesAge)

	// array function
	fmt.Println("Arr length", len(peoplesName))
	fmt.Println("[1]'s name", peoplesName[1])
}
