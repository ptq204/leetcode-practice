// Link: 

func numSubarrayProductLessThanK(nums []int, k int) int {
    i := 0
    j := 0
    ans := 0
    product := 1
    for j < len(nums) {
        product *= nums[j]
        for i <= j && product >= k {
            product /= nums[i]
            i++
        }
        ans += j - i + 1
        j++
    }
    return ans
}

// Time complexity: O(N)
// Space complexity: O(1)