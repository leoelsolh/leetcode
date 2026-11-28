# **Solution: Contains Duplicate**

With this problem, I need first create a new map where I can set a key for the integers, so that I can then loop through the provided slice, and check whether the value passed has already been seen before. So each pass checks that the `n` hasn't already been set to a key in the `seen` list.

Coming from python, my inital move was to write `if x in y` then quickly realised that, in this case I'll need to `make` a map like this; `map[int]bool` so I can perform a boolean check (comma ok idiom) with each pass from the `for` loop. 

After finishing my program I realised I has read the assignment wrong, and that the function was to return *true* if the slice *did* contain a duplicate. My initial solution has opposite returns... Easy fix though, and luckily no other bugs where uncovered from flipping the return statements. Next time read it thouroughly before submitting.

### **Solution Code**
```go 
func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)
	for _, n := range nums {
		_, ok := seen[n]
		if !ok {
			seen[n] = true
		} else {
			return true
		}
	}
	return false
}
```
### **Example Output:**
```
true 
false 
true
```
***