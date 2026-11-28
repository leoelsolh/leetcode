# **Solution: Max Consecutive Ones**

For this problem I need to loop through the slice, incrementing to a count for every pass where `n` is `1`. One count will be for the length of the current run, and the other will be for the final max.  

At first my approuch was to just increment to a single count variable, and reset the count every time I hit a zero. But I quickly realised that would not work since it meant the count would forget whatever it had already passed meaning the max count of 1's could be lost in the middle. 

The solution was to have two seperate variables as mentioned above. One for the so-far-best count, and one for the current count. If the current count is larder than the running maximum, then the current count *becomes* the running maximum. 

### **Solution Code:**
```go
func findMaxConsecutiveOnes(nums []int) int {
	var best int
	var curr int
	for _, n := range nums {
		if n == 1 {
			curr++
			if curr > best {
				best = curr
			}
		} else {
			curr = 0
		}
	}
	return best
}
```