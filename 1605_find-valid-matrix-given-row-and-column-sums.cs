/*
 * @lc app=leetcode.cn id=find-valid-matrix-given-row-and-column-sums lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1605] 给定行和列的和求可行矩阵
 *
 * https://leetcode.cn/problems/find-valid-matrix-given-row-and-column-sums/description/
 *
 * algorithms
 * Medium (76.31%)
 * Total Accepted:    6.3K
 * Total Submissions: 8.1K
 * Testcase Example:  '[3,8]\n[4,7]'
 *
 * 给你两个非负整数数组 rowSum 和 colSum ，其中 rowSum[i] 是二维矩阵中第 i 行元素的和， colSum[j] 是第 j
 * 列元素的和。换言之你不知道矩阵里的每个元素，但是你知道每一行和每一列的和。
 * 
 * 请找到大小为 rowSum.length x colSum.length 的任意 非负整数 矩阵，且该矩阵满足 rowSum 和 colSum
 * 的要求。
 * 
 * 请你返回任意一个满足题目要求的二维矩阵，题目保证存在 至少一个 可行矩阵。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：rowSum = [3,8], colSum = [4,7]
 * 输出：[[3,0],
 * ⁠     [1,7]]
 * 解释：
 * 第 0 行：3 + 0 = 3 == rowSum[0]
 * 第 1 行：1 + 7 = 8 == rowSum[1]
 * 第 0 列：3 + 1 = 4 == colSum[0]
 * 第 1 列：0 + 7 = 7 == colSum[1]
 * 行和列的和都满足题目要求，且所有矩阵元素都是非负的。
 * 另一个可行的矩阵为：[[1,2],
 * ⁠                 [3,5]]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：rowSum = [5,7,10], colSum = [8,6,8]
 * 输出：[[0,5,0],
 * ⁠     [6,1,0],
 * ⁠     [2,0,8]]
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：rowSum = [14,9], colSum = [6,9,8]
 * 输出：[[0,9,5],
 * ⁠     [6,0,3]]
 * 
 * 
 * 示例 4：
 * 
 * 
 * 输入：rowSum = [1,0], colSum = [1]
 * 输出：[[1],
 * ⁠     [0]]
 * 
 * 
 * 示例 5：
 * 
 * 
 * 输入：rowSum = [0], colSum = [0]
 * 输出：[[0]]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= rowSum.length, colSum.length <= 500
 * 0 <= rowSum[i], colSum[i] <= 10^8
 * sum(rowSum) == sum(colSum)
 * 
 * 
 */
public class Solution
{
    public int[][] RestoreMatrix(int[] rowSum, int[] colSum)
    {
        var n = rowSum.Length;
        var m = colSum.Length;
        var ans = new int[n][];
        for (var i = 0; i < n; i++)
        {
            ans[i] = new int[m];
        }
        var rowInfo = new List<(int, int)>();
        var colInfo = new List<(int, int)>();
        for (var i = 0; i < n; i++)
        {
            if (rowSum[i] > 0)
            {
                rowInfo.Add((i, rowSum[i]));
            }
        }
        for (var i = 0; i < m; i++)
        {
            if (colSum[i] > 0)
            {
                colInfo.Add((i, colSum[i]));
            }
        }
        rowInfo.Sort((x, y) => x.Item2.CompareTo(y.Item2));
        colInfo.Sort((x, y) => x.Item2.CompareTo(y.Item2));
        var p = 0;
        var q = 0;
        while (p < rowInfo.Count() || q < colInfo.Count())
        {
            var (i, r) = rowInfo[p];
            var (j, c) = colInfo[q];
            var v = Math.Min(r, c);
            ans[i][j] = v;
            rowInfo[p] = (i, r - v);
            colInfo[q] = (j, c - v);
            if (rowInfo[p].Item2 == 0)
            {
                p++;
            }
            if (colInfo[q].Item2 == 0)
            {
                q++;
            }
        }
        return ans;
    }
}
