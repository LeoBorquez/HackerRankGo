package main

import (
	"fmt"
	"math"
)

func main() {
	// var declaration
	var msg string
	msg = "Hello" // msg:= "Hello" shortcut

	fmt.Println(msg) // print message

	const pi = 1.618 // constant declaration

	// var numbers [5]int
	numbers := [...]int{0, 0, 0, 0, 0} // arrays have a fixed size

	slice := []int{2, 3, 4} // slices have a DYNAMIC SIZE, unlike arrays
	sliceString := []string{"Hello"}

	fmt.Println(pi, numbers, slice, sliceString)

	b := *getPointer()
	fmt.Println("Value is", b)

	// maps declaration
	var names = make(map[int]string)
	names[0] = "John"
	names[1] = "Jane"
	fmt.Println(names)
}

// Pointers point to a memory location of a variable.
func getPointer() (myPointer *int) { // return/read a pointer *int
	a := 234
	return &a // & point to
}

func flowControl() float64 {
	lim := 10.5
	if v := math.Pow(3, 2); v < lim {
		return v
	} else {
		fmt.Println("return 10")
	}
	return lim
}
