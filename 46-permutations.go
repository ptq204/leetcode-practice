// Link: https://leetcode.com/problems/permutations/
// Given an array nums of distinct integers, return all the possible permutations.
// You can return the answer in any order.

func solve(check *[]bool, curList []int, ans *[][]int, nums []int) {
    n := len(nums)
    if len(curList) == n {
        tmp := make([]int, n)
        copy(tmp, curList)
        *ans = append(*ans, tmp)
        return
    }
    for i, num := range nums {
        if !(*check)[i] {
            curList = append(curList, num)
            (*check)[i] = true
            solve(check, curList, ans, nums)
            (*check)[i] = false
            curList = curList[:len(curList)-1]
        }
    }
}

func permute(nums []int) [][]int {
    check := make([]bool, len(nums))
    ans := make([][]int, 0)
    solve(&check, []int{}, &ans, nums)
    return ans
}