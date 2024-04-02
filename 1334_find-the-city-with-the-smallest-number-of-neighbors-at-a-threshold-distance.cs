/*
 * @lc app=leetcode.cn id=find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1334] 阈值距离内邻居最少的城市
 *
 * https://leetcode.cn/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/description/
 *
 * algorithms
 * Medium (53.42%)
 * Total Accepted:    14K
 * Total Submissions: 25.3K
 * Testcase Example:  '4\n[[0,1,3],[1,2,1],[1,3,4],[2,3,1]]\n4'
 *
 * 有 n 个城市，按从 0 到 n-1 编号。给你一个边数组 edges，其中 edges[i] = [fromi, toi, weighti] 代表
 * fromi 和 toi 两个城市之间的双向加权边，距离阈值是一个整数 distanceThreshold。
 * 
 * 返回能通过某些路径到达其他城市数目最少、且路径距离 最大 为 distanceThreshold 的城市。如果有多个这样的城市，则返回编号最大的城市。
 * 
 * 注意，连接城市 i 和 j 的路径的距离等于沿该路径的所有边的权重之和。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 
 * 输入：n = 4, edges = [[0,1,3],[1,2,1],[1,3,4],[2,3,1]], distanceThreshold = 4
 * 输出：3
 * 解释：城市分布图如上。
 * 每个城市阈值距离 distanceThreshold = 4 内的邻居城市分别是：
 * 城市 0 -> [城市 1, 城市 2] 
 * 城市 1 -> [城市 0, 城市 2, 城市 3] 
 * 城市 2 -> [城市 0, 城市 1, 城市 3] 
 * 城市 3 -> [城市 1, 城市 2] 
 * 城市 0 和 3 在阈值距离 4 以内都有 2 个邻居城市，但是我们必须返回城市 3，因为它的编号最大。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 
 * 
 * 输入：n = 5, edges = [[0,1,2],[0,4,8],[1,2,3],[1,4,2],[2,3,1],[3,4,1]],
 * distanceThreshold = 2
 * 输出：0
 * 解释：城市分布图如上。 
 * 每个城市阈值距离 distanceThreshold = 2 内的邻居城市分别是：
 * 城市 0 -> [城市 1] 
 * 城市 1 -> [城市 0, 城市 4] 
 * 城市 2 -> [城市 3, 城市 4] 
 * 城市 3 -> [城市 2, 城市 4]
 * 城市 4 -> [城市 1, 城市 2, 城市 3] 
 * 城市 0 在阈值距离 2 以内只有 1 个邻居城市。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 2 
 * 1 
 * edges[i].length == 3
 * 0 i < toi < n
 * 1 i, distanceThreshold 
 * 所有 (fromi, toi) 都是不同的。
 * 
 * 
 */
public class Solution
{
    public int FindTheCity(int n, int[][] edges, int distanceThreshold)
    {
        const int Maxn = 0x3f3f3f3f;
        var G = new List<(int, int)>[n];
        for (var i = 0; i < n; i++) { G[i] = new(); }
        foreach (var edge in edges)
        {
            var (u, v, w) = (edge[0], edge[1], edge[2]);
            G[u].Add((v, w));
            G[v].Add((u, w));
        }
        var (ans, count) = (-1, Maxn);
        for (var src = 0; src < n; src++)
        {
            // dijkstra
            var dists = new int[n];
            var visits = new bool[n];
            Array.Fill(dists, Maxn);
            dists[src] = 0;
            for (var i = 0; i < n; i++)
            {
                var u = -1;
                for (var v = 0; v < n; v++)
                {
                    if (!visits[v] && (u is -1 || dists[v] < dists[u])) { u = v; }
                }
                foreach (var (v, w) in G[u])
                {
                    if (dists[u] + w < dists[v]) { dists[v] = dists[u] + w; }
                }
                visits[u] = true;
            }
            var c = dists.Count(d => d <= distanceThreshold);
            if (c <= count) { (ans, count) = (src, c); }
        }
        return ans;
    }
}
