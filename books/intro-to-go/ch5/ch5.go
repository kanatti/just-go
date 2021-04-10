package ch5

import "fmt"

func Run() {
	fmt.Print("\n- Arrays - \n")
	var x [5]int
	x[4] = 100
	fmt.Println("x is a fixed length array with value", x)

	grades := [5]float64{56, 75, 80, 99, 76}
	var total float64 = 0

	for _, grade := range grades {
		total += grade
	}

	fmt.Printf("Avg grade for %v is %.2f\n", grades, total/float64(len(grades)))

	// Slices can change size. Has capacity and Length
	fmt.Print("\n- Slices - \n")
	var x1 []int32
	x2 := make([]int32, 5) // length 5
	x3 := make([]int32, 5, 10) // length 5 and capacity 10

	printSliceDetailsInt("x1", x1)
	printSliceDetailsInt("x2", x2)
	printSliceDetailsInt("x3", x3)
	printSliceDetailsFloat("grades", grades) // Either pass as slice or funtion needs to know size

	// Types of arr and slice
	fmt.Printf("%T\n",grades)
	fmt.Printf("%T\n",x2)

	// Maps
	fmt.Print("\n- Maps - \n")
	ages := make(map[string]int)
	ages["SpiderMan"] = 26
	ages["Batman"] = 32
	fmt.Printf("%v\n", ages)

	delete(ages, "Batman")

	if age, ok := ages["SpiderMan"]; ok {
		fmt.Println("Got SpiderMans age", age)
	}

	if age, ok := ages["Batman"]; ok {
		fmt.Println("Got Batmans age", age)
	}

	// shorter way to create maps

	numbers := map[int]string {
		1: "One",
		2: "Two",
	}

	fmt.Printf("Numbers mapping %v\n", numbers);
}

func printSliceDetailsInt(prefix string, arr []int32) {
	fmt.Printf("%s is %v, len - %d, cap - %d\n", prefix, arr, len(arr), cap(arr))
}

func printSliceDetailsFloat(prefix string, arr [5]float64) {
	fmt.Printf("%s is %v, len - %d, cap - %d\n", prefix, arr, len(arr), cap(arr))
}