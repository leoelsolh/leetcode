# **Solution: Richest Customer Wealth**

LeetCode problem 1672 is about finding the wealthiest customer. Given an integer grid, I need to exract the wealth of the customer with the highest total. With customers having multiple accounts each, the solution includes nested for loops, `current` and `best` variables, greater-than checks, etc.

So, to get started, Ill need to declare a `best` variable of type `int` which I can assign the greatest sum to. It'll be placed outside the loop so as not to be reset each pass. Then I loop through the accounts and declare a new `curr` variable, also of type `int` to sum the current passing account wealth in the following loop. 

Then comes the nested loop I mentioned above, which is now ranging through one customer, and all their accounts, writing the sum to the `curr` variable. 

If the current account total it greater-than the `best` value, the `best` *becomes* the current sum. Then the function returns that `best` value.

### **Solution Code:**
```go 
func maximumWealth(accounts [][]int) int {
	var best int
	for _, acc := range accounts {
		var curr int
		for _, a := range acc {
			curr = curr + a
		}
		if curr > best {
			best = curr
		}
	}
	return best
}
```
### **Example Output:**
```
6
10
17
```
***