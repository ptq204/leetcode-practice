// Link: https://leetcode.com/problems/relative-sort-array/

func relativeSortArray(arr1 []int, arr2 []int) []int {
    hmap := make(map[int]int)
    for _, k := range arr2 {
        hmap[k] = 0
    }
    tmp := make([]int, 0)
    for _, k := range arr1 {
        if _, ok := hmap[k]; !ok {
            tmp = append(tmp, k)
        } else {
            hmap[k]++
        }
    }
    sort.Slice(tmp, func (i,j int) bool {
        return tmp[i] < tmp[j]
    })
    i := 0
    for _, k := range arr2 {
        v, _ := hmap[k]
        for j:=0; j<v; i,j=i+1,j+1 {
            arr1[i] = k
        }
    }
    j := i
    for i < len(arr1) {
        arr1[i] = tmp[i-j]
        i++
    }
    return arr1
}

// Time complexity: max(O(N), O(MlogM)) where M = numbers of elements not appear in arr2 
// Space complexity: O(len(arr2))