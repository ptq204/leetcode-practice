// Link: https://leetcode.com/problems/longest-palindromic-substring/
// Given a string s, return the longest palindromic substring in s.

func solve(i,j int, s string, dp *[][]int) int {
    if (*dp)[i][j] != -1 {
        return (*dp)[i][j]
    }
    if i == j {
        return 1
    }
    if i > j {
        return 0
    }
    if s[i] != s[j] {
        (*dp)[i][j] = -len(s)
    } else {
        (*dp)[i][j] = 2 + solve(i+1, j-1, s, dp)
    }
    return (*dp)[i][j]
}

func longestPalindrome(s string) string {
    n := len(s)
    if n == 1 {
        return s
    }
    dp := make([][]int, n)
    for i:=0; i<n; i++ {
        dp[i] = make([]int, n)
        for j:=0; j<n; j++ {
            dp[i][j] = -1
        }
    }

    ans := ""
    maxL := 0
    for i:=0; i<n; i++ {
        for j:=0; j<=i; j++ {
            l := solve(j, i, s, &dp)
            if l > maxL {
                ans = s[j:i+1]
                maxL = l
            }
        }
    }
    return ans
}

// Time complexity: O(N^2)
// Space complexity: O(N^2)