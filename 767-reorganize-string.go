// Link: https://leetcode.com/problems/reorganize-string/

func reorganizeString(S string) string {
    hmap := make(map[rune]int)
    maxChar := rune(S[0])
    maxCnt := 0
    for _, c := range S {
        if _, ok := hmap[c]; !ok {
            hmap[c] = 1
        } else {
            hmap[c]++
        }
        if hmap[c] > maxCnt {
            maxCnt = hmap[c]
            maxChar = c
        }
    }
    ans := make([]rune, len(S))
    i := 0
    cnt := 0
    for cnt < maxCnt && i < len(S) {
        ans[i] = maxChar
        i+=2
        cnt++
    }
    if cnt < maxCnt {
        return ""
    }
    delete(hmap, maxChar)
    for k, v := range hmap {
        for j:=0; j<v; j++ {
            if i >= len(S) {
                i = 1
            }
            ans[i] = k
            if ans[i] == ans[i-1] {
                return ""
            }
            i+=2
        }
    }
    return string(ans)
}

// Time complexity: O(N)
// Space complexity: O(N)