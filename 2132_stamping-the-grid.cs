/*
 * @lc app=leetcode.cn id=stamping-the-grid lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [2132] 用邮票贴满网格图
 *
 * https://leetcode.cn/problems/stamping-the-grid/description/
 *
 * algorithms
 * Hard (31.91%)
 * Total Accepted:    4.8K
 * Total Submissions: 12.2K
 * Testcase Example:  '[[1,0,0,0],[1,0,0,0],[1,0,0,0],[1,0,0,0],[1,0,0,0]]\n4\n3'
 *
 * 给你一个 m x n 的二进制矩阵 grid ，每个格子要么为 0 （空）要么为 1 （被占据）。
 * 
 * 给你邮票的尺寸为 stampHeight x stampWidth 。我们想将邮票贴进二进制矩阵中，且满足以下 限制 和 要求 ：
 * 
 * 
 * 覆盖所有 空 格子。
 * 不覆盖任何 被占据 的格子。
 * 我们可以放入任意数目的邮票。
 * 邮票可以相互有 重叠 部分。
 * 邮票不允许 旋转 。
 * 邮票必须完全在矩阵 内 。
 * 
 * 
 * 如果在满足上述要求的前提下，可以放入邮票，请返回 true ，否则返回 false 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 输入：grid = [[1,0,0,0],[1,0,0,0],[1,0,0,0],[1,0,0,0],[1,0,0,0]], stampHeight =
 * 4, stampWidth = 3
 * 输出：true
 * 解释：我们放入两个有重叠部分的邮票（图中标号为 1 和 2），它们能覆盖所有与空格子。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 
 * 输入：grid = [[1,0,0,0],[0,1,0,0],[0,0,1,0],[0,0,0,1]], stampHeight = 2,
 * stampWidth = 2 
 * 输出：false 
 * 解释：没办法放入邮票覆盖所有的空格子，且邮票不超出网格图以外。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * m == grid.length
 * n == grid[r].length
 * 1 <= m, n <= 10^5
 * 1 <= m * n <= 2 * 10^5
 * grid[r][c] 要么是 0 ，要么是 1 。
 * 1 <= stampHeight, stampWidth <= 10^5
 * 
 * 
 */
// 参考: https://leetcode.cn/problems/stamping-the-grid/solutions/1199642/wu-nao-zuo-fa-er-wei-qian-zhui-he-er-wei-zwiu/?envType=daily-question&envId=2023-12-14
public class Solution
{
    public bool PossibleToStamp(int[][] G, int h, int w)
    {
        var (n, m) = (G.Length, G[0].Length);
        var S = new int[n + 1][];
        for (var i = 0; i < n + 1; i++) { S[i] = new int[m + 1]; }
        for (var i = 0; i < n; i++)
        {
            for (var j = 0; j < m; j++)
            {
                S[i + 1][j + 1] = S[i + 1][j] + S[i][j + 1] - S[i][j] + G[i][j];
            }
        }
        var d = new int[n + 2][];
        for (var i = 0; i < n + 2; i++) { d[i] = new int[m + 2]; }
        for (var i = h - 1; i < n; i++)
        {
            for (var j = w - 1; j < m; j++)
            {
                var (x, y) = (i - h + 1, j - w + 1);
                if (S[i + 1][j + 1] - S[i + 1][y] - S[x][j + 1] + S[x][y] != 0)
                { continue; }
                d[x + 1][y + 1]++;
                d[i + 2][y + 1]--;
                d[x + 1][j + 2]--;
                d[i + 2][j + 2]++;
            }
        }
        for (var i = 0; i < n; i++)
        {
            for (var j = 0; j < m; j++)
            {
                d[i + 1][j + 1] += d[i + 1][j] + d[i][j + 1] - d[i][j];
                if ((G[i][j], d[i + 1][j + 1]) == (0, 0)) { return false; }
            }
        }
        return true;
    }
}
