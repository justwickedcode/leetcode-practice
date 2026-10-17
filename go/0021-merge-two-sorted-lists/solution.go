// 21. Merge Two Sorted Lists (Easy)
// https://leetcode.com/problems/merge-two-sorted-lists/
// Runtime: 0 ms (beats 100.0%) · Memory: 4.2 MB (beats 100.0%)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := &ListNode{}
    head := dummy

    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            head.Next = list1
            list1 = list1.Next
        } else {
            head.Next = list2
            list2 = list2.Next
        }

        head = head.Next
    }

    if list1 != nil {
        head.Next = list1
    } else {
        head.Next = list2
    }

    return dummy.Next
}
