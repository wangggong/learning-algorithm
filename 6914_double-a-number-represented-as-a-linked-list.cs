/*
 * @lc app=leetcode.cn id=double-a-number-represented-as-a-linked-list lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6914] 翻倍以链表形式表示的数字
 *
 * https://leetcode.cn/problems/double-a-number-represented-as-a-linked-list/description/
 *
 * algorithms
 * Medium (53.15%)
 * Total Accepted:    4.2K
 * Total Submissions: 7.9K
 * Testcase Example:  '[1,8,9]'
 *
 * 给你一个 非空 链表的头节点 head ，表示一个不含前导零的非负数整数。
 * 
 * 将链表 翻倍 后，返回头节点 head 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：head = [1,8,9]
 * 输出：[3,7,8]
 * 解释：上图中给出的链表，表示数字 189 。返回的链表表示数字 189 * 2 = 378 。
 * 
 * 示例 2：
 * 
 * 
 * 输入：head = [9,9,9]
 * 输出：[1,9,9,8]
 * 解释：上图中给出的链表，表示数字 999 。返回的链表表示数字 999 * 2 = 1998 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 链表中节点的数目在范围 [1, 10^4] 内
 * 0 <= Node.val <= 9
 * 生成的输入满足：链表表示一个不含前导零的数字，除了数字 0 本身。
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
    public ListNode DoubleIt(ListNode head)
    {
        ListNode reverse(ListNode head)
        {
            var (prev, curr) = (null as ListNode, head);
            while (curr is not null)
            {
                var next = curr.next;
                curr.next = prev;
                (prev, curr) = (curr, next);
            }
            return prev;
        }
        head = reverse(head);
        var dummy = new ListNode(-1);
        var curr = dummy;
        void setNext(ListNode node) => (curr.next, curr) = (node, node);
        var C = 0;
        for (; head is not null; head = head.next)
        {
            var v = head.val * 2 + C;
            C = v / 10;
            setNext(new ListNode(v % 10));
        }
        if (C > 0) { setNext(new ListNode(C)); }
        return reverse(dummy.next);
    }
}
