// Link: https://leetcode.com/problems/number-of-subarrays-with-bounded-maximum/

func inRange(num, left, right int) bool {
    return num >= left && num <= right
}

func numSubarrayBoundedMax(nums []int, left int, right int) int {
    i := 0
    j := 0
    ans := 0
    lastValid := 0
    for idx, k := range nums {
        if inRange(k, left, right) {
            ans += j - i + 1
            lastValid = j - i + 1
        } else if k > right {
            i = idx+1
            j = idx+1
            lastValid = 0
            continue
        } else {
            ans += lastValid
        }
        j++
    }
    return ans
}

// Time complexity: O(N)
// Space complexity: O(1)