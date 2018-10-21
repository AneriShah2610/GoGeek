package calculator

import (
	"fmt"
)

func SelectOne() {
	var string1 string
	fmt.Printf("Select 1 from given list : \n 1. Calculator \n 2. Fibboseries \n 3. Factorial \n 4. AtmMachine \n 5. MergeSort \n 6. Cipher Techniques\n")
	fmt.Println("Select any one from above list:")
	fmt.Scanln(&string1)

	switch {
	case string1 == "Calculator" || string1 == "1":
		Calculate()
	case string1 == "Fiboseries" || string1 == "2":
		Fibo()
	case string1 == "Factorial" || string1 == "3":
		Factorial()
	case string1 == "AtmMachine" || string1 == "4":
		AtmMachine()
	case string1 == "MergeSort" || string1 == "5":
		MergeSort()
	case string1 == "Cipher" || string1 == "6":
		CipherTechnique()
	default:
		fmt.Println("Invalid delection")
	}
}
