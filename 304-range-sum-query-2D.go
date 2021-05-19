// Link: https://leetcode.com/problems/range-sum-query-2d-immutable/

type NumMatrix struct {
    dp [][]int
}

func Constructor(matrix [][]int) NumMatrix {
    n := len(matrix)
    m := len(matrix[0])
    dp := make([][]int, n+1)
    for i:=0; i<=n; i++ {
        dp[i] = make([]int, m+1)
        for j:=0; j<=m; j++ {
            if i == 0 || j == 0 {
                dp[i][j] = 0
                continue
            }
            dp[i][j] = matrix[i-1][j-1] + dp[i][j-1] + dp[i-1][j] - dp[i-1][j-1]
        }
    }
    return NumMatrix{dp}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    r1, c1, r2, c2 := row1+1, col1+1, row2+1, col2+1
    return this.dp[r2][c2] - this.dp[r2][c1-1] - this.dp[r1-1][c2] + this.dp[r1-1][c1-1]
}

// Time complexity: O(1)
// Space complexity: O(N*M)