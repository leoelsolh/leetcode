package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome(" "))

	// Output:
	// : true
	// : false
	// : true
}

func isPalindrome(s string) bool {
	re, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		fmt.Println("Error compiling regex:", err)
	}
	res := re.ReplaceAllString(s, "")
	lres := strings.ToLower(res)

	l := 0
	r := len(res) - 1
	for l < r {
		if lres[l] == lres[r] {
			l++
			r--
		} else {
			return false
		}
	}
	return true
}
