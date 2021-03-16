// Link: https://leetcode.com/problems/two-sum/
// Given an array of integers nums and an integer target,
// return indices of the two numbers such that they add up to target.

func twoSum(nums []int, target int) []int {
    hmap := make(map[int]int)
    for i, k := range nums {
        if _, ok := hmap[target-k]; ok {
            return []int{i, hmap[target-k]}
        }
        hmap[k] = i
    }
    return []int{}
}

// Time complexity: O(N)
// Space complexity: O(N)