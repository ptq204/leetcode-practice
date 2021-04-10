// Link: https://leetcode.com/problems/climbing-stairs/

func climbStairs(n int) int {
    dp := make([]int, n+2)
    dp[1] = 1
    dp[2] = 2
    if n <= 2 {
        return dp[n]
    }
    for i:=3; i<=n; i++ {
        dp[i] = dp[i-1] + dp[i-2]
    }
    return dp[n]
}

// Time complexity: O(N)
// Space complexity: O(N)