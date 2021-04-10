// Link: https://leetcode.com/problems/tuple-with-same-product/

func tupleSameProduct(nums []int) int {
    hmap := make(map[int]int)
    for i:=0; i<len(nums); i++ {
        for j:=i+1; j<len(nums); j++ {
            k := nums[i] * nums[j]
            if _, ok := hmap[k]; !ok {
                hmap[k] = 1
            } else {
                hmap[k]++
            }
        }
    }
    ans := 0
    for _, v := range hmap {
        if v > 1 {
            ans += 4 * v * (v - 1)
        }
    }
    return ans
}

// Time complexity: O(N^2)
// Space complexity: O(N^2)