// Link: https://leetcode.com/problems/array-nesting/

func arrayNesting(nums []int) int {
    n := len(nums)
    ans := 0
    for k:=0; k<n; k++ {
        curr := nums[k]
        if curr == -1 {
            continue
        }
        leng := 0
        for nums[curr] != -1 {
            tmp := curr
            curr = nums[curr]
            nums[tmp] = -1
            leng++
        }
        if ans < leng {
            ans = leng
        }
    }
    return ans
}

// Time complexity: O(N)
// Space complexity: O(1)