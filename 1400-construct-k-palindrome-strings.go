// Link: https://leetcode.com/problems/construct-k-palindrome-strings/

func canConstruct(s string, k int) bool {
    if k > len(s) {
        return false
    }
    hmap := make(map[string]int)
    for _, c := range s {
        t := string(c)
        if _, ok := hmap[t]; !ok {
            hmap[t] = 1
        } else {
            hmap[t]++
        }
    }
    if len(hmap) == k {
        return true
    }
    oddCnt := 0
    for _, v := range hmap {
        if v % 2 == 1 {
            oddCnt++
        }
    }
    if oddCnt > k {
        return false
    }
    return true
}

// Time complexity: O(N)
// Space complexity: O(1)