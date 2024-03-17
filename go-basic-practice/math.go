package main

import "fmt"

func main() {
	var a = 10
	var b = 10

	var plus = a + b
	var min = a - b
	var multiply = a * b
	var division = a / b
	var modulo = a % b

	fmt.Println("plus =", plus)
	fmt.Println("min =", min)
	fmt.Println("multiply =", multiply)
	fmt.Println("division =", division)
	fmt.Println("modulo =", modulo)

	// augmented assignment
	a += 10
	fmt.Println("augmented a plus", a)

	a -= 10
	fmt.Println("augmented a min", a)

	a *= 10
	fmt.Println("augmented a multiply", a)

	a /= 10
	fmt.Println("augmented a division", a)

	a %= 10
	fmt.Println("augmented a modulo", a)

	// unary operator
	var i = 1
	i++
	i++

	fmt.Println("i value =", i)

	i--
	i--

	fmt.Println("i value =", i)
}
