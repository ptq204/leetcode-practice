// Link: https://leetcode.com/problems/remove-nth-node-from-end-of-list/
// Given the head of a linked list, remove the nth node from the end of the list and return its head.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head == nil || head.Next == nil {
        return nil
    }
    head1 := head
    head2 := head
    prev := head
    cnt := 1
    for head2 != nil && head2.Next != nil {
        if cnt >= n {
            prev = head1
            head1 = head1.Next
        }
        head2 = head2.Next
        cnt++
    }
    if prev == head1 {
        head = prev.Next
    } else {
        prev.Next = head1.Next
    }
    return head
}