// Link: https://leetcode.com/problems/longest-univalue-path/

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func solve(node *TreeNode, curLen, prevVal int, maxLen *int) int {
    if node == nil {
        return 0
    }
    if node.Val == prevVal {
        curLen++
        left := solve(node.Left, curLen, node.Val, maxLen)
        right := solve(node.Right, curLen, node.Val, maxLen)
        *maxLen = max(*maxLen, curLen + left)
        *maxLen = max(*maxLen, curLen + right)
        *maxLen = max(*maxLen, left + right)
        return 1 + max(left, right)
    }
    curLen = 0
    left := solve(node.Left, curLen, node.Val, maxLen)
    right := solve(node.Right, curLen, node.Val, maxLen)
    *maxLen = max(*maxLen, left + right)
    return 0
}

func longestUnivaluePath(root *TreeNode) int {
    var ans int
    solve(root, 0, -1001, &ans)
    return ans
}

// Time complexity: O(N)
// Space complexity: O(H) where H is the height of the tree