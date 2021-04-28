// Link: https://leetcode.com/problems/count-binary-substrings/

func countBinarySubstrings(s string) int {
    zero := 0
    one := 0
    isZero := false
    ans := 0
    for _, c := range s {
        if c == '0' {
            if isZero {
                zero++
            } else {
                zero=1
            }
            isZero = true
        } else {
            if !isZero {
                one++
            } else {
                one=1
            }
            isZero = false
        }
        if zero > 0 && one > 0 {
            ans++
            if one == zero {
                if isZero {
                    one = 0
                } else {
                    zero = 0
                }
            }
        }
    }
    return ans
}

// Time complexity: O(N)
// Space complexity: O(1)