/*
 * @lc app=leetcode.cn id=distant-barcodes lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1054] 距离相等的条形码
 *
 * https://leetcode.cn/problems/distant-barcodes/description/
 *
 * algorithms
 * Medium (39.88%)
 * Total Accepted:    14.1K
 * Total Submissions: 33.8K
 * Testcase Example:  '[1,1,1,2,2,2]'
 *
 * 在一个仓库里，有一排条形码，其中第 i 个条形码为 barcodes[i]。
 * 
 * 请你重新排列这些条形码，使其中任意两个相邻的条形码不能相等。 你可以返回任何满足该要求的答案，此题保证存在答案。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：barcodes = [1,1,1,2,2,2]
 * 输出：[2,1,2,1,2,1]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：barcodes = [1,1,1,1,2,2,3,3]
 * 输出：[1,3,1,3,2,1,2,1]
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= barcodes.length <= 10000
 * 1 <= barcodes[i] <= 10000
 * 
 * 
 */


/*
 * // 大根堆
 * public class Solution
 * {
 *     public int[] RearrangeBarcodes(int[] barcodes)
 *     {
 *         var Q = new PriorityQueue<int, int>();
 *         foreach (var g in barcodes.GroupBy(x => x))
 *         {
 *             Q.Enqueue(g.Key, -g.Count());
 *         }
 *         var ans = new List<int>();
 *         while (Q.Count >= 2)
 *         {
 *             Q.TryDequeue(out var p, out var c1);
 *             Q.TryDequeue(out var q, out var c2);
 *             ans.Add(p);
 *             ans.Add(q);
 *             if (c1 + 1 < 0)
 *             {
 *                 Q.Enqueue(p, c1 + 1);
 *             }
 *             if (c2 + 1 < 0)
 *             {
 *                 Q.Enqueue(q, c2 + 1);
 *             }
 *         }
 *         if (Q.Count > 0)
 *         {
 *             var v = Q.Dequeue();
 *             ans.Add(v);
 *         }
 *         return ans.ToArray();
 *     }
 * }
 */

// 参考: https://leetcode.cn/problems/distant-barcodes/solution/python3javacgotypescript-yi-ti-yi-jie-ji-3or2/
// 
// 计数排序, 更像是把排好序的数组 shuffle 了一下.
public class Solution
{
    public int[] RearrangeBarcodes(int[] barcodes)
    {
        var (n, k) = (barcodes.Length, 0);
        var ans = new int[n];
        foreach (var g in barcodes.GroupBy(x => x).OrderBy(g => -g.Count()))
        {
            for (var (i, c) = (0, g.Count()); i < c; i++)
            {
                ans[k] = g.Key;
                k = k + 2 >= n ? 1 : k + 2;
            }
        }
        return ans;
    }
}
