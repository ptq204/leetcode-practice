// Link: https://leetcode.com/problems/integer-break/

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}

func integerBreak(n int) int {
    if n == 1 || n == 2 {
        return 1
    }
    if n == 3 {
        return 2
    }
    dp := make([]int, n+1)
    dp[0] = 1
    dp[1] = 1
    for i:=0; i<=n; i++ {
        for j:=1; j<=i; j++ {
            dp[i] = max(dp[i], j*dp[i-j])
        } 
    }
    return dp[n]
}

// Time complexity: O(N^2)
// Space complexity: O(N)