package main

import "fmt"

func main() {
	fmt.Println(runningSum([]int{1, 2, 3, 4}))     // [1,3,6,10]
	fmt.Println(runningSum([]int{1, 1, 1, 1, 1}))  // [1,2,3,4,5]
	fmt.Println(runningSum([]int{3, 1, 2, 10, 1})) // [3,4,6,16,17]
}

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
