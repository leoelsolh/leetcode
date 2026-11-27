package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}

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
