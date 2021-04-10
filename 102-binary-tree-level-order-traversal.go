// Link: https://leetcode.com/problems/binary-tree-level-order-traversal/

func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    queue := make([]*TreeNode, 0)
    queue = append(queue, root)
    ans := make([][]int, 0)
    for len(queue) > 0 {
        size := len(queue)
        tmp := make([]int, 0)
        for size > 0 {
            node := queue[0]
            tmp = append(tmp, node.Val)
            queue = queue[1:]
            size--
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        ans = append(ans, tmp)
    }
    return ans
}

// Time complexity: O(N)
// Space complexity: O(N)