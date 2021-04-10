// Link: https://leetcode.com/problems/sum-root-to-leaf-numbers/

func solve(cur int, root *TreeNode) int {
    if root == nil {
        return 0
    }
    cur = cur*10 + root.Val
    if root.Left == nil && root.Right == nil {
        return cur
    }
    return solve(cur, root.Right) + solve(cur, root.Left) 
}

func sumNumbers(root *TreeNode) int {
    cur := 0
    return solve(cur, root)
}

// Time complexity: O(N)
// Space complexity: O(1)