// Link: https://leetcode.com/problems/array-of-doubled-pairs/

func canReorderDoubled(arr []int) bool {
    hmap := make(map[int]int)
    for _, k := range arr {
        if _, ok := hmap[k]; !ok {
            hmap[k] = 1
        } else {
            hmap[k]++
        }
    }
    sort.Slice(arr, func (i,j int) bool {
        return arr[i] < arr[j]
    })
    for _, k := range arr {
        if k % 2 == 1 && hmap[k] > 0 {
            if _, ok := hmap[k*2]; ok {
                if hmap[k*2] > 0 {
                    hmap[k]--
                    hmap[k*2]--
                }
            }
            continue
        }
        if k % 2 == 0 && hmap[k] > 0 {
            if _, ok := hmap[k/2]; ok {
                if hmap[k/2] > 0 {
                    hmap[k]--
                    hmap[k/2]--
                }
            }
        }
    }
    for _, v := range hmap {
        if v != 0 {
            return false
        }
    }
    return true
}

// Time complexity: O(NlogN)
// Space complexity: O(N)