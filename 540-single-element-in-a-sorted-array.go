// Link: https://leetcode.com/problems/single-element-in-a-sorted-array/

func singleNonDuplicate(nums []int) int {
    i := 0
    j := len(nums)-1
    for i < j {
        m := (i+j)/2
        if (m%2==0 && nums[m+1]==nums[m]) ||
           (m%2==1 && nums[m-1]==nums[m]) {
            i = m+1
        } else {
            j = m
        }
    }
    return nums[i]
}

// Time complexity: O(LogN)
// Space complexity: O(1)