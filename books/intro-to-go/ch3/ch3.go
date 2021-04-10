package ch3

import "fmt"

const pi float64 = 3.14

func Run() {
	var x string = "Hello"
	y := "World"
	fmt.Println(x, y);

	var (
		start = "Value of pi is"
		end = "!!!"
	)

	fmt.Println(start, pi, end)
}