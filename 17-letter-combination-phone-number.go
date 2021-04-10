// Link: https://leetcode.com/problems/letter-combinations-of-a-phone-number/

func letterCombinations(digits string) []string {
    var ans []string
    if len(digits) == 0 {
        return ans
    }
    phone := make(map[int]string)
    phone[2] = "abc"
    phone[3] = "def"
    phone[4] = "ghi"
    phone[5] = "jkl"
    phone[6] = "mno"
    phone[7] = "pqrs"
    phone[8] = "tuv"
    phone[9] = "wxyz"
    for i, d := range digits {
        curr := phone[int(d-'0')]
        if i == 0 {
            ans = make([]string, 0)
            for _, c := range curr {
                ans = append(ans, string(c))
            }
            continue
        }
        currList := make([]string, 0)
        for _, k := range ans {
            for _, t := range curr {
                currList = append(currList, k + string(t))
            }
        }
        ans = currList
    }
    return ans
}

// Time complexity: O(3^N * 4^M) where N is number of digits which have length 3 letter , M is number of digits which have length 4 letter.
// Space complexity: O(3^N * 4^M)