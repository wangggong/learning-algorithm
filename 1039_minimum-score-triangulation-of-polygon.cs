/*
 * @lc app=leetcode.cn id=minimum-score-triangulation-of-polygon lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1039] 多边形三角剖分的最低得分
 *
 * https://leetcode.cn/problems/minimum-score-triangulation-of-polygon/description/
 *
 * algorithms
 * Medium (57.60%)
 * Total Accepted:    9K
 * Total Submissions: 15.1K
 * Testcase Example:  '[1,2,3]'
 *
 * 你有一个凸的 n 边形，其每个顶点都有一个整数值。给定一个整数数组 values ，其中 values[i] 是第 i 个顶点的值（即 顺时针顺序
 * ）。
 * 
 * 假设将多边形 剖分 为 n - 2 个三角形。对于每个三角形，该三角形的值是顶点标记的乘积，三角剖分的分数是进行三角剖分后所有 n - 2
 * 个三角形的值之和。
 * 
 * 返回 多边形进行三角剖分后可以得到的最低分 。
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
 * 输入：values = [1,2,3]
 * 输出：6
 * 解释：多边形已经三角化，唯一三角形的分数为 6。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 
 * 
 * 输入：values = [3,7,4,5]
 * 输出：144
 * 解释：有两种三角剖分，可能得分分别为：3*7*5 + 4*5*7 = 245，或 3*4*5 + 3*4*7 = 144。最低分数为 144。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 
 * 
 * 输入：values = [1,3,1,4,1,5]
 * 输出：13
 * 解释：最低分数三角剖分的得分情况为 1*1*3 + 1*1*4 + 1*1*5 + 1*1*1 = 13。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * n == values.length
 * 3 <= n <= 50
 * 1 <= values[i] <= 100
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/minimum-score-triangulation-of-polygon/solution/shi-pin-jiao-ni-yi-bu-bu-si-kao-dong-tai-aty6/
// 
// DP 都写不对了... 傻逼...
public class Solution
{
    public int MinScoreTriangulation(int[] values)
    {
        var memo = new Dictionary<(int, int), int>();
        int traversal(int p, int q)
        {
            var key = (p, q);
            if (!memo.ContainsKey(key))
            {
                if (p + 1 == q)
                {
                    memo[key] = 0;
                    return memo[key];
                }
                memo[key] = int.MaxValue;
                for (var i = p + 1; i < q; i++)
                {
                    memo[key] = Math.Min(memo[key],
                        traversal(p, i) + traversal(i, q) + values[p] * values[i] * values[q]);
                }
            }
            return memo[key];
        }
        return traversal(0, values.Length - 1);
    }
}

/*
 * // 超时版本. 自己想想为啥吧...
 * public class Solution
 * {
 *     public int MinScoreTriangulation(int[] values)
 *     {
 *         var n = values.Length;
 *         var memo = new Dictionary<(long, int), int>();
 *         int traversal(long visit, int left)
 *         {
 *             if (left == 0)
 *             {
 *                 return 0;
 *             }
 *             var key = (visit, left);
 *             if (!memo.ContainsKey(key))
 *             {
 *                 memo[key] = int.MaxValue;
 *                 for (var i = 0; i < n; i++)
 *                 {
 *                     if ((visit & (1 << i)) == 0)
 *                     {
 *                         var (p, q) = (i - 1, i + 1);
 *                         for (; (visit & (1 << ((p + n) % n))) != 0; p--) { }
 *                         for (; (visit & (1 << (q % n))) != 0; q++) { }
 *                         var cur = values[(p + n) % n] * values[i] * values[q % n];
 *                         memo[key] = Math.Min(memo[key], traversal(visit | (1 << i), left - 1) + cur);
 *                     }
 *                 }
 *             }
 *             return memo[key];
 *         }
 *         return traversal(0, n - 2);
 *     }
 * }
 */
