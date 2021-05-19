// Link: https://leetcode.com/problems/unique-paths-ii/

func solve(i,j,n,m int, grid [][]int, dp *[][]int) int {
    if i == n-1 && j == m-1 && grid[i][j] == 0 {
        return 1
    }
    if i<0 || i>n-1 || j<0 || j>m-1 || grid[i][j] == 1 {
        return 0
    }
    if (*dp)[i][j] != -1 {
        return (*dp)[i][j]
    }
    (*dp)[i][j] = solve(i+1, j, n, m, grid, dp) + solve(i, j+1, n, m, grid, dp)
    return (*dp)[i][j]
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    n := len(obstacleGrid)
    m := len(obstacleGrid[0])
    dp := make([][]int, n)
    for i:=0; i<n; i++ {
        dp[i] = make([]int, m)
        for j:=0; j<m; j++ {
            dp[i][j] = -1
        }
    }
    return solve(0, 0, n, m, obstacleGrid, &dp)
}

// Time complexity: O(N*M)
// Space complexity: O(N*M)