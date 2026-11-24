# **Solution: Reverse String**

Since Go has different arrays when compares to python or JSON, we will not be using `["h","e","l","l","o"]` as an input, instead we will use `[]byte{'h', 'e', 'l', 'l', 'o'}`. So setting this solution code up, that was the first thing I ran into that made me pause.

So approuching this problem, we need to think of a two-pointer solution. This is where i take two points, one at the first char and one at the last, then you swap them, like this:

```
h e l l o
^       ^
```

Then move inward one step and do the same, it will look something like this:

```
o e l l h
  ^   ^
```

This is done by simply switching them out; `a, b = b, a`. But intead of `a, b` we'll have positions. 

So, for the solution, first we declare a left-most position and a right-most position, these are the starting point and we will move inward from here. The left position is set to `0` while the right position is set to the length of the slice, minus one to take into account the index number starting from `0`.

After that, we use a `for` to loop until the condition is false (l < r), swapping the letters every itteration, then jumping inwards a step, until the left position meets the right.

### **Solution Code**
```go
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
	l := 0
	r := len(s) - 1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}
```

Example output:

```
$ go run . 
 ⇄ hello
 ⇄ olleh
```

***

(for the '⇄' symbols: [EmojiDB.org](https://emojidb.org/reverse-symbol-emojis))