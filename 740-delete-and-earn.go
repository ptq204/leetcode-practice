// Link: https://leetcode.com/problems/delete-and-earn/

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}

func deleteAndEarn(nums []int) int {
    maxx := 0
    for _, k := range nums {
        if k > maxx {
            maxx = k
        }
    }
    amount := make([]int, maxx+1)
    for _, k := range nums {
        if amount[k] == 0 {
            amount[k] = 1
        } else {
            amount[k]++
        }
    }
    a := 0
    b := amount[1]
    c := 0
    for i:=2; i<=maxx; i++ {
        c = max(c, a + amount[i]*i)
        c = max(c, b)
        a = b
        b = c
    }
    return b
}

// Time complexity: O(N)
// Space complexity: O(1)