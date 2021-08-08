// Link: https://leetcode.com/problems/find-duplicate-subtrees/

import (
    "fmt"
    "strconv"
)

func solve(root *TreeNode, hmap map[string]int, ans *[]*TreeNode) string {
    if root == nil {
        return "n"
    }
    left := solve(root.Left, hmap, ans)
    right := solve(root.Right, hmap, ans)
    res := left + " " + right + " " + strconv.Itoa(root.Val)
    if _, ok := hmap[res]; ok {
        if hmap[res] == 1 {
            *ans = append(*ans, root)
        }
        hmap[res]++
    } else {
        hmap[res] = 1
    }
    return res
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
    freq := make(map[string]int)
    ans := make([]*TreeNode, 0)
    solve(root, freq, &ans)
    return ans
}

// Time complexity: O(N)
// Space complexity: O(N)