package main

import (
	"fmt"
)

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Printf(" ⇄ %v\n", string(s))
	reverseString(s)
	fmt.Printf(" ⇄ %v\n", string(s))
}

func reverseString(s []byte) {
	l := len(s) - len(s)
	r := len(s) - 1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}
