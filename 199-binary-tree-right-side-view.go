// Link: https://leetcode.com/problems/binary-tree-right-side-view/

func rightSideView(root *TreeNode) []int {
    ans := make([]int, 0)
    if root == nil {
        return ans
    }
    queue := make([]*TreeNode, 0)
    queue = append(queue, root)
    for len(queue) > 0 {
        size := len(queue)
        var front *TreeNode
        for size > 0 {
            front = queue[0]
            queue = queue[1:]
            size--
            if front.Left != nil {
                queue = append(queue, front.Left)
            }
            if front.Right != nil {
                queue = append(queue, front.Right)
            }
        }
        ans = append(ans, front.Val)
    }
    return ans
}

// Time complexity: O(N)
// Space complexity: O(N)