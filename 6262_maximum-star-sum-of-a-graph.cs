/*
 * @lc app=leetcode.cn id=maximum-star-sum-of-a-graph lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6262] 图中最大星和
 *
 * https://leetcode.cn/problems/maximum-star-sum-of-a-graph/description/
 *
 * algorithms
 * Medium (35.78%)
 * Total Accepted:    1.9K
 * Total Submissions: 5.4K
 * Testcase Example:  '[1,2,3,4,10,-10,-20]\n[[0,1],[1,2],[1,3],[3,4],[3,5],[3,6]]\n2'
 *
 * 给你一个 n 个点的无向图，节点从 0 到 n - 1 编号。给你一个长度为 n 下标从 0 开始的整数数组 vals ，其中 vals[i] 表示第
 * i 个节点的值。
 * 
 * 同时给你一个二维整数数组 edges ，其中 edges[i] = [ai, bi] 表示节点 ai 和 bi 之间有一条双向边。
 * 
 * 星图 是给定图中的一个子图，它包含一个中心节点和 0 个或更多个邻居。换言之，星图是给定图中一个边的子集，且这些边都有一个公共节点。
 * 
 * 下图分别展示了有 3 个和 4 个邻居的星图，蓝色节点为中心节点。
 * 
 * 
 * 
 * 星和 定义为星图中所有节点值的和。
 * 
 * 给你一个整数 k ，请你返回 至多 包含 k 条边的星图中的 最大星和 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 输入：vals = [1,2,3,4,10,-10,-20], edges =
 * [[0,1],[1,2],[1,3],[3,4],[3,5],[3,6]], k = 2
 * 输出：16
 * 解释：上图展示了输入示例。
 * 最大星和对应的星图在上图中用蓝色标出。中心节点是 3 ，星图中还包含邻居 1 和 4 。
 * 无法得到一个和大于 16 且边数不超过 2 的星图。
 * 
 * 
 * 示例 2：
 * 
 * 输入：vals = [-5], edges = [], k = 0
 * 输出：-5
 * 解释：只有一个星图，就是节点 0 自己。
 * 所以我们返回 -5 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * n == vals.length
 * 1 <= n <= 10^5
 * -10^4 <= vals[i] <= 10^4
 * 0 <= edges.length <= min(n * (n - 1) / 2, 10^5)
 * edges[i].length == 2
 * 0 <= ai, bi <= n - 1
 * ai != bi
 * 0 <= k <= n - 1
 * 
 * 
 */
public class Solution
{
    public int MaxStarSum(int[] vals, int[][] edges, int k)
    {
        var n = vals.Length;
        var G = new List<int>[n];
        for (var i = 0; i < n; i++) { G[i] = new(); }
        foreach (var e in edges)
        {
            G[e[0]].Add(e[1]);
            G[e[1]].Add(e[0]);
        }
        var ans = int.MinValue;
        for (var i = 0; i < n; i++)
        {
            G[i].Sort((x, y) => -vals[x].CompareTo(vals[y]));
            var cur = vals[i];
            for (var j = 0; j < k && j < G[i].Count; j++)
            {
                if (vals[G[i][j]] > 0) { cur += vals[G[i][j]]; }
            }
            ans = Math.Max(ans, cur);
        }
        return ans;
    }
}
