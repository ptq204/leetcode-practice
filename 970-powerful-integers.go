// Link: https://leetcode.com/problems/powerful-integers/

func min(a,b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}

func powerfulIntegers(x int, y int, bound int) []int {
    a := min(x, y)
    if a == 1 {
        a = max(x, y)
    }
    k := 0
    for a != 1 && int(math.Pow(float64(a), float64(k))) < bound {
        k++
    }
    ans := make([]int, 0)
    hmap := make(map[int]bool)
    for i:=0; i<=k; i++ {
        for j:=0; j<=k; j++ {
            m := int(math.Pow(float64(x), float64(i)))
            n := int(math.Pow(float64(y), float64(j)))
            if m < 0 || n < 0 {
                continue
            }
            val := m + n
            if val <= bound {
                if _, ok := hmap[val]; !ok {
                    hmap[val] = true
                    ans = append(ans, val)
                }
            }
        }
    }
    return ans
}

// Time complexity: O(K^2) 
// Space complexity: O(N)