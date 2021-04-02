// Link: 

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func solve(idx, n int, dp *[]int, nums []int) int {
    if idx >= n-1 {
        return 0
    }
    if (*dp)[idx] != -1 {
        return (*dp)[idx]
    }
    steps := nums[idx]
    res := n+1
    for i:=1; i<=steps; i++ {
        res = min(res, 1 + solve(idx+i, n, dp, nums))
    }
    (*dp)[idx] = res
    return res
}

func jump(nums []int) int {
    n := len(nums)
    dp := make([]int, n)
    for i:=0; i<n; i++ {
        dp[i] = -1
    }
    return solve(0, n, &dp, nums)
}

// Time complexity: O(N^2)
// Space complexity: O(N)
