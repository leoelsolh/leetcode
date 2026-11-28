# **Solution: Valid Anagram**

Problem 242 involves determining whether a given string is a *valid anagram*, and anagram being, "a word or phrase formed by rearranging the letters of a different word or phrase using all the original letters exactly once. [Wikipedia | Anagram](https://en.wikipedia.org/wiki/Anagram)"

So, to write a program which determines if something is a valid anagram, the program will have to do the following. 

First, just to get it out the way; make sure both `s` and `t` are the same length. Secondly, create a map to count the letters. Third it will have to loop over each character in `s` and add `+1` to that letter's count, then do the same with `t`; loop ever every character, but this time *subtract*. And finnaly, if any count is *not* zero, it means it's not an anagram.

Something that tripped me up; the rune value. Looping over a string gives you a `rune`, not a string character. This did not occur to me at the beginning, meaning my map was initially `map[string]int` when it needed to be `map[rune]int` to properly assign the character to its count. I also initially reached for the comma ok idiom inside the loop at first, but it isn't needed here since reading a missing key from a map returns the type's *zero value*, so `count[c]++` works even the first time a letter is seen.

### **Solution Code:**
```go 
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

```