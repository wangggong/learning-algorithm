/*
 * @lc app=leetcode.cn id=shortest-cycle-in-a-graph lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6330] 图中的最短环
 *
 * https://leetcode.cn/problems/shortest-cycle-in-a-graph/description/
 *
 * algorithms
 * Hard (35.26%)
 * Total Accepted:    2K
 * Total Submissions: 5.2K
 * Testcase Example:  '7\n[[0,1],[1,2],[2,0],[3,4],[4,5],[5,6],[6,3]]'
 *
 * 现有一个含 n 个顶点的 双向 图，每个顶点按从 0 到 n - 1 标记。图中的边由二维整数数组 edges 表示，其中 edges[i] =
 * [ui, vi] 表示顶点 ui 和 vi 之间存在一条边。每对顶点最多通过一条边连接，并且不存在与自身相连的顶点。
 * 
 * 返回图中 最短 环的长度。如果不存在环，则返回 -1 。
 * 
 * 环 是指以同一节点开始和结束，并且路径中的每条边仅使用一次。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：n = 7, edges = [[0,1],[1,2],[2,0],[3,4],[4,5],[5,6],[6,3]]
 * 输出：3
 * 解释：长度最小的循环是：0 -> 1 -> 2 -> 0 
 * 
 * 
 * 示例 2：
 * 
 * 输入：n = 4, edges = [[0,1],[0,2]]
 * 输出：-1
 * 解释：图中不存在循环
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 2 <= n <= 1000
 * 1 <= edges.length <= 1000
 * edges[i].length == 2
 * 0 <= ui, vi < n
 * ui != vi
 * 不存在重复的边
 * 
 * 
 */
public class Solution
{
    public int FindShortestCycle(int n, int[][] edges)
	{
		var G = new List<int>[n];
		for (var i = 0; i < n; i++)
		{
			G[i] = new();
		}
		foreach (var e in edges)
		{
			var (u, v) = (e[0], e[1]);
			G[u].Add(v);
			G[v].Add(u);
		}
		int bfs(int k)
		{
			var ans = n + 1;
			var Q = new Queue<(int, int)>();
			var dist = Enumerable.Range(0, n).Select(_ => -1).ToArray();
			Q.Enqueue((k, -1));
			dist[k] = 0;
			while (Q.Count > 0)
			{
				var (u, p) = Q.Dequeue();
				foreach (var v in G[u])
				{
					if (dist[v] < 0)
					{
						Q.Enqueue((v, u));
						dist[v] = dist[u] + 1;
					}
					else if (v != p)
					{
						ans = Math.Min(ans, dist[u] + dist[v] + 1);
					}
				}
			}
			return ans;
		}
		var ans = Enumerable.Range(0, n).Select(i => bfs(i)).Min();
		return ans > n ? -1 : ans;
    }
}
