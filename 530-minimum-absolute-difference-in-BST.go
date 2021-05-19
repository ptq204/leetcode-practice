// Link: https://leetcode.com/problems/minimum-absolute-difference-in-bst/

func min(a,b int) int {
    if a < b {
        return a
    }
    return b
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}

func solve(node *TreeNode, ans, prev *int) {
    if node == nil {
        return
    }
    solve(node.Left, ans, prev)
    if *prev != -1 {
        *ans = min(*ans, abs(node.Val-*prev))
    }
    *prev = node.Val
    solve(node.Right, ans, prev)
}

func getMinimumDifference(root *TreeNode) int {
    ans := math.MaxInt64
    prev := -1
    solve(root, &ans, &prev)
    return ans
}

// Time complexity: O(N)
// Space complexity: O(N)