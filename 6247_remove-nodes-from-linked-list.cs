/*
 * @lc app=leetcode.cn id=remove-nodes-from-linked-list lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6247] 从链表中移除节点
 *
 * https://leetcode.cn/problems/remove-nodes-from-linked-list/description/
 *
 * algorithms
 * Medium (65.43%)
 * Total Accepted:    3.6K
 * Total Submissions: 5.5K
 * Testcase Example:  '[5,2,13,3,8]'
 *
 * 给你一个链表的头节点 head 。
 * 
 * 对于列表中的每个节点 node ，如果其右侧存在一个具有 严格更大 值的节点，则移除 node 。
 * 
 * 返回修改后链表的头节点 head 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 
 * 输入：head = [5,2,13,3,8]
 * 输出：[13,8]
 * 解释：需要移除的节点是 5 ，2 和 3 。
 * - 节点 13 在节点 5 右侧。
 * - 节点 13 在节点 2 右侧。
 * - 节点 8 在节点 3 右侧。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：head = [1,1,1,1]
 * 输出：[1,1,1,1]
 * 解释：每个节点的值都是 1 ，所以没有需要移除的节点。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 给定列表中的节点数目在范围 [1, 10^5] 内
 * 1 <= Node.val <= 10^5
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
    public ListNode RemoveNodes(ListNode head)
    {
        Stack<ListNode> S = new();
        for (; head != null; head = head.next)
        {
            while (S.Count > 0 && S.Peek().val < head.val) { S.Pop(); }
            S.Push(head);
        }
        var dummy = new ListNode();
        head = dummy;
        for (; S.Count > 0; S.Pop())
        {
            head.next = S.Peek();
            head = head.next;
            head.next = null;
        }
        ListNode prev = null;
        var curr = dummy.next;
        while (curr != null)
        {
            var next = curr.next;
            curr.next = prev;
            prev = curr;
            curr = next;
        }
        return prev;
    }
}
