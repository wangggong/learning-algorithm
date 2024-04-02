/*
 * @lc app=leetcode.cn id=check-knight-tour-configuration lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6322] 检查骑士巡视方案
 *
 * https://leetcode.cn/problems/check-knight-tour-configuration/description/
 *
 * algorithms
 * Medium (53.15%)
 * Total Accepted:    4.2K
 * Total Submissions: 7.9K
 * Testcase Example:  '[[0,11,16,5,20],[17,4,19,10,15],[12,1,8,21,6],[3,18,23,14,9],[24,13,2,7,22]]'
 *
 * 骑士在一张 n x n 的棋盘上巡视。在有效的巡视方案中，骑士会从棋盘的 左上角 出发，并且访问棋盘上的每个格子 恰好一次 。
 * 
 * 给你一个 n x n 的整数矩阵 grid ，由范围 [0, n * n - 1] 内的不同整数组成，其中 grid[row][col] 表示单元格
 * (row, col) 是骑士访问的第 grid[row][col] 个单元格。骑士的行动是从下标 0 开始的。
 * 
 * 如果 grid 表示了骑士的有效巡视方案，返回 true；否则返回 false。
 * 
 * 注意，骑士行动时可以垂直移动两个格子且水平移动一个格子，或水平移动两个格子且垂直移动一个格子。下图展示了骑士从某个格子出发可能的八种行动路线。
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：grid =
 * [[0,11,16,5,20],[17,4,19,10,15],[12,1,8,21,6],[3,18,23,14,9],[24,13,2,7,22]]
 * 输出：true
 * 解释：grid 如上图所示，可以证明这是一个有效的巡视方案。
 * 
 * 
 * 示例 2：
 * 
 * 输入：grid = [[0,3,6],[5,8,1],[2,7,4]]
 * 输出：false
 * 解释：grid 如上图所示，考虑到骑士第 7 次行动后的位置，第 8 次行动是无效的。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * n == grid.length == grid[i].length
 * 3 <= n <= 7
 * 0 <= grid[row][col] < n * n
 * grid 中的所有整数 互不相同
 * 
 * 
 */

/*
 * public class Solution
 * {
 *     public bool CheckValidGrid(int[][] G)
 *     {
 *         var n = G.Length;
 *         var m = G[0].Length;
 *         var positions = new (int, int)[n * m];
 *         for (var i = 0; i < n; i++)
 *         {
 *             for (var j = 0; j < m; j++)
 *             {
 *                 positions[G[i][j]] = (i, j);
 *             }
 *         }
 *         if ((0, 0) != positions[0])
 *         {
 *             return false;
 *         }
 *         for (var i = 1; i < n * m; i++)
 *         {
 *             var (nx, ny) = positions[i];
 *             var (px, py) = positions[i - 1];
 *             var dx = Math.Abs(nx - px);
 *             var dy = Math.Abs(ny - py);
 *             if (!((dx == 1 && dy == 2) || (dx == 2 && dy == 1)))
 *             {
 *                 return false;
 *             }
 *         }
 *         return true;
 *     }
 * }
 */

public class Solution
{
    public bool CheckValidGrid(int[][] grid)
    {
        var positions = grid.SelectMany((row, x) => row
                .Select((v, y) => (v, x, y)))
            .OrderBy(x => x.v)
            .Select(x => (x.x, x.y))
            .ToList();
        return positions.First() is (0, 0) && Enumerable
            .Range(0, positions.Count() - 1)
            .All(i => (Math.Abs(positions[i + 1].x - positions[i].x),
                    Math.Abs(positions[i + 1].y - positions[i].y))
                is (1, 2) or (2, 1));
    }
}
