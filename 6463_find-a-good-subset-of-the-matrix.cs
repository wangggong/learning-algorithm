/*
 * @lc app=leetcode.cn id=find-a-good-subset-of-the-matrix lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6463] 找到矩阵中的好子集
 *
 * https://leetcode.cn/problems/find-a-good-subset-of-the-matrix/description/
 *
 * algorithms
 * Hard (48.73%)
 * Total Accepted:    727
 * Total Submissions: 1.5K
 * Testcase Example:  '[[0,1,1,0],[0,0,0,1],[1,1,1,1]]'
 *
 * 给你一个下标从 0 开始大小为 m x n 的二进制矩阵 grid 。
 * 
 * 从原矩阵中选出若干行构成一个行的 非空 子集，如果子集中任何一列的和至多为子集大小的一半，那么我们称这个子集是 好子集。
 * 
 * 更正式的，如果选出来的行子集大小（即行的数量）为 k，那么每一列的和至多为 floor(k / 2) 。
 * 
 * 请你返回一个整数数组，它包含好子集的行下标，请你将子集中的元素 升序 返回。
 * 
 * 如果有多个好子集，你可以返回任意一个。如果没有好子集，请你返回一个空数组。
 * 
 * 一个矩阵 grid 的行 子集 ，是删除 grid 中某些（也可能不删除）行后，剩余行构成的元素集合。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：grid = [[0,1,1,0],[0,0,0,1],[1,1,1,1]]
 * 输出：[0,1]
 * 解释：我们可以选择第 0 和第 1 行构成一个好子集。
 * 选出来的子集大小为 2 。
 * - 第 0 列的和为 0 + 0 = 0 ，小于等于子集大小的一半。
 * - 第 1 列的和为 1 + 0 = 1 ，小于等于子集大小的一半。
 * - 第 2 列的和为 1 + 0 = 1 ，小于等于子集大小的一半。
 * - 第 3 列的和为 0 + 1 = 1 ，小于等于子集大小的一半。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：grid = [[0]]
 * 输出：[0]
 * 解释：我们可以选择第 0 行构成一个好子集。
 * 选出来的子集大小为 1 。
 * - 第 0 列的和为 0 ，小于等于子集大小的一半。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：grid = [[1,1,1],[1,1,1]]
 * 输出：[]
 * 解释：没有办法得到一个好子集。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * m == grid.length
 * n == grid[i].length
 * 1 <= m <= 10^4
 * 1 <= n <= 5
 * grid[i][j] 要么是 0 ，要么是 1 。
 * 
 * 
 */
public class Solution
{
    public IList<int> GoodSubsetofBinaryMatrix(int[][] G)
    {
        int getDigit(int[] row) => row
            .Select((v, i) => v * (1 << i))
            .Sum();
        var n = G[0].Length;
        var d = Enumerable
            .Range(0, 1 << n)
            .ToDictionary(v => v, _ => new List<int>());
        for (var (i, m) = (0, G.Length); i < m; i++)
        {
            d[getDigit(G[i])].Add(i);
        }
        for (var i = 0; i < (1 << n); i++)
        {
            if (!d[i].Any())
            {
                continue;
            }
            if (i == 0)
            {
                return d[i];
            }
            for (var j = 0; j < (1 << n); j++)
            {
                if ((i & j) == 0 && d[j].Any())
                {
                    return new List<int>{ d[i][0], d[j][0], }
                        .OrderBy(x => x)
                        .ToList();
                }
            }
        }
        return new List<int>();
    }
}
