/*
 * @lc app=leetcode.cn id=shortest-path-in-binary-matrix lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1091] 二进制矩阵中的最短路径
 *
 * https://leetcode.cn/problems/shortest-path-in-binary-matrix/description/
 *
 * algorithms
 * Medium (38.73%)
 * Total Accepted:    64.5K
 * Total Submissions: 164.8K
 * Testcase Example:  '[[0,1],[1,0]]'
 *
 * 给你一个 n x n 的二进制矩阵 grid 中，返回矩阵中最短 畅通路径 的长度。如果不存在这样的路径，返回 -1 。
 * 
 * 二进制矩阵中的 畅通路径 是一条从 左上角 单元格（即，(0, 0)）到 右下角 单元格（即，(n - 1, n -
 * 1)）的路径，该路径同时满足下述要求：
 * 
 * 
 * 路径途经的所有单元格都的值都是 0 。
 * 路径中所有相邻的单元格应当在 8 个方向之一 上连通（即，相邻两单元之间彼此不同且共享一条边或者一个角）。
 * 
 * 
 * 畅通路径的长度 是该路径途经的单元格总数。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：grid = [[0,1],[1,0]]
 * 输出：2
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：grid = [[0,0,0],[1,1,0],[1,1,0]]
 * 输出：4
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：grid = [[1,0,0],[1,1,0],[1,1,0]]
 * 输出：-1
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * n == grid.length
 * n == grid[i].length
 * 1 
 * grid[i][j] 为 0 或 1
 * 
 * 
 */
public class Solution
{
    public int ShortestPathBinaryMatrix(int[][] G)
    {
        var (n, m) = (G.Length, G[0].Length);
        var visit = Enumerable.Range(0, n).Select(_ => new bool[m]).ToArray();
        var Q = new Queue<(int, int)>();
        if (G[0][0] == 0)
        {
            Q.Enqueue((0, 0));
            visit[0][0] = true;
        }
        for (var step = 1; Q.Count > 0; step++)
        {
            for (var c = Q.Count; c > 0; c--)
            {
                var (x, y) = Q.Dequeue();
                if ((n - 1 - x, m - 1 - y) is (0, 0))
                {
                    return step;
                }
                for (var dx = -1; dx <= 1; dx++)
                {
                    for (var dy = -1; dy <= 1; dy++)
                    {
                        if ((dx, dy) is (0, 0))
                        {
                            continue;
                        }
                        var (nx, ny) = (x + dx, y + dy);
                        if (0 <= nx && nx < n
                            && 0 <= ny && ny < m
                            && G[nx][ny] == 0 && !visit[nx][ny])
                        {
                            Q.Enqueue((nx, ny));
                            visit[nx][ny] = true;
                        }
                    }
                }
            }
        }
        return -1;
    }
}
