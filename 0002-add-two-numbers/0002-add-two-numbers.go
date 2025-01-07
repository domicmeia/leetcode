/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    head := &ListNode{}
    current := head
    carry := 0

    for l1 != nil || l2 != nil || carry != 0 {
        n1, n2 := 0, 0

        if l1 != nil {
            n1 = l1.Val
            l1 = l1.Next
        }

        if l2 != nil {
            n2 = l2.Val
            l2 = l2.Next
        }

        sum := n1 + n2 + carry
        current.Next = &ListNode{Val: sum % 10}
        carry = sum / 10
        current = current.Next
    }

    return head.Next
}