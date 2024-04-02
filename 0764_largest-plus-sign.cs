/*
 * @lc app=leetcode.cn id=largest-plus-sign lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [764] 最大加号标志
 *
 * https://leetcode.cn/problems/largest-plus-sign/description/
 *
 * algorithms
 * Medium (49.77%)
 * Total Accepted:    20.7K
 * Total Submissions: 38.9K
 * Testcase Example:  '5\n[[4,2]]'
 *
 * 在一个 n x n 的矩阵 grid 中，除了在数组 mines 中给出的元素为 0，其他每个元素都为 1。mines[i] = [xi, yi]表示
 * grid[xi][yi] == 0
 * 
 * 返回  grid 中包含 1 的最大的 轴对齐 加号标志的阶数 。如果未找到加号标志，则返回 0 。
 * 
 * 一个 k 阶由 1 组成的 “轴对称”加号标志 具有中心网格 grid[r][c] == 1 ，以及4个从中心向上、向下、向左、向右延伸，长度为
 * k-1，由 1 组成的臂。注意，只有加号标志的所有网格要求为 1 ，别的网格可能为 0 也可能为 1 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 
 * 输入: n = 5, mines = [[4, 2]]
 * 输出: 2
 * 解释: 在上面的网格中，最大加号标志的阶只能是2。一个标志已在图中标出。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 
 * 
 * 输入: n = 1, mines = [[0, 0]]
 * 输出: 0
 * 解释: 没有加号标志，返回 0 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 500
 * 1 <= mines.length <= 5000
 * 0 <= xi, yi < n
 * 每一对 (xi, yi) 都 不重复​​​​​​​
 * 
 * 
 */
public class Solution
{
    private int LowerBound(List<int> list, int k)
    {
        int p = 0, q = list.Count();
        while (p < q)
        {
            int mid = (p + q + 1) >> 1;
            if (list[mid] <= k) { p = mid; }
            else { q = mid - 1; }
        }
        return p;
    }

    public int OrderOfLargestPlusSign(int n, int[][] mines)
    {
        int ans = 0;
        List<int>[] rows = new List<int>[n];
        List<int>[] cols = new List<int>[n];
        for (int i = 0; i < n; i++)
        {
            rows[i] = new List<int>();
            cols[i] = new List<int>();
        }
        foreach (var mine in mines)
        {
            rows[mine[0]].Add(mine[1]);
            cols[mine[1]].Add(mine[0]);
        }
        for (int i = 0; i < n; i++)
        {
            rows[i].Add(-1);
            rows[i].Add(n);
            rows[i].Sort();
            cols[i].Add(-1);
            cols[i].Add(n);
            cols[i].Sort();
        }
        for (int i = 0; i < n; i++)
        {
            for (int j = 0; j < n; j++)
            {
                int l = LowerBound(rows[i], j);
                int u = LowerBound(cols[j], i);
                int cur = Math.Min(
                    Math.Min(j - rows[i][l], rows[i][l + 1] - j),
                    Math.Min(i - cols[j][u], cols[j][u + 1] - i));
                ans = Math.Max(ans, cur);
            }
        }
        return ans;
    }
}
