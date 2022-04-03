/*
 * @lc app=leetcode.cn id=linked-list-cycle-ii lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [142] 环形链表 II
 *
 * https://leetcode-cn.com/problems/linked-list-cycle-ii/description/
 *
 * algorithms
 * Medium (55.57%)
 * Total Accepted:    349.8K
 * Total Submissions: 629.3K
 * Testcase Example:  '[3,2,0,-4]\n1'
 *
 * 给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
 *
 * 如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos
 * 来表示链表尾连接到链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环。注意：pos
 * 不作为参数进行传递，仅仅是为了标识链表的实际情况。
 *
 * 不允许修改 链表。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：head = [3,2,0,-4], pos = 1
 * 输出：返回索引为 1 的链表节点
 * 解释：链表中有一个环，其尾部连接到第二个节点。
 *
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入：head = [1,2], pos = 0
 * 输出：返回索引为 0 的链表节点
 * 解释：链表中有一个环，其尾部连接到第一个节点。
 *
 *
 * 示例 3：
 *
 *
 *
 *
 * 输入：head = [1], pos = -1
 * 输出：返回 null
 * 解释：链表中没有环。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表中节点的数目范围在范围 [0, 10^4] 内
 * -10^5 <= Node.val <= 10^5
 * pos 的值为 -1 或者链表中的一个有效索引
 *
 *
 *
 *
 * 进阶：你是否可以使用 O(1) 空间解决此题？
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
 * // Hash:
 * func detectCycle(h *ListNode) *ListNode {
 * 	m := make(map[*ListNode]struct{})
 * 	for h != nil {
 * 		if _, ok := m[h]; ok {
 * 			return h
 * 		}
 *
 * 		m[h] = struct{}{}
 * 		h = h.Next
 * 	}
 *
 * 	return nil
 * }
 */

// 快慢指针: 快指针一次走两步, 慢指针一次走一步, 有环则相遇.
// 如果相遇, 可以推出慢指针到入环点距离与起点到入环点距离与环长同余.
func detectCycle(h *ListNode) *ListNode {
	p, q := h, h
	for q != nil {
		p = p.Next
		if q.Next == nil {
			return nil
		}
		q = q.Next.Next
		if p == q {
			r := h
			for p != r {
				p, r = p.Next, r.Next
			}
			return r
		}
	}
	return nil
}
