# **Solution: Length of Last Word**

Starting it off, I'll use the `strings` package to slice `s` into substrings using the `strings.Split()` function. This works when we take the first example; `"Hello World"` because it will split at the whitespace. 

But when it comes to example 2; `"   fly me   to   the moon  "` because `Split` cuts at every single space, so consecutive spaces produce empty strings ("") in the slice. With trailing spaces, the last element is "", which has length 0. So in this case we would get an answer like: `length of last string = 0` because its counting the empty space to be the last word.

I took this flawed approuch at first, before testing the edge cases and realising the `Split()` function would not work in this case. We need `strings.Fields()`. 

The `strings.Fields()` function splits the string s around each instance of one or more consecutive white space characters, meaning we dont have to worry about double spaces. Alteretively, I could have stripped away extra spaces before splitting the string, but `Fields` is more reliable in my opinion.

### **Solution Code:**
```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))                 // 5
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  ")) // 4
	fmt.Println(lengthOfLastWord("luffy is still joyboy"))       // 6
}

func lengthOfLastWord(s string) int {
	slce := strings.Fields(s)
	lwrd := slce[len(slce)-1]
	return len(lwrd)
}
```