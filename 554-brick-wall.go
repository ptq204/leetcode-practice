// Link: https://leetcode.com/problems/brick-wall/

func min(a,b int) int {
    if a < b {
        return a
    }
    return b
}

func leastBricks(wall [][]int) int {
    maxCol := 0
    for k, r := range wall {
        s := r[0]
        for i:=1; i<len(r); i++ {
            s += r[i]
            wall[k][i] = s
        }
        if len(r) > maxCol {
            maxCol = len(r)
        }
    }
    ans := len(wall)
    hmap := make(map[int]int)
    notOverlap := 0
    for _, r := range wall {
        for i:=0; i<len(r)-1; i++ {
            leng := r[i]
            if _, ok := hmap[leng]; !ok {
                hmap[leng] = 1
            } else {
                hmap[leng]++
            }
            if hmap[leng] > notOverlap {
                notOverlap = hmap[leng]
            }
        }
    }
    return min(ans, len(wall)-notOverlap)
}

// Time complexity: O(N*M) where N is the number of rows, M is the medium number of bricks per row.
// Space complexity: O(N)