// Link: https://leetcode.com/problems/serialize-and-deserialize-binary-tree/

type Codec struct {
    
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    if root == nil {
        return ""
    }
    stk := make([]*TreeNode, 0)
    stk = append(stk, root)
    res := ""
    for len(stk) > 0 {
        front := stk[len(stk)-1]
        stk = stk[:len(stk)-1]
        if front == nil {
            res += "n|"
            continue
        }
        res += (strconv.Itoa(front.Val) + "|")
        stk = append(stk, front.Right)
        stk = append(stk, front.Left)
    }
    return res[:len(res)-1]
}

// Deserializes your encoded data to tree.
func constructTree(i *int, n int, tokens []string) *TreeNode {
    if *i >= n || tokens[*i] == "n" {
        *i++
        return nil
    }
    val, _ := strconv.Atoi(tokens[*i])
    *i++
    return &TreeNode{val, constructTree(i, n, tokens), constructTree(i, n, tokens)}
}

func (this *Codec) deserialize(data string) *TreeNode {
    if len(data) == 0 {
        return nil
    }
    tokens := strings.Split(data, "|")
    i := 0
    return constructTree(&i, len(tokens), tokens)
}

// Time complexity: O(N)
// Space complexity: O(N)