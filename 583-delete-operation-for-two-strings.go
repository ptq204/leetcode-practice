// Link: https://leetcode.com/problems/delete-operation-for-two-strings/

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}

func minDistance(word1 string, word2 string) int {
    n := len(word1)
    m := len(word2)
    dp := make([][]int, n+1)
    for i:=0; i<=n; i++ {
        dp[i] = make([]int, m+1)
    }
    for i:=1; i<=n; i++ {
        for j:=1; j<=m; j++ {
            if word1[i-1] == word2[j-1] {
                dp[i][j] = 1 + dp[i-1][j-1]
            } else {
                dp[i][j] = max(dp[i-1][j], dp[i][j-1])
            }
        }
    }
    return n + m - 2 * dp[n][m]
}

// Time complexity: O(N*M)
// Space complexity: O(N*M)