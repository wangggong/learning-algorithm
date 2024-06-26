/*
 * @lc app=leetcode.cn id=unique-paths-iii lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [980] 不同路径 III
 *
 * https://leetcode.cn/problems/unique-paths-iii/description/
 *
 * algorithms
 * Hard (74.13%)
 * Total Accepted:    22.3K
 * Total Submissions: 29.9K
 * Testcase Example:  '[[1,0,0,0],[0,0,0,0],[0,0,2,-1]]'
 *
 * 在二维网格 grid 上，有 4 种类型的方格：
 * 
 * 
 * 1 表示起始方格。且只有一个起始方格。
 * 2 表示结束方格，且只有一个结束方格。
 * 0 表示我们可以走过的空方格。
 * -1 表示我们无法跨越的障碍。
 * 
 * 
 * 返回在四个方向（上、下、左、右）上行走时，从起始方格到结束方格的不同路径的数目。
 * 
 * 每一个无障碍方格都要通过一次，但是一条路径中不能重复通过同一个方格。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：[[1,0,0,0],[0,0,0,0],[0,0,2,-1]]
 * 输出：2
 * 解释：我们有以下两条路径：
 * 1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2)
 * 2. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2)
 * 
 * 示例 2：
 * 
 * 输入：[[1,0,0,0],[0,0,0,0],[0,0,0,2]]
 * 输出：4
 * 解释：我们有以下四条路径： 
 * 1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2),(2,3)
 * 2. (0,0),(0,1),(1,1),(1,0),(2,0),(2,1),(2,2),(1,2),(0,2),(0,3),(1,3),(2,3)
 * 3. (0,0),(1,0),(2,0),(2,1),(2,2),(1,2),(1,1),(0,1),(0,2),(0,3),(1,3),(2,3)
 * 4. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2),(2,3)
 * 
 * 示例 3：
 * 
 * 输入：[[0,1],[2,0]]
 * 输出：0
 * 解释：
 * 没有一条路能完全穿过每一个空的方格一次。
 * 请注意，起始和结束方格可以位于网格中的任意位置。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= grid.length * grid[0].length <= 20
 * 
 * 
 */
public class Solution
{
    public int UniquePathsIII(int[][] G)
    {
        const int Start = 1;
        const int End = 2;
        var directions = new int[] { 0, 1, 0, -1, 0, };
        var (n, m) = (G.Length, G[0].Length);
        bool isValidPosition(int x, int y) => 0 <= x && x < n
            && 0 <= y && y < m
            && G[x][y] >= 0;
        var (start, end, S) = (0, 0, 0);
        for (var i = 0; i < n; i++)
        {
            for (var j = 0; j < m; j++)
            {
                var position = i * m + j;
                if (G[i][j] >= 0) { S |= 1 << position; }
                if (G[i][j] is Start) { start = position; }
                if (G[i][j] is End) { end = position; }
            }
        }
        var memo = new Dictionary<(int, int), int>();
        int dfs(int position, int mask)
        {
            var key = (position, mask);
            if (memo.ContainsKey(key)) { return memo[key]; }
            if (position == end) { return memo[key] = mask == S ? 1 : 0; }
            memo[key] = 0;
            var (x, y) = (position / m, position % m);
            for (var d = 0; d + 1 < directions.Length; d++)
            {
                var (nx, ny) = (x + directions[d], y + directions[d + 1]);
                var next = nx * m + ny;
                if (isValidPosition(nx, ny) && (mask & (1 << next)) is 0)
                {
                    memo[key] += dfs(next, mask | (1 << next));
                }
            }
            return memo[key];
        }
        return dfs(start, 1 << start);
    }
}
