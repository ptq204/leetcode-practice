// Link: https://leetcode.com/problems/maximum-subarray/
// Given an integer array nums, find the contiguous subarray (containing at least one number)
// which has the largest sum and return its sum.

func max(i,j int) int {
    if i < j {
        return j
    }
    return i
}

func maxSubArray(nums []int) int {
    n := len(nums)
    i := 0
    j := 0
    ans := math.MinInt64
    sum := 0
    for j < n {
        sum += nums[j]
        if sum < nums[j] {
            for i < j && sum < nums[j] {
                sum -= nums[i]
                ans = max(ans, sum)
                i++
            }
        } else {
            ans = max(ans, sum)
        }
        j++
    }
    return ans
}

// Time complexity: O(N)
// Space complexity: O(1)