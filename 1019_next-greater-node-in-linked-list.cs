/*
 * @lc app=leetcode.cn id=next-greater-node-in-linked-list lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1019] 链表中的下一个更大节点
 *
 * https://leetcode.cn/problems/next-greater-node-in-linked-list/description/
 *
 * algorithms
 * Medium (60.99%)
 * Total Accepted:    30.7K
 * Total Submissions: 50.2K
 * Testcase Example:  '[2,1,5]'
 *
 * 给定一个长度为 n 的链表 head
 * 
 * 对于列表中的每个节点，查找下一个 更大节点 的值。也就是说，对于每个节点，找到它旁边的第一个节点的值，这个节点的值 严格大于 它的值。
 * 
 * 返回一个整数数组 answer ，其中 answer[i] 是第 i 个节点( 从1开始 )的下一个更大的节点的值。如果第 i
 * 个节点没有下一个更大的节点，设置 answer[i] = 0 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 
 * 输入：head = [2,1,5]
 * 输出：[5,5,0]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 
 * 
 * 输入：head = [2,7,4,3,5]
 * 输出：[7,0,5,5,0]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 链表中节点数为 n
 * 1 <= n <= 10^4
 * 1 <= Node.val <= 10^9
 * 
 * 
 */
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     public int val;
 *     public ListNode next;
 *     public ListNode(int val=0, ListNode next=null) {
 *         this.val = val;
 *         this.next = next;
 *     }
 * }
 */
public class Solution
{
    public int[] NextLargerNodes(ListNode head)
	{
		var ans = new List<int>();
		var S = new Stack<(int, int)>();
		for (var (i, node) = (0, head); node != null; (i, node) = (i + 1, node.next))
		{
			ans.Add(0);
			while (S.Count > 0 && node.val > S.Peek().Item2)
			{
				var (j, _) = S.Pop();
				ans[j] = node.val;
			}
			S.Push((i, node.val));
		}
		return ans.ToArray();
    }
}
