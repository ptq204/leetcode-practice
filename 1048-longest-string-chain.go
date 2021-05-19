// Link: https://leetcode.com/problems/longest-string-chain/

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func longestStrChain(words []string) int {
    sort.Slice(words, func (i,j int) bool {
        return len(words[i]) < len(words[j])
    })
    n := len(words)
    dp := make(map[string]int, n)
    for _, w := range words {
        dp[w] = 1
    }
    ans := 0
    for _, w := range words {
        for i:=0; i<len(w); i++ {
            predecessor := w[:i] + w[i+1:]
            if _, ok := dp[predecessor]; ok {
                dp[w] = max(dp[w], 1 + dp[predecessor])
            }
        }
        ans = max(ans, dp[w])
    }
    return ans
}

// Time complexity: O(N*W) where W is length of word
// Space complexity: O(N)