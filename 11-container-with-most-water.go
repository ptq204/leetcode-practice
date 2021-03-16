// Link: https://leetcode.com/problems/container-with-most-water/

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func maxArea(height []int) int {
    ans := 0
    n := len(height)
    i := 0
    j := n-1
    for i < j {
        val := min(height[i], height[j]) * (j - i)
        if val > ans {
            ans = val
        }
        if height[i] > height[j] {
            j--
        } else {
            i++
        }
    }
    return ans
}

// Time complexity: O(N)
// Space complexity: O(1)