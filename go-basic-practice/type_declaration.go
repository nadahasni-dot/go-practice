package main

import "fmt"

func main() {
	type NoKTP string

	var myKtp NoKTP = "1234123412341234"

	var example string = "4321432143214321"
	var exampleKtp NoKTP = NoKTP(myKtp)

	fmt.Println(example)
	fmt.Println(exampleKtp)
}
