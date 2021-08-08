// Link: https://leetcode.com/problems/couples-holding-hands/

func minSwapsCouples(row []int) int {
    hmap := make(map[int]int)
    n := len(row)
    for i, k := range row {
        hmap[k] = i
    }
    ans := 0
    for i:=0; i<n; i+=2 {
        lover := row[i]+1
        if row[i] % 2 == 1 {
            lover = row[i]-1
        }
        if row[i+1] == lover {
            continue
        }
        idx := hmap[lover]
        hmap[row[i+1]] = idx
        hmap[lover] = i+1
        row[i+1], row[idx] = row[idx], row[i+1]
        ans++
    }
    return ans
}

// Time complexity: O(N)
// Space complexity: O(N)