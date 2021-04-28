// Link: https://leetcode.com/problems/sum-of-left-leaves/

func sumOfLeftLeaves(root *TreeNode) int {
    if root == nil {
        return 0
    }
    stk := make([]*TreeNode, 0)
    stk = append(stk, root)
    ans := 0
    for len(stk) > 0 {
        top := stk[len(stk)-1]
        stk = stk[:len(stk)-1]
        if top.Left != nil {
            stk = append(stk, top.Left)
            if top.Left.Left == nil && top.Left.Right == nil {
                ans += top.Left.Val
            }
        }
        if top.Right != nil {
            stk = append(stk, top.Right)
        }
    }
    return ans
}