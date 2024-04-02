/*
 * @lc app=leetcode.cn id=number-of-black-blocks lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6928] 黑格子的数目
 *
 * https://leetcode.cn/problems/number-of-black-blocks/description/
 *
 * algorithms
 * Medium (36.73%)
 * Total Accepted:    1.1K
 * Total Submissions: 3K
 * Testcase Example:  '3\n3\n[[0,0]]'
 *
 * 给你两个整数 m 和 n ，表示一个下标从 0 开始的 m x n 的网格图。
 * 
 * 给你一个下标从 0 开始的二维整数矩阵 coordinates ，其中 coordinates[i] = [x, y] 表示坐标为 [x, y]
 * 的格子是 黑色的 ，所有没出现在 coordinates 中的格子都是 白色的。
 * 
 * 一个块定义为网格图中 2 x 2 的一个子矩阵。更正式的，对于左上角格子为 [x, y] 的块，其中 0 <= x < m - 1 且 0 <= y <
 * n - 1 ，包含坐标为 [x, y] ，[x + 1, y] ，[x, y + 1] 和 [x + 1, y + 1] 的格子。
 * 
 * 请你返回一个下标从 0 开始长度为 5 的整数数组 arr ，arr[i] 表示恰好包含 i 个 黑色 格子的块的数目。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：m = 3, n = 3, coordinates = [[0,0]]
 * 输出：[3,1,0,0,0]
 * 解释：网格图如下：
 * 
 * 只有 1 个块有一个黑色格子，这个块是左上角为 [0,0] 的块。
 * 其他 3 个左上角分别为 [0,1] ，[1,0] 和 [1,1] 的块都有 0 个黑格子。
 * 所以我们返回 [3,1,0,0,0] 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：m = 3, n = 3, coordinates = [[0,0],[1,1],[0,2]]
 * 输出：[0,2,2,0,0]
 * 解释：网格图如下：
 * 
 * 有 2 个块有 2 个黑色格子（左上角格子分别为 [0,0] 和 [0,1]）。
 * 左上角为 [1,0] 和 [1,1] 的两个块，都有 1 个黑格子。
 * 所以我们返回 [0,2,2,0,0] 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 2 <= m <= 10^5
 * 2 <= n <= 10^5
 * 0 <= coordinates.length <= 10^4
 * coordinates[i].length == 2
 * 0 <= coordinates[i][0] < m
 * 0 <= coordinates[i][1] < n
 * coordinates 中的坐标对两两互不相同。
 * 
 * 
 */
public class Solution
{
    private static (int, int)[] Directions = new[]
    {
        (-1, -1),
        (-1, 1),
        (1, 1),
        (1, -1),
    };

    public long[] CountBlackBlocks(int n, int m, int[][] coordinates)
    {
        var ans = new long[Directions.Length + 1];
        ans[0] = (long)(n - 1) * (long)(m - 1);
        var G = Enumerable.Range(0, n).Select(_ => new HashSet<int>()).ToArray();
        bool inBoard(int x, int y) => 0 <= x && x < n && 0 <= y && y < m;
        int getTotal(int x, int y, int dx, int dy) => new (int, int)[]
        {
            (x, y),
            (x, y + dy),
            (x + dx, y),
            (x + dx, y + dy),
        }
        .Where(p => G[p.Item1].Contains(p.Item2))
        .Count();
        foreach (var c in coordinates.ToHashSet())
        {
            var (x, y) = (c[0], c[1]);
            foreach (var (dx, dy) in Directions)
            {
                if (!inBoard(x + dx, y + dy)) { continue; }
                var k = getTotal(x, y, dx, dy);
                ans[k]--;
                ans[k + 1]++;
            }
            G[x].Add(y);
        }
        return ans;
    }
}
