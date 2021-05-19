// Link: https://leetcode.com/problems/maximum-points-you-can-obtain-from-cards/

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}

func maxScore(cardPoints []int, k int) int {
    n := len(cardPoints)
    i := 0
    j := n-k-1
    prefixSum := make([]int, n)
    sum := 0
    for i, p := range cardPoints {
        sum += p
        prefixSum[i] = sum 
    }
    if k == n {
        return sum
    }
    ans := 0
    for j < n {
        ans = max(ans, sum - prefixSum[j] + prefixSum[i] - cardPoints[i])
        j++
        i++
    }
    return ans
}

// Time complexity: O(N)
// Space complexity: O(N)