/*
 * @lc app=leetcode.cn id=sum-of-matrix-after-queries lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6472] 查询后矩阵的和
 *
 * https://leetcode.cn/problems/sum-of-matrix-after-queries/description/
 *
 * algorithms
 * Medium (22.18%)
 * Total Accepted:    1.9K
 * Total Submissions: 8.8K
 * Testcase Example:  '3\n[[0,0,1],[1,2,2],[0,2,3],[1,0,4]]'
 *
 * 给你一个整数 n 和一个下标从 0 开始的 二维数组 queries ，其中 queries[i] = [typei, indexi, vali] 。
 * 
 * 一开始，给你一个下标从 0 开始的 n x n 矩阵，所有元素均为 0 。每一个查询，你需要执行以下操作之一：
 * 
 * 
 * 如果 typei == 0 ，将第 indexi 行的元素全部修改为 vali ，覆盖任何之前的值。
 * 如果 typei == 1 ，将第 indexi 列的元素全部修改为 vali ，覆盖任何之前的值。
 * 
 * 
 * 请你执行完所有查询以后，返回矩阵中所有整数的和。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 输入：n = 3, queries = [[0,0,1],[1,2,2],[0,2,3],[1,0,4]]
 * 输出：23
 * 解释：上图展示了每个查询以后矩阵的值。所有操作执行完以后，矩阵元素之和为 23 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 
 * 输入：n = 3, queries = [[0,0,4],[0,1,2],[1,0,1],[0,2,3],[1,2,1]]
 * 输出：17
 * 解释：上图展示了每一个查询操作之后的矩阵。所有操作执行完以后，矩阵元素之和为 17 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 10^4
 * 1 <= queries.length <= 5 * 10^4
 * queries[i].length == 3
 * 0 <= typei <= 1
 * 0 <= indexi < n
 * 0 <= vali <= 10^5
 * 
 * 
 */
public class Solution
{
    public long MatrixSumQueries(int n, int[][] queries)
    {
        long ans = 0;
        var (rows, cols) = (new bool[n], new bool[n]);
        var (rowCount, colCount) = (0, 0);
        foreach (var query in queries.Reverse())
        {
            var (t, i, v) = (query[0], query[1], query[2]);
            if (t == 0)
            {
                if (!rows[i])
                {
                    (rows[i], rowCount) = (true, rowCount + 1);
                    ans += (long)((n - colCount) * v);
                }
            }
            else
            {
                if (!cols[i])
                {
                    (cols[i], colCount) = (true, colCount + 1);
                    ans += (long)((n - rowCount) * v);
                }
            }
        }
        return ans;
    }
}
