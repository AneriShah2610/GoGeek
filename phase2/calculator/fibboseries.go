package calculator

import (
	"fmt"
)

// use this method1 when we use for loop
/*func fibonacci(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonacci(i-1) + fibonacci(i-2)
}*/

// method 2
func fibonacci(num1, num2, num int) {

	if num >= 0 {
		fmt.Println(num1)
		/*a += b
		b = a - b*/
		sum := num1 + num2
		num1 = num2
		num2 = sum
		fibonacci(num1, num2, num-1)
	}

}

func Fibo() {
	var num int
	fmt.Print("Enter number for which u want to find fibonacci series:")
	fmt.Scanln(&num)
	// For method1
	/*for i := 0; i < num; i++ {
		fmt.Println(fibonacci(i))
	}*/
	// For method2
	num1, num2 := 0, 1
	fmt.Println("Fibonacci series:")
	fibonacci(num1, num2, num-1)
}
