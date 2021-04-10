// Link: https://leetcode.com/problems/3sum-closest/
// find three integers in nums such that the sum is closest to target

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}

func threeSumClosest(nums []int, target int) int {
    sort.Slice(nums, func (i, j int) bool {
        return nums[i] < nums[j]
    })
    n := len(nums)
    minDiff := math.MaxInt64
    ans := 0
    for i:=0; i<n; i++ {
        l := i+1
        h := n-1
        for l < h {
            sum := nums[i]+nums[l]+nums[h]
            diff := abs(sum - target)
            if diff == 0 {
                return target
            }
            if diff < minDiff {
                minDiff = diff
                ans = sum
            }
            if sum < target {
                l++
            } else {
                h--
            }
        }
    }
    return ans
}

// Time complexity: O(N^2)
// Space complexity: O(1)