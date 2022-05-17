/*
 * @lc app=leetcode.cn id=sort-list lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [148] 排序链表
 *
 * https://leetcode-cn.com/problems/sort-list/description/
 *
 * algorithms
 * Medium (66.53%)
 * Total Accepted:    276.3K
 * Total Submissions: 415.3K
 * Testcase Example:  '[4,2,1,3]'
 *
 * 给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [4,2,1,3]
 * 输出：[1,2,3,4]
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = [-1,5,3,4,0]
 * 输出：[-1,0,3,4,5]
 *
 *
 * 示例 3：
 *
 *
 * 输入：head = []
 * 输出：[]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表中节点的数目在范围 [0, 5 * 10^4] 内
 * -10^5 <= Node.val <= 10^5
 *
 *
 *
 *
 * 进阶：你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// mergeSort
//
// 5.16 后记: 你觉得为什么 `sortList` 要返回 `*ListNode` 结构呢? 如果不返回怎么样? 需要怎么处理?
// 提示: Golang 中 `head` 作为参数传入的是 **指针值**.
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p, q := head, head.Next
	for q != nil && q.Next != nil {
		p, q = p.Next, q.Next.Next
	}
	next := p.Next
	p.Next = nil
	head1, head2 := sortList(head), sortList(next)
	dummy := &ListNode{Val: -1}
	curr := dummy
	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			curr.Next = head1
			head1 = head1.Next
		} else {
			curr.Next = head2
			head2 = head2.Next
		}
		curr = curr.Next
	}
	if head1 != nil {
		curr.Next = head1
	} else {
		curr.Next = head2
	}
	return dummy.Next
}
