// Link: https://leetcode.com/problems/check-array-formation-through-concatenation/

func canFormArray(arr []int, pieces [][]int) bool {
    hmap := make(map[int][]int)
    for _, p := range pieces {
        hmap[p[0]] = p
    }
    i := 0
    n := len(arr)
    for i < n {
        if _, ok := hmap[arr[i]]; !ok {
            return false
        }
        p := hmap[arr[i]]
        for _, e := range p {
            if e != arr[i] {
                return false
            }
            i++
        }
    }
    return true
}