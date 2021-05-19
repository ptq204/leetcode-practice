// Link: https://leetcode.com/problems/target-sum/solution/

func solve(idx, n, target, currSum int, nums []int, dp *[][]int) int {
    if idx >= n {
        if currSum == target {
            return 1
        }
        return 0
    }
    scaleSum := currSum + 1000
    if (*dp)[scaleSum][idx] != -1 {
        return (*dp)[scaleSum][idx]
    }
    (*dp)[scaleSum][idx] = solve(idx+1, n, target, currSum+nums[idx], nums, dp) +
    solve(idx+1, n, target, currSum-nums[idx], nums, dp)
    return (*dp)[scaleSum][idx]
}

func findTargetSumWays(nums []int, target int) int {
    n := len(nums)
    dp := make([][]int, 2001)
    for i:=0; i<2001; i++ {
        dp[i] = make([]int, n)
        for j:=0; j<n; j++ {
            dp[i][j] = -1
        }
    }
    return solve(0, n, target, 0, nums, &dp)
}

// Time complexity: O(S*N) where S is the sum of all numbers
// Space complexity: O(S*N)