// Link: 

func flatten(root *TreeNode) {
    if root == nil {
        return
    }
    // Find last node in the left
    left := root.Left
    for left != nil {
        if left.Right == nil && left.Left == nil {
            break
        }
        if left.Right != nil {
            left = left.Right
        } else if left.Left != nil {
            left = left.Left
        }
    }
    flatten(root.Left)
	// Concatenate right list to left list
    if left != nil && root.Right != nil {
        left.Right = root.Right
    }
    flatten(root.Right)
    if root.Left != nil {
        root.Right = root.Left
        root.Left = nil
    }
}

// Time complexity: O(N)
// Space complexity: O(1)