// Link: https://leetcode.com/problems/swap-nodes-in-pairs/
// Given a linked list, swap every two adjacent nodes and return its head.

func swapPairs(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    prev := &ListNode{}
    prev.Next = head
    ans := prev
    cur := head
    for cur != nil && cur.Next != nil {
        tmp := cur.Next
        prev.Next = cur.Next
        cur.Next = tmp.Next
        tmp.Next = cur
        prev = cur
        cur = cur.Next
    }
    return ans.Next
}

// Time complexity: O(N)
// Space complexity: O(1)