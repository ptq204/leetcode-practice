// Link: https://leetcode.com/problems/combination-sum-iv/

func combinationSum4(nums []int, target int) int {
    dp := make([]int, target+1)

    for t:=1; t<=target; t++ {
        for _, k := range nums {
            if k > t {
                continue
            }
            if k == t {
                dp[t] += 1
            } else {
                dp[t] += dp[t-k]
            }
        }
    }
    return dp[target]
}

// Time complexity: O(target*N)
// Space complexity: O(N)