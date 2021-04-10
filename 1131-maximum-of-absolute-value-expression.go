// Link: https://leetcode.com/problems/maximum-of-absolute-value-expression/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findMax(arr1 []int, arr2 []int) int {
    max1 := math.MinInt64
    min1 := math.MaxInt64
    
    max2 := math.MinInt64
    min2 := math.MaxInt64
    
    max3 := math.MinInt64
    min3 := math.MaxInt64
    
    max4 := math.MinInt64
    min4 := math.MaxInt64
    
    for i,_ := range arr1 {
        max1 = max(max1, arr1[i]+arr2[i]+i)
        min1 = min(min1, arr1[i]+arr2[i]+i)
        
        max2 = max(max2, arr1[i]+arr2[i]-i)
        min2 = min(min2, arr1[i]+arr2[i]-i)
        
        max3 = max(max3, arr1[i]-arr2[i]+i)
        min3 = min(min3, arr1[i]-arr2[i]+i)
        
        max4 = max(max4, arr1[i]-arr2[i]-i)
        min4 = min(min4, arr1[i]-arr2[i]-i)
    }
    
    ans := math.MinInt64
    ans = max((max1-min1), (max2-min2))
    ans = max(ans, (max3-min3))
    ans = max(ans, (max4-min4))
    return ans
}

func maxAbsValExpr(arr1 []int, arr2 []int) int {
    return max(findMax(arr1, arr2), findMax(arr2, arr1))
}

// Time complexity: O(N)
// Space complexity: O(1)
