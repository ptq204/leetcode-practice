// Link: https://leetcode.com/problems/merge-intervals/

func max(i,j int) int {
    if i < j {
        return j
    }
    return i
}

func min(i,j int) int {
    if i < j {
        return i
    }
    return j
}

func overlap(a, b []int) bool {
    return a[0] <= b[0] && a[1] >= b[0]
}

func merge(intervals [][]int) [][]int {
    ans := make([][]int, 0)
    sort.Slice(intervals, func (i,j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    cur := intervals[0]
    n := len(intervals)
    for i:=1; i<n; i++ {
        if overlap(cur, intervals[i]) || overlap(intervals[i], cur) {
            cur[0] = min(cur[0], intervals[i][0])
            cur[1] = max(cur[1], intervals[i][1])
        } else {
            ans = append(ans, cur)
            cur = intervals[i]
        }
    }
    ans = append(ans, cur)
    return ans
}

// Time complexity: O(NlogN)
// Space complexity: O(1)