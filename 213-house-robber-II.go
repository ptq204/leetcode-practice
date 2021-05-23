// Link: https://leetcode.com/problems/house-robber-ii/

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}

func rob(nums []int) int {
    n := len(nums)
    if n == 1 {
        return nums[0]
    }
    a := 0
    b := nums[0]
    c := 0
    d := 0
    k := nums[0]
    t := nums[1]
    for i:=1; i<n; i++ {
        if i >= 2 {
            if i < n-1 {
                k = max(k, a+nums[i])
            }
            t = max(t, c+nums[i])
        }
        if i < n-1 {
            k = max(k, b)
        }
        t = max(t, d)
        a = b
        b = k
        c = d
        d = t
    }
    return max(b, d)
}

// Time complexity: O(N)
// Space complexity: O(1)