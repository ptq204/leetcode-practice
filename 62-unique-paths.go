// Link: https://leetcode.com/problems/unique-paths/

func solve(i, j, m, n int, dp *[][]int) int {
    if i >= m || j >= n {
        return 0
    }
    if (*dp)[i][j] != -1 {
        return (*dp)[i][j]
    }
    if i == m-1 && j == n-1 {
        return 1
    }
    ans := 0
    ans += (solve(i+1, j, m, n, dp) + solve(i, j+1, m, n, dp))
    (*dp)[i][j] = ans
    return ans
}

func uniquePaths(m int, n int) int {
    dp := make([][]int, m)
    for i:=0; i<m; i++ {
        dp[i] = make([]int, n)
        for j:=0; j<n; j++ {
            dp[i][j] = -1
        }
    }
    return solve(0, 0, m, n, &dp)
}

// Time complexity: O(M*N)
// Space complexity: O(M*N)