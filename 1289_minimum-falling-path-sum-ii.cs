/*
 * @lc app=leetcode.cn id=minimum-falling-path-sum-ii lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1289] 下降路径最小和  II
 *
 * https://leetcode.cn/problems/minimum-falling-path-sum-ii/description/
 *
 * algorithms
 * Hard (58.32%)
 * Total Accepted:    22.7K
 * Total Submissions: 37.4K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * 给你一个 n x n 整数矩阵 grid ，请你返回 非零偏移下降路径 数字和的最小值。
 * 
 * 非零偏移下降路径 定义为：从 grid 数组中的每一行选择一个数字，且按顺序选出来的数字中，相邻数字不在原数组的同一列。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 
 * 输入：grid = [[1,2,3],[4,5,6],[7,8,9]]
 * 输出：13
 * 解释：
 * 所有非零偏移下降路径包括：
 * [1,5,9], [1,5,7], [1,6,7], [1,6,8],
 * [2,4,8], [2,4,9], [2,6,7], [2,6,8],
 * [3,4,8], [3,4,9], [3,5,7], [3,5,9]
 * 下降路径中数字和最小的是 [1,5,7] ，所以答案是 13 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：grid = [[7]]
 * 输出：7
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * n == grid.length == grid[i].length
 * 1 <= n <= 200
 * -99 <= grid[i][j] <= 99
 * 
 * 
 */
public class Solution
{
    public int MinFallingPathSum(int[][] G)
    {
        var n = G.Length;
        var dp = new int[n];
        for (var i = 0; i < n; i++) { dp[i] = G[0][i]; }
        for (var j = 1; j < n; j++)
        {
            // 灵佬表示这里可以直接舍弃记录 index, 见下面.
            var (minVal, _2ndMinVal) = (int.MaxValue, int.MaxValue);
            for (var i = 0; i < n; i++)
            {
                if (dp[i] < minVal) { (minVal, _2ndMinVal) = (dp[i], minVal); }
                else if (dp[i] < _2ndMinVal) { _2ndMinVal = dp[i]; }
                else { }
            }
            for (var i = 0; i < n; i++)
            {
                dp[i] = G[j][i] + (dp[i] == minVal ? _2ndMinVal : minVal);
            }
        }
        return dp.Min();
    }
}
