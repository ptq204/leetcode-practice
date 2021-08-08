// Link:

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func findClosestElements(arr []int, k int, x int) []int {
    n := len(arr)
    i := 0
    j := n-1
    idx := 0
    for i <= j {
        m := (i+j)/2
        if arr[m] == x {
            idx = m
            break
        } else if arr[m] > x {
            j = m-1
        } else {
            i = m+1
        }
    }
    idx = i
    i = max(0, idx-k)
    j = min(n-1, idx+k)
    for j - i >= k {
        if x-arr[i] <= arr[j]-x {
            j--
        } else {
            i++
        }
    }
    return arr[i:j+1]
}

// Time complexity: O(logN + K)
// Space complexity: O(1)