// Link: https://leetcode.com/problems/palindrome-linked-list/

func reverse(head *ListNode) *ListNode {
    var prev *ListNode
    curr := head
    for curr != nil {
        tmp := curr.Next
        curr.Next = prev
        prev = curr
        curr = tmp
    }
    return prev
}

func isPalindrome(head *ListNode) bool {
    slow := head
    fast := head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    mid := slow
    if fast != nil {
        slow = slow.Next
    }
    head2 := reverse(slow)
    head1 := head
    for head1 != mid && head1 != nil {
        if head1.Val != head2.Val {
            return false
        }
        head1 = head1.Next
        head2 = head2.Next
    }
    return true
}

// Time complexity: O(N)
// Space complexity: O(1)