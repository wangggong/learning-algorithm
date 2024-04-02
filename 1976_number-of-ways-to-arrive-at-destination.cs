/*
 * @lc app=leetcode.cn id=number-of-ways-to-arrive-at-destination lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1976] 到达目的地的方案数
 *
 * https://leetcode.cn/problems/number-of-ways-to-arrive-at-destination/description/
 *
 * algorithms
 * Medium (37.29%)
 * Total Accepted:    11.7K
 * Total Submissions: 27.6K
 * Testcase Example:  '7\n' +
  '[[0,6,7],[0,1,2],[1,2,3],[1,3,3],[6,3,3],[3,5,1],[6,5,1],[2,5,1],[0,4,5],[4,6,2]]'
 *
 * 你在一个城市里，城市由 n 个路口组成，路口编号为 0 到 n - 1 ，某些路口之间有 双向
 * 道路。输入保证你可以从任意路口出发到达其他任意路口，且任意两个路口之间最多有一条路。
 * 
 * 给你一个整数 n 和二维整数数组 roads ，其中 roads[i] = [ui, vi, timei] 表示在路口 ui 和 vi
 * 之间有一条需要花费 timei 时间才能通过的道路。你想知道花费 最少时间 从路口 0 出发到达路口 n - 1 的方案数。
 * 
 * 请返回花费 最少时间 到达目的地的 路径数目 。由于答案可能很大，将结果对 10^9 + 7 取余 后返回。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：n = 7, roads =
 * [[0,6,7],[0,1,2],[1,2,3],[1,3,3],[6,3,3],[3,5,1],[6,5,1],[2,5,1],[0,4,5],[4,6,2]]
 * 输出：4
 * 解释：从路口 0 出发到路口 6 花费的最少时间是 7 分钟。
 * 四条花费 7 分钟的路径分别为：
 * - 0 ➝ 6
 * - 0 ➝ 4 ➝ 6
 * - 0 ➝ 1 ➝ 2 ➝ 5 ➝ 6
 * - 0 ➝ 1 ➝ 3 ➝ 5 ➝ 6
 * 
 * 
 * 示例 2：
 * 
 * 输入：n = 2, roads = [[1,0,10]]
 * 输出：1
 * 解释：只有一条从路口 0 到路口 1 的路，花费 10 分钟。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 200
 * n - 1 <= roads.length <= n * (n - 1) / 2
 * roads[i].length == 3
 * 0 <= ui, vi <= n - 1
 * 1 <= timei <= 10^9
 * ui != vi
 * 任意两个路口之间至多有一条路。
 * 从任意路口出发，你能够到达其他任意路口。
 * 
 * 
 */
public class Solution
{
    public int CountPaths(int n, int[][] edges)
    {
        const long Mod = (long)1e9 + 7;
        const long Maxn = (long)1e12;
        var G = new List<(int, long)>[n];
        for (var i = 0; i < n; i++)
        {
            G[i] = new();
        }
        foreach (var e in edges)
        {
            var (u, v, w) = (e[0], e[1], e[2]);
            G[u].Add((v, w));
            G[v].Add((u, w));
        }
        #region dijkstra
        var dists = new long[n];
        Array.Fill(dists, Maxn);
        dists[0] = 0;
        var visits = new bool[n];
        var counts = new long[n];
        counts[0] = 1;
        for (var i = 1; i < n; i++)
        {
            var u = -1;
            for (var v = 0; v < n; v++)
            {
                if (!visits[v] && (u is -1 || dists[u] > dists[v]))
                {
                    u = v;
                }
            }
            foreach (var (v, w) in G[u])
            {
                var d = dists[u] + w;
                if (d <= dists[v])
                {
                    if (d < dists[v])
                    {
                        counts[v] = 0;
                    }
                    counts[v] = (counts[v] + counts[u]) % Mod;
                    dists[v] = d;
                }
            }
            visits[u] = true;
        }
        #endregion
        return (int)counts.Last();
    }
}
