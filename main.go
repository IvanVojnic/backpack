package main

import "fmt"

// MaxWeight define limit for a backpack
const MaxWeight = 9

func main() {
	cost := [3]int{4, 9, 16}
	weight := [3]int{2, 3, 5}
	dp := make([]int, MaxWeight+1)
	items := make([]int, MaxWeight+1)
	items[0] = -1
	for i := 0; i < len(dp); i++ {
		dp[i] = 0
	}
	for i := 1; i <= MaxWeight; i++ {
		for j := 0; j < len(weight); j++ {
			if weight[j] <= i {
				dp[i] = max(dp[i], dp[i-weight[j]]+cost[j])
				items[i] = j
			}
		}
	}

	for i := 0; i < len(dp); i++ {
		fmt.Printf("\nmax cost in any backpack - %v", dp[i])
	}
	fmt.Printf("\n\nSolution - %v", dp[MaxWeight])
	result := make([]int, 0)
	maxW := MaxWeight
	for i := items[MaxWeight]; i != -1; i = items[MaxWeight] {
		result = append(result, i)
		maxW -= weight[i]
	}
	for i := 0; i < len(result); i++ {
		fmt.Printf("\nitems in backpack - %v", result[i])
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
