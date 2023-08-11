给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
示例 1：
输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
示例 2：

输入：head = [1], n = 1
输出：[]
示例 3：

输入：head = [1,2], n = 1
输出：[1]
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{Val: 0, Next: head} 
    first := dummy
    second := dummy


    for i := 1; i <= n+1; i++ {
        first = first.Next
    }


    for first != nil {
        first = first.Next
        second = second.Next
    }

    second.Next = second.Next.Next

    return dummy.Next
}
思路：双指针，找到倒数第n个前一个节点，修改指针方向
