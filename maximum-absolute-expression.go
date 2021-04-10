package main

import (
	"fmt"
	"math"
)
// Find maximum |A[i] - A[j]| + |i - j|

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findMaxAbsolute(nums []int) int {
	max1 := math.MinInt64
	max2 := math.MinInt64
	min1 := math.MaxInt64
	min2 := math.MaxInt64
	ans := math.MinInt64
	for i, _ := range nums {
		max1 = max(max1, nums[i]+i)
		min1 = min(min1, nums[i]+i)
		max2 = max(max2, nums[i]-i)
		min2 = min(min2, nums[i]-i)
		ans = max(ans, max1-min1)
		ans = max(ans, max2-min2)
	}
	return ans
}

func main() {
	nums := []int{1, 10, 3, 5, 4, 3, 9}
	fmt.Println(findMaxAbsolute(nums))
}