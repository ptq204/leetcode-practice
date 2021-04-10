// Link: https://leetcode.com/problems/number-of-different-integers-in-a-string/

func numDifferentIntegers(word string) int {
    newWord := ""
    for _, c := range word {
        if c >= '0' && c <= '9' {
            newWord += string(c)
        } else {
            newWord += " "
        }
    }
    nums := strings.Fields(newWord)
    hmap := make(map[int]bool)
    for _, num := range nums {
        k := 0
        for _, d := range num {
            k = 10 * k + int(d - '0')
        }
        if _, ok := hmap[k]; !ok {
            hmap[k] = true
        }
    }
    return len(hmap)
}

// Time complexity: O(N)
// Space complexity: O(N)