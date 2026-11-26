package main

import "fmt"

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
}

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := make(map[rune]int)
	for _, c := range s {
		count[c]++
	}
	for _, c := range t {
		count[c]--
	}
	for _, n := range count {
		if n != 0 {
			return false
		}
	}
	return true
}
