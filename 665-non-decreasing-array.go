// Link: https://leetcode.com/problems/non-decreasing-array/

func checkPossibility(nums []int) bool {
    breakPnt := 0
    cntBreak := 0
    for i, k := range nums {
        if i > 0 && k < nums[i-1] {
            breakPnt = i
            cntBreak++
        }
    }
    if cntBreak > 1 {
        return false
    }
    if breakPnt >= 2 {
        if nums[breakPnt-2] > nums[breakPnt] {
            if breakPnt < len(nums)-1 && nums[breakPnt+1] < nums[breakPnt-1] {
                return false
            }
        }
    }
    return true
}

// Time complexity: O(N)
// Space complexity: O(1)