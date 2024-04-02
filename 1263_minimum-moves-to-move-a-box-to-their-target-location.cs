/*
 * @lc app=leetcode.cn id=minimum-moves-to-move-a-box-to-their-target-location lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1263] 推箱子
 *
 * https://leetcode.cn/problems/minimum-moves-to-move-a-box-to-their-target-location/description/
 *
 * algorithms
 * Hard (43.93%)
 * Total Accepted:    9.4K
 * Total Submissions: 18.2K
 * Testcase Example:  '[["#","#","#","#","#","#"],["#","T","#","#","#","#"],["#",".",".","B",".","#"],["#",".","#","#",".","#"],["#",".",".",".","S","#"],["#","#","#","#","#","#"]]'
 *
 * 「推箱子」是一款风靡全球的益智小游戏，玩家需要将箱子推到仓库中的目标位置。
 * 
 * 游戏地图用大小为 m x n 的网格 grid 表示，其中每个元素可以是墙、地板或者是箱子。
 * 
 * 现在你将作为玩家参与游戏，按规则将箱子 'B' 移动到目标位置 'T' ：
 * 
 * 
 * 玩家用字符 'S' 表示，只要他在地板上，就可以在网格中向上、下、左、右四个方向移动。
 * 地板用字符 '.' 表示，意味着可以自由行走。
 * 墙用字符 '#' 表示，意味着障碍物，不能通行。 
 * 箱子仅有一个，用字符 'B' 表示。相应地，网格上有一个目标位置 'T'。
 * 玩家需要站在箱子旁边，然后沿着箱子的方向进行移动，此时箱子会被移动到相邻的地板单元格。记作一次「推动」。
 * 玩家无法越过箱子。
 * 
 * 
 * 返回将箱子推到目标位置的最小 推动 次数，如果无法做到，请返回 -1。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 
 * 输入：grid = [["#","#","#","#","#","#"],
 * ⁠            ["#","T","#","#","#","#"],
 * ["#",".",".","B",".","#"],
 * ["#",".","#","#",".","#"],
 * ["#",".",".",".","S","#"],
 * ["#","#","#","#","#","#"]]
 * 输出：3
 * 解释：我们只需要返回推箱子的次数。
 * 
 * 示例 2：
 * 
 * 
 * 输入：grid = [["#","#","#","#","#","#"],
 * ⁠            ["#","T","#","#","#","#"],
 * ["#",".",".","B",".","#"],
 * ["#","#","#","#",".","#"],
 * ["#",".",".",".","S","#"],
 * ["#","#","#","#","#","#"]]
 * 输出：-1
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：grid = [["#","#","#","#","#","#"],
 * ["#","T",".",".","#","#"],
 * ["#",".","#","B",".","#"],
 * ["#",".",".",".",".","#"],
 * ["#",".",".",".","S","#"],
 * ["#","#","#","#","#","#"]]
 * 输出：5
 * 解释：向下、向左、向左、向上再向上。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * m == grid.length
 * n == grid[i].length
 * 1 <= m, n <= 20
 * grid 仅包含字符 '.', '#',  'S' , 'T', 以及 'B'。
 * grid 中 'S', 'B' 和 'T' 各只能出现一个。
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/minimum-moves-to-move-a-box-to-their-target-location/solution/tui-xiang-zi-by-leetcode-solution-spzi/
// 需要双端队列 or 两个队列来回捯.
public class Solution
{
    private int[] Directions = new[] { 0, -1, 0, 1, 0, };

    public int MinPushBox(char[][] G)
    {
        var (n, m) = (G.Length, G[0].Length);
        bool isValid(int x, int y) => 0 <= x && x < n && 0 <= y && y < m && G[x][y] != '#';
        int gv(int x, int y) => x * m + y;
        var (isx, isy, ibx, iby) = (0, 0, 0, 0);
        for (var i = 0; i < n; i++)
        {
            for (var j = 0; j < m; j++)
            {
                if (G[i][j] == 'S')
                {
                    (isx, isy) = (i, j);
                }
                if (G[i][j] == 'B')
                {
                    (ibx, iby) = (i, j);
                }
            }
        }
        var Q = new Queue<(int, int, int, int)>();
        var dist = new int[n * m][];
        for (var i = 0; i < n * m; i++)
        {
            dist[i] = new int[n * m];
            Array.Fill(dist[i], int.MaxValue);
        }
        Q.Enqueue((isx, isy, ibx, iby));
        dist[gv(isx, isy)][gv(ibx, iby)] = 0;
        while (Q.Count > 0)
        {
            var nextQ = new Queue<(int, int, int, int)>();
            while (Q.Count > 0)
            {
                var (sx, sy, bx, by) = Q.Dequeue();
                if (G[bx][by] == 'T')
                {
                    return dist[gv(sx, sy)][gv(bx, by)];
                }
                for (var i = 0; i < 4; i++)
                {
                    var (dx, dy) = (Directions[i], Directions[i + 1]);
                    var (nsx, nsy) = (sx + dx, sy + dy);
                    if (!isValid(nsx, nsy))
                    {
                        continue;
                    }
                    if ((nsx, nsy) == (bx, by))
                    {
                        var (nbx, nby) = (bx + dx, by + dy);
                        if (!isValid(nbx, nby)
                            || dist[gv(nsx, nsy)][gv(nbx, nby)] <= dist[gv(sx, sy)][gv(bx, by)] + 1)
                        {
                            continue;
                        }
                        nextQ.Enqueue((nsx, nsy, nbx, nby));
                        dist[gv(nsx, nsy)][gv(nbx, nby)] = dist[gv(sx, sy)][gv(bx, by)] + 1;
                    }
                    else
                    {
                        if (dist[gv(nsx, nsy)][gv(bx, by)] <= dist[gv(sx, sy)][gv(bx, by)])
                        {
                            continue;
                        }
                        Q.Enqueue((nsx, nsy, bx, by));
                        dist[gv(nsx, nsy)][gv(bx, by)] = dist[gv(sx, sy)][gv(bx, by)];
                    }
                }
            }
            Q = nextQ;
        }
        return -1;
    }
}
