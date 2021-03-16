// Link: https://leetcode.com/problems/valid-parentheses/
// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

/*
An input string is valid if:
- Open brackets must be closed by the same type of brackets.
- Open brackets must be closed in the correct order.
*/

func isValid(s string) bool {
    openStack := make([]rune, 0)
    for _, c := range s {
        if c == '(' || c == '[' || c == '{' {
            openStack = append(openStack, c)
            continue
        }
        n := len(openStack)
        if n == 0 {
            return false
        }
        if c == ')' && openStack[n-1] != '('  {
            return false
        }
        if c == ']' && openStack[n-1] != '[' {
            return false
        }
        if c == '}' && openStack[n-1] != '{'  {
            return false
        }
        openStack = openStack[:(n-1)]
    }
    return len(openStack) == 0
}

// Time complexity: O(N)
// Space complexity: O(N/2)