// Link: https://leetcode.com/problems/add-one-row-to-tree/submissions/

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
    if depth == 1 {
        node := &TreeNode{val, root, nil}
        return node
    }
    d := 0
    queue := make([]*TreeNode, 0)
    queue = append(queue, root)
    for len(queue) > 0 {
        size := len(queue)
        d++
        for size > 0 {
            front := queue[0]
            queue = queue[1:]
            if d == depth - 1 {
                front.Left = &TreeNode{val, front.Left, nil}
                front.Right = &TreeNode{val, nil, front.Right}
            }
            if front.Left != nil {
                queue = append(queue, front.Left)
            }
            if front.Right != nil {
                queue = append(queue, front.Right)
            }
            size--
        }
        if d == depth - 1 {
            return root
        }
    }
    return root
}

