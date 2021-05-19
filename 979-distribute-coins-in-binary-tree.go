// Link: https://leetcode.com/problems/distribute-coins-in-binary-tree/

func abs(x int) int {
    if x >= 0 {
        return x
    }
    return -x
}

func solve(root *TreeNode, ans *int) int {
    if root == nil {
        return 0
    }
    left := solve(root.Left, ans)
    right := solve(root.Right, ans)
    diff := root.Val - 1
    need := diff + left + right
    *ans += abs(need)
    return need
}

func distributeCoins(root *TreeNode) int {
    ans := 0
    solve(root, &ans)
    return ans
}

// Time complexity: O(N)
// Space complexity: O(n)