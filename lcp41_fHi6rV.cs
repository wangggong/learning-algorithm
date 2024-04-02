/*
 * Medium
 * 
 * 在 n*m 大小的棋盘中，有黑白两种棋子，黑棋记作字母 "X", 白棋记作字母 "O"，空余位置记作 "."。当落下的棋子与其他相同颜色的棋子在行、列或对角线完全包围（中间不存在空白位置）另一种颜色的棋子，则可以翻转这些棋子的颜色。
 * 
 * 
 * 
 * 「力扣挑战赛」黑白翻转棋项目中，将提供给选手一个未形成可翻转棋子的棋盘残局，其状态记作 chessboard。若下一步可放置一枚黑棋，请问选手最多能翻转多少枚白棋。
 * 
 * 注意：
 * 
 * 若翻转白棋成黑棋后，棋盘上仍存在可以翻转的白棋，将可以 继续 翻转白棋
 * 输入数据保证初始棋盘状态无可以翻转的棋子且存在空余位置
 * 示例 1：
 * 
 * 输入：chessboard = ["....X.","....X.","XOOO..","......","......"]
 * 
 * 输出：3
 * 
 * 解释：
 * 可以选择下在 [2,4] 处，能够翻转白方三枚棋子。
 * 
 * 示例 2：
 * 
 * 输入：chessboard = [".X.",".O.","XO."]
 * 
 * 输出：2
 * 
 * 解释：
 * 可以选择下在 [2,2] 处，能够翻转白方两枚棋子。
 * 
 * 
 * 示例 3：
 * 
 * 输入：chessboard = [".......",".......",".......","X......",".O.....","..O....","....OOX"]
 * 
 * 输出：4
 * 
 * 解释：
 * 可以选择下在 [6,3] 处，能够翻转白方四枚棋子。
 * 
 * 
 * 提示：
 * 
 * 1 <= chessboard.length, chessboard[i].length <= 8
 * chessboard[i] 仅包含 "."、"O" 和 "X"
 */

/*
 * public class Solution
 * {
 *     public int FlipChess(string[] chessboard)
 *     {
 *         var (n, m) = (chessboard.Length, chessboard[0].Length);
 *         var init = chessboard.SelectMany(row => row.Where(c => c is 'O')).Count();
 *         var ans = 0;
 *         bool onboard(int x, int y) => 0 <= x && x < n && 0 <= y && y < m;
 *         void handle(char[][] B)
 *         {
 *             while (true)
 *             {
 *                 var valid = false;
 *                 for (var i = 0; i < n; i++)
 *                 {
 *                     for (var j = 0; j < m; j++)
 *                     {
 *                         if (B[i][j] is not 'X')
 *                         {
 *                             continue;
 *                         }
 *                         for (var di = -1; di <= 1; di++)
 *                         {
 *                             for (var dj = -1; dj <= 1; dj++)
 *                             {
 *                                 if (di is 0 && dj is 0)
 *                                 {
 *                                     continue;
 *                                 }
 *                                 var k = 1;
 *                                 for (; onboard(i + k * di, j + k * dj) && B[i + k * di][j + k * dj] is 'O'; k++) { }
 *                                 if (k > 1 && onboard(i + k * di, j + k * dj) && B[i + k * di][j + k * dj] is 'X')
 *                                 {
 *                                     for (valid = true; k > 0; k--)
 *                                     {
 *                                         B[i + k * di][j + k * dj] = 'X';
 *                                     }
 *                                 }
 *                             }
 *                         }
 *                     }
 *                 }
 *                 if (!valid)
 *                 {
 *                     break;
 *                 }
 *             }
 *             ans = Math.Max(ans, init - B.SelectMany(row => row.Where(c => c is 'O')).Count());
 *         }
 *         for (var i = 0; i < n; i++)
 *         {
 *             for (var j = 0; j < m; j++)
 *             {
 *                 if (chessboard[i][j] is '.')
 *                 {
 *                     var B = chessboard.Select(row => row.ToCharArray()).ToArray();
 *                     B[i][j] = 'X';
 *                     handle(B);
 *                 }
 *             }
 *         }
 *         return ans;
 *     }
 * }
 */

// 参考: https://leetcode.cn/problems/fHi6rV/solution/python3javacgo-yi-ti-yi-jie-bfs-by-lcbin-e4vo/
public class Solution
{
    private IEnumerable<(int, int)> Directions()
    {
        for (var dx = -1; dx <= 1; dx++)
        {
            for (var dy = -1; dy <= 1; dy++)
            {
                if (dx is 0 && dy is 0) { continue; }
                yield return (dx, dy);
            }
        }
    }

    private int Bfs(string[] chessboard, int x, int y)
    {
        var B = chessboard.Select(row => row.ToCharArray()).ToArray();
        int countO() => B.SelectMany(row => row.Where(c => c is 'O')).Count();
        var (n, m) = (B.Length, B[0].Length);
        bool isValue(int x, int y, char c) => 0 <= x && x < n
            && 0 <= y && y < m
            && B[x][y] == c;
        var init = countO();
        var Q = new Queue<(int, int)>();
        Q.Enqueue((x, y));
        B[x][y] = 'X';
        while (Q.Count > 0)
        {
            (x, y) = Q.Dequeue();
            foreach (var (dx, dy) in Directions())
            {
                var k = 1;
                for (; isValue(x + k * dx, y + k * dy, 'O'); k++) { }
                if (isValue(x + k * dx, y + k * dy, 'X'))
                {
                    for (k--; k > 0; k--)
                    {
                        Q.Enqueue((x + k * dx, y + k * dy));
                        B[x + k * dx][y + k * dy] = 'X';
                    }
                }
            }
        }
        return init - countO();
    }

    public int FlipChess(string[] chessboard) => Enumerable
        .Range(0, chessboard.Length)
        .SelectMany(i => Enumerable
            .Range(0, chessboard[0].Length)
            .Where(j => chessboard[i][j] is '.')
            .Select(j => Bfs(chessboard, i, j)))
        .Max();
}
