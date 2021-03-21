// Link: https://leetcode.com/problems/increasing-triplet-subsequence/
// Given an integer array nums, return true if there exists a triple of indices (i, j, k) such that i < j < k 
// and nums[i] < nums[j] < nums[k]. If no such indices exists, return false.

func increasingTriplet(nums []int) bool {
    minF := math.MaxInt64
    minS := math.MaxInt64
    for _, num := range nums {
        if num < minF {
            minF = num
        }
        if num > minF && num < minS {
            minS = num
        }
        if num > minS {
            return true
        }
    }
    return false
}

// Time complexity: O(N)
// Space complexity: O(1)