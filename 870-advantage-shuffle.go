// Link: https://leetcode.com/problems/advantage-shuffle/

func advantageCount(nums1 []int, nums2 []int) []int {
    n := len(nums1)
    nums2Sort := make([][]int, 0)
    for i, k := range nums2 {
        nums2Sort = append(nums2Sort, []int{i, k})
    }
    sort.Slice(nums2Sort, func (i, j int) bool {
        return nums2Sort[i][1] < nums2Sort[j][1]
    })
    sort.Slice(nums1, func (i, j int) bool {
        return nums1[i] < nums1[j]
    })
    idx := 0
    ans := make([]int, n)
    for i:=0; i<n; i++ {
        ans[i] = -1
    }
    tmp := make([]int, 0)
    
    for _, k := range nums2Sort {
        for i:=idx; i<n; i++ {
            if nums1[i] > k[1] {
                ans[k[0]] = nums1[i]
                idx = i + 1
                break
            } else {
                tmp = append(tmp, nums1[i])
            }
        }
    }
    idx = 0
    for i, k := range ans {
        if k == -1 {
            ans[i] = tmp[idx]
            idx++
        }
        
    }
    return ans
}

// Time complexity: O(N)
// Space complexity: O(N)