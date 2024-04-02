/*
 * @lc app=leetcode.cn id=minimum-time-to-complete-all-tasks lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6318] 完成所有任务的最少时间
 *
 * https://leetcode.cn/problems/minimum-time-to-complete-all-tasks/description/
 *
 * algorithms
 * Hard (31.50%)
 * Total Accepted:    1.6K
 * Total Submissions: 5.1K
 * Testcase Example:  '[[2,3,1],[4,5,1],[1,5,2]]'
 *
 * 你有一台电脑，它可以 同时 运行无数个任务。给你一个二维整数数组 tasks ，其中 tasks[i] = [starti, endi,
 * durationi] 表示第 i 个任务需要在 闭区间 时间段 [starti, endi] 内运行 durationi
 * 个整数时间点（但不需要连续）。
 * 
 * 当电脑需要运行任务时，你可以打开电脑，如果空闲时，你可以将电脑关闭。
 * 
 * 请你返回完成所有任务的情况下，电脑最少需要运行多少秒。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：tasks = [[2,3,1],[4,5,1],[1,5,2]]
 * 输出：2
 * 解释：
 * - 第一个任务在闭区间 [2, 2] 运行。
 * - 第二个任务在闭区间 [5, 5] 运行。
 * - 第三个任务在闭区间 [2, 2] 和 [5, 5] 运行。
 * 电脑总共运行 2 个整数时间点。
 * 
 * 
 * 示例 2：
 * 
 * 输入：tasks = [[1,3,2],[2,5,3],[5,6,2]]
 * 输出：4
 * 解释：
 * - 第一个任务在闭区间 [2, 3] 运行
 * - 第二个任务在闭区间 [2, 3] 和 [5, 5] 运行。
 * - 第三个任务在闭区间 [5, 6] 运行。
 * 电脑总共运行 4 个整数时间点。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= tasks.length <= 2000
 * tasks[i].length == 3
 * 1 <= starti, endi <= 2000
 * 1 <= durationi <= endi - starti + 1 
 * 
 * 
 */

/*
 * // 参考: https://leetcode.cn/problems/minimum-time-to-complete-all-tasks/solution/tan-xin-pythonjavacgo-by-endlesscheng-w3k3/
 * // 想了半天左端点排序...
 * public class Solution
 * {
 *     private const int N = 2000;
 * 
 *     public int FindMinimumTime(int[][] tasks)
 *     {
 *         var visit = new bool[N + 1];
 *         foreach (var task in tasks.OrderBy(x => x[1]))
 *         {
 *             var (l, r, d) = (task[0], task[1], task[2]);
 *             for (var i = r; i >= l && d > 0; i--)
 *             {
 *                 if (visit[i])
 *                 {
 *                     d--;
 *                 }
 *             }
 *             if (d > 0)
 *             {
 *                 for (var i = r; i >= l && d > 0; i--)
 *                 {
 *                     if (!visit[i])
 *                     {
 *                         visit[i] = true;
 *                         d--;
 *                     }
 *                 }
 *             }
 *         }
 *         return visit.Where(x => x).Count();
 *     }
 * }
 */

// 参考: https://leetcode.cn/problems/minimum-time-to-complete-all-tasks/solution/chai-fen-yue-shu-by-tsreaper-swhm/
// 奇怪的 SPFA. 建图逻辑:
// 1. `a_{i + 1} - a_{i} >= 0`
// 2. `a_{i + 1} - a_{i} <= 1` => `a_{i} - a_{i + 1} + 1 >= 0`
// 3. `a_{end} - a_{start - 1} >= d` => `a_{end} - a_{start - 1} - d >= 0`
public class Solution
{
    public int FindMinimumTime(int[][] tasks)
    {
        // 建图
        var max = tasks.Select(x => x[1]).Max();
        // 超级源点
        var src = max + 1;
        var G = new List<(int, int)>[src + 1];
        for (var i = 0; i <= src; i++)
        {
            G[i] = new();
        }
        for (var i = 1; i <= max; i++)
        {
            G[i].Add((i - 1, 0));
        }
        for (var i = 0; i < max; i++)
        {
            G[i].Add((i + 1, 1));
        }
        foreach (var task in tasks)
        {
            var (l, r, d) = (task[0], task[1], task[2]);
            G[r].Add((l - 1, -d));
        }
        for (var i = 0; i <= max; i++)
        {
            G[src].Add((i, 0));
        }
        // SPFA 求最短路
        var Q = new Queue<int>();
        var visit = new bool[src + 1];
        var dist = new int[src + 1];
        for (var i = 0; i <= src; i++)
        {
            dist[i] = int.MaxValue;
        }
        Q.Enqueue(src);
        visit[src] = true;
        dist[src] = 0;
        while (Q.Count > 0)
        {
            var u = Q.Dequeue();
            visit[u] = false;
            foreach (var (v, d) in G[u])
            {
                if (dist[v] > dist[u] + d)
                {
                    dist[v] = dist[u] + d;
                    if (!visit[v])
                    {
                        Q.Enqueue(v);
                        visit[v] = true;
                    }
                }
            }
        }
        // 结果的含义: `max{a_i} - min{a_i}`
        return dist[max] - dist[0];
    }
}
