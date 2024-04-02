/*
 * @lc app=leetcode.cn id=reachable-nodes-in-subdivided-graph lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [882] 细分图中的可到达节点
 *
 * https://leetcode.cn/problems/reachable-nodes-in-subdivided-graph/description/
 *
 * algorithms
 * Hard (49.90%)
 * Total Accepted:    7.1K
 * Total Submissions: 11.9K
 * Testcase Example:  '[[0,1,10],[0,2,1],[1,2,2]]\n6\n3'
 *
 * 给你一个无向图（原始图），图中有 n 个节点，编号从 0 到 n - 1 。你决定将图中的每条边 细分 为一条节点链，每条边之间的新节点数各不相同。
 * 
 * 图用由边组成的二维数组 edges 表示，其中 edges[i] = [ui, vi, cnti] 表示原始图中节点 ui 和 vi
 * 之间存在一条边，cnti 是将边 细分 后的新节点总数。注意，cnti == 0 表示边不可细分。
 * 
 * 要 细分 边 [ui, vi] ，需要将其替换为 (cnti + 1) 条新边，和 cnti 个新节点。新节点为 x1, x2, ..., xcnti
 * ，新边为 [ui, x1], [x1, x2], [x2, x3], ..., [xcnti+1, xcnti], [xcnti, vi] 。
 * 
 * 现在得到一个 新的细分图 ，请你计算从节点 0 出发，可以到达多少个节点？如果节点间距离是 maxMoves 或更少，则视为 可以到达 。
 * 
 * 给你原始图和 maxMoves ，返回 新的细分图中从节点 0 出发 可到达的节点数 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：edges = [[0,1,10],[0,2,1],[1,2,2]], maxMoves = 6, n = 3
 * 输出：13
 * 解释：边的细分情况如上图所示。
 * 可以到达的节点已经用黄色标注出来。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：edges = [[0,1,4],[1,2,6],[0,2,8],[1,3,1]], maxMoves = 10, n = 4
 * 输出：23
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：edges = [[1,2,4],[1,4,5],[1,3,1],[2,3,4],[3,4,5]], maxMoves = 17, n = 5
 * 输出：1
 * 解释：节点 0 与图的其余部分没有连通，所以只有节点 0 可以到达。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 0 <= edges.length <= min(n * (n - 1) / 2, 10^4)
 * edges[i].length == 3
 * 0 <= ui < vi < n
 * 图中 不存在平行边
 * 0 <= cnti <= 10^4
 * 0 <= maxMoves <= 10^9
 * 1 <= n <= 3000
 * 
 * 
 */
// 参考: https://leetcode.cn/problems/reachable-nodes-in-subdivided-graph/solution/tu-jie-zhuan-huan-cheng-dan-yuan-zui-dua-6l8o/
//
// Dijkstra 应用题
public class Solution
{
    public int ReachableNodes(int[][] edges, int maxMoves, int n)
    {
        var G = new List<(int, int)>[n];
        for (int i = 0; i < n; i++) { G[i] = new(); }
        foreach (var e in edges)
        {
            int u = e[0], v = e[1], w = e[2];
            G[u].Add((v, w + 1));
            G[v].Add((u, w + 1));
        }
        // Dijkstra 非堆优化版
        int[] Dijkstra(int start)
        {
            var dist = new int[n];
            var vis = new bool[n];
            for (int i = 0; i < n; i++) { dist[i] = 0x3f3f3f3f; }
            dist[start] = 0;
            for (int i = 1; i < n; i++)
            {
                var u = -1;
                for (int v = 0; v < n; v++)
                {
                    if (!vis[v] && (u == -1 || dist[u] > dist[v])) { u = v; }
                }
                foreach (var e in G[u]) { dist[e.Item1] = Math.Min(dist[e.Item1], dist[u] + e.Item2); }
                vis[u] = true;
            }
            return dist;
        }
        var dist = Dijkstra(0);
        var ans = 0;
        foreach (var d in dist) { ans += d <= maxMoves ? 1 : 0; }
        foreach (var e in edges)
        {
            int u = e[0], v = e[1], w = e[2];
            int a = Math.Max(maxMoves - dist[u], 0), b = Math.Max(maxMoves - dist[v], 0);
            ans += Math.Min(a + b, w);
        }
        return ans;
    }
}
