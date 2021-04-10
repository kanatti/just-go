package main

import (
	"fmt"
	"github.com/kanatti/just-go/books/intro-to-go/ch2"
	"github.com/kanatti/just-go/books/intro-to-go/ch3"
	"github.com/kanatti/just-go/books/intro-to-go/ch4"
	"github.com/kanatti/just-go/books/intro-to-go/ch5"
	"github.com/kanatti/just-go/books/intro-to-go/ch6"
	"github.com/kanatti/just-go/books/intro-to-go/ch7"
)

func main() {
	fmt.Println(" --- Chapter 2 : Types --- ")
	ch2.Run()
	fmt.Println("\n --- Chapter 3 : Variables --- ")
	ch3.Run()
	fmt.Println("\n --- Chapter 4 : Control Structures --- ")
	ch4.Run()
	fmt.Println("\n --- Chapter 5 : Collections --- ")
	ch5.Run()
	fmt.Println("\n --- Chapter 6 : Functions --- ")
	ch6.Run()
	fmt.Println("\n --- Chapter 6 : Structs and Interfaces --- ")
	ch7.Run()
}
