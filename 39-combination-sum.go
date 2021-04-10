// Link: https://leetcode.com/problems/combination-sum/

func solve(idx, n, sum, target int, curList []int, ans *[][]int, candidates []int) {
    if sum == target {
        tmp := make([]int, len(curList))
        copy(tmp, curList)
        *ans = append(*ans, tmp)
        return
    }
    if sum > target {
        return
    }
    for i:=idx; i<n; i++ {
        curList = append(curList, candidates[i])
        solve(i, n, sum+candidates[i], target, curList, ans, candidates)
        curList = curList[:len(curList)-1]
    }
}

func combinationSum(candidates []int, target int) [][]int {
    ans := make([][]int, 0)
    solve(0, len(candidates), 0, target, []int{}, &ans, candidates)
    return ans
}

// Time complexity: 
// Space complexity: 