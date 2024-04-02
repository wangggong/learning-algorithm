/*
 * @lc app=leetcode.cn id=minimum-edge-weight-equilibrium-queries-in-a-tree lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [100018] 边权重均等查询
 *
 * https://leetcode.cn/problems/minimum-edge-weight-equilibrium-queries-in-a-tree/description/
 *
 * algorithms
 * Hard (36.79%)
 * Total Accepted:    1.4K
 * Total Submissions: 3.1K
 * Testcase Example:  '7\n' +
  '[[0,1,1],[1,2,1],[2,3,1],[3,4,2],[4,5,2],[5,6,2]]\n' +
  '[[0,3],[3,6],[2,6],[0,6]]'
 *
 * 现有一棵由 n 个节点组成的无向树，节点按从 0 到 n - 1 编号。给你一个整数 n 和一个长度为 n - 1 的二维整数数组 edges ，其中
 * edges[i] = [ui, vi, wi] 表示树中存在一条位于节点 ui 和节点 vi 之间、权重为 wi 的边。
 * 
 * 另给你一个长度为 m 的二维整数数组 queries ，其中 queries[i] = [ai, bi] 。对于每条查询，请你找出使从 ai 到 bi
 * 路径上每条边的权重相等所需的 最小操作次数 。在一次操作中，你可以选择树上的任意一条边，并将其权重更改为任意值。
 * 
 * 注意：
 * 
 * 
 * 查询之间 相互独立 的，这意味着每条新的查询时，树都会回到 初始状态 。
 * 从 ai 到 bi的路径是一个由 不同 节点组成的序列，从节点 ai 开始，到节点 bi 结束，且序列中相邻的两个节点在树中共享一条边。
 * 
 * 
 * 返回一个长度为 m 的数组 answer ，其中 answer[i] 是第 i 条查询的答案。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：n = 7, edges = [[0,1,1],[1,2,1],[2,3,1],[3,4,2],[4,5,2],[5,6,2]], queries
 * = [[0,3],[3,6],[2,6],[0,6]]
 * 输出：[0,0,1,3]
 * 解释：第 1 条查询，从节点 0 到节点 3 的路径中的所有边的权重都是 1 。因此，答案为 0 。
 * 第 2 条查询，从节点 3 到节点 6 的路径中的所有边的权重都是 2 。因此，答案为 0 。
 * 第 3 条查询，将边 [2,3] 的权重变更为 2 。在这次操作之后，从节点 2 到节点 6 的路径中的所有边的权重都是 2 。因此，答案为 1 。
 * 第 4 条查询，将边 [0,1]、[1,2]、[2,3] 的权重变更为 2 。在这次操作之后，从节点 0 到节点 6 的路径中的所有边的权重都是 2
 * 。因此，答案为 3 。
 * 对于每条查询 queries[i] ，可以证明 answer[i] 是使从 ai 到 bi 的路径中的所有边的权重相等的最小操作次数。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：n = 8, edges = [[1,2,6],[1,3,4],[2,4,6],[2,5,3],[3,6,6],[3,0,8],[7,0,2]],
 * queries = [[4,6],[0,4],[6,5],[7,4]]
 * 输出：[1,2,2,3]
 * 解释：第 1 条查询，将边 [1,3] 的权重变更为 6 。在这次操作之后，从节点 4 到节点 6 的路径中的所有边的权重都是 6 。因此，答案为 1
 * 。
 * 第 2 条查询，将边 [0,3]、[3,1] 的权重变更为 6 。在这次操作之后，从节点 0 到节点 4 的路径中的所有边的权重都是 6 。因此，答案为
 * 2 。
 * 第 3 条查询，将边 [1,3]、[5,2] 的权重变更为 6 。在这次操作之后，从节点 6 到节点 5 的路径中的所有边的权重都是 6 。因此，答案为
 * 2 。
 * 第 4 条查询，将边 [0,7]、[0,3]、[1,3] 的权重变更为 6 。在这次操作之后，从节点 7 到节点 4 的路径中的所有边的权重都是 6
 * 。因此，答案为 3 。
 * 对于每条查询 queries[i] ，可以证明 answer[i] 是使从 ai 到 bi 的路径中的所有边的权重相等的最小操作次数。 
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 10^4
 * edges.length == n - 1
 * edges[i].length == 3
 * 0 <= ui, vi < n
 * 1 <= wi <= 26
 * 生成的输入满足 edges 表示一棵有效的树
 * 1 <= queries.length == m <= 2 * 10^4
 * queries[i].length == 2
 * 0 <= ai, bi < n
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/minimum-edge-weight-equilibrium-queries-in-a-tree/solutions/2424060/lca-mo-ban-by-endlesscheng-j54b/
//
// 关键词: 树上倍增, LCA
// 相关题目: 1483_kth-ancestor-of-a-tree-node.cs
public class Solution
{
    public int[] MinOperationsQueries(int n, int[][] edges, int[][] queries)
    {
        const int N = 26;
        const int D = 14;
        // 建图
        var G = Enumerable.Range(0, n)
            .Select(_ => new List<(int, int)>())
            .ToList();
        foreach (var edge in edges)
        {
            var (u, v, w) = (edge[0], edge[1], edge[2] - 1);
            G[u].Add((v, w));
            G[v].Add((u, w));
        }
        // DFS
        var parents = Enumerable.Range(0, n)
            .Select(_ => new int[D])
            .ToList();
        var counts = Enumerable.Range(0, n)
            .Select(_ => Enumerable.Range(0, D)
                .Select(_ => new int[N])
                .ToList())
            .ToList();
        var depths = new int[n];
        void build(int u, int p, int d)
        {
            parents[u][0] = p;
            depths[u] = d;
            foreach (var (v, w) in G[u])
            {
                if (v == p) { continue; }
                counts[v][0][w] = 1;
                build(v, u, d + 1);
            }
        }
        build(0, -1, 0);
        // 倍增
        for (var d = 0; d + 1 < D; d++)
        {
            for (var u = 0; u < n; u++)
            {
                var p = parents[u][d];
                if (p is -1)
                {
                    parents[u][d + 1] = -1;
                    continue;
                }
                parents[u][d + 1] = parents[p][d];
                for (var i = 0; i < N; i++)
                {
                    counts[u][d + 1][i] = counts[u][d][i] + counts[p][d][i];
                }
            }
        }
        // LCA
        return queries.Select(q =>
        {
            var (u, v) = (q[0], q[1]);
            if (depths[u] > depths[v]) { (u, v) = (v, u); }
            var ans = depths[u] + depths[v];
            var cnts = new int[N];
            for (var d = 0; d < D; d++)
            {
                if ((((depths[v] - depths[u]) >> d) & 1) is 0) { continue; }
                for (var i = 0; i < N; i++) { cnts[i] += counts[v][d][i]; }
                v = parents[v][d];
            }
            if (u != v)
            {
                for (var d = D - 1; d >= 0; d--)
                {
                    if (parents[u][d] == parents[v][d]) { continue; }
                    for (var i = 0; i < N; i++)
                    {
                        cnts[i] += counts[u][d][i] + counts[v][d][i];
                    }
                    (u, v) = (parents[u][d], parents[v][d]);
                }
                for (var i = 0; i < N; i++)
                {
                    cnts[i] += counts[u][0][i] + counts[v][0][i];
                }
                (u, v) = (parents[u][0], parents[v][0]);
            }
            ans -= 2 * depths[u];
            return ans - cnts.Max();
        }).ToArray();
    }
}
