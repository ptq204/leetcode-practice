// Link: https://leetcode.com/problems/generate-parentheses/
// Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

func solve(n, openNum, closeNum int, bracket rune, curr string, ans *[]string) {
    curr += string(bracket)
    if openNum == n && closeNum == n {
        *ans = append(*ans, curr)
        return
    }
    if openNum < n {
        solve(n, openNum+1, closeNum, '(', curr, ans)
    }
    if closeNum < openNum {
        solve(n, openNum, closeNum+1, ')', curr, ans)
    }
}

func generateParenthesis(n int) []string {
    ans := make([]string, 0)
    solve(n, 1, 0, '(', "", &ans)
    return ans
}