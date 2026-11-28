# **Solution: Two Sum**

Two Sum. Known by all. Except me, until now... Two Sum is a problem that includes algortihmic challanges, where you must identify elements in an array whose sum equals a target number. 

To solve this leetcode problem, ill first need to create a map to which I can assign the numbers already seen, along with their indexies. After declaring the map with `make` I need to the loop through the provided array and set a complement variable. 

The complement variable will be checked as a key in the map, and if it is present in the map, we return the array index `i` taken from the `for` loop, along with the value from the comma-ok idiom check. If *not* present, we simply add a key with the value of `s[n] = i` 

In this case my final code does not produce output identical to LeetCode's example simply because in Go, thats how slices are printed. So intead of `[2,1]` I get `[2 1]` which is fine in my opinion. If I wanted to get the output identical I could use `strings.ReplaceAll()`. 

### **Solution Code:**
```go
func twoSum(nums []int, target int) []int {
	s := make(map[int]int)
	for i, n := range nums {
		comp := target - n
		val, ok := s[comp]
		if ok {
			return []int{i, val}
		} 
		s[n] = i
	}
	return nil
}
```