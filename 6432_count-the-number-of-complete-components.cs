/*
 * @lc app=leetcode.cn id=count-the-number-of-complete-components lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6432] 统计完全连通分量的数量
 *
 * https://leetcode.cn/problems/count-the-number-of-complete-components/description/
 *
 * algorithms
 * Medium (63.15%)
 * Total Accepted:    2.1K
 * Total Submissions: 3.2K
 * Testcase Example:  '6\n[[0,1],[0,2],[1,2],[3,4]]'
 *
 * 给你一个整数 n 。现有一个包含 n 个顶点的 无向 图，顶点按从 0 到 n - 1 编号。给你一个二维整数数组 edges 其中 edges[i]
 * = [ai, bi] 表示顶点 ai 和 bi 之间存在一条 无向 边。
 * 
 * 返回图中 完全连通分量 的数量。
 * 
 * 如果在子图中任意两个顶点之间都存在路径，并且子图中没有任何一个顶点与子图外部的顶点共享边，则称其为 连通分量 。
 * 
 * 如果连通分量中每对节点之间都存在一条边，则称其为 完全连通分量 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 
 * 输入：n = 6, edges = [[0,1],[0,2],[1,2],[3,4]]
 * 输出：3
 * 解释：如上图所示，可以看到此图所有分量都是完全连通分量。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 
 * 
 * 输入：n = 6, edges = [[0,1],[0,2],[1,2],[3,4],[3,5]]
 * 输出：1
 * 解释：包含节点 0、1 和 2 的分量是完全连通分量，因为每对节点之间都存在一条边。
 * 包含节点 3 、4 和 5 的分量不是完全连通分量，因为节点 4 和 5 之间不存在边。
 * 因此，在图中完全连接分量的数量是 1 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 50
 * 0 <= edges.length <= n * (n - 1) / 2
 * edges[i].length == 2
 * 0 <= ai, bi <= n - 1
 * ai != bi
 * 不存在重复的边
 * 
 * 
 */

/*
 * public class Solution
 * {
 *     public int CountCompleteComponents(int n, int[][] edges)
 *     {
 *         var S = edges.Select(edge => (edge[0], edge[1])).ToHashSet();
 *         var pa = Enumerable.Range(0, n).Select(k => k).ToArray();
 *         int query(int k) => k == pa[k] ? k : (pa[k] = query(pa[k]));
 *         void merge(int p, int q) => pa[query(p)] = query(q);
 *         foreach (var edge in edges)
 *         {
 *             merge(edge[0], edge[1]);
 *         }
 *         var nodesMap = new Dictionary<int, List<int>>();
 *         for (var i = 0; i < n; i++)
 *         {
 *             var p = query(i);
 *             if (!nodesMap.ContainsKey(p))
 *             {
 *                 nodesMap[p] = new();
 *             }
 *             nodesMap[p].Add(i);
 *         }
 *         bool valid(List<int> list)
 *         {
 *             foreach (var u in list)
 *             {
 *                 foreach (var v in list)
 *                 {
 *                     if (u != v && !S.Contains((u, v)) && !S.Contains((v, u)))
 *                     {
 *                         return false;
 *                     }
 *                 }
 *             }
 *             return true;
 *         }
 *         return nodesMap.Where(kv => valid(kv.Value)).Count();
 *     }
 * }
 */

public class Solution
{
    public int CountCompleteComponents(int n, int[][] edges)
    {
        var pa = Enumerable.Range(0, n).Select(i => i).ToArray();
        int query(int k) => k == pa[k] ? k : (pa[k] = query(pa[k]));
        void merge(int p, int q) => pa[query(p)] = query(q);
        var degrees = new int[n];
        foreach (var edge in edges)
        {
            var (u, v) = (edge[0], edge[1]);
            merge(u, v);
            degrees[u]++;
            degrees[v]++;
        }
        return Enumerable
            .Range(0, n)
            .GroupBy(i => query(i))
            .Where(g => g.All(v => degrees[v] == g.Count() - 1))
            .Count();
    }
}
