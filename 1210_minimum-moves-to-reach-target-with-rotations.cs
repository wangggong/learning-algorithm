/*
 * @lc app=leetcode.cn id=minimum-moves-to-reach-target-with-rotations lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1210] 穿过迷宫的最少移动次数
 *
 * https://leetcode.cn/problems/minimum-moves-to-reach-target-with-rotations/description/
 *
 * algorithms
 * Hard (47.13%)
 * Total Accepted:    3.9K
 * Total Submissions: 7.9K
 * Testcase Example:  '[[0,0,0,0,0,1],[1,1,0,0,1,0],[0,0,0,0,1,1],[0,0,1,0,1,0],[0,1,1,0,0,0],[0,1,1,0,0,0]]\r'
 *
 * 你还记得那条风靡全球的贪吃蛇吗？
 * 
 * 我们在一个 n*n 的网格上构建了新的迷宫地图，蛇的长度为 2，也就是说它会占去两个单元格。蛇会从左上角（(0, 0) 和 (0,
 * 1)）开始移动。我们用 0 表示空单元格，用 1 表示障碍物。蛇需要移动到迷宫的右下角（(n-1, n-2) 和 (n-1, n-1)）。
 * 
 * 每次移动，蛇可以这样走：
 * 
 * 
 * 如果没有障碍，则向右移动一个单元格。并仍然保持身体的水平／竖直状态。
 * 如果没有障碍，则向下移动一个单元格。并仍然保持身体的水平／竖直状态。
 * 如果它处于水平状态并且其下面的两个单元都是空的，就顺时针旋转 90 度。蛇从（(r, c)、(r, c+1)）移动到 （(r, c)、(r+1,
 * c)）。
 * 
 * 如果它处于竖直状态并且其右面的两个单元都是空的，就逆时针旋转 90 度。蛇从（(r, c)、(r+1, c)）移动到（(r, c)、(r,
 * c+1)）。
 * 
 * 
 * 
 * 返回蛇抵达目的地所需的最少移动次数。
 * 
 * 如果无法到达目的地，请返回 -1。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 输入：grid = [[0,0,0,0,0,1],
 * ⁠              [1,1,0,0,1,0],
 * [0,0,0,0,1,1],
 * [0,0,1,0,1,0],
 * [0,1,1,0,0,0],
 * [0,1,1,0,0,0]]
 * 输出：11
 * 解释：
 * 一种可能的解决方案是 [右, 右, 顺时针旋转, 右, 下, 下, 下, 下, 逆时针旋转, 右, 下]。
 * 
 * 
 * 示例 2：
 * 
 * 输入：grid = [[0,0,1,1,1,1],
 * [0,0,0,0,1,1],
 * [1,1,0,0,0,1],
 * [1,1,1,0,0,1],
 * [1,1,1,0,0,1],
 * [1,1,1,0,0,0]]
 * 输出：9
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 2 <= n <= 100
 * 0 <= grid[i][j] <= 1
 * 蛇保证从空单元格开始出发。
 * 
 * 
 */
public class Solution
{
    private record Position(int X0, int Y0, int X1, int Y1);

    public int MinimumMoves(int[][] G)
    {
        var n = G.Length;
        bool isFinish(Position pos) => pos.Y0 == n - 1 && pos.Y1 == n - 1 && (pos.X0 + pos.X1 == 2 * n - 3);
        bool isValidPos(int x, int y) => 0 <= x && x < n && 0 <= y && y < n && G[y][x] == 0;
        IEnumerable<Position> nextPositions(Position pos)
        {
            yield return new Position(pos.X0 + 1, pos.Y0, pos.X1 + 1, pos.Y1);
            yield return new Position(pos.X0, pos.Y0 + 1, pos.X1, pos.Y1 + 1);
            if (pos.Y0 == pos.Y1 && isValidPos(pos.X0, pos.Y0 + 1) && isValidPos(pos.X1, pos.Y1 + 1))
            {
                yield return new Position(pos.X0, pos.Y0, pos.X0, pos.Y0 + 1);
            }
            if (pos.X0 == pos.X1 && isValidPos(pos.X0 + 1, pos.Y0) && isValidPos(pos.X1 + 1, pos.Y1))
            {
                yield return new Position(pos.X0, pos.Y0, pos.X0 + 1, pos.Y0);
            }
        }
        var Q = new Queue<Position>();
        var visit = new Dictionary<Position, int>();
        var initPos = new Position(0, 0, 1, 0);
        Q.Enqueue(initPos);
        visit[initPos] = 0;
        while (Q.Count > 0)
        {
            var pos = Q.Dequeue();
            var steps = visit[pos];
            if (isFinish(pos))
            {
                return steps;
            }
            foreach (var nextPos in nextPositions(pos))
            {
                if (isValidPos(nextPos.X0, nextPos.Y0) && isValidPos(nextPos.X1, nextPos.Y1) && !visit.ContainsKey(nextPos))
                {
                    Q.Enqueue(nextPos);
                    visit[nextPos] = steps + 1;
                }
            }
        }
        return -1;
    }
}
