package main

import "fmt"

func main() {
	adultAge := 21
	completionScoreMinimum := 75

	myScore := 80
	myAge := 24

	isScorePassed := myScore >= completionScoreMinimum
	isAgePassed := myAge >= adultAge
	isChild := !isAgePassed

	isAllPassed := isAgePassed && isScorePassed
	isPartiallyPassed := isAgePassed || isScorePassed

	fmt.Println("Is adult", isAgePassed)
	fmt.Println("Is child", isChild)
	fmt.Println("Is score passed", isScorePassed)
	fmt.Println("Is all passed", isAllPassed)
	fmt.Println("Is partially passed", isPartiallyPassed)
}
