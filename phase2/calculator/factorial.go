package calculator

import (
	"fmt"
)

func factorial(num int) int {
	if num <= 1 {
		return 1
	}
	return num * factorial(num-1)
}
func factorialForFloat(floatNum float64) float64 {
	if floatNum <= 0 {
		return 1.0
	} else {
		return floatNum * factorialForFloat(floatNum-0.1)
	}
}
func Factorial() {
	var num int
	fmt.Println("enter number:")
	fmt.Scanln(&num)
	floatNum := float64(num)
	fmt.Println("Factorial of float number:",factorialForFloat(floatNum))
	fmt.Println("Factorial of integer number:",factorial(num))
}
