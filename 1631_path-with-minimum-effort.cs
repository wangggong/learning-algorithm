/*
 * @lc app=leetcode.cn id=path-with-minimum-effort lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1631] 最小体力消耗路径
 *
 * https://leetcode.cn/problems/path-with-minimum-effort/description/
 *
 * algorithms
 * Medium (50.97%)
 * Total Accepted:    49.5K
 * Total Submissions: 95.9K
 * Testcase Example:  '[[1,2,2],[3,8,2],[5,3,5]]'
 *
 * 你准备参加一场远足活动。给你一个二维 rows x columns 的地图 heights ，其中 heights[row][col] 表示格子
 * (row, col) 的高度。一开始你在最左上角的格子 (0, 0) ，且你希望去最右下角的格子 (rows-1, columns-1) （注意下标从
 * 0 开始编号）。你每次可以往 上，下，左，右 四个方向之一移动，你想要找到耗费 体力 最小的一条路径。
 * 
 * 一条路径耗费的 体力值 是路径上相邻格子之间 高度差绝对值 的 最大值 决定的。
 * 
 * 请你返回从左上角走到右下角的最小 体力消耗值 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 
 * 输入：heights = [[1,2,2],[3,8,2],[5,3,5]]
 * 输出：2
 * 解释：路径 [1,3,5,3,5] 连续格子的差值绝对值最大为 2 。
 * 这条路径比路径 [1,2,2,2,5] 更优，因为另一条路径差值最大值为 3 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 
 * 
 * 输入：heights = [[1,2,3],[3,8,4],[5,3,5]]
 * 输出：1
 * 解释：路径 [1,2,3,4,5] 的相邻格子差值绝对值最大为 1 ，比路径 [1,3,5,3,5] 更优。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：heights = [[1,2,1,1,1],[1,2,1,2,1],[1,2,1,2,1],[1,2,1,2,1],[1,1,1,2,1]]
 * 输出：0
 * 解释：上图所示路径不需要消耗任何体力。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * rows == heights.length
 * columns == heights[i].length
 * 1 
 * 1 
 * 
 * 
 */
public class Solution
{
    public int MinimumEffortPath(int[][] heights)
    {
        const int Maxn = 1_000_000;
        var directions = new (int, int)[]
        {
            (0, 1),
            (1, 0),
            (0, -1),
            (-1, 0),
        };
        bool check(int k)
        {
            var (n, m) = (heights.Length, heights[0].Length);
            var Q = new Queue<int>();
            var visits = new bool[n * m];
            Q.Enqueue(0);
            visits[0] = true;
            while (Q.Count > 0)
            {
                var q = Q.Dequeue();
                if (q == n * m - 1) { return true; }
                var (x, y) = (q / m, q % m);
                foreach (var (dx, dy) in directions)
                {
                    var (nx, ny) = (x + dx, y + dy);
                    var nq = nx * m + ny;
                    if (!(0 <= nx && nx < n
                        && 0 <= ny && ny < m
                        && !visits[nq]
                        && Math.Abs(heights[nx][ny] - heights[x][y]) <= k))
                    { continue; }
                    Q.Enqueue(nq);
                    visits[nq] = true;
                }
            }
            return false;
        }
        var (p, q) = (0, Maxn);
        while (p < q)
        {
            var mid = (p + q) >> 1;
            if (check(mid)) { q = mid; }
            else { p = mid + 1; }
        }
        return p;
    }
}
