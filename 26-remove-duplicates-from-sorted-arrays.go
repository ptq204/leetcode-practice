// Link: https://leetcode.com/problems/remove-duplicates-from-sorted-array/
// Given a sorted array nums, remove the duplicates in-place such that each element appears only once and returns the new length.

func removeDuplicates(nums []int) int {
    n := len(nums)
    for i:=0; i < n; i++ {
        cur := nums[i]
        j := i+1
        cnt := 0
        for j < n && nums[j] == cur {
            cnt++
            j++
        }
        for j < n {
            nums[j-cnt] = nums[j]
            j++
        }
        n -= cnt
    }
    nums = nums[:n]
    return len(nums)
}

// Time complexity: O(N)
// Space complexity: O(1)