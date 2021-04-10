// Link: https://leetcode.com/problems/find-common-characters/

func findCommon(hmap map[string]int, s string) map[string]int {
    res := make(map[string]int)
    for _, c := range s {
        k := string(c)
        if _, ok := hmap[k]; ok && hmap[k] > 0 {
            if _, ok := res[k]; !ok {
                res[k] = 1
            } else {
                res[k]++
            }
            hmap[k]--
        }
    }
    return res
}

func commonChars(A []string) []string {
    hmap := make(map[string]int)
    for _, c := range A[0] {
        k := string(c)
        if _, ok := hmap[k]; !ok {
            hmap[k] = 1
        } else {
            hmap[k]++
        }
    }
    for i:=1; i<len(A); i++ {
        hmap = findCommon(hmap, A[i])
    }
    ans := make([]string, 0)
    for k, v := range hmap {
        for i:=0; i<v; i++ {
            ans = append(ans, k)
        }
    }
    return ans
}

// Time complexity: O(N*M) where M is length of string
// Space complexity: O(N*M)