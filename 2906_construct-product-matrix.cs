/*
 * Medium
 * 
 * 给你一个下标从 0 开始、大小为 n * m 的二维整数矩阵 grid ，定义一个下标从 0 开始、大小为 n * m 的的二维矩阵 p。如果满足以下条件，则称 p 为 grid 的 乘积矩阵 ：
 * 
 * 对于每个元素 p[i][j] ，它的值等于除了 grid[i][j] 外所有元素的乘积。乘积对 12345 取余数。
 * 返回 grid 的乘积矩阵。
 * 
 *  
 * 
 * 示例 1：
 * 
 * 输入：grid = [[1,2],[3,4]]
 * 输出：[[24,12],[8,6]]
 * 解释：p[0][0] = grid[0][1] * grid[1][0] * grid[1][1] = 2 * 3 * 4 = 24
 * p[0][1] = grid[0][0] * grid[1][0] * grid[1][1] = 1 * 3 * 4 = 12
 * p[1][0] = grid[0][0] * grid[0][1] * grid[1][1] = 1 * 2 * 4 = 8
 * p[1][1] = grid[0][0] * grid[0][1] * grid[1][0] = 1 * 2 * 3 = 6
 * 所以答案是 [[24,12],[8,6]] 。
 * 示例 2：
 * 
 * 输入：grid = [[12345],[2],[1]]
 * 输出：[[2],[0],[0]]
 * 解释：p[0][0] = grid[0][1] * grid[0][2] = 2 * 1 = 2
 * p[0][1] = grid[0][0] * grid[0][2] = 12345 * 1 = 12345. 12345 % 12345 = 0 ，所以 p[0][1] = 0
 * p[0][2] = grid[0][0] * grid[0][1] = 12345 * 2 = 24690. 24690 % 12345 = 0 ，所以 p[0][2] = 0
 * 所以答案是 [[2],[0],[0]] 。
 *  
 * 
 * 提示：
 * 
 * 1 <= n == grid.length <= 105
 * 1 <= m == grid[i].length <= 105
 * 2 <= n * m <= 105
 * 1 <= grid[i][j] <= 109
 * 通过次数 2.7K
 * 提交次数 8.2K
 * 通过率 32.5%
 * 
 */

// 参考: https://leetcode.cn/problems/construct-product-matrix/solutions/2483137/zhou-sai-chang-kao-qian-hou-zhui-fen-jie-21hr/
//
// 前后缀分解板子题.
public class Solution
{
    public int[][] ConstructProductMatrix(int[][] M)
    {
        const long Mod = 12345;
        var (n, m) = (M.Length, M[0].Length);
        var ans = new int[n][];
        for (var i = 0; i < n; i++) { ans[i] = new int[m]; }
        var prefixes = new long[n * m + 1];
        var suffixes = new long[n * m + 1];
        (prefixes[0], suffixes[^1]) = (1, 1);
        for (var i = 0; i < n * m; i++)
        {
            var (x, y) = (i / m, i % m);
            prefixes[i + 1] = prefixes[i] * (long)M[x][y] % Mod;
        }
        for (var i = n * m - 1; i >= 0; i--)
        {
            var (x, y) = (i / m, i % m);
            suffixes[i] = suffixes[i + 1] * (long)M[x][y] % Mod;
        }
        for (var i = 0; i < n * m; i++)
        {
            var (x, y) = (i / m, i % m);
            ans[x][y] = (int)(prefixes[i] * suffixes[i + 1] % Mod);
        }
        return ans;
    }
}
