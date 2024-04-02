/*
 * @lc app=leetcode.cn id=disconnect-path-in-a-binary-matrix-by-at-most-one-flip lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6305] 二进制矩阵中翻转最多一次使路径不连通
 *
 * https://leetcode.cn/problems/disconnect-path-in-a-binary-matrix-by-at-most-one-flip/description/
 *
 * algorithms
 * Medium (25.34%)
 * Total Accepted:    1.1K
 * Total Submissions: 4.1K
 * Testcase Example:  '[[1,1,1],[1,0,0],[1,1,1]]'
 *
 * 给你一个下标从 0 开始的 m x n 二进制 矩阵 grid 。你可以从一个格子 (row, col) 移动到格子 (row + 1, col) 或者
 * (row, col + 1) ，前提是前往的格子值为 1 。如果从 (0, 0) 到 (m - 1, n - 1) 没有任何路径，我们称该矩阵是 不连通
 * 的。
 * 
 * 你可以翻转 最多一个 格子的值（也可以不翻转）。你 不能翻转 格子 (0, 0) 和 (m - 1, n - 1) 。
 * 
 * 如果可以使矩阵不连通，请你返回 true ，否则返回 false 。
 * 
 * 注意 ，翻转一个格子的值，可以使它的值从 0 变 1 ，或从 1 变 0 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 
 * 输入：grid = [[1,1,1],[1,0,0],[1,1,1]]
 * 输出：true
 * 解释：按照上图所示我们翻转蓝色格子里的值，翻转后从 (0, 0) 到 (2, 2) 没有路径。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 
 * 
 * 输入：grid = [[1,1,1],[1,0,1],[1,1,1]]
 * 输出：false
 * 解释：无法翻转至多一个格子，使 (0, 0) 到 (2, 2) 没有路径。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * m == grid.length
 * n == grid[i].length
 * 1 <= m, n <= 1000
 * 1 <= m * n <= 10^5
 * grid[0][0] == grid[m - 1][n - 1] == 1
 * 
 * 
 */

/*
 * // 参考: https://leetcode.cn/problems/disconnect-path-in-a-binary-matrix-by-at-most-one-flip/solution/ji-shu-by-tsreaper-m01b/
 * //
 * // 转化为计数问题
 * public class Solution
 * {
 *     private const long Mod = (long)1e9 + 7;
 * 
 *     public bool IsPossibleToCutPath(int[][] G)
 *     {
 *         var n = G.Length;
 *         var m = G[0].Length;
 *         var CF = new long[n, m];
 *         var CG = new long[n, m];
 *         CF[0, 0] = 1;
 *         CG[n - 1, m - 1] = 1;
 *         for (int i = 0; i < n; i++)
 *         {
 *             for (int j = 0; j < m; j++)
 *             {
 *                 if (i == 0 && j == 0)
 *                 {
 *                     continue;
 *                 }
 *                 if (G[i][j] == 0)
 *                 {
 *                     continue;
 *                 }
 *                 if (i > 0 && G[i - 1][j] == 1)
 *                 {
 *                     CF[i, j] = (CF[i, j] + CF[i - 1, j]) % Mod;
 *                 }
 *                 if (j > 0 && G[i][j - 1] == 1)
 *                 {
 *                     CF[i, j] = (CF[i, j] + CF[i, j - 1]) % Mod;
 *                 }
 *             }
 *         }
 *         for (int i = n - 1; i >= 0; i--)
 *         {
 *             for (int j = m - 1; j >= 0; j--)
 *             {
 *                 if (i == n - 1 && j == m - 1)
 *                 {
 *                     continue;
 *                 }
 *                 if (G[i][j] == 0)
 *                 {
 *                     continue;
 *                 }
 *                 if (i < n - 1 && G[i + 1][j] == 1)
 *                 {
 *                     CG[i, j] = (CG[i, j] + CG[i + 1, j]) % Mod;
 *                 }
 *                 if (j < m - 1 && G[i][j + 1] == 1)
 *                 {
 *                     CG[i, j] = (CG[i, j] + CG[i, j + 1]) % Mod;
 *                 }
 *             }
 *         }
 *         var c = (CF[n - 1, m - 1] * CG[n - 1, m - 1]) % Mod;
 *         if (c == 0)
 *         {
 *             return true;
 *         }
 *         for (var i = 0; i < n; i++)
 *         {
 *             for (var j = 0; j < m; j++)
 *             {
 *                 if ((i, j) == (0, 0) || (i, j) == (n - 1, m - 1))
 *                 {
 *                     continue;
 *                 }
 *                 var v = (CF[i, j] * CG[i, j]) % Mod;
 *                 if (v == c)
 *                 {
 *                     return true;
 *                 }
 *             }
 *         }
 *         return false;
 *     }
 * }
 */

// 参考: https://leetcode.cn/problems/disconnect-path-in-a-binary-matrix-by-at-most-one-flip/solution/zhuan-huan-cheng-qiu-lun-kuo-shi-fou-xia-io8x/
public class Solution
{
    private int n;
    private int m;

    private bool dfs(int[][] G, int x, int y)
    {
        if ((x, y) == (n - 1, m - 1))
        {
            return true;
        }
        G[x][y] = 0;
        return (x + 1 < n && G[x + 1][y] == 1 && dfs(G, x + 1, y))
            || (y + 1 < m && G[x][y + 1] == 1 && dfs(G, x, y + 1));
    }

    public bool IsPossibleToCutPath(int[][] G)
    {
        n = G.Length;
        m = G[0].Length;
        if (n * m <= 2)
        {
            return false;
        }
        return n == 1 || m == 1
            || G[0][1] == 0 || G[1][0] == 0
            || G[n - 1][m - 2] == 0 || G[n - 2][m - 1] == 0
            || !dfs(G, 1, 0) || !dfs(G, 0, 1);
    }
}
