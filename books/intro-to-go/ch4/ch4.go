package ch4

import "fmt"

func Run() {
	i := 1
	for i <= 5 {
		var joiner string
		if i % 2 == 0 {
			joiner = "not "
		} else {
			joiner = ""
		}

		var translation string
		switch i {
		case 1: translation = "One"
		case 2: translation = "Two"
		case 3: translation = "Three"
		default: translation = "Other"
		}

		fmt.Printf("%d is %sodd and translates to %s\n", i, joiner, translation)
		i += 1
	}
	fmt.Println("")
}