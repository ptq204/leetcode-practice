// Link: https://leetcode.com/problems/verifying-an-alien-dictionary/

func isAlienSorted(words []string, order string) bool {
    letters := make(map[rune]int)
    for i, l := range order {
        letters[l] = i
    }
    for i:=1; i<len(words); i++ {
        cur := words[i]
        prev := words[i-1]
        i := 0
        j := 0
        isGreater := false
        for i<len(cur) && j<len(prev) {
            if letters[rune(cur[i])] < letters[rune(prev[j])] {
                return false
            } else if letters[rune(cur[i])] > letters[rune(prev[j])] {
                isGreater = true
                break
            }
            i++
            j++
        }
        if j<len(prev) && !isGreater {
            return false
        }
    }
    return true
}