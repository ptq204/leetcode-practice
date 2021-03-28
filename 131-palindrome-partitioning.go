// Link: https://leetcode.com/problems/palindrome-partitioning/

func isPalindrome(s string) bool {
    n := len(s)
    for i:=0; i<n/2; i++ {
        if s[i] != s[n-i-1] {
            return false
        } 
    }
    return true
}

func solve(idx, n int, cur []string, ans *[][]string, s string) {
    if idx >= n {
        tmp := make([]string, len(cur))
        copy(tmp, cur)
        *ans = append(*ans, tmp)
        return
    }
    for leng:=1; leng<=n-idx; leng++ {
        if leng == 1 || isPalindrome(s[idx:idx+leng]) {
            cur = append(cur, s[idx:idx+leng])
            solve(idx+leng, n, cur, ans, s)
            cur = cur[:len(cur)-1]
        }
    }
}

func partition(s string) [][]string {
    ans := make([][]string, 0)
    solve(0, len(s), []string{}, &ans, s)
    return ans
}

// Time complexity: O(N*2^N)
// Space complexity: