/*
 * @lc app=leetcode.cn id=add-two-numbers-ii lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [445] 两数相加 II
 *
 * https://leetcode.cn/problems/add-two-numbers-ii/description/
 *
 * algorithms
 * Medium (60.22%)
 * Total Accepted:    130.8K
 * Total Submissions: 215.2K
 * Testcase Example:  '[7,2,4,3]\n[5,6,4]'
 *
 * 给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。
 * 
 * 你可以假设除了数字 0 之外，这两个数字都不会以零开头。
 * 
 * 
 * 
 * 示例1：
 * 
 * 
 * 
 * 
 * 输入：l1 = [7,2,4,3], l2 = [5,6,4]
 * 输出：[7,8,0,7]
 * 
 * 
 * 示例2：
 * 
 * 
 * 输入：l1 = [2,4,3], l2 = [5,6,4]
 * 输出：[8,0,7]
 * 
 * 
 * 示例3：
 * 
 * 
 * 输入：l1 = [0], l2 = [0]
 * 输出：[0]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 链表的长度范围为 [1, 100]
 * 0 <= node.val <= 9
 * 输入数据保证链表代表的数字无前导 0
 * 
 * 
 * 
 * 
 * 进阶：如果输入链表不能翻转该如何解决？
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
    private ListNode Reverse(ListNode head)
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

    public ListNode AddTwoNumbers(ListNode l1, ListNode l2, bool reversed = false)
    {
        if (!reversed)
        {
            return Reverse(AddTwoNumbers(Reverse(l1), Reverse(l2), true));
        }
        var dummy = new ListNode();
        var curr = dummy;
        var C = 0;
        while (l1 is not null || l2 is not null)
        {
            var (p, q) = (l1?.val ?? 0, l2?.val ?? 0);
            curr.next = new ListNode((p + q + C) % 10);
            C = (p + q + C) / 10;
            (l1, l2, curr) = (l1?.next ?? null, l2?.next ?? null, curr.next);
        }
        if (C is not 0)
        {
            curr.next = new ListNode(C);
        }
        return dummy.next;
    }
}
