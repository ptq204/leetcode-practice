// Link: https://leetcode.com/problems/palindrome-number/
// Given an integer x, return true if x is palindrome integer.

func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    d := make([]int, 0)
    for x > 0 {
        d = append(d, x % 10)
        x /= 10
    }
    i := 0
    j := len(d)-1

    for i < j {
        if d[i] != d[j] {
            return false
        }
        i++
        j--
    }
    return true
}

// number of digits: N
// Time complexity: O(N)
// Space complexity: O(N)