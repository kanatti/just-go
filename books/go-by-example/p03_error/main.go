package p03_error

import (
	"errors"
	"fmt"
)

func Run() {
	fmt.Println(" ---- Errors ----")

	fmt.Println(" -- Simple Error --")
	printDivisionResults(5, 3)
	printDivisionResults(10, 0)

	fmt.Println(" -- Custom Error --")
	_, e := process(42)
	if argErr, ok := e.(*argError); ok {
		fmt.Println(argErr.arg)
		fmt.Println(argErr.prob)
	}
}

func divide(n int, d int) (float64, error) {
	fmt.Printf("Dividing %d by %d\n", n, d)
	if d == 0 {
		return 0, errors.New("Cant divide by zero")
	}
	return float64(n)/float64(d), nil
}

func printDivisionResults(n int, d int) {
	if result, error := divide(n, d); error != nil {
		fmt.Println("Got error", error)
	} else {
		fmt.Printf("Got result %0.2f \n", result)
	}
}

type argError struct {
	arg int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func process(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't process 42"}
	}
	return arg + 1, nil
}
