package main

import (
	"fmt"
	"math"
	"runtime"
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
	// statement if
	lim := 10.5
	if v := math.Pow(3, 2); v < lim {
		return v
	} else {
		fmt.Println("return 10")
	}

	// for loop
	for count := 0; count <= 10; count++ {
		fmt.Println("My counter is at", count)
	}

	// for range loop
	entry := []string{"Jack", "John", "Jones"}
	for i, val := range entry {
		fmt.Printf("At position %d, the character %s is present\n", i, val)
	}

	// while loop
	n := 0
	x := 42
	for n != x {
		n++
	}

	// switch
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	return lim
}
