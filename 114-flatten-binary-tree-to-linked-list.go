// Link: https://leetcode.com/problems/flatten-binary-tree-to-linked-list/

func flatten(root *TreeNode) {
    if root == nil {
        return
    }
    flatten(root.Left)
    flatten(root.Right)
    left := root.Left
    right := root.Right
    // find tmp which is the last node in the left branch
    tmp := left
    if left != nil {
        for tmp.Right != nil {
            tmp = tmp.Right
        }
        tmp.Right = right
        root.Right = left
    }
    root.Left = nil
}

// Time complexity: O(N)
// Space complexity: O(N)