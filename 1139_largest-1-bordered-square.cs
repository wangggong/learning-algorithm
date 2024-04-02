/*
 * @lc app=leetcode.cn id=largest-1-bordered-square lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1139] 最大的以 1 为边界的正方形
 *
 * https://leetcode.cn/problems/largest-1-bordered-square/description/
 *
 * algorithms
 * Medium (49.49%)
 * Total Accepted:    24K
 * Total Submissions: 44K
 * Testcase Example:  '[[1,1,1],[1,0,1],[1,1,1]]'
 *
 * 给你一个由若干 0 和 1 组成的二维网格 grid，请你找出边界全部由 1 组成的最大 正方形 子网格，并返回该子网格中的元素数量。如果不存在，则返回
 * 0。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：grid = [[1,1,1],[1,0,1],[1,1,1]]
 * 输出：9
 * 
 * 
 * 示例 2：
 * 
 * 输入：grid = [[1,1,0,0]]
 * 输出：1
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= grid.length <= 100
 * 1 <= grid[0].length <= 100
 * grid[i][j] 为 0 或 1
 * 
 * 
 */
public class Solution
{
    public int Largest1BorderedSquare(int[][] G)
    {
        var n = G.Length;
        var m = G[0].Length;
        var S = new int[n + 1][];
        for (var i = 0; i <= n; i++)
        {
            S[i] = new int[m + 1];
        }
        for (var i = 0; i < n; i++)
        {
            for (var j = 0; j < m; j++)
            {
                S[i + 1][j + 1] = S[i + 1][j] + S[i][j + 1] - S[i][j] + G[i][j];
            }
        }
        var ans = 0;
        for (var i = 0; i < n; i++)
        {
            for (var j = 0; j < m; j++)
            {
                if (G[i][j] == 0)
                {
                    continue;
                }
                for (var k = 1; i + k <= n && j + k <= m; k++)
                {
                    if (S[i + 1][j + k] - S[i][j + k] - S[i + 1][j] + S[i][j] == k
                            && S[i + k][j + 1] - S[i + k][j] - S[i][j + 1] + S[i][j] == k
                            && S[i + k][j + k] - S[i + k - 1][j + k] - S[i + k][j] + S[i + k - 1][j] == k
                            && S[i + k][j + k] - S[i + k][j + k - 1] - S[i][j + k] + S[i][j + k - 1] == k)
                    {
                        ans = Math.Max(ans, k * k);
                    }
                }
            }
        }
        return ans;
    }
}
