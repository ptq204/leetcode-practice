// Link: https://leetcode.com/problems/trim-a-binary-search-tree/

func trimBinarySearchTree(root *TreeNode, low, high int) *TreeNode {
    if root == nil {
        return nil
    }
    for root != nil && (root.Val < low || root.Val > high) {
        if root.Val < low {
            root = root.Right
        } else if root.Val > high {
            root = root.Left
        }
    }
    if root != nil {
        root.Left = trimBinarySearchTree(root.Left, low, high)
        root.Right = trimBinarySearchTree(root.Right, low, high)
    }
    return root
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
    return trimBinarySearchTree(root, low, high)
}

// Time complexity: O(N)
// Space complexity: O(1)