/*
 * @lc app=leetcode.cn id=walking-robot-simulation lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [874] 模拟行走机器人
 *
 * https://leetcode.cn/problems/walking-robot-simulation/description/
 *
 * algorithms
 * Medium (43.13%)
 * Total Accepted:    42.5K
 * Total Submissions: 90.3K
 * Testcase Example:  '[4,-1,3]\n[]'
 *
 * 机器人在一个无限大小的 XY 网格平面上行走，从点 (0, 0) 处开始出发，面向北方。该机器人可以接收以下三种类型的命令 commands
 * ：
 * 
 * 
 * -2 ：向左转 90 度
 * -1 ：向右转 90 度
 * 1  ：向前移动 x 个单位长度
 * 
 * 
 * 在网格上有一些格子被视为障碍物 obstacles 。第 i 个障碍物位于网格点  obstacles[i] = (xi, yi) 。
 * 
 * 机器人无法走到障碍物上，它将会停留在障碍物的前一个网格方块上，但仍然可以继续尝试进行该路线的其余部分。
 * 
 * 返回从原点到机器人所有经过的路径点（坐标为整数）的最大欧式距离的平方。（即，如果距离为 5 ，则返回 25 ）
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * 注意：
 * 
 * 
 * 北表示 +Y 方向。
 * 东表示 +X 方向。
 * 南表示 -Y 方向。
 * 西表示 -X 方向。
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：commands = [4,-1,3], obstacles = []
 * 输出：25
 * 解释：
 * 机器人开始位于 (0, 0)：
 * 1. 向北移动 4 个单位，到达 (0, 4)
 * 2. 右转
 * 3. 向东移动 3 个单位，到达 (3, 4)
 * 距离原点最远的是 (3, 4) ，距离为 3^2 + 4^2 = 25
 * 
 * 示例 2：
 * 
 * 
 * 输入：commands = [4,-1,4,-2,4], obstacles = [[2,4]]
 * 输出：65
 * 解释：机器人开始位于 (0, 0)：
 * 1. 向北移动 4 个单位，到达 (0, 4)
 * 2. 右转
 * 3. 向东移动 1 个单位，然后被位于 (2, 4) 的障碍物阻挡，机器人停在 (1, 4)
 * 4. 左转
 * 5. 向北走 4 个单位，到达 (1, 8)
 * 距离原点最远的是 (1, 8) ，距离为 1^2 + 8^2 = 65
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * commands[i] is one of the values in the list [-2,-1,1,2,3,4,5,6,7,8,9].
 * 0 
 * -3 * 10^4 i, yi 
 * 答案保证小于 2^31
 * 
 * 
 */
public class Solution
{
    private int[] directions = new[] { 0, 1, 0, -1, 0, };
    private const int D = 4;

    public int RobotSim(int[] commands, int[][] obstacles)
    {
        var S = obstacles.Select(o => (o[0], o[1])).ToHashSet();
        var (x, y) = (0, 0);
        var d = 0;
        var ans = 0;
        foreach (var c in commands)
        {
            switch (c)
            {
                case -2:
                    d = (d + D - 1) % D;
                    break;
                case -1:
                    d = (d + 1) % D;
                    break;
                default:
                    for (var i = 0; i < c; i++)
                    {
                        var (nx, ny) = (x + directions[d], y + directions[(d + 1) % D]);
                        if (S.Contains((nx, ny))) { break; }
                        (x, y) = (nx, ny);
                    }
                    break;
            }
            ans = Math.Max(ans, x * x + y * y);
        }
        return ans;
    }
}
