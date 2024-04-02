/*
 * @lc app=leetcode.cn id=minimum-edge-reversals-so-every-node-is-reachable lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [100041] 可以到达每一个节点的最少边反转次数
 *
 * https://leetcode.cn/problems/minimum-edge-reversals-so-every-node-is-reachable/description/
 *
 * algorithms
 * Hard (58.98%)
 * Total Accepted:    563
 * Total Submissions: 949
 * Testcase Example:  '4\n[[2,0],[2,1],[1,3]]'
 *
 * 给你一个 n 个点的 简单有向图 （没有重复边的有向图），节点编号为 0 到 n - 1 。如果这些边是双向边，那么这个图形成一棵 树 。
 * 
 * 给你一个整数 n 和一个 二维 整数数组 edges ，其中 edges[i] = [ui, vi] 表示从节点 ui 到节点 vi 有一条 有向边
 * 。
 * 
 * 边反转 指的是将一条边的方向反转，也就是说一条从节点 ui 到节点 vi 的边会变为一条从节点 vi 到节点 ui 的边。
 * 
 * 对于范围 [0, n - 1] 中的每一个节点 i ，你的任务是分别 独立 计算 最少 需要多少次 边反转 ，从节点 i 出发经过 一系列有向边
 * ，可以到达所有的节点。
 * 
 * 请你返回一个长度为 n 的整数数组 answer ，其中 answer[i]表示从节点 i 出发，可以到达所有节点的 最少边反转 次数。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 
 * 输入：n = 4, edges = [[2,0],[2,1],[1,3]]
 * 输出：[1,1,0,2]
 * 解释：上图表示了与输入对应的简单有向图。
 * 对于节点 0 ：反转 [2,0] ，从节点 0 出发可以到达所有节点。
 * 所以 answer[0] = 1 。
 * 对于节点 1 ：反转 [2,1] ，从节点 1 出发可以到达所有节点。
 * 所以 answer[1] = 1 。
 * 对于节点 2 ：不需要反转就可以从节点 2 出发到达所有节点。
 * 所以 answer[2] = 0 。
 * 对于节点 3 ：反转 [1,3] 和 [2,1] ，从节点 3 出发可以到达所有节点。
 * 所以 answer[3] = 2 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 
 * 
 * 输入：n = 3, edges = [[1,2],[2,0]]
 * 输出：[2,0,1]
 * 解释：上图表示了与输入对应的简单有向图。
 * 对于节点 0 ：反转 [2,0] 和 [1,2] ，从节点 0 出发可以到达所有节点。
 * 所以 answer[0] = 2 。
 * 对于节点 1 ：不需要反转就可以从节点 2 出发到达所有节点。
 * 所以 answer[1] = 0 。
 * 对于节点 2 ：反转 [1,2] ，从节点 2 出发可以到达所有节点。
 * 所以 answer[2] = 1 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 2 <= n <= 10^5
 * edges.length == n - 1
 * edges[i].length == 2
 * 0 <= ui == edges[i][0] < n
 * 0 <= vi == edges[i][1] < n
 * ui != vi
 * 输入保证如果边是双向边，可以得到一棵树。
 * 
 * 
 */
public class Solution
{
    public int[] MinEdgeReversals(int n, int[][] edges)
    {
        var G = new List<(int, bool)>[n];
        for (var i = 0; i < n; i++) { G[i] = new(); }
        foreach (var edge in edges)
        {
            var (u, v) = (edge[0], edge[1]);
            G[u].Add((v, true));
            G[v].Add((u, false));
        }
        var counts = new int[n];
        var reals = new int[n];
        void dfs(int u, int p, Action<int, int, bool> pre, Action<int, int, bool> post)
        {
            foreach (var (v, realEdge) in G[u])
            {
                if (v == p) { continue; }
                if (pre is not null) { pre(u, v, realEdge); }
                dfs(v, u, pre, post);
                if (post is not null) { post(u, v, realEdge); }
            }
        }
        dfs(0, -1, null, (u, v, realEdge) => 
        {
            counts[u] += counts[v];
            reals[u] += reals[v];
            counts[u]++;
            if (realEdge) { reals[u]++; }
        });
        var ans = new int[n];
        ans[0] = counts[0] - reals[0];
        dfs(0, -1, (u, v, realEdge) => ans[v] = ans[u] + (realEdge ? 1 : -1), null);
        return ans;
    }
}
