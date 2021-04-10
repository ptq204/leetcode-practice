// Link: https://leetcode.com/problems/longest-string-chain/

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func isPredecessor(prev, cur string) bool {
    if len(cur) <= len(prev) {
        return false
    }
    for i:=0; i<len(cur); i++ {
        if cur[:i] + cur[i+1:] == prev {
            return true
        }
    }
    return false
}

func solve(idx, n int, words []string, dp *[]int) int {
    if idx >= n {
        return 0
    }
    if (*dp)[idx] != -1 {
        return (*dp)[idx]
    }
    
    ans := 0
    cur := words[idx]
    for i:=idx+1; i<n && len(words[i]) <= len(cur)+1; i++ {
        if isPredecessor(cur, words[i]) {
            ans = max(ans, 1 + solve(i, n, words, dp))
        }
    }
    (*dp)[idx] = ans
    return ans
}

func longestStrChain(words []string) int {
    sort.Slice(words, func (i, j int) bool {
        return len(words[i]) < len(words[j])
    })
    n := len(words)
    dp := make([]int, n+1)
    for i:=0; i<n; i++ {
        dp[i] = -1
    }
    ans := 0
    for i:=0; i<n; i++ {
        ans = max(ans, 1 + solve(i, n, words, &dp)) 
    }
    return ans
}

// Time complexity: O(N^2)
// Space complexity: O(N)