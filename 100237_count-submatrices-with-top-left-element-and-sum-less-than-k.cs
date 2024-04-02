/*
 * @lc app=leetcode.cn id=count-submatrices-with-top-left-element-and-sum-less-than-k lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [100237] 元素和小于等于 k 的子矩阵的数目
 *
 * https://leetcode.cn/problems/count-submatrices-with-top-left-element-and-sum-less-than-k/description/
 *
 * algorithms
 * Medium (58.07%)
 * Total Accepted:    2.7K
 * Total Submissions: 4.7K
 * Testcase Example:  '[[7,6,3],[6,6,1]]\n18'
 *
 * 给你一个下标从 0 开始的整数矩阵 grid 和一个整数 k。
 * 
 * 返回包含 grid 左上角元素、元素和小于或等于 k 的 子矩阵 的数目。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：grid = [[7,6,3],[6,6,1]], k = 18
 * 输出：4
 * 解释：如上图所示，只有 4 个子矩阵满足：包含 grid 的左上角元素，并且元素和小于或等于 18 。
 * 
 * 示例 2：
 * 
 * 
 * 输入：grid = [[7,2,9],[1,5,0],[2,6,6]], k = 20
 * 输出：6
 * 解释：如上图所示，只有 6 个子矩阵满足：包含 grid 的左上角元素，并且元素和小于或等于 20 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * m == grid.length 
 * n == grid[i].length
 * 1 <= n, m <= 1000 
 * 0 <= grid[i][j] <= 1000
 * 1 <= k <= 10^9
 * 
 * 
 */
public class Solution
{
    public int CountSubmatrices(int[][] G, int k)
    {
        var (n, m) = (G.Length, G[0].Length);
        var S = new int[n + 1][];
        for (var i = 0; i <= n; i++)
        {
            S[i] = new int[m + 1];
        }
        var ans = 0;
        for (var i = 0; i < n; i++)
        {
            for (var j = 0; j < m; j++)
            {
                S[i + 1][j + 1] = S[i + 1][j] + S[i][j + 1] - S[i][j] + G[i][j];
                if (S[i + 1][j + 1] <= k)
                {
                    ans++;
                }
            }
        }
        return ans;
    }
}
