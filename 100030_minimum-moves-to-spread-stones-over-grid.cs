/*
 * @lc app=leetcode.cn id=minimum-moves-to-spread-stones-over-grid lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [100030] 将石头分散到网格图的最少移动次数
 *
 * https://leetcode.cn/problems/minimum-moves-to-spread-stones-over-grid/description/
 *
 * algorithms
 * Medium (33.54%)
 * Total Accepted:    2.3K
 * Total Submissions: 6.8K
 * Testcase Example:  '[[1,1,0],[1,1,1],[1,2,1]]'
 *
 * 给你一个大小为 3 * 3 ，下标从 0 开始的二维整数矩阵 grid ，分别表示每一个格子里石头的数目。网格图中总共恰好有 9
 * 个石头，一个格子里可能会有 多个 石头。
 * 
 * 每一次操作中，你可以将一个石头从它当前所在格子移动到一个至少有一条公共边的相邻格子。
 * 
 * 请你返回每个格子恰好有一个石头的 最少移动次数 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 
 * 输入：grid = [[1,1,0],[1,1,1],[1,2,1]]
 * 输出：3
 * 解释：让每个格子都有一个石头的一个操作序列为：
 * 1 - 将一个石头从格子 (2,1) 移动到 (2,2) 。
 * 2 - 将一个石头从格子 (2,2) 移动到 (1,2) 。
 * 3 - 将一个石头从格子 (1,2) 移动到 (0,2) 。
 * 总共需要 3 次操作让每个格子都有一个石头。
 * 让每个格子都有一个石头的最少操作次数为 3 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 
 * 
 * 输入：grid = [[1,3,0],[1,0,0],[1,0,3]]
 * 输出：4
 * 解释：让每个格子都有一个石头的一个操作序列为：
 * 1 - 将一个石头从格子 (0,1) 移动到 (0,2) 。
 * 2 - 将一个石头从格子 (0,1) 移动到 (1,1) 。
 * 3 - 将一个石头从格子 (2,2) 移动到 (1,2) 。
 * 4 - 将一个石头从格子 (2,2) 移动到 (2,1) 。
 * 总共需要 4 次操作让每个格子都有一个石头。
 * 让每个格子都有一个石头的最少操作次数为 4 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * grid.length == grid[i].length == 3
 * 0 <= grid[i][j] <= 9
 * grid 中元素之和为 9 。
 * 
 * 
 */
public class Solution
{
    public int MinimumMoves(int[][] grid)
    {
        const int Maxn = 0x3f3f3f3f;
        const int N = 3;
        var froms = new List<int>();
        var tos = new bool[N * N];
        for (var i = 0; i < N * N; i++)
        {
            var (x, y) = (i / N, i % N);
            for (var k = grid[x][y]; k > 1; k--) { froms.Add(i); }
            if (grid[x][y] is 0) { tos[i] = true; }
        }
        var dists = Enumerable.Range(0, N * N)
            .Select(i => Enumerable.Range(0, N * N)
                .Select(j => Math.Abs(i / N - j / N) + Math.Abs(i % N - j % N))
                .ToList())
            .ToList();
        var ans = Maxn;
        void dfs(int k, int m)
        {
            if (k == froms.Count())
            {
                ans = Math.Min(ans, m);
                return;
            }
            for (var i = 0; i < N * N; i++)
            {
                if (!tos[i]) { continue; }
                tos[i] = false;
                dfs(k + 1, m + dists[froms[k]][i]);
                tos[i] = true;
            }
        }
        dfs(0, 0);
        return ans;
    }
}
