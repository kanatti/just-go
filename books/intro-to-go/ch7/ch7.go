package ch7

import "fmt"

func Run() {
	point1 := new(Point) // allocates memory and returns a ptr
	fmt.Printf("%v\n", point1)
	point2 := Point{x: 0, y: 0, label: "origin"}
	point3 := &Point{0, 1, "unit-y"}
	fmt.Printf("%v\n", point2)
	fmt.Printf("%v\n", point3)

	vector1 := makeVectorFromPoints(&point2, point3)
	fmt.Printf("%v\n", vector1)

	dog1 := &Dog{Animal{true}, "Brownie"}
	fmt.Printf("%v\n", dog1)
	dog1.introduce()
	dog1.breathe()

	printLabels(point3, vector1, dog1)
	printLabels(point3, vector1, dog1)
}

func makeVectorFromPoints(origin *Point, dest *Point) *Vector {
	return &Vector{ origin, dest, (*origin).label + " -> " + (*dest).label } // deferencing explicitly is not required though
}

type Point struct {
	x, y int32 // These fields are only visible within the package as not starting with capital
	label string
}

func (p *Point) getLabel() string {
	return p.label
}

func (p *Point) setLabel(label string) {
	p.label = label
}

type Vector struct {
	origin *Point
	dest *Point
	label string
}

func (v *Vector) getLabel() string {
	return v.label
}

func (v *Vector) setLabel(label string) {
	v.label = label
}

func (v *Vector) String() string {
	return fmt.Sprintf("{%v %v %s}", v.origin, (*v).dest, v.label) // v.origin works same as (*v).origin
}

// Embedded types for is-a relation instead of has-a relation
// Basically composition over inheritance

type Animal struct {
	isAlive bool
}

func (animal *Animal) breathe() {
	if animal.isAlive {
		fmt.Println("breathing ... ")
	} else {
		fmt.Println("Sorry, cant breathe ...")
	}
}

type Dog struct {
	Animal // Keeping fields without name allows to proxy method calls to owner
	name string
}


func (dog* Dog) introduce() {
	fmt.Println("Hi, this dog's name is", dog.name)
}

func (dog *Dog) getLabel() string {
	return dog.name
}

func (dog *Dog) setLabel(name string) {
	dog.name = name
}

type LabelledItem interface {
	getLabel() string
	setLabel(string)
}

func printLabels(items ...LabelledItem) {
	details := ""
	for _, item := range items {
		details = details + item.getLabel() + ","
		item.setLabel(item.getLabel() + "[OK]")
	}
	fmt.Println(details)
}
