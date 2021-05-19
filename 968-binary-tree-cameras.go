// Link: https://leetcode.com/problems/binary-tree-cameras/

func solve(root *TreeNode, parent *TreeNode, cam *int, check map[*TreeNode]bool) {
    if root == nil {
        return
    }
    if _, ok := check[root]; ok {
        return
    }
    solve(root.Left, root, cam, check)
    solve(root.Right, root, cam, check)
    // leaf node
    if root.Left == nil && root.Right == nil {
        if parent != nil {
            return
        }
        (*cam)++
        check[root] = true
        return
    }
    _, checkLeft := check[root.Left]
    _, checkRight := check[root.Right]
    _, checkCurr := check[root]
    if !checkLeft || !checkRight || (parent == nil && !checkCurr) {
        (*cam)++
        check[root.Left] = true
        check[root.Right] = true
        check[parent] = true
        check[root] = true
    }
}

func minCameraCover(root *TreeNode) int {
    if root == nil {
        return 0
    }
    check := make(map[*TreeNode]bool)
    check[nil] = true
    var cam int
    solve(root, nil, &cam, check)
    return cam
}

// Time complexity: O(N)
// Space complexity: O(N)