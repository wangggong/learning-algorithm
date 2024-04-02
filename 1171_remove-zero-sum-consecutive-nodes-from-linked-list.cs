/*
 * @lc app=leetcode.cn id=remove-zero-sum-consecutive-nodes-from-linked-list lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1171] 从链表中删去总和值为零的连续节点
 *
 * https://leetcode.cn/problems/remove-zero-sum-consecutive-nodes-from-linked-list/description/
 *
 * algorithms
 * Medium (48.45%)
 * Total Accepted:    22.9K
 * Total Submissions: 46K
 * Testcase Example:  '[1,2,-3,3,1]'
 *
 * 给你一个链表的头节点 head，请你编写代码，反复删去链表中由 总和 值为 0 的连续节点组成的序列，直到不存在这样的序列为止。
 * 
 * 删除完毕后，请你返回最终结果链表的头节点。
 * 
 * 
 * 
 * 你可以返回任何满足题目要求的答案。
 * 
 * （注意，下面示例中的所有序列，都是对 ListNode 对象序列化的表示。）
 * 
 * 示例 1：
 * 
 * 输入：head = [1,2,-3,3,1]
 * 输出：[3,1]
 * 提示：答案 [1,2,1] 也是正确的。
 * 
 * 
 * 示例 2：
 * 
 * 输入：head = [1,2,3,-3,4]
 * 输出：[1,2,4]
 * 
 * 
 * 示例 3：
 * 
 * 输入：head = [1,2,3,-3,-2]
 * 输出：[1]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 给你的链表中可能有 1 到 1000 个节点。
 * 对于链表中的每个节点，节点的值：-1000 <= node.val <= 1000.
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

/*
 * public class Solution
 * {
 *     public ListNode RemoveZeroSumSublists(ListNode head)
 *     {
 *         var dummy = new ListNode(0, head);
 *         var S = new Stack<ListNode>();
 *         var visit = new HashSet<int>();
 *         S.Push(dummy);
 *         visit.Add(0);
 *         for (var (sum, cur) = (0, head); cur is not null; cur = cur.next)
 *         {
 *             var next = cur.next;
 *             sum += cur.val;
 *             for (var v = sum; visit.Contains(v); )
 *             {
 *                 sum -= cur.val;
 *                 cur = S.Pop();
 *                 visit.Remove(sum);
 *             }
 *             visit.Add(sum);
 *             S.Push(cur);
 *             S.Peek().next = next;
 *         }
 *         return dummy.next;
 *     }
 * }
 */

// 参考: https://leetcode.cn/problems/remove-zero-sum-consecutive-nodes-from-linked-list/solution/cong-lian-biao-zhong-shan-qu-zong-he-zhi-h18o/
//
// 官解的思路更直观一些. 也好调试.
public class Solution
{
    public ListNode RemoveZeroSumSublists(ListNode head)
    {
        var dummy = new ListNode(0, head);
        var S = 0;
        var d = new Dictionary<int, ListNode>();
        for (var cur = dummy; cur is not null; cur = cur.next)
        {
            S += cur.val;
            d[S] = cur;
        }
        S = 0;
        for (var cur = dummy; cur is not null; cur = cur.next)
        {
            cur.next = d[S].next;
            S += cur.next?.val ?? 0;
        }
        return dummy.next;
    }
}
