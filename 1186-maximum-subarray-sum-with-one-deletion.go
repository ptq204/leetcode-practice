// Link: https://leetcode.com/problems/maximum-subarray-sum-with-one-deletion/

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}

func maximumSum(arr []int) int {
    n := len(arr)
    dp1 := make([]int, n)
    dp2 := make([]int, n)
    dp1[0] = arr[0]
    dp2[n-1] = arr[n-1]
    ans := max(dp1[0], dp2[n-1])
    
    for i:=1; i<n; i++ {
        dp1[i] = max(arr[i], dp1[i-1] + arr[i])
        ans = max(ans, dp1[i])
    }
    for i:=n-2; i>=0; i-- {
        dp2[i] = max(arr[i], dp2[i+1] + arr[i])
    }
    for i:=1; i<n-1; i++ {
        if arr[i] < 0 {
            ans = max(ans, dp1[i-1] + dp2[i+1])
        }
    }
    return ans
}

// Time complexity: O(N)
// Space complexity: O(N)