// Link:

func min(a,b int) int {
    if a < b {
        return a
    }
    return b
}

func numRollsToTarget(d int, f int, target int) int {
    dp := make([][]int, d+1)
    mod := int(math.Pow(10,9) + 7)
    for i:=0; i<=d; i++ {
        dp[i] = make([]int, target+1)
    }
    dp[0][0] = 1
    for k:=1; k<=target; k++ {
        for i:=1; i<=d; i++ {
            limit := min(k, f)
            for j:=1; j<=limit; j++ {
                if k >= j {
                    dp[i][k] += dp[i-1][k-j]
                    dp[i][k] %= mod
                }
            }
        }
    }
    return dp[d][target]
}

// Time complexity: O(d*f*target)
// Space complexity: O(d*target)