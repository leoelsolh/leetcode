package main

import "fmt"

func main() {
	fmt.Println(checkString("aaabbb"))
}

func checkString(s string) bool {
	for i, c := range s {
		if c == 'b' {
			if i == len(s)-1 {
				return true
			}
			if s[i+1] == 'a' {
				return false
			}
		}
	}
	return true
}
