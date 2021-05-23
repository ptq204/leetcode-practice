// Link: https://leetcode.com/problems/minimum-swaps-to-make-sequences-increasing/

func min(a,b int) int {
    if a < b {
        return a
    }
    return b
}

func solve(idx,n int, nums1, nums2 *[]int, swapped int, dp *[][]int) int {
    if idx >= n {
        return 0
    }
    if (*dp)[swapped][idx] != -1 {
        return (*dp)[swapped][idx]
    }    
    res := 1001
    if idx == 0 || (idx > 0 && (*nums1)[idx] > (*nums1)[idx-1] && (*nums2)[idx] > (*nums2)[idx-1]) {
        val := solve(idx+1, n, nums1, nums2, 0, dp)
        res = min(res, val)
    }
    if idx == 0 || (idx > 0 && (*nums2)[idx] > (*nums1)[idx-1] && (*nums1)[idx] > (*nums2)[idx-1]) {
        (*nums1)[idx], (*nums2)[idx] = (*nums2)[idx], (*nums1)[idx]
        val := solve(idx+1, n, nums1, nums2, 1, dp)
        res = min(res, val+1)
        (*nums1)[idx], (*nums2)[idx] = (*nums2)[idx], (*nums1)[idx]
    }
    (*dp)[swapped][idx] = res
    return res
}

func minSwap(nums1 []int, nums2 []int) int {
    n := len(nums1)
    dp := make([][]int, 2)
    for i:=0; i<2; i++ {
        dp[i] = make([]int, n)
        for j:=0; j<n; j++ {
            dp[i][j] = -1
        }
    }
    ans := solve(0, n, &nums1, &nums2, 0, &dp)
    return ans
}

// Time complexity: O(N)
// Space complexity: O(N)