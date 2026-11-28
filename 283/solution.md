# **Solution: Move Zeroes**

*Move Zeroes* is a problem similar to *Remove Element* in the way you loop over the arrays numbers, shifting the unwanted further and further to the end. The only difference with the two, is that *Move Zeroes* solution includes setting the value to `0` on every `num` from `count` to `len(nums)`, aka, the unwanted numbers...

Writing the solution, I took heavy *inspiration* from my previous code on the remove element problem. To complete it all I needed was a simple three-part `for` loop. 

### **Solution Code:**
```go
package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	var count int
	for _, num := range nums {
		if num != 0 {
			nums[count] = num
			count++
		}
	}
	for i := count; i < len(nums); i++ {
		nums[i] = 0
	}
}
```
