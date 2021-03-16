// Link: https://leetcode.com/problems/merge-k-sorted-lists/
// You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.
// Merge all the linked-lists into one sorted linked-list and return it.

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    head := &ListNode{}
    head1 := head
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            head.Next = l1
            l1 = l1.Next
        } else {
            head.Next = l2
            l2 = l2.Next
        }
        head = head.Next
    }
    if l1 != nil {
        head.Next = l1
    }
    if l2 != nil {
        head.Next = l2
    }
    return head1.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
    k := len(lists)
    if k == 0 {
        return nil
    }
    ans := lists[0]
    for i:=1; i<k; i++ {
        ans = mergeTwoLists(ans, lists[i])
    }
    return ans
}

// Time complexity: O(K*N) where K is number of linked lists, N is the length of merged list
// Space complexity: O(1)