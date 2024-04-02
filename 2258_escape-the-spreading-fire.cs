/*
 * @lc app=leetcode.cn id=escape-the-spreading-fire lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [2258] 逃离火灾
 *
 * https://leetcode.cn/problems/escape-the-spreading-fire/description/
 *
 * algorithms
 * Hard (36.99%)
 * Total Accepted:    13.5K
 * Total Submissions: 26K
 * Testcase Example:  '[[0,2,0,0,0,0,0],[0,0,0,2,2,1,0],[0,2,0,0,1,2,0],[0,0,2,2,2,0,2],[0,0,0,0,0,0,0]]'
 *
 * 给你一个下标从 0 开始大小为 m x n 的二维整数数组 grid ，它表示一个网格图。每个格子为下面 3 个值之一：
 * 
 * 
 * 0 表示草地。
 * 1 表示着火的格子。
 * 2 表示一座墙，你跟火都不能通过这个格子。
 * 
 * 
 * 一开始你在最左上角的格子 (0, 0) ，你想要到达最右下角的安全屋格子 (m - 1, n - 1) 。每一分钟，你可以移动到 相邻
 * 的草地格子。每次你移动 之后 ，着火的格子会扩散到所有不是墙的 相邻 格子。
 * 
 * 请你返回你在初始位置可以停留的 最多 分钟数，且停留完这段时间后你还能安全到达安全屋。如果无法实现，请你返回 -1 。如果不管你在初始位置停留多久，你
 * 总是 能到达安全屋，请你返回 10^9 。
 * 
 * 注意，如果你到达安全屋后，火马上到了安全屋，这视为你能够安全到达安全屋。
 * 
 * 如果两个格子有共同边，那么它们为 相邻 格子。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 输入：grid =
 * [[0,2,0,0,0,0,0],[0,0,0,2,2,1,0],[0,2,0,0,1,2,0],[0,0,2,2,2,0,2],[0,0,0,0,0,0,0]]
 * 输出：3
 * 解释：上图展示了你在初始位置停留 3 分钟后的情形。
 * 你仍然可以安全到达安全屋。
 * 停留超过 3 分钟会让你无法安全到达安全屋。
 * 
 * 示例 2：
 * 
 * 
 * 
 * 输入：grid = [[0,0,0,0],[0,1,2,0],[0,2,0,0]]
 * 输出：-1
 * 解释：上图展示了你马上开始朝安全屋移动的情形。
 * 火会蔓延到你可以移动的所有格子，所以无法安全到达安全屋。
 * 所以返回 -1 。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 
 * 输入：grid = [[0,0,0],[2,2,0],[1,2,0]]
 * 输出：1000000000
 * 解释：上图展示了初始网格图。
 * 注意，由于火被墙围了起来，所以无论如何你都能安全到达安全屋。
 * 所以返回 10^9 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * m == grid.length
 * n == grid[i].length
 * 2 <= m, n <= 300
 * 4 <= m * n <= 2 * 10^4
 * grid[i][j] 是 0 ，1 或者 2 。
 * grid[0][0] == grid[m - 1][n - 1] == 0
 * 
 * 
 */
public class Solution
{
    public int MaximumMinutes(int[][] G)
    {
        const int Maxn = (int)1e9;
        var directions = new (int, int)[]
        {
            (0, 1),
            (1, 0),
            (0, -1),
            (-1, 0),
        };
        var (n, m) = (G.Length, G[0].Length);
        bool onGrid(int x, int y) => 0 <= x && x < n && 0 <= y && y < m && G[x][y] is 0;
        bool check(int k)
        {
            var Qf = new Queue<(int, int)>();
            var onFire = Enumerable
                .Range(0, n)
                .Select(_ => new bool[m])
                .ToArray();
            void spread()
            {
                for (var c = Qf.Count; c > 0; c--)
                {
                    var (x, y) = Qf.Dequeue();
                    foreach (var (dx, dy) in directions)
                    {
                        var (nx, ny) = (x + dx, y + dy);
                        if (!(onGrid(nx, ny) && !onFire[nx][ny])) { continue; }
                        Qf.Enqueue((nx, ny));
                        onFire[nx][ny] = true;
                    }
                }
            }
            for (var i = 0; i < n; i++)
            {
                for (var j = 0; j < m; j++)
                {
                    if (G[i][j] is not 1) { continue; }
                    Qf.Enqueue((i, j));
                    onFire[i][j] = true;
                }
            }
            while (k-- > 0) { spread(); }
            var Qp = new Queue<(int, int)>();
            var visits = Enumerable
                .Range(0, n)
                .Select(_ => new bool[m])
                .ToArray();
            Qp.Enqueue((0, 0));
            visits[0][0] = true;
            for (; Qp.Count > 0; spread())
            {
                for (var c = Qp.Count; c > 0; c--)
                {
                    var (x, y) = Qp.Dequeue();
                    if ((x, y) == (n - 1, m - 1)) { return true; }
                    if (onFire[x][y]) { continue; }
                    foreach (var (dx, dy) in directions)
                    {
                        var (nx, ny) = (x + dx, y + dy);
                        if (!(onGrid(nx, ny) && !onFire[nx][ny] && !visits[nx][ny])) { continue; }
                        Qp.Enqueue((nx, ny));
                        visits[nx][ny] = true;
                    }
                }
            }
            return false;
        }
        var (p, q) = (0, n * m + 1);
        while (p < q)
        {
            var mid = (p + q) >> 1;
            if (!check(mid)) { q = mid; }
            else { p = mid + 1; }
        }
        return p > m * n ? Maxn : p - 1;
    }
}
