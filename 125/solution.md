# **Solution: Valid Palindrome**

So breaking it down, we first need to do some work on the provided string since we cant have a proper palindrome with commas etc. I'll use `regexp` and `strings` to do the work on the string before the palindrome check happens. 

Then, to start on the palindrome check. What I need to do is similar to the *Reverse String* problem in the way that we have to loop over two positions. 
But in this case, instead of reversing the characters positions, we ensure they are both identical before moving inwards a step. 

If both the left-most character and the right-most character are equal, we move inwards a step and perform the same check until, either `l` and `r` meet in the middle, or until one of them dont match.

The `for l < r { }` will loop until `l` is larger than `r`, which is when they meet in the middle. And the `else` branch will return; breaking out of the loop if the characters dont match.


### **Solution Code**

```go
package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) // true
	fmt.Println(isPalindrome("race a car"))                     // false
	fmt.Println(isPalindrome(" "))                              // true
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
```
### **Example output:**
```
$ go run .
true
false
true
```
***