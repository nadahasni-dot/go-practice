package main

import "fmt"

func main() {
	names := [...]string{"Nada", "Hasni", "Muhammad", "Eko", "Sandhika", "Kurniawan", "Budi", "Siti"} //array
	// names := []string{"Nada", "Hasni", "Muhammad", "Eko", "Sandhika", "Kurniawan", "Budi", "Siti"} //slice

	slice := names[4:6]
	fmt.Println(slice)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[4:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)

	fmt.Println("slice len", len(slice4))
	fmt.Println("slice capacity", cap(slice4))

	slice5 := append(slice4, "Rafi")
	slice5[0] = "Nada Baru"
	slice5[1] = "Hasni Baru"
	fmt.Println("New name", slice5)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Nada"
	newSlice[1] = "Hasni"
	// newSlice[2] = "Muhammad" error: should use append because exceed length array

	newSlice2 := append(newSlice, "Muhammad")

	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	fromSlice := names[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)
}
