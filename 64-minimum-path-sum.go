// Link: https://leetcode.com/problems/minimum-path-sum/

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func solve(i, j, m, n int, dp *[][]int, grid [][]int) int {
    if i == m-1 && j == n-1 {
        return grid[i][j]
    }
    if (*dp)[i][j] != -1 {
        return (*dp)[i][j]
    }
    sum := math.MaxInt64
    if i < m-1 {
        sum = min(sum, solve(i+1, j, m, n, dp, grid))
    }
    if j < n-1 {
        sum = min(sum, solve(i, j+1, m, n, dp, grid))
    }
    (*dp)[i][j] = grid[i][j] + sum
    return (*dp)[i][j]
}

func minPathSum(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    dp := make([][]int, m)
    for i:=0; i<m; i++ {
        dp[i] = make([]int, n)
        for j:=0; j<n; j++ {
            dp[i][j] = -1
        }
    }
    return solve(0, 0, m, n, &dp, grid)
}

// Time complexity: O(M*N)
// Space complexity: O(M*N)