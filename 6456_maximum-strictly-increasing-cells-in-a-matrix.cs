/*
 * @lc app=leetcode.cn id=maximum-strictly-increasing-cells-in-a-matrix lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6456] 矩阵中严格递增的单元格数
 *
 * https://leetcode.cn/problems/maximum-strictly-increasing-cells-in-a-matrix/description/
 *
 * algorithms
 * Hard (17.71%)
 * Total Accepted:    806
 * Total Submissions: 4K
 * Testcase Example:  '[[3,1],[3,4]]'
 *
 * 给你一个下标从 1 开始、大小为 m x n 的整数矩阵 mat，你可以选择任一单元格作为 起始单元格 。
 * 
 * 从起始单元格出发，你可以移动到 同一行或同一列 中的任何其他单元格，但前提是目标单元格的值 严格大于 当前单元格的值。
 * 
 * 你可以多次重复这一过程，从一个单元格移动到另一个单元格，直到无法再进行任何移动。
 * 
 * 请你找出从某个单元开始访问矩阵所能访问的 单元格的最大数量 。
 * 
 * 返回一个表示可访问单元格最大数量的整数。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 输入：mat = [[3,1],[3,4]]
 * 输出：2
 * 解释：上图展示了从第 1 行、第 2 列的单元格开始，可以访问 2 个单元格。可以证明，无论从哪个单元格开始，最多只能访问 2 个单元格，因此答案是 2
 * 。 
 * 
 * 
 * 示例 2：
 * 
 * 
 * 
 * 输入：mat = [[1,1],[1,1]]
 * 输出：1
 * 解释：由于目标单元格必须严格大于当前单元格，在本示例中只能访问 1 个单元格。 
 * 
 * 
 * 示例 3：
 * 
 * 
 * 
 * 输入：mat = [[3,1,6],[-9,5,7]]
 * 输出：4
 * 解释：上图展示了从第 2 行、第 1 列的单元格开始，可以访问 4 个单元格。可以证明，无论从哪个单元格开始，最多只能访问 4 个单元格，因此答案是 4
 * 。  
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * m == mat.length 
 * n == mat[i].length 
 * 1 <= m, n <= 10^5
 * 1 <= m * n <= 10^5
 * -10^5 <= mat[i][j] <= 10^5
 * 
 * 
 */
// 参考: https://leetcode.cn/problems/maximum-strictly-increasing-cells-in-a-matrix/solution/dong-tai-gui-hua-you-hua-pythonjavacgo-b-axv0/
public class Solution
{
    public int MaxIncreasingCells(int[][] M)
    {
        var (n, m) = (M.Length, M[0].Length);
        var dpRows = new int[n];
        var dpCols = new int[m];
        foreach (var g in M
            .SelectMany((row, i) => row.Select((v, j) => (i, j, v)))
            .GroupBy(x => x.v)
            .OrderBy(g => g.Key))
        {
            foreach (var kv in g
                .Select(x =>
                {
                    var (i, j, v) = x;
                    return ((i, j), Math.Max(dpRows[i], dpCols[j]) + 1);
                })
                .ToDictionary(x => x.Item1, x => x.Item2))
            {
                var (i, j) = kv.Key;
                dpRows[i] = Math.Max(dpRows[i], kv.Value);
                dpCols[j] = Math.Max(dpCols[j], kv.Value);
            }
        }
        return Math.Max(dpRows.Max(), dpCols.Max());
    }
}
