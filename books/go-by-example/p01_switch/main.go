package p01_switch

import "fmt"

func Run() {
	fmt.Println(" ---- SWITCH ----")
	fmt.Println(" -- normalSwitch -- ")
	helloSuperheros("peter parker")
	helloSuperheros("tim")

	fmt.Println("\n -- switchWithMultipleMatch -- ")
	isWeekend("saturday")
	isWeekend("tuesday")

	fmt.Println("\n -- switchWithConditions -- ")
	passOrFail(80, 20)
	passOrFail(72, 65)

	fmt.Println("\n -- normalSwitch -- ")
	switchOnTypes(int(10))
	switchOnTypes(true)
	switchOnTypes(1.2)
}

func helloSuperheros(name string) {
	fmt.Printf("%s - ", name)
	switch name {
	case "bruce wayne":
		fmt.Println("Hello batman")
	case "peter parker":
		fmt.Println("Hello spiderman")
	case "tony stark":
		fmt.Println("Hello ironman")
	default:
		fmt.Println("Hello", name)
	}
}

func isWeekend(day string) {
	fmt.Printf("day is %s ", day)
	switch day {
	case "saturday", "sunday":
		fmt.Println("Its weekend!")
	default:
		fmt.Println("Its a working day")
	}
}

func passOrFail(grade int8, attendancePercent int8) {
	fmt.Printf("Grade - %d, attendance - %d%% ", grade, attendancePercent)
	switch {
	case grade < 40 || attendancePercent < 60:
		fmt.Println("Sorry you failed")
	case grade > 90:
		fmt.Println("You got A grade")
	case grade > 80:
		fmt.Println("You got B grade")
	case grade > 70:
		fmt.Println("You got C grade")
	default:
		fmt.Println("You passed")
	}
}

func switchOnTypes(i interface{}) {
	fmt.Printf("Got %v ", i)
	switch v := i.(type) { // v is not type, v is value with type casted
	case bool:
		fmt.Println("Its a boolean")
	case int:
		fmt.Println("Its an int")
	default:
		fmt.Printf("Unknown type %T\n", v)
	}
}