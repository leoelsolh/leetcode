# **Solution: Running Sum of 1d Array**

[A running sum](https://en.wikipedia.org/wiki/Running_total), also known as a rolling sum, is the summation of a sequence of numbers which gets updated after each new number is added. This is done by adding the value of the new number to the previouse running total. That means to follow the rules of a running sum, ill need to consider skip the addition when passing the first number. Since its running sum is just itself, all we have to do is append it as-is

Now that that's out of the way, lets get into the solution. 

There are two paths to go here; either I can build a new slice and return that, *or* I can overwrite the input slice in place, like I did in the *Reverse String* problem. Both have their downsides, but I think I'll go with building a new slice since memory usage wont be a worry, and Its a good habit to not mutate provided input by default. And since I'm creating a new slice, I'll need to remember to append the previouse sum from the new slice instead of the provided slice (see `res = append(res, res[i-1]+nums[i])`). 

When declaring the new slice using `make` I decided to initialize it with a length of `0` and then append to it.

When going through this problem, for a long time I held a flawed approuch of trying to add the current pass number, to the number infront of it. This caused a few panic's since I was indexing out of range on the final pass. The solution was to add the current itteration to the numer *before* it. This took a while to land for me but I got there in the end...

I think the lesson that really stuck for me with this problem, was the two scenarios I'd previously never considered:

- Looking forward (`s[i+1]`) breaks on the last index, so I gaurd the end.
- Looking backwards (`s[i-1]`) breaks at the *first* index. So I gaurd the start.

Heres my final solution to this LeetCode problem:

### **Solution Code:**
```go
func runningSum(nums []int) []int {
	res := make([]int, 0)
	for i := range nums {
		if i == 0 {
			res = append(res, nums[i])
			continue
		}
		res = append(res, res[i-1]+nums[i])
	}
	return res
}
```
### **Output:**
```
$ go run .
[1 3 6 10]
[1 2 3 4 5]
[3 4 6 16 17]
```
