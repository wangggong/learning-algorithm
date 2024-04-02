/*
 * @lc app=leetcode.cn id=minimum-cost-to-merge-stones lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1000] 合并石头的最低成本
 *
 * https://leetcode.cn/problems/minimum-cost-to-merge-stones/description/
 *
 * algorithms
 * Hard (45.05%)
 * Total Accepted:    15.5K
 * Total Submissions: 30K
 * Testcase Example:  '[3,2,4,1]\n2'
 *
 * 有 N 堆石头排成一排，第 i 堆中有 stones[i] 块石头。
 * 
 * 每次移动（move）需要将连续的 K 堆石头合并为一堆，而这个移动的成本为这 K 堆石头的总数。
 * 
 * 找出把所有石头合并成一堆的最低成本。如果不可能，返回 -1 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：stones = [3,2,4,1], K = 2
 * 输出：20
 * 解释：
 * 从 [3, 2, 4, 1] 开始。
 * 合并 [3, 2]，成本为 5，剩下 [5, 4, 1]。
 * 合并 [4, 1]，成本为 5，剩下 [5, 5]。
 * 合并 [5, 5]，成本为 10，剩下 [10]。
 * 总成本 20，这是可能的最小值。
 * 
 * 
 * 示例 2：
 * 
 * 输入：stones = [3,2,4,1], K = 3
 * 输出：-1
 * 解释：任何合并操作后，都会剩下 2 堆，我们无法再进行合并。所以这项任务是不可能完成的。.
 * 
 * 
 * 示例 3：
 * 
 * 输入：stones = [3,5,1,2,6], K = 3
 * 输出：25
 * 解释：
 * 从 [3, 5, 1, 2, 6] 开始。
 * 合并 [5, 1, 2]，成本为 8，剩下 [3, 8, 6]。
 * 合并 [3, 8, 6]，成本为 17，剩下 [17]。
 * 总成本 25，这是可能的最小值。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= stones.length <= 30
 * 2 <= K <= 30
 * 1 <= stones[i] <= 100
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/minimum-cost-to-merge-stones/solution/tu-jie-qu-jian-dpzhuang-tai-she-ji-yu-yo-ppv0/
public class Solution
{
    public int MergeStones(int[] stones, int k)
    {
        var n = stones.Length;
        if ((n - 1) % (k - 1) != 0)
        {
            return -1;
        }
        var S = new int[n + 1];
        for (var i = 0; i < n; i++)
        {
            S[i + 1] = S[i] + stones[i];
        }
        var dp = new int[n][];
        for (var i = 0; i < n; i++)
        {
            dp[i] = new int[n];
        }
        for (var i = n - 1; i >= 0; i--)
        {
            for (var j = i + 1; j < n; j++)
            {
                dp[i][j] = int.MaxValue;
                for (var l = i; l < j; l += k - 1)
                {
                    dp[i][j] = Math.Min(dp[i][j], dp[i][l] + dp[l + 1][j]);
                }
                if ((j - i) % (k - 1) == 0)
                {
                    dp[i][j] += S[j + 1] - S[i];
                }
            }
        }
        return dp[0][n - 1];
    }
}

/*
 * // 记忆化搜索, 可以看到这个 DFS 并不机械, 甚至有点匪夷所思...
 * public class Solution
 * {
 *     public int MergeStones(int[] stones, int k)
 *     {
 *         var n = stones.Length;
 *         if ((n - 1) % (k - 1) != 0)
 *         {
 *             return -1;
 *         }
 *         var S = new int[n + 1];
 *         for (var i = 0; i < n; i++)
 *         {
 *             S[i + 1] = S[i] + stones[i];
 *         }
 *         var memo = new Dictionary<(int, int, int), int>();
 *         int traversal(int p, int q, int m)
 *         {
 *             var key = (p, q, m);
 *             if (memo.ContainsKey(key))
 *             {
 *                 return memo[key];
 *             }
 *             if (m == 1)
 *             {
 *                 return memo[key] = p == q ? 0 : (traversal(p, q, k) + S[q + 1] - S[p]);
 *             }
 *             memo[key] = int.MaxValue;
 *             for (var i = p; i < q; i += k - 1)
 *             {
 *                 memo[key] = Math.Min(memo[key], traversal(p, i, 1) + traversal(i + 1, q, m - 1));
 *             }
 *             return memo[key];
 *         }
 *         return traversal(0, n - 1, 1);
 *     }
 * }
 */
