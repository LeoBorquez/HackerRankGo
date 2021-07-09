package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Repeated String")
	//fmt.Println(repeatedStrings("cfimaakj", 554045874191))
	fmt.Println(repeatedStrings("aba", 10))
}

func repeatedStrings(s string, n int64) int64 {
	var r, t, ir int64

	if len(s) <= 1 && s == "a"{
		t = n
	} else {
		for _, v := range strings.Split(s, "") {
			if v == "a" {
				r++
			}
		}

		full := int(n) / len(s)
		rest := int(n) % len(s)

		sub := s[0:int(rest)]

		for _, v := range strings.Split(sub, "") {
			if v == "a" {
				ir++
			}
		}

		t = (r * int64(full)) + ir
	}

	return t
}
