// Link: https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/

func findLCA(root *TreeNode, x, y int, c1, c2 *bool) *TreeNode {
    if root == nil {
        return nil
    }
    if root.Val == x {
        *c1 = true
        return root
    }
    if root.Val == y {
        *c2 = true
        return root
    }
    left := findLCA(root.Left, x, y, c1, c2)
    right := findLCA(root.Right, x, y, c1, c2)
    
    if left != nil && right != nil {
        return root
    }
    if left != nil {
        return left
    }
    return right
}

func find(root *TreeNode, x int) bool {
    if root == nil {
        return false
    }
    if root.Val == x {
        return true
    }
    return find(root.Left, x) || find(root.Right, x)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    var c1, c2 bool
    lca := findLCA(root, p.Val, q.Val, &c1, &c2)
    if c1 && c2 || c1 && find(root, q.Val) || c2 && find(root, p.Val) {
        return lca
    }
    return &TreeNode{}
}

// Time complexity: O(N)
// Space complexity: O(1)