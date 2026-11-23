# **Leetcode Problem 2124**

The problem given in this leetcode challange, is, given a string consisting of only the characters 'a' and 'b', return true if every 'a' appears before every 'b'. Otherwise return false. 

Now this was a fun one because I've never really worked with indexing strings like this, ranging over the string, checking length of index, etc. 

So, heres my code solution. 

```go
package main

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
```

My though process going through this:

The first thing I did was print the runecount of the string, which I quickly realised wouldnt do anything for me since I needed to focus on *order* rather than count. 

So, if I count a `b` in the string, then see an `a` at any index after that initial `b`, I need to return false. This means the solution would be to say *if c is 'b', check the next index, if the string on the next index is 'a' then return false. But if its 'b', then continue, and repeat until string ends*.

The problem holding me back was that the loop didnt care about the length of the string. Meaning the `if s[i+1]` would try to read a non-existant index after the last character. Casuing a panic. 

The solution was to check if we we're on the last itteration. If so, just return true. Becuase the `if c == 'b'` gaurd above ensures the last char can't be `a`, and since its the last index; the result is true. 