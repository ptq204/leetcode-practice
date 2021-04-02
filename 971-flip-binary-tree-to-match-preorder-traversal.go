// Link: https://leetcode.com/problems/flip-binary-tree-to-match-preorder-traversal/

func solve(root *TreeNode, idx *int, ans *[]int, voyage []int) {
    if root == nil || *idx >= len(voyage) {
        return
    }
    if root.Val != voyage[*idx] {
        *ans = []int{-1}
        return
    }
    (*idx)++
    if *idx < len(voyage) && root.Left != nil && root.Left.Val != voyage[*idx] {
        *ans = append(*ans, root.Val)
        solve(root.Right, idx, ans, voyage)
        solve(root.Left, idx, ans, voyage)
    } else {
        solve(root.Left, idx, ans, voyage)
        solve(root.Right, idx, ans, voyage)
    }
}

func flipMatchVoyage(root *TreeNode, voyage []int) []int {
    ans := make([]int, 0)
    idx := 0
    solve(root, &idx, &ans, voyage)
    if len(ans) > 0 && ans[0] == -1 {
        return []int{-1}
    }
    return ans
}

// Time complexity: O(N)
// Space complexity: O(N)