/*
 * @lc app=leetcode.cn id=find-champion-i lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [100115] 找到冠军 I
 *
 * https://leetcode.cn/problems/find-champion-i/description/
 *
 * algorithms
 * Easy (77.96%)
 * Total Accepted:    3.9K
 * Total Submissions: 5K
 * Testcase Example:  '[[0,1],[0,0]]'
 *
 * 一场比赛中共有 n 支队伍，按从 0 到  n - 1 编号。
 * 
 * 给你一个下标从 0 开始、大小为 n * n 的二维布尔矩阵 grid 。对于满足 0 <= i, j <= n - 1 且 i != j 的所有 i,
 * j ：如果 grid[i][j] == 1，那么 i 队比 j 队 强 ；否则，j 队比 i 队 强 。
 * 
 * 在这场比赛中，如果不存在某支强于 a 队的队伍，则认为 a 队将会是 冠军 。
 * 
 * 返回这场比赛中将会成为冠军的队伍。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：grid = [[0,1],[0,0]]
 * 输出：0
 * 解释：比赛中有两支队伍。
 * grid[0][1] == 1 表示 0 队比 1 队强。所以 0 队是冠军。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：grid = [[0,0,1],[1,0,1],[0,0,0]]
 * 输出：1
 * 解释：比赛中有三支队伍。
 * grid[1][0] == 1 表示 1 队比 0 队强。
 * grid[1][2] == 1 表示 1 队比 2 队强。
 * 所以 1 队是冠军。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * n == grid.length
 * n == grid[i].length
 * 2 <= n <= 100
 * grid[i][j] 的值为 0 或 1
 * 对于满足 i != j 的所有 i, j ，grid[i][j] != grid[j][i] 均成立
 * 生成的输出满足：如果 a 队比 b 队强，b 队比 c 队强，那么 a 队比 c 队强
 * 
 * 
 */
// 根据 https://leetcode.cn/problems/find-champion-i/solutions/2513102/xiao-yang-xiao-en-bao-li-chu-qi-ji-wei-s-qgav/, 这里一定有解; 因为两两比赛过.
public class Solution
{
    public int FindChampion(int[][] G) => Enumerable
        .Range(0, G.Length)
        .Where(n => Enumerable
            .Range(0, G.Length)
            .All(m => m == n || G[m][n] is 0))
        .First();
}
