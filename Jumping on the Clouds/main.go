package main

import "fmt"

func main() {
	c := []int32{0, 0, 0, 0, 1, 0}
	jumpingOnClouds(c)
}

func jumpingOnClouds(c []int32) int32 {
	fmt.Println(c)
	var counter int
	for i, val := range c {
		if i != 0 && val == 0{
			counter++
		}
	}
	fmt.Println(counter)
	return 3
}
