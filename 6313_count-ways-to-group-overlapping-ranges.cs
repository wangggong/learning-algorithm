/*
 * @lc app=leetcode.cn id=count-ways-to-group-overlapping-ranges lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6313] 统计将重叠区间合并成组的方案数
 *
 * https://leetcode.cn/problems/count-ways-to-group-overlapping-ranges/description/
 *
 * algorithms
 * Medium (31.93%)
 * Total Accepted:    2.9K
 * Total Submissions: 9K
 * Testcase Example:  '[[6,10],[5,15]]'
 *
 * 给你一个二维整数数组 ranges ，其中 ranges[i] = [starti, endi] 表示 starti 到 endi
 * 之间（包括二者）的所有整数都包含在第 i 个区间中。
 * 
 * 你需要将 ranges 分成 两个 组（可以为空），满足：
 * 
 * 
 * 每个区间只属于一个组。
 * 两个有 交集 的区间必须在 同一个 组内。
 * 
 * 
 * 如果两个区间有至少 一个 公共整数，那么这两个区间是 有交集 的。
 * 
 * 
 * 比方说，区间 [1, 3] 和 [2, 5] 有交集，因为 2 和 3 在两个区间中都被包含。
 * 
 * 
 * 请你返回将 ranges 划分成两个组的 总方案数 。由于答案可能很大，将它对 10^9 + 7 取余 后返回。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：ranges = [[6,10],[5,15]]
 * 输出：2
 * 解释：
 * 两个区间有交集，所以它们必须在同一个组内。
 * 所以有两种方案：
 * - 将两个区间都放在第 1 个组中。
 * - 将两个区间都放在第 2 个组中。
 * 
 * 
 * 示例 2：
 * 
 * 输入：ranges = [[1,3],[10,20],[2,5],[4,8]]
 * 输出：4
 * 解释：
 * 区间 [1,3] 和 [2,5] 有交集，所以它们必须在同一个组中。
 * 同理，区间 [2,5] 和 [4,8] 也有交集，所以它们也必须在同一个组中。
 * 所以总共有 4 种分组方案：
 * - 所有区间都在第 1 组。
 * - 所有区间都在第 2 组。
 * - 区间 [1,3] ，[2,5] 和 [4,8] 在第 1 个组中，[10,20] 在第 2 个组中。
 * - 区间 [1,3] ，[2,5] 和 [4,8] 在第 2 个组中，[10,20] 在第 1 个组中。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= ranges.length <= 10^5
 * ranges[i].length == 2
 * 0 <= starti <= endi <= 10^9
 * 
 * 
 */

// 并查集
/*
 * public class Solution
 * {
 *     private const long Mod = (long)1e9 + 7;
 * 
 *     public int CountWays(int[][] ranges)
 *     {
 *         var n = ranges.Length;
 *         Array.Sort(ranges, (x, y) => x[0].CompareTo(y[0]));
 *         var pa = new int[n];
 *         for (var i = 0; i < n; i++)
 *         {
 *             pa[i] = i;
 *         }
 *         int query(int k) => k == pa[k] ? pa[k] : (pa[k] = query(pa[k]));
 *         void merge(int p, int q) => pa[query(p)] = query(q);
 *         var cur = ranges[0][1];
 *         for (int i = 0; i + 1 < n; i++)
 *         {
 *             if (cur >= ranges[i + 1][0])
 *             {
 *                 merge(i, i + 1);
 *             }
 *             cur = Math.Max(cur, ranges[i + 1][1]);
 *         }
 *         var S = new HashSet<int>();
 *         for (var i = 0; i < n; i++)
 *         {
 *             S.Add(query(i));
 *         }
 *         long ans = 1;
 *         for (var i = 0; i < S.Count(); i++)
 *         {
 *             ans = (ans << 1) % Mod;
 *         }
 *         return (int)ans;
 *     }
 * }
 */

public class Solution
{
    public int CountWays(int[][] ranges)
    {
        const long Mod = (long)1e9 + 7;
        var (count, right) = (0, -1);
        foreach (var range in ranges.OrderBy(x => x[0]))
        {
            var (l, r) = (range[0], range[1]);
            if (l > right)
            {
                count++;
            }
            right = Math.Max(right, r);
        }
        var ans = 1l;
        for (var k = 2l; count > 0; count >>= 1)
        {
            if ((count & 1) != 0)
            {
                ans = ans * k % Mod;
            }
            k = k * k % Mod;
        }
        return (int)ans;
    }
}
