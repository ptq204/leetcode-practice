// Link: https://leetcode.com/problems/push-dominoes/

func pushDominoes(dominoes string) string {
    i := 0
    j := 1
    n := len(dominoes)
    ans := []byte(dominoes)
    for j < n {
        for j < n-1 && ans[j] == '.' {
            j++
        }
        if ans[i] != 'R' && ans[j] == 'L' {
            for i < j {
                ans[i] = 'L'
                i++
            }
        } else if ans[i] != 'R' && ans[j] == 'R' {
            i = j
        } else if ans[i] == 'R' && ans[j] != 'L' {
            for i < j {
                ans[i] = 'R'
                i++
            }
            ans[j] = 'R'
        } else if ans[i] == 'R' && ans[j] == 'L' {
            right := i+1
            left := j-1
            for right < left {
                ans[right] = 'R'
                ans[left] = 'L'
                right++
                left--
            }
            i = j+1
        }
        j++
    }
    return string(ans)
}

// Time complexity: O(N)
// Space complexity: O(N)