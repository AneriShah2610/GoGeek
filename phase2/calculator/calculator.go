package calculator

import (
	"fmt"
	"math"
)

// interface that contains method
type calculatorResult interface {
	addition() float64
	subtraction() float64
	multiplication() float64
	diviison() float64
	sin() float64
	cos() float64
	power() float64
}

// struct
type calculatorStruct struct {
	a, b float64
}

// addition mehod
func (c *calculatorStruct) addition() float64 {
	return c.a + c.b
}

// subtraction
func (c *calculatorStruct) subtraction() float64 {
	return c.a - c.b
}

// multipliaction
func (c *calculatorStruct) multiplication() float64 {
	return c.a * c.b
}

// division
func (c *calculatorStruct) diviison() float64 {
	return c.a / c.b
}

// power
// power method1
/*func power(a, b float64) float64 {
	if b == 0 {
		return 1
	} else if b > 0 {
		return a * power(a, b-1)
	} else {
		return 1 / power(a, -b)
	}
}
*/
// power method2
/*func power(a, b float64) float64 {
	return math.Pow(a, b)
}*/
// power method3
func power(a, b float64) float64 {
	if b == 0 {
		return 1
	}
	if int(b)%2 == 0 {
		i := power(a, b/2)
		return i * i
	} else {
		return a * power(a, b-1)
	}
}

// sin
func (c *calculatorStruct) sin(a float64) float64 {
	return math.Sin(a)
}

// cosine
func (c *calculatorStruct) cos(a float64) float64 {
	return math.Cos(a)
}

// tan
func (c *calculatorStruct) tan(a float64) float64 {
	return math.Tan(a)
}

func Calculate() {
	var a, b float64
	fmt.Println("Enter num1:")
	fmt.Scanln(&a)
	fmt.Println("Enter num2:")
	fmt.Scanln(&b)
	nums := &calculatorStruct{a, b}
	fmt.Println(nums)
	fmt.Println("please select functions:")
	var functions string
	fmt.Println("Enter functions")
	fmt.Scanln(&functions)
	switch {
	case functions == "add" || functions == "ADD" || functions == "+":
		fmt.Println("Addition:", nums.addition())
	case functions == "sub" || functions == "SUB" || functions == "-":
		fmt.Println("Subtraction", nums.subtraction())
	case functions == "mul" || functions == "MUL" || functions == "*":
		fmt.Println("Multiplication:", nums.multiplication())
	case functions == "div" || functions == "DIV" || functions == "/":
		fmt.Println("Division:", nums.diviison())
	case functions == "pow" || functions == "POW" || functions == "^":
		fmt.Println("Power:", power(a, b))
	case functions == "sin" || functions == "SIN":
		fmt.Println("Sin of a:", nums.sin(a))
	case functions == "cos" || functions == "COS":
		fmt.Println("cos of a:", nums.cos(a))
	case functions == "tan" || functions == "TAN":
		fmt.Println("tan of a:", nums.tan(a))
	default:
		fmt.Println("Invalid selection")
	}
}
