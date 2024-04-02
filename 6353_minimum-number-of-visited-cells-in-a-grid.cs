/*
 * @lc app=leetcode.cn id=minimum-number-of-visited-cells-in-a-grid lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6353] 网格图中最少访问的格子数
 *
 * https://leetcode.cn/problems/minimum-number-of-visited-cells-in-a-grid/description/
 *
 * algorithms
 * Hard (31.87%)
 * Total Accepted:    1.8K
 * Total Submissions: 5.6K
 * Testcase Example:  '[[3,4,2,1],[4,2,3,1],[2,1,0,0],[2,4,0,0]]'
 *
 * 给你一个下标从 0 开始的 m x n 整数矩阵 grid 。你一开始的位置在 左上角 格子 (0, 0) 。
 * 
 * 当你在格子 (i, j) 的时候，你可以移动到以下格子之一：
 * 
 * 
 * 满足 j < k <= grid[i][j] + j 的格子 (i, k) （向右移动），或者
 * 满足 i < k <= grid[i][j] + i 的格子 (k, j) （向下移动）。
 * 
 * 
 * 请你返回到达 右下角 格子 (m - 1, n - 1) 需要经过的最少移动格子数，如果无法到达右下角格子，请你返回 -1 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 输入：grid = [[3,4,2,1],[4,2,3,1],[2,1,0,0],[2,4,0,0]]
 * 输出：4
 * 解释：上图展示了到达右下角格子经过的 4 个格子。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 
 * 输入：grid = [[3,4,2,1],[4,2,1,1],[2,1,1,0],[3,4,1,0]]
 * 输出：3
 * 解释：上图展示了到达右下角格子经过的 3 个格子。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 
 * 输入：grid = [[2,1,0],[1,0,0]]
 * 输出：-1
 * 解释：无法到达右下角格子。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * m == grid.length
 * n == grid[i].length
 * 1 <= m, n <= 10^5
 * 1 <= m * n <= 10^5
 * 0 <= grid[i][j] < m * n
 * grid[m - 1][n - 1] == 0
 * 
 * 
 */

/*
 * // 参考: https://leetcode.cn/problems/minimum-number-of-visited-cells-in-a-grid/solution/bfs-bing-cha-ji-by-xjliang-3ito/
 * // 
 * // 暴力 BFS, 但需要倒着枚举.
 * public class Solution
 * {
 *     public int MinimumVisitedCells(int[][] G)
 *     {
 *         var (n, m) = (G.Length, G[0].Length);
 *         var dist = new int[n][];
 *         for (var i = 0; i < n; i++)
 *         {
 *             dist[i] = new int[m];
 *             Array.Fill(dist[i], int.MaxValue);
 *         }
 *         var Q = new Queue<(int, int)>();
 *         Q.Enqueue((0, 0));
 *         for (var step = 1; Q.Count > 0; step++)
 *         {
 *             for (var c = Q.Count; c > 0; c--)
 *             {
 *                 var (x, y) = Q.Dequeue();
 *                 if ((x, y) == (n - 1, m - 1))
 *                 {
 *                     return step;
 *                 }
 *                 if (step >= dist[x][y])
 *                 {
 *                     continue;
 *                 }
 *                 dist[x][y] = step;
 *                 for (var i = G[x][y]; i > 0; i--)
 *                 {
 *                     var (nx, ny) = (x + i, y);
 *                     if (nx < n && dist[nx][ny] > step + 1)
 *                     {
 *                         Q.Enqueue((nx, ny));
 *                     }
 *                     (nx, ny) = (x, y + i);
 *                     if (ny < m && dist[nx][ny] > step + 1)
 *                     {
 *                         Q.Enqueue((nx, ny));
 *                     }
 *                 }
 *             }
 *         }
 *         return dist[n - 1][m - 1] >= int.MaxValue ? -1 : dist[n - 1][m - 1];
 *     }
 * }
 */

// 参考: https://leetcode.cn/problems/minimum-number-of-visited-cells-in-a-grid/solution/m-n-ge-you-xian-dui-lie-by-zerotrac2-d9rg/
// 
// 优先队列
public class Solution
{
    public int MinimumVisitedCells(int[][] G)
    {
        var (n, m) = (G.Length, G[0].Length);
        var dist = new int[n][];
        for (var i = 0; i < n; i++)
        {
            dist[i] = new int[m];
            Array.Fill(dist[i], int.MaxValue);
        }
        var Qs = new PriorityQueue<int, int>[n + m];
        for (var i = 0; i < n + m; i++)
        {
            Qs[i] = new();
        }
        dist[0][0] = 1;
        for (var i = 0; i < n; i++)
        {
            for (var j = 0; j < m; j++)
            {
                var (rowQ, colQ) = (Qs[i], Qs[j + n]);
                while (rowQ.Count > 0 && rowQ.Peek() + G[i][rowQ.Peek()] < j)
                {
                    rowQ.Dequeue();
                }
                if (rowQ.Count > 0)
                {
                    dist[i][j] = Math.Min(dist[i][j], dist[i][rowQ.Peek()] + 1);
                }
                while (colQ.Count > 0 && colQ.Peek() + G[colQ.Peek()][j] < i)
                {
                    colQ.Dequeue();
                }
                if (colQ.Count > 0)
                {
                    dist[i][j] = Math.Min(dist[i][j], dist[colQ.Peek()][j] + 1);
                }
                if (dist[i][j] < int.MaxValue)
                {
                    rowQ.Enqueue(j, dist[i][j]);
                    colQ.Enqueue(i, dist[i][j]);
                }
            }
        }
        return dist[n - 1][m - 1] >= int.MaxValue ? -1 : dist[n - 1][m - 1];
    }
}
