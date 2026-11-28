package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
	fmt.Println(sortedSquares([]int{-7, -3, 2, 3, 11}))
}

func sortedSquares(nums []int) []int {
	srted := make([]int, 0)
	for _, num := range nums {
		srted = append(srted, num*num)
	}
	slices.Sort(srted)
	return srted
}
