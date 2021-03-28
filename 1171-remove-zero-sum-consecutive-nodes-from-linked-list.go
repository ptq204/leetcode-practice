// Link: https://leetcode.com/problems/remove-zero-sum-consecutive-nodes-from-linked-list/

func removeZeroSumSublists(head *ListNode) *ListNode {
    hmap := make(map[int]*ListNode)
    hmap[0] = &ListNode{Next: head}
    cur := head
    sum := 0
    for cur != nil {
        sum += cur.Val
        if _, ok := hmap[sum]; ok {
            // clear nodes which sum to 0 in map
            tmpSum := sum
            tmpCur := hmap[sum].Next
            for tmpCur != cur {
                tmpSum += tmpCur.Val
                if tmpSum != 0 {
                    delete(hmap, tmpSum)
                }
                tmpCur = tmpCur.Next
            }
            hmap[sum].Next = cur.Next
        } else {
            hmap[sum] = cur
        }
        cur = cur.Next
    }
    return hmap[0].Next
}

// Time complexity: O(N)
// Space complexity: O(N)