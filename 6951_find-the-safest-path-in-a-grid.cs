/*
 * @lc app=leetcode.cn id=find-the-safest-path-in-a-grid lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6951] 找出最安全路径
 *
 * https://leetcode.cn/problems/find-the-safest-path-in-a-grid/description/
 *
 * algorithms
 * Medium (19.51%)
 * Total Accepted:    2.4K
 * Total Submissions: 10K
 * Testcase Example:  '[[1,0,0],[0,0,0],[0,0,1]]'
 *
 * 给你一个下标从 0 开始、大小为 n x n 的二维矩阵 grid ，其中 (r, c) 表示：
 * 
 * 
 * 如果 grid[r][c] = 1 ，则表示一个存在小偷的单元格
 * 如果 grid[r][c] = 0 ，则表示一个空单元格
 * 
 * 
 * 你最开始位于单元格 (0, 0) 。在一步移动中，你可以移动到矩阵中的任一相邻单元格，包括存在小偷的单元格。
 * 
 * 矩阵中路径的 安全系数 定义为：从路径中任一单元格到矩阵中任一小偷所在单元格的 最小 曼哈顿距离。
 * 
 * 返回所有通向单元格 (n - 1, n - 1) 的路径中的 最大安全系数 。
 * 
 * 单元格 (r, c) 的某个 相邻 单元格，是指在矩阵中存在的 (r, c + 1)、(r, c - 1)、(r + 1, c) 和 (r - 1,
 * c) 之一。
 * 
 * 两个单元格 (a, b) 和 (x, y) 之间的 曼哈顿距离 等于 | a - x | + | b - y | ，其中 |val| 表示 val
 * 的绝对值。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：grid = [[1,0,0],[0,0,0],[0,0,1]]
 * 输出：0
 * 解释：从 (0, 0) 到 (n - 1, n - 1) 的每条路径都经过存在小偷的单元格 (0, 0) 和 (n - 1, n - 1) 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：grid = [[0,0,1],[0,0,0],[0,0,0]]
 * 输出：2
 * 解释：
 * 上图所示路径的安全系数为 2：
 * - 该路径上距离小偷所在单元格（0，2）最近的单元格是（0，0）。它们之间的曼哈顿距离为 | 0 - 0 | + | 0 - 2 | = 2 。
 * 可以证明，不存在安全系数更高的其他路径。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：grid = [[0,0,0,1],[0,0,0,0],[0,0,0,0],[1,0,0,0]]
 * 输出：2
 * 解释：
 * 上图所示路径的安全系数为 2：
 * - 该路径上距离小偷所在单元格（0，3）最近的单元格是（1，2）。它们之间的曼哈顿距离为 | 0 - 1 | + | 3 - 2 | = 2 。
 * - 该路径上距离小偷所在单元格（3，0）最近的单元格是（3，2）。它们之间的曼哈顿距离为 | 3 - 3 | + | 0 - 2 | = 2 。
 * 可以证明，不存在安全系数更高的其他路径。
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= grid.length == n <= 400
 * grid[i].length == n
 * grid[i][j] 为 0 或 1
 * grid 至少存在一个小偷
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/find-the-safest-path-in-a-grid/solution/jie-jin-on2-de-zuo-fa-duo-yuan-bfsdao-xu-r5um/
//
// 想到了二分答案, 没想到预处理距离.
public class Solution
{
    private const int Maxn = 0x3f3f3f3f;

    public int MaximumSafenessFactor(IList<IList<int>> G)
    {
        var directions = new int[] { 0, 1, 0, -1, 0, };
        var n = G.Count();
        bool inGraph(int x, int y) => 0 <= x && x < n && 0 <= y && y < n;
        var dists = new int[n][];
        for (var i = 0; i < n; i++) { dists[i] = new int[n]; }
        for (var i = 0; i < n; i++) { Array.Fill(dists[i], Maxn); }
        var Q = new Queue<int>();
        for (var index = 0; index < n * n; index++)
        {
            var (i, j) = (index / n, index % n);
            if (G[i][j] is 0) { continue; }
            Q.Enqueue(index);
            dists[i][j] = 0;
        }
        for (var step = 1; Q.Count > 0; step++)
        {
            for (var c = Q.Count; c > 0; c--)
            {
                var index = Q.Dequeue();
                var (i, j) = (index / n, index % n);
                for (var d = 0; d < directions.Length - 1; d++)
                {
                    var (ni, nj) = (i + directions[d], j + directions[d + 1]);
                    if (inGraph(ni, nj) && dists[ni][nj] > step)
                    {
                        dists[ni][nj] = step;
                        Q.Enqueue(ni * n + nj);
                    }
                }
            }
        }
        bool check(int k)
        {
            var Q = new Queue<int>();
            var S = new HashSet<int>();
            if (dists[0][0] >= k)
            {
                Q.Enqueue(0);
                S.Add(0);
            }
            while (Q.Count > 0)
            {
                var index = Q.Dequeue();
                var (i, j) = (index / n, index % n);
                if (i == n - 1 && j == n - 1) { return true; }
                for (var d = 0; d < directions.Length - 1; d++)
                {
                    var (ni, nj) = (i + directions[d], j + directions[d + 1]);
                    var nIndex = ni * n + nj;
                    if (inGraph(ni, nj) && !S.Contains(nIndex) && dists[ni][nj] >= k)
                    {
                        Q.Enqueue(nIndex);
                        S.Add(nIndex);
                    }
                }
            }
            return false;
        }
        var (p, q) = (1, n * n + 1);
        while (p < q)
        {
            var mid = (p + q) >> 1;
            if (!check(mid)) { q = mid; }
            else { p = mid + 1; }
        }
        return p - 1;
    }
}
