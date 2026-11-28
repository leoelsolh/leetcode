<div align="center">

# **Golang Leetcode Solutions**
[![LeetCode](https://img.shields.io/badge/LeetCode-elsolh-FFA116?style=for-the-badge&logo=leetcode&logoColor=white)](https://leetcode.com/u/elsolh/)
![Go version](https://img.shields.io/github/go-mod/go-version/leoelsolh/leetcode?style=for-the-badge&logo=go&logoColor=white)
![Solved](https://img.shields.io/badge/solved-8-2ea44f?style=for-the-badge)

</div>

Working through LeetCode in Go, focusing on fundamentals and common patterns. Each solution has notes on thought process, some include what didn't work. 
But most of the notes are just personal thoughts, intuition, and things too look back on. 

And, Check out my [leetcode profile](https://leetcode.com/u/elsolh) if you get the chance.

***

### **Problem Index**
| #    | Problem                                    | Notes                       |
|------|--------------------------------------------|-----------------------------|
| 1    | Two Sum                                    | [notes](./1/solution.md)    |
| 27   | Remove Element                             | [notes](./27/solution.md)   |
| 58   | Length of Last Word                        | [notes](./58/solution.md)   |
| 125  | Valid Palindrome                           | [notes](./125/solution.md)  |
| 217  | Contains Duplicate                         | [notes](./217/solution.md)  |
| 242  | Valid Anagram                              | [notes](./242/solution.md)  |
| 344  | Reverse String                             | [notes](./344/solution.md)  |
| 485  | Maximum Consecutive Ones                   | [notes](./485/solution.md)  |
| 977  | Squares of a Sorted Array                  | [notes](./977/solution.md)  |
| 1480 | Running Sum of 1d Array                    | [notes](./1480/solution.md) |
| 1672 | Richest Customer Wealth                    | [notes](./1672/solution.md) | 
| 2124 | Check if All A's Appears Before All B's    | [notes](./2124/solution.md) |

***

### **Revisit:**
| #    | Reason                                          |
|------|-------------------------------------------------|
| 2124 | Redundant check and a last-index special case   |
| 125  | Should use `regexp.MustCompile`, keep `res`/`lres` consistent |
| 58   | Solve without `strings.Fields`                  |
| 242  | `[26]int` instead of a map                      |
| 283  | Single pass with swaps                          |
| 977  | O(n) follow-up with two pointers                |