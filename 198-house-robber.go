// Link: https://leetcode.com/problems/house-robber/

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}

func rob(nums []int) int {
    n := len(nums)
    if n == 1 {
        return nums[0]
    }
    a := nums[n-1]
    b := max(nums[n-2], nums[n-1])
    c := math.MinInt64
    for i:=n-3; i>=0; i-- {
        c = math.MinInt64
        c = max(c, nums[i]+a)
        c = max(c, b)
        a = b
        b = c
    }
    return max(a, b)
}

// Time complexity: O(N)
// Space complexity: O(1)