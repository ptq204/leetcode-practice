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
    prefixSum := make([]int, 0)
    sum := 0
    ans := math.MinInt64
    for i:=0; i<n; i++ {
        if ans < nums[i] {
            ans = nums[i]
        }
        sum += nums[i]
        prefixSum = append(prefixSum, sum)
    }
    i := -1
    j := 0
    for j < n && i < j {
        if i == -1 {
            sum = prefixSum[j]
        }  else {
            sum = prefixSum[j] - prefixSum[i]
        }
        if sum < nums[j] {
            for sum < nums[j] && i < j {
                i++
                sum = prefixSum[j] - prefixSum[i]
                ans = max(ans, sum)
            }
        } else {
            ans = max(ans, sum)
        }
        j++
    }
    return ans
}

// Time complexity: O(N)
// Space complexity: O(N)