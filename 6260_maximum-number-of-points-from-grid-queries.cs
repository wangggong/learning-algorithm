/*
 * @lc app=leetcode.cn id=maximum-number-of-points-from-grid-queries lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6260] 矩阵查询可获得的最大分数
 *
 * https://leetcode.cn/problems/maximum-number-of-points-from-grid-queries/description/
 *
 * algorithms
 * Hard (27.02%)
 * Total Accepted:    1.2K
 * Total Submissions: 4.3K
 * Testcase Example:  '[[1,2,3],[2,5,7],[3,5,1]]\n[5,6,2]'
 *
 * 给你一个大小为 m x n 的整数矩阵 grid 和一个大小为 k 的数组 queries 。
 * 
 * 找出一个大小为 k 的数组 answer ，且满足对于每个整数 queres[i] ，你从矩阵 左上角 单元格开始，重复以下过程：
 * 
 * 
 * 如果 queries[i] 严格 大于你当前所处位置单元格，如果该单元格是第一次访问，则获得 1 分，并且你可以移动到所有 4
 * 个方向（上、下、左、右）上任一 相邻 单元格。
 * 否则，你不能获得任何分，并且结束这一过程。
 * 
 * 
 * 在过程结束后，answer[i] 是你可以获得的最大分数。注意，对于每个查询，你可以访问同一个单元格 多次 。
 * 
 * 返回结果数组 answer 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：grid = [[1,2,3],[2,5,7],[3,5,1]], queries = [5,6,2]
 * 输出：[5,8,1]
 * 解释：上图展示了每个查询中访问并获得分数的单元格。
 * 
 * 示例 2：
 * 
 * 输入：grid = [[5,2,1],[1,1,2]], queries = [3]
 * 输出：[0]
 * 解释：无法获得分数，因为左上角单元格的值大于等于 3 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * m == grid.length
 * n == grid[i].length
 * 2 <= m, n <= 1000
 * 4 <= m * n <= 10^5
 * k == queries.length
 * 1 <= k <= 10^4
 * 1 <= grid[i][j], queries[i] <= 10^6
 * 
 * 
 */
// 参考: https://leetcode.cn/problems/maximum-number-of-points-from-grid-queries/solution/by-endlesscheng-qeei/
public class Solution
{
    public int[] MaxPoints(int[][] grid, int[] queries)
    {
        var directions = new List<(int, int)>
        {
            (0, 1),
            (1, 0),
            (0, -1),
            (-1, 0),
        };
        var n = grid.Length;
        var m = grid[0].Length;
        var pa = new int[n * m];
        var size = new int[n * m];
        var infos = new List<(int, int)>();
        for (var i = 0; i < n; i++)
        {
            for (var j = 0; j < m; j++)
            {
                var v = i * m + j;
                pa[v] = v;
                size[v] = 1;
                infos.Add((grid[i][j], v));
            }
        }
        infos.Sort((x, y) => x.Item1.CompareTo(y.Item1));
        int query(int k)
        {
            if (k != pa[k]) { pa[k] = query(pa[k]); }
            return pa[k];
        }
        void merge(int p, int q)
        {
            var qp = query(p);
            var qq = query(q);
            if (size[qp] > size[qq])
            {
                merge(q, p);
                return;
            }
            size[qq] += size[qp];
            pa[qp] = qq;
            return;
        }
        var qn = queries.Length;
        var queryInfos = new List<(int, int)>();
        var ans = new int[qn];
        for (var i = 0; i < qn; i++) { queryInfos.Add((queries[i], i)); }
        var cur = 0;
        foreach ((var k, var i) in queryInfos.OrderBy(t => t.Item1))
        {
            if (k <= grid[0][0]) { continue; }
            for (; cur < n * m && infos[cur].Item1 < k; cur++)
            {
                var v = infos[cur].Item2;
                var x = v / m;
                var y = v % m;
                foreach ((var dx, var dy) in directions)
                {
                    var nx = x + dx;
                    var ny = y + dy;
                    var nv = nx * m + ny;
                    if (nx < 0 || nx >= n || ny < 0 || ny >= m || grid[nx][ny] >= k || query(v) == query(nv)) { continue; }
                    merge(v, nv);
                }
            }
            ans[i] = size[query(0)];
        }
        return ans;
    }
}
