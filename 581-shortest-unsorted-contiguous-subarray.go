// Link: https://leetcode.com/problems/shortest-unsorted-continuous-subarray/

func findUnsortedSubarray(nums []int) int {
    n := len(nums)
    max := math.MinInt64
    min := math.MaxInt64
    l, r := n-1, 0
    for i:=0; i<n; i++ {
        if nums[i] < max {
            r = i
        } else {
            max = nums[i]
        }
        if nums[n-1-i] > min {
            l = n-1-i
        } else {
            min = nums[n-1-i]
        }
    }
    if r == 0 || r <= l {
        return 0
    }
    return r-l+1
}

// Time complexity: O(N)
// Space complexity: O(1)