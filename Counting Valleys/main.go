package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Counting Valleys")
	fmt.Println(CountingValleys(12, "DDUUDDUDUUUD"))
}

func CountingValleys(steps int32, path string) int32 {
	var l, c int32 = 0, 0
	for _, val := range strings.Split(path, "") {
		if val == "U" {
			if l < 0 && l+1 == 0 {
				c++
			}
			l++
		} else if val == "D" {
			l--
		}
	}
	return c
}
