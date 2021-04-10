// Link: https://leetcode.com/problems/longest-common-prefix/
// Write a function to find the longest common prefix string amongst an array of strings.

func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    tmp := ""
    for i, s := range strs {
        if i == 0 {
            tmp = s
        } else if len(tmp) == 0 {
            return tmp
        } else {
            j := 0
            for j < len(s) && j < len(tmp) {
                if tmp[j] != s[j] {
                    break
                }
                j++
            }
            tmp = tmp[:j]
        }
    }
    return tmp
}

// Time complexity: O(N)
// Space complexity: O(k) where k is the length of the first string