/*
 * @lc app=leetcode.cn id=painting-the-walls lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6447] 给墙壁刷油漆
 *
 * https://leetcode.cn/problems/painting-the-walls/description/
 *
 * algorithms
 * Hard (19.08%)
 * Total Accepted:    740
 * Total Submissions: 3.8K
 * Testcase Example:  '[1,2,3,2]\n[1,2,3,2]'
 *
 * 给你两个长度为 n 下标从 0 开始的整数数组 cost 和 time ，分别表示给 n 堵不同的墙刷油漆需要的开销和时间。你有两名油漆匠：
 * 
 * 
 * 一位需要 付费 的油漆匠，刷第 i 堵墙需要花费 time[i] 单位的时间，开销为 cost[i] 单位的钱。
 * 一位 免费 的油漆匠，刷 任意 一堵墙的时间为 1 单位，开销为 0 。但是必须在付费油漆匠 工作 时，免费油漆匠才会工作。
 * 
 * 
 * 请你返回刷完 n 堵墙最少开销为多少。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：cost = [1,2,3,2], time = [1,2,3,2]
 * 输出：3
 * 解释：下标为 0 和 1 的墙由付费油漆匠来刷，需要 3 单位时间。同时，免费油漆匠刷下标为 2 和 3 的墙，需要 2 单位时间，开销为 0
 * 。总开销为 1 + 2 = 3 。
 * 
 * 
 * 示例 2：
 * 
 * 输入：cost = [2,3,4,2], time = [1,1,1,1]
 * 输出：4
 * 解释：下标为 0 和 3 的墙由付费油漆匠来刷，需要 2 单位时间。同时，免费油漆匠刷下标为 1 和 2 的墙，需要 2 单位时间，开销为 0
 * 。总开销为 2 + 2 = 4 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= cost.length <= 500
 * cost.length == time.length
 * 1 <= cost[i] <= 10^6
 * 1 <= time[i] <= 500
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/painting-the-walls/solution/xuan-huo-bu-xuan-de-dian-xing-si-lu-by-e-ulcd/
//
// 傻逼! DP 不会写了!
public class Solution
{
    private const int Maxn = 0x3f3f3f3f;

    public int PaintWalls(int[] cost, int[] time)
    {
        var n = cost.Length;
        var memo = new Dictionary<(int, int), int>();
        int dfs(int k, int p)
        {
            var key = (k, p);
            if (!memo.ContainsKey(key))
            {
                memo[key] = k >= n
                    ? (p >= 0 ? 0 : Maxn)
                    : Math.Min(
                        dfs(k + 1, p + time[k]) + cost[k],
                        dfs(k + 1, p - 1));
            }
            return memo[key];
        }
        return dfs(0, 0);
    }

    /*
     * // 0-1 背包, 巧妙
     * public int PaintWalls(int[] cost, int[] time)
     * {
     *     var n = cost.Length;
     *     var dp = Enumerable.Range(0, n + 1)
     *         .Select(_ => Maxn)
     *         .ToArray();
     *     dp[0] = 0;
     *     foreach (var (c, t) in cost.Zip(time))
     *     {
     *         for (var j = n; j >= 0; j--)
     *         {
     *             dp[j] = Math.Min(dp[j], dp[Math.Max(j - t - 1, 0)] + c);
     *         }
     *     }
     *     return dp[n];
     * }
     */
}
