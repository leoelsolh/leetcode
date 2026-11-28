# **Solution: Squares of a Sorted Array**

Sqares of a sorted array. To solve this problem I first need to create a *sorted* slice to which I can assign the squares to. To find the squares, all I'll need to really do is `num * num`, and to trim it down ill perform the multiplication inside the append signature.

Quick and easy.

### **Solution Code:**
```go
func sortedSquares(nums []int) []int {
	srted := make([]int, 0)
	for _, num := range nums {
		srted = append(srted, num*num)
	}
	slices.Sort(srted)
	return srted
}
```