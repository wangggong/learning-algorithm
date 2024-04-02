/*
 * @lc app=leetcode.cn id=find-if-path-exists-in-graph lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1971] 寻找图中是否存在路径
 *
 * https://leetcode.cn/problems/find-if-path-exists-in-graph/description/
 *
 * algorithms
 * Easy (45.50%)
 * Total Accepted:    10.6K
 * Total Submissions: 21.5K
 * Testcase Example:  '3\n[[0,1],[1,2],[2,0]]\n0\n2'
 *
 * 有一个具有 n 个顶点的 双向 图，其中每个顶点标记从 0 到 n - 1（包含 0 和 n - 1）。图中的边用一个二维整数数组 edges
 * 表示，其中 edges[i] = [ui, vi] 表示顶点 ui 和顶点 vi 之间的双向边。 每个顶点对由 最多一条
 * 边连接，并且没有顶点存在与自身相连的边。
 * 
 * 请你确定是否存在从顶点 source 开始，到顶点 destination 结束的 有效路径 。
 * 
 * 给你数组 edges 和整数 n、source 和 destination，如果从 source 到 destination 存在 有效路径 ，则返回
 * true，否则返回 false 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：n = 3, edges = [[0,1],[1,2],[2,0]], source = 0, destination = 2
 * 输出：true
 * 解释：存在由顶点 0 到顶点 2 的路径:
 * - 0 → 1 → 2 
 * - 0 → 2
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：n = 6, edges = [[0,1],[0,2],[3,5],[5,4],[4,3]], source = 0, destination =
 * 5
 * 输出：false
 * 解释：不存在由顶点 0 到顶点 5 的路径.
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 2 * 10^5
 * 0 <= edges.length <= 2 * 10^5
 * edges[i].length == 2
 * 0 <= ui, vi <= n - 1
 * ui != vi
 * 0 <= source, destination <= n - 1
 * 不存在重复边
 * 不存在指向顶点自身的边
 * 
 * 
 */

/*
 * public class Solution
 * {
 *     public bool ValidPath(int n, int[][] edges, int source, int destination)
 *     {
 *         var G = new List<int>[n];
 *         for (var i = 0; i < n; i++)
 *         {
 *             G[i] = new();
 *         }
 *         var visit = new bool[n];
 *         foreach (var e in edges)
 *         {
 *             G[e[0]].Add(e[1]);
 *             G[e[1]].Add(e[0]);
 *         }
 *         var Q = new Queue<int>();
 *         Q.Enqueue(source);
 *         visit[source] = true;
 *         while (Q.Count > 0)
 *         {
 *             var q = Q.Dequeue();
 *             foreach (var v in G[q])
 *             {
 *                 if (!visit[v])
 *                 {
 *                     Q.Enqueue(v);
 *                     visit[v] = true;
 *                 }
 *             }
 *         }
 *         return visit[destination];
 *     }
 * }
 */

public class Solution
{
    public bool ValidPath(int n, int[][] edges, int source, int destination)
    {
        var pa = new int[n];
        for (var i = 0; i < n; i++)
        {
            pa[i] = i;
        }
        int query(int k) => k == pa[k] ? pa[k] : (pa[k] = query(pa[k]));
        void merge(int p, int q) => pa[query(p)] = query(q);
        foreach (var e in edges)
        {
            merge(e[0], e[1]);
        }
        return query(source) == query(destination);
    }
}
