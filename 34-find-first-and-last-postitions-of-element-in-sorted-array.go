// Link: https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
// Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.

func searchRange(nums []int, target int) []int {
    n := len(nums)
    i := 0
    j := n-1
    for i <= j {
        m := (i+j) / 2
        if nums[m] == target {
            k := m-1
            ans := []int{m, m}
            for k >= 0 && nums[k] == target {
                ans[0] = k
                k--
            }
            k = m+1
            for k < n && nums[k] == target {
                ans[1] = k
                k++
            }
            return ans
        }
        if nums[m] < target {
            i = m + 1
        } else {
            j = m - 1
        }
    }
    return []int{-1, -1}
}