package main

import "fmt"

func main() {
	fmt.Println(maximumWealth([][]int{{1, 2, 3}, {3, 2, 1}}))
	fmt.Println(maximumWealth([][]int{{1, 5}, {7, 3}, {3, 5}}))
	fmt.Println(maximumWealth([][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}))
}

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
