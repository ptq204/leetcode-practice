// Link: https://leetcode.com/problems/insert-interval/

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

func insert(intervals [][]int, newInterval []int) [][]int {
    ans := make([][]int, 0)
    n := len(intervals)
    for i:=0; i<n; i++ {
        if overlap(newInterval, intervals[i]) || overlap(intervals[i], newInterval) {
            newInterval[0] = min(newInterval[0], intervals[i][0])
            newInterval[1] = max(newInterval[1], intervals[i][1])
        } else {
            if intervals[i][0] <= newInterval[0] {
                ans = append(ans, intervals[i])
            } else {
                ans = append(ans, newInterval)
                newInterval = intervals[i]
            }
        }
    }
    ans = append(ans, newInterval)
    return ans
}

// Time complexity: O(N)
// Space complexity: O(1)