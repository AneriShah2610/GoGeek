package calculator

import (
	"fmt"
)

type atmmacine interface {
	FoundNotes()
}
type atmmachinestruct struct {
	currencyValue int
}

func findNotes(currencyValue int) map[int]int {
	notes := []int{2000, 1000, 500, 200, 100, 50, 20, 10, 5, 2, 1}
	note := make(map[int]int)
	if currencyValue <= 0 {
		return map[int]int{}
	}
	for i := 0; i < len(notes); i++ {
		denomination := 0
		if currencyValue >= notes[i] {
			denomination = currencyValue / notes[i]
			currencyValue -= denomination * notes[i]
		}
		if denomination != 0 {
			note[notes[i]] = denomination

		}
		/*if denomination != 0 {
			fmt.Println(notes[i], " : ", denomination)
		}*/
		//note1 := reflect.ValueOf(note).MapKeys()
	}
	return note
}
func AtmMachine() {
	var currencyValue int
	fmt.Println("Enter amount:")
	fmt.Scanln(&currencyValue)
	fmt.Println(findNotes(currencyValue))
}
