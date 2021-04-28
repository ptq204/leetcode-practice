// Link: https://leetcode.com/problems/min-cost-climbing-stairs/

func min(a,b int) int {
    if a < b {
        return a
    }
    return b
}

func minCostClimbingStairs(cost []int) int {
    n := len(cost)
    a := 0
    b := 0
    for i:=2; i<=n; i++ {
        val := min(b+cost[i-1], a+cost[i-2])
        a = b
        b = val
    }
    return b
}

// Time complexity: O(N)
// Space complexity: O(1)