// Link: https://leetcode.com/problems/largest-sum-of-averages/

func max(a,b float64) float64 {
    if a > b {
        return a
    }
    return b
}

func solve(idx, n, k int, nums, prefixSum []int, dp *[][]float64) float64 {
    if idx >= n {
        return 0
    }
    if k == 1 {
        return float64(prefixSum[n-1] - prefixSum[idx] + nums[idx]) / float64(n - idx)
    }
    if (*dp)[idx][k] != -1.0 {
        return (*dp)[idx][k]
    }
    var res float64
    for i:=0; i<n-idx; i++ {
        nextIdx := idx + i
        average := float64(prefixSum[nextIdx] - prefixSum[idx] + nums[idx]) / float64(nextIdx-idx+1)
        res = max(res, average + solve(idx+i+1, n, k-1, nums, prefixSum, dp))
    }
    (*dp)[idx][k] = res
    return res
}

func largestSumOfAverages(nums []int, k int) float64 {
    n := len(nums)
    dp := make([][]float64, n)
    for i:=0; i<n; i++ {
        dp[i] = make([]float64, k+1)
        for j:=0; j<=k; j++ {
            dp[i][j] = float64(-1)
        }
    }
    prefixSum := make([]int, n)
    sum := 0
    for i, v := range nums {
        sum += v
        prefixSum[i] = sum
    }
    return solve(0, n, k, nums, prefixSum, &dp)
}

// Time complexity: O(N*K)
// Space complexity: O(N*K)