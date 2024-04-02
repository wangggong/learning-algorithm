/*
 * @lc app=leetcode.cn id=minimum-time-to-visit-a-cell-in-a-grid lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6366] 在网格图中访问一个格子的最少时间
 *
 * https://leetcode.cn/problems/minimum-time-to-visit-a-cell-in-a-grid/description/
 *
 * algorithms
 * Hard (31.36%)
 * Total Accepted:    1.6K
 * Total Submissions: 4.9K
 * Testcase Example:  '[[0,1,3,2],[5,1,2,5],[4,3,8,6]]'
 *
 * 给你一个 m x n 的矩阵 grid ，每个元素都为 非负 整数，其中 grid[row][col] 表示可以访问格子 (row, col) 的 最早
 * 时间。也就是说当你访问格子 (row, col) 时，最少已经经过的时间为 grid[row][col] 。
 * 
 * 你从 最左上角 出发，出发时刻为 0 ，你必须一直移动到上下左右相邻四个格子中的 任意 一个格子（即不能停留在格子上）。每次移动都需要花费 1
 * 单位时间。
 * 
 * 请你返回 最早 到达右下角格子的时间，如果你无法到达右下角的格子，请你返回 -1 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 
 * 输入：grid = [[0,1,3,2],[5,1,2,5],[4,3,8,6]]
 * 输出：7
 * 解释：一条可行的路径为：
 * - 时刻 t = 0 ，我们在格子 (0,0) 。
 * - 时刻 t = 1 ，我们移动到格子 (0,1) ，可以移动的原因是 grid[0][1] <= 1 。
 * - 时刻 t = 2 ，我们移动到格子 (1,1) ，可以移动的原因是 grid[1][1] <= 2 。
 * - 时刻 t = 3 ，我们移动到格子 (1,2) ，可以移动的原因是 grid[1][2] <= 3 。
 * - 时刻 t = 4 ，我们移动到格子 (1,1) ，可以移动的原因是 grid[1][1] <= 4 。
 * - 时刻 t = 5 ，我们移动到格子 (1,2) ，可以移动的原因是 grid[1][2] <= 5 。
 * - 时刻 t = 6 ，我们移动到格子 (1,3) ，可以移动的原因是 grid[1][3] <= 6 。
 * - 时刻 t = 7 ，我们移动到格子 (2,3) ，可以移动的原因是 grid[2][3] <= 7 。
 * 最终到达时刻为 7 。这是最早可以到达的时间。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 
 * 
 * 输入：grid = [[0,2,4],[3,2,1],[1,0,4]]
 * 输出：-1
 * 解释：没法从左上角按题目规定走到右下角。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * m == grid.length
 * n == grid[i].length
 * 2 <= m, n <= 1000
 * 4 <= m * n <= 10^5
 * 0 <= grid[i][j] <= 10^5
 * grid[0][0] == 0
 * 
 * 
 */
// 参考: https://leetcode.cn/problems/minimum-time-to-visit-a-cell-in-a-grid/solution/er-fen-da-an-bfspythonjavacgo-by-endless-j10w/
public class Solution
{
    private const int MaxN = (int)1e6;
    private (int, int)[] Directions = new[]
    {
        (0, 1),
        (0, -1),
        (1, 0),
        (-1, 0),
    };
    public int MinimumTime(int[][] G)
    {
        var n = G.Length;
        var m = G[0].Length;
        if (G[0][1] > 1 && G[1][0] > 1)
        {
            return -1;
        }
        bool canReach(int k)
        {
            if (k < m + n - 2 || k < G[n - 1][m - 1])
            {
                return false;
            }
            var Q = new Queue<(int, int)>();
            var visit = new bool[n][];
            for (var i = 0; i < n; i++)
            {
                visit[i] = new bool[m];
            }
            Q.Enqueue((n - 1, m - 1));
            visit[n - 1][m - 1] = true;
            for (var t = k - 1; Q.Count > 0; t--)
            {
                for (var count = Q.Count; count > 0; count--)
                {
                    var (x, y) = Q.Dequeue();
                    foreach (var (dx, dy) in Directions)
                    {
                        var (nx, ny) = (x + dx, y + dy);
                        if (0 <= nx && nx < n && 0 <= ny && ny < m && !visit[nx][ny] && G[nx][ny] <= t)
                        {
                            if ((nx, ny) is (0, 0))
                            {
                                return true;
                            }
                            Q.Enqueue((nx, ny));
                            visit[nx][ny] = true;
                        }
                    }
                }
            }
            return false;
        }
        var p = 0;
        var q = MaxN;
        while (p < q)
        {
            var mid = (p + q) >> 1;
            if (canReach(mid))
            {
                q = mid;
            }
            else
            {
                p = mid + 1;
            }
        }
        return p + (m + n + p) % 2;
    }
}
