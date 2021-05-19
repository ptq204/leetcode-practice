// Link: https://leetcode.com/problems/minimum-size-subarray-sum/

func min(a,b int) int {
    if a < b {
        return a
    }
    return b
}

func minSubArrayLen(target int, nums []int) int {
    i := 0
    j := 0
    n := len(nums)
    sum := 0
    ans := math.MaxInt64
    for j < n {
        sum += nums[j]
        if sum >= target {
            for i <= j && sum >= target {
                fmt.Println(sum, i, j)
                ans = min(ans, j-i+1)
                sum -= nums[i]
                i++
            }
        }
        j++
    }
    if ans == math.MaxInt64 {
        ans = 0
    }
    return ans
}

// Time complexity: O(N)
// Space complexity: O(1)