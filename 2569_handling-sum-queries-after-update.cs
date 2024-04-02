/*
 * @lc app=leetcode.cn id=handling-sum-queries-after-update lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [2569] 更新数组后处理求和查询
 *
 * https://leetcode.cn/problems/handling-sum-queries-after-update/description/
 *
 * algorithms
 * Hard (39.19%)
 * Total Accepted:    2.6K
 * Total Submissions: 6.7K
 * Testcase Example:  '[1,0,1]\n[0,0,0]\n[[1,1,1],[2,1,0],[3,0,0]]'
 *
 * 给你两个下标从 0 开始的数组 nums1 和 nums2 ，和一个二维数组 queries 表示一些操作。总共有 3 种类型的操作：
 * 
 * 
 * 操作类型 1 为 queries[i] = [1, l, r] 。你需要将 nums1 从下标 l 到下标 r 的所有 0 反转成 1 或将 1 反转成
 * 0 。l 和 r 下标都从 0 开始。
 * 操作类型 2 为 queries[i] = [2, p, 0] 。对于 0 <= i < n 中的所有下标，令 nums2[i] = nums2[i]
 * + nums1[i] * p 。
 * 操作类型 3 为 queries[i] = [3, 0, 0] 。求 nums2 中所有元素的和。
 * 
 * 
 * 请你返回一个数组，包含所有第三种操作类型的答案。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums1 = [1,0,1], nums2 = [0,0,0], queries = [[1,1,1],[2,1,0],[3,0,0]]
 * 输出：[3]
 * 解释：第一个操作后 nums1 变为 [1,1,1] 。第二个操作后，nums2 变成 [1,1,1] ，所以第三个操作的答案为 3 。所以返回 [3]
 * 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums1 = [1], nums2 = [5], queries = [[2,0,0],[3,0,0]]
 * 输出：[5]
 * 解释：第一个操作后，nums2 保持不变为 [5] ，所以第二个操作的答案是 5 。所以返回 [5] 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums1.length,nums2.length <= 10^5
 * nums1.length = nums2.length
 * 1 <= queries.length <= 10^5
 * queries[i].length = 3
 * 0 <= l <= r <= nums1.length - 1
 * 0 <= p <= 10^6
 * 0 <= nums1[i] <= 1
 * 0 <= nums2[i] <= 10^9
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/handling-sum-queries-after-update/solution/xian-duan-shu-by-endlesscheng-vx80/
//
// 线段树, lazy, 左闭右开
public class Solution
{
    public long[] HandleQuery(int[] nums1, int[] nums2, int[][] queries) 
    {
        var n = nums1.Length;
        var counts = new long[n << 2];
        var flips = new bool[n << 2];
        
        void maintain(int o) => counts[o] = counts[o << 1] + counts[o << 1 | 1];
        void tag(int o, int l, int r) => (counts[o], flips[o]) = ((long)(r - l) - counts[o], !flips[o]);
        void update(int o, int l, int r, int L, int R)
        {
            if (L <= l && r <= R)
            {
                tag(o, l, r);
                return;
            }
            var m = (l + r) >> 1;
            if (flips[o])
            {
                tag(o << 1, l, m);
                tag(o << 1 | 1, m, r);
                flips[o] = false;
            }
            if (L < m)
            {
                update(o << 1, l, m, L, R);
            }
            if (m < R)
            {
                update(o << 1 | 1, m, r, L, R);
            }
            maintain(o);
        }
        
        for (var i = 0; i < n; i++)
        {
            if (nums1[i] is 1)
            {
                update(1, 0, n, i, i + 1);
            }
        }
        var sum = nums2.Select(v => (long)v).Sum();
        var ans = new List<long>();
        foreach (var q in queries)
        {
            switch (q[0])
            {
                case 1:
                    update(1, 0, n, q[1], q[2] + 1);
                    break;
                case 2:
                    sum += (long)q[1] * counts[1];
                    break;
                default:
                    ans.Add(sum);
                    break;
            }
        }
        return ans.ToArray();
    }
}
