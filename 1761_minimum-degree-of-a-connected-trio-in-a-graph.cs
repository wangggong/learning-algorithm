/*
 * @lc app=leetcode.cn id=minimum-degree-of-a-connected-trio-in-a-graph lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1761] 一个图中连通三元组的最小度数
 *
 * https://leetcode.cn/problems/minimum-degree-of-a-connected-trio-in-a-graph/description/
 *
 * algorithms
 * Hard (46.68%)
 * Total Accepted:    4.4K
 * Total Submissions: 9.2K
 * Testcase Example:  '6\n[[1,2],[1,3],[3,2],[4,1],[5,2],[3,6]]'
 *
 * 给你一个无向图，整数 n 表示图中节点的数目，edges 数组表示图中的边，其中 edges[i] = [ui, vi] ，表示 ui 和 vi
 * 之间有一条无向边。
 * 
 * 一个 连通三元组 指的是 三个 节点组成的集合且这三个点之间 两两 有边。
 * 
 * 连通三元组的度数 是所有满足此条件的边的数目：一个顶点在这个三元组内，而另一个顶点不在这个三元组内。
 * 
 * 请你返回所有连通三元组中度数的 最小值 ，如果图中没有连通三元组，那么返回 -1 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：n = 6, edges = [[1,2],[1,3],[3,2],[4,1],[5,2],[3,6]]
 * 输出：3
 * 解释：只有一个三元组 [1,2,3] 。构成度数的边在上图中已被加粗。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：n = 7, edges = [[1,3],[4,1],[4,3],[2,5],[5,6],[6,7],[7,5],[2,6]]
 * 输出：0
 * 解释：有 3 个三元组：
 * 1) [1,4,3]，度数为 0 。
 * 2) [2,5,6]，度数为 2 。
 * 3) [5,6,7]，度数为 2 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 2 
 * edges[i].length == 2
 * 1 
 * 1 i, vi 
 * ui != vi
 * 图中没有重复的边。
 * 
 * 
 */

/*
 * public class Solution
 * {
 *     public int MinTrioDegree(int n, int[][] edges)
 *     {
 *         var degrees = new int[n];
 *         var G = new bool[n][];
 *         for (var i = 0; i < n; i++) { G[i] = new bool[n]; }
 *         foreach (var edge in edges)
 *         {
 *             var (u, v) = (edge[0] - 1, edge[1] - 1);
 *             (G[u][v], G[v][u]) = (true, true);
 *             degrees[u]++;
 *             degrees[v]++;
 *         }
 *         var ans = int.MaxValue;
 *         for (var i = 0; i < n; i++)
 *         {
 *             for (var j = i + 1; j < n; j++)
 *             {
 *                 for (var k = j + 1; k < n; k++)
 *                 {
 *                     if (G[i][j] && G[j][k] && G[k][i])
 *                     {
 *                         ans = Math.Min(ans, degrees[i] + degrees[j] + degrees[k] - 6);
 *                     }
 *                 }
 *             }
 *         }
 *         return ans == int.MaxValue ? -1 : ans;
 *     }
 * }
 */

// 参考: https://leetcode.cn/problems/minimum-degree-of-a-connected-trio-in-a-graph/solutions/2417898/yi-ge-tu-zhong-lian-tong-san-yuan-zu-de-wuv8o/
//
// 官解指出了一种技巧, 将无向边转为有向边, 从而将复杂度降至 `O(E^(3/2))`. 但对稠密图没啥大用.
public class Solution
{
    public int MinTrioDegree(int n, int[][] edges)
    {
        var degrees = new int[n];
        foreach (var edge in edges)
        {
            var (u, v) = (edge[0] - 1, edge[1] - 1);
            degrees[u]++;
            degrees[v]++;
        }
        var G = Enumerable.Range(0, n)
            .Select(_ => new HashSet<int>())
            .ToArray();
        foreach (var edge in edges)
        {
            var (u, v) = (edge[0] - 1, edge[1] - 1);
            if (degrees[v] < degrees[u] || (degrees[v] == degrees[u] && v < u))
            {
                (u, v) = (v, u);
            }
            G[u].Add(v);
        }
        var ans = int.MaxValue;
        for (var u = 0; u < n; u++)
        {
            foreach (var v in G[u])
            {
                foreach (var w in G[v])
                {
                    if (G[u].Contains(w))
                    {
                        ans = Math.Min(ans, degrees[u] + degrees[v] + degrees[w] - 6);
                    }
                }
            }
        }
        return ans == int.MaxValue ? -1 : ans;
    }
}
