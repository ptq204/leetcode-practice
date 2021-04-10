// Link: https://leetcode.com/problems/max-chunks-to-make-sorted/

func maxChunksToSorted(arr []int) int {
    sum := 0
    bound := 0
    ans := 0
    for i, k := range arr {
        sum += k
        tmp := (i+1)*(i+2)/2 - bound*(bound+1)/2 - (i+1)
        if sum == tmp {
            ans++
            sum = 0
            bound = i
        }
    }
    return ans
}

// Time complexity: O(N)
// Space complexity: O(1)