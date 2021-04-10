package ch6

import "fmt"

const PREFIX string = "ch6: "

func Run() {
	grades := []float64 { 60, 70, 82, 95 }
	avg, sum, size := averageSum(grades)

	fmt.Printf("Grades: %v, avg: %.2f, sum: %.0f, size: %d\n", grades, avg, sum, size)

	printWithPrefix("Sum of 1,2,3 is", add(1, 2, 3))

	getSquares := makeSquaresGenerator()
	fmt.Println(getSquares())
	fmt.Println(getSquares())
	fmt.Println(getSquares())

	// defer, panic, recover
	deferAndRecover()

	// Pointers
	x := int(0)
	fmt.Println("x before increment", x)
	increment(&x)
	fmt.Println("x after increment", x)
}

func averageSum(xs []float64) (float64, float64, int) {
	size := len(xs)
	total := add(xs...)
	return total/float64(size), total, size
}

// sum with variadic parameter
func add(xs ...float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total
}

// Go doesnt support mix of regular and variadic. Need to merge and create a new variadic and pass it
func printWithPrefix(a ...interface{}) {
	fmt.Println(append([]interface{}{PREFIX}, a...)...)
}

// Returns func. i is closure
func makeSquaresGenerator() func() uint32 {
	i := uint32(0)
	return func() (ret uint32) { // Named return
		ret = i
		i += 2
		return
	}
}

func deferAndRecover() {
	defer func() { // anonymous function
		str := recover()
		fmt.Println("recovered from panic", str)
	}()
	panic("PANIC!!")
}

// * is used both as ptr type and deference operation
func increment(xPtr *int) {
	*xPtr = *xPtr + 1
}