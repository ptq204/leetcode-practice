// Link: https://leetcode.com/problems/group-anagrams/

func groupAnagrams(strs []string) [][]string {
    cache := make(map[string][]string)
    for _, str := range strs {
        tmp := []rune(str)
        sort.Slice(tmp, func (i, j int) bool {
            return tmp[i] < tmp[j]
        })
        s := string(tmp)
        if _, ok := cache[s]; !ok {
            cache[s] = make([]string, 0)
        }
        cache[s] = append(cache[s], str)
    }
    ans := make([][]string, 0)
    for _, anagrams := range cache {
        ans = append(ans, anagrams)
    }
    return ans
}

// Time complexity: O(N*(klogk)) where k is the length of each string
// Space complexity: O(N)