/*
 * @lc app=leetcode.cn id=modify-graph-edge-weights lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6442] 修改图中的边权
 *
 * https://leetcode.cn/problems/modify-graph-edge-weights/description/
 *
 * algorithms
 * Hard (10.30%)
 * Total Accepted:    397
 * Total Submissions: 3.1K
 * Testcase Example:  '5\n[[4,1,-1],[2,0,-1],[0,3,-1],[4,3,-1]]\n0\n1\n5'
 *
 * 给你一个 n 个节点的 无向带权连通 图，节点编号为 0 到 n - 1 ，再给你一个整数数组 edges ，其中 edges[i] = [ai,
 * bi, wi] 表示节点 ai 和 bi 之间有一条边权为 wi 的边。
 * 
 * 部分边的边权为 -1（wi = -1），其他边的边权都为 正 数（wi > 0）。
 * 
 * 你需要将所有边权为 -1 的边都修改为范围 [1, 2 * 10^9] 中的 正整数 ，使得从节点 source 到节点 destination 的
 * 最短距离 为整数 target 。如果有 多种 修改方案可以使 source 和 destination 之间的最短距离等于 target
 * ，你可以返回任意一种方案。
 * 
 * 如果存在使 source 到 destination 最短距离为 target
 * 的方案，请你按任意顺序返回包含所有边的数组（包括未修改边权的边）。如果不存在这样的方案，请你返回一个 空数组 。
 * 
 * 注意：你不能修改一开始边权为正数的边。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 
 * 输入：n = 5, edges = [[4,1,-1],[2,0,-1],[0,3,-1],[4,3,-1]], source = 0,
 * destination = 1, target = 5
 * 输出：[[4,1,1],[2,0,1],[0,3,3],[4,3,1]]
 * 解释：上图展示了一个满足题意的修改方案，从 0 到 1 的最短距离为 5 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 
 * 
 * 输入：n = 3, edges = [[0,1,-1],[0,2,5]], source = 0, destination = 2, target =
 * 6
 * 输出：[]
 * 解释：上图是一开始的图。没有办法通过修改边权为 -1 的边，使得 0 到 2 的最短距离等于 6 ，所以返回一个空数组。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 
 * 
 * 输入：n = 4, edges = [[1,0,4],[1,2,3],[2,3,5],[0,3,-1]], source = 0,
 * destination = 2, target = 6
 * 输出：[[1,0,4],[1,2,3],[2,3,5],[0,3,1]]
 * 解释：上图展示了一个满足题意的修改方案，从 0 到 2 的最短距离为 6 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 100
 * 1 <= edges.length <= n * (n - 1) / 2
 * edges[i].length == 3
 * 0 <= ai, bi < n
 * wi = -1 或者 1 <= wi <= 10^7
 * ai != bi
 * 0 <= source, destination < n
 * source != destination
 * 1 <= target <= 10^9
 * 输入的图是连通图，且没有自环和重边。
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/modify-graph-edge-weights/solution/bi-wan-hou-kan-liao-xia-tao-lun-qu-shun-1foic/
public class Solution
{
    private const int Maxn = 0x3f3f3f3f;

    public int[][] ModifiedGraphEdges(int n, int[][] edges, int source, int destination, int target)
    {
        var G = Enumerable.Range(0, n)
            .Select(_ => Enumerable.Range(0, n).Select(_ => Maxn).ToArray())
            .ToArray();
        for (var i = 0; i < n; i++)
        {
            G[i][i] = 0;
        }
        foreach (var edge in edges)
        {
            var (u, v, w) = (edge[0], edge[1], edge[2]);
            (G[u][v], G[v][u]) = (w, w);
        }
        var dist = new int[n];
        var visit = new bool[n];
        void dijkstra()
        {
            for (var i = 0; i < n; i++)
            {
                (dist[i], visit[i]) = (Maxn, false);
            }
            dist[source] = 0;
            for (var i = 0; i < n; i++)
            {
                var u = -1;
                for (var v = 0; v < n; v++)
                {
                    if (!visit[v] && (u == -1 || dist[u] > dist[v]))
                    {
                        u = v;
                    }
                }
                for (var v = 0; v < n; v++)
                {
                    var w = G[u][v] == -1 ? Maxn : G[u][v];
                    if (dist[v] > dist[u] + w)
                    {
                        dist[v] = dist[u] + w;
                    }
                }
                visit[u] = true;
            }
        }
        dijkstra();
        if (dist[destination] < target)
        {
            return new int[0][];
        }
        var valid = false;
        foreach (var edge in edges)
        {
            var (u, v, w) = (edge[0], edge[1], edge[2]);
            if (w == -1)
            {
                (G[u][v], G[v][u]) = (1, 1);
                dijkstra();
                if (dist[destination] <= target)
                {
                    var weight = G[u][v] + (target - dist[destination]);
                    (G[u][v], G[v][u]) = (weight, weight);
                    valid = true;
                    break;
                }
            }
        }
        if (!valid && dist[destination] > target)
        {
            return new int[0][];
        }
        for (var (i, m) = (0, edges.Length); i < m; i++)
        {
            var edge = edges[i];
            var (u, v, w) = (edge[0], edge[1], edge[2]);
            if (w == -1)
            {
                edges[i][2] = G[u][v] == -1 ? Maxn : G[u][v];
            }
        }
        return edges;
    }
}
