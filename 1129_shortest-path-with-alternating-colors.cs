/*
 * @lc app=leetcode.cn id=shortest-path-with-alternating-colors lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1129] 颜色交替的最短路径
 *
 * https://leetcode.cn/problems/shortest-path-with-alternating-colors/description/
 *
 * algorithms
 * Medium (40.83%)
 * Total Accepted:    11.7K
 * Total Submissions: 28K
 * Testcase Example:  '3\n[[0,1],[1,2]]\n[]'
 *
 * 在一个有向图中，节点分别标记为 0, 1, ..., n-1。图中每条边为红色或者蓝色，且存在自环或平行边。
 * 
 * red_edges 中的每一个 [i, j] 对表示从节点 i 到节点 j 的红色有向边。类似地，blue_edges 中的每一个 [i, j]
 * 对表示从节点 i 到节点 j 的蓝色有向边。
 * 
 * 返回长度为 n 的数组 answer，其中 answer[X] 是从节点 0 到节点 X
 * 的红色边和蓝色边交替出现的最短路径的长度。如果不存在这样的路径，那么 answer[x] = -1。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：n = 3, red_edges = [[0,1],[1,2]], blue_edges = []
 * 输出：[0,1,-1]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：n = 3, red_edges = [[0,1]], blue_edges = [[2,1]]
 * 输出：[0,1,-1]
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：n = 3, red_edges = [[1,0]], blue_edges = [[2,1]]
 * 输出：[0,-1,-1]
 * 
 * 
 * 示例 4：
 * 
 * 
 * 输入：n = 3, red_edges = [[0,1]], blue_edges = [[1,2]]
 * 输出：[0,1,2]
 * 
 * 
 * 示例 5：
 * 
 * 
 * 输入：n = 3, red_edges = [[0,1],[0,2]], blue_edges = [[1,0]]
 * 输出：[0,1,1]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 100
 * red_edges.length <= 400
 * blue_edges.length <= 400
 * red_edges[i].length == blue_edges[i].length == 2
 * 0 <= red_edges[i][j], blue_edges[i][j] < n
 * 
 * 
 */

public class Solution
{
    public int[] ShortestAlternatingPaths(int n, int[][] redEdges, int[][] blueEdges)
    {
        var G = new List<(int, bool)>[n];
        for (var i = 0; i < n; i++)
        {
            G[i] = new();
        }
        foreach (var re in redEdges)
        {
            G[re[0]].Add((re[1], true));
        }
        foreach (var be in blueEdges)
        {
            G[be[0]].Add((be[1], false));
        }
        var ans = new int[n];
        for (var i = 0; i < n; i++)
        {
            ans[i] = int.MaxValue;
        }
        var visit = new bool[n, 2];
        var Q = new Queue<(int, int, bool)>();
        Q.Enqueue((0, 0, true));
        Q.Enqueue((0, 0, false));
        ans[0] = 0;
        visit[0, 0] = true;
        visit[0, 1] = true;
        while (Q.Count > 0)
        {
            var (u, d, isRed) = Q.Dequeue();
            foreach (var (v, ir) in G[u])
            {
                if (ir == isRed || visit[v, ir ? 1 : 0])
                {
                    continue;
                }
                Q.Enqueue((v, d + 1, ir));
                ans[v] = Math.Min(ans[v], d + 1);
                visit[v, ir ? 1 : 0] = true;
            }
        }
        for (var i = 0; i < n; i++)
        {
            if (ans[i] == int.MaxValue)
            {
                ans[i] = -1;
            }
        }
        return ans;
    }
}
