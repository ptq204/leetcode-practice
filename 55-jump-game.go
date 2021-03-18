// Link: https://leetcode.com/problems/jump-game/

func canJump(nums []int) bool {
    pos := 0
    maxDist := nums[0]
    for pos <= maxDist {
        if pos >= len(nums)-1 {
            return true
        }
        if maxDist <= pos + nums[pos] {
            maxDist = pos + nums[pos]
        }
        pos++
    }
    return false
}

// Time complexity: O(N)
// Space complexity: O(1)