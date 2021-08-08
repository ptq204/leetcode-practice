// Link: https://leetcode.com/problems/maximize-distance-to-closest-person/

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a,b int) int {
    if a < b {
        return a
    }
    return b
}

func maxDistToClosest(seats []int) int {
    ans := 0
    l := 0
    r := 0
    n := len(seats)
    for r < n {
        for r < n && seats[r] == 0 {
            r++
        }
        if r == n {
            return max(ans, n-l-1)
        }
        if seats[l] == 0 {
            ans = max(ans, r-l)
        } else {
            for k:=l+1; k<r; k++ {
                dist := min(k-l, r-k)
                ans = max(ans, dist)
            }
        }
        l = r
        r++
    }
    return ans
}

// Time complexity: O(N)
// Space complexity: O(1)