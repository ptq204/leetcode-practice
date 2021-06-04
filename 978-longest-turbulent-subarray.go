// Link: https://leetcode.com/problems/longest-turbulent-subarray/

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func maxTurbulenceSize(arr []int) int {
    if len(arr) == 1 {
        return 1
    }
    ans := 2
    cnt := 2
    n := len(arr)
    isIncr := false
    if arr[1] > arr[0] {
        isIncr = true
    } else if arr[1] == arr[0] {
        cnt = 1
    }
    ans = cnt
    i := 2
    for i < n {
        if (arr[i] > arr[i-1] && !isIncr) || (arr[i] < arr[i-1] && isIncr) {
            cnt++
        } else if arr[i] == arr[i-1] {
            cnt = 1
        } else {
            cnt = 2
        }
        ans = max(ans, cnt)
        isIncr = arr[i] > arr[i-1]
        i++
    }
    return ans
}

// Time complexity: O(N)
// Space complexity: O(1)