// Link: https://leetcode.com/problems/longest-palindromic-subsequence/

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}

func solve(i,j,n int, s string, dp *[][]int) int {
    if i == j {
        return 1
    }
    if i > j {
        return 0
    }
    if (*dp)[i][j] != -1 {
        return (*dp)[i][j] 
    }
    res := -1
    if s[i] == s[j] {
        res = 2 + solve(i+1, j-1, n, s, dp)
    } else {
        res = max(solve(i+1, j, n, s, dp), solve(i, j-1, n, s, dp))
    }
    (*dp)[i][j] = res
    return res
}

func longestPalindromeSubseq(s string) int {
    n := len(s)
    dp := make([][]int, n)
    for i:=0; i<n; i++ {
        dp[i] = make([]int, n)
        for j:=0; j<n; j++ {
            dp[i][j] = -1
        }
    }
    return solve(0, n-1, n, s, &dp)
}

// Time complexity: O(N^2)
// Space complexity: O(N^2)