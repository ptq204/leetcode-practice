// Link: https://leetcode.com/problems/champagne-tower/

func min(a,b float64) float64 {
    if a < b {
        return a
    }
    return b
}

func champagneTower(poured int, query_row int, query_glass int) float64 {
    // max_row := 100
    dp := make([][]float64, query_row+1)
    for i:=0; i<=query_row; i++ {
        dp[i] = make([]float64, query_row+1)
    }
    dp[0][0] = float64(poured)
    for i:=0; i<query_row; i++ {
        for j:=0; j<=i; j++ {
            if dp[i][j] >= 1 {
                overflow := float64(dp[i][j]-1)/2.0
                dp[i+1][j] += overflow
                dp[i+1][j+1] += overflow
                dp[i][j] = 1.0
            }
        }
    }
    return min(dp[query_row][query_glass], 1.0)
}

// Time complexity: O(R^2) where R is query_row
// Space complexity: O(R^2)