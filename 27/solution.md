# **Solution: Remove Element**

LeetCode problem #27 includes removing all occurrences of `val` in a given integer array. To solve this problem I must range over the array, skip the pass if the value matches `val`, and increment `k` if not. Not only do I have to increment but I have to also copy over each *good* value (in this case a *good* value is one that does not match `val`) to `nums[k]`.

Something that took a while for me to wrap my head around, and which required me reading the instructions twice to understand; the function returns `k`, the number of values that aren't `val`. LeetCode only checks the first `k` spots of nums and ignores the rest. So no need to remove the values matching `val`, they sit past the point of what is checked, so they either get ignored, or overwritten.

So instead of deleting, which was my first approuch, I instead had to copy each *good* value to the front, with `k` being both the index where the next good value goes, *and* the final `k` count which is returned. So; *write a value to nums[k], then k++*.

# **Solution Code:**
```go 
func removeElement(nums []int, val int) int {
	var k int
	for _, n := range nums {
		if n == val {
			continue
		}
		nums[k] = n
		k++
	}
	fmt.Printf("%d, nums = %v\n", k, nums)
	return k
}
```