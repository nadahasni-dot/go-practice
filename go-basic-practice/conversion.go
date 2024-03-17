package main

import "fmt"

func main() {
	var value32 int32 = 32768
	var value64 int64 = int64(value32)
	var value16 int16 = int16(value32)

	fmt.Println(value32)
	fmt.Println(value64)
	fmt.Println(value16) // this will result in -32768, because int16 can only store value between -32768 till 32767

	var name = "Nada Hasni Muhammad"
	var n = name[0]
	var nString = string(n)

	fmt.Println("My name is", name)
	fmt.Println("First letter of my name is", n)
	fmt.Println("First letter of my name in string is", nString)
}
