/*
 * @lc app=leetcode.cn id=linked-list-random-node lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [382] 链表随机节点
 *
 * https://leetcode-cn.com/problems/linked-list-random-node/description/
 *
 * algorithms
 * Medium (71.61%)
 * Total Accepted:    40.5K
 * Total Submissions: 56.5K
 * Testcase Example:  '["Solution","getRandom","getRandom","getRandom","getRandom","getRandom"]\n' +
  '[[[1,2,3]],[],[],[],[],[]]'
 *
 * 给你一个单链表，随机选择链表的一个节点，并返回相应的节点值。每个节点 被选中的概率一样 。
 *
 * 实现 Solution 类：
 *
 *
 * Solution(ListNode head) 使用整数数组初始化对象。
 * int getRandom() 从链表中随机选择一个节点并返回该节点的值。链表中所有节点被选中的概率相等。
 *
 *
 *
 *
 * 示例：
 *
 *
 * 输入
 * ["Solution", "getRandom", "getRandom", "getRandom", "getRandom",
 * "getRandom"]
 * [[[1, 2, 3]], [], [], [], [], []]
 * 输出
 * [null, 1, 3, 2, 2, 3]
 *
 * 解释
 * Solution solution = new Solution([1, 2, 3]);
 * solution.getRandom(); // 返回 1
 * solution.getRandom(); // 返回 3
 * solution.getRandom(); // 返回 2
 * solution.getRandom(); // 返回 2
 * solution.getRandom(); // 返回 3
 * // getRandom() 方法应随机返回 1、2、3中的一个，每个元素被返回的概率相等。
 *
 *
 *
 * 提示：
 *
 *
 * 链表中的节点数在范围 [1, 10^4] 内
 * -10^4 <= Node.val <= 10^4
 * 至多调用 getRandom 方法 10^4 次
 *
 *
 *
 *
 * 进阶：
 *
 *
 * 如果链表非常大且长度未知，该怎么处理？
 * 你能否在不使用额外空间的情况下解决此问题？
 *
 *
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
import "math/rand"

/*
 * // 蓄水池算法
 * type Solution struct {
 * 	ln *ListNode
 * }
 *
 * func Constructor(head *ListNode) Solution {
 * 	return Solution{ln: head}
 * }
 *
 * func (this *Solution) GetRandom() int {
 * 	ans := -1
 * 	for c, index := this.ln, 1; c != nil && index >= 0; index++ {
 * 		if rand.Intn(index) == 0 {
 * 			ans = c.Val
 * 		}
 * 		c = c.Next
 * 	}
 * 	return ans
 * }
 */

type Solution struct {
	nodes []int
	n     int
}

func Constructor(head *ListNode) Solution {
	s := Solution{}
	for head != nil {
		s.nodes = append(s.nodes, head.Val)
		head = head.Next
		s.n++
	}
	return s
}

func (s *Solution) GetRandom() int {
	return s.nodes[rand.Intn(s.n)]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
