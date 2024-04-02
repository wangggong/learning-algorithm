/*
 * @lc app=leetcode.cn id=circle-and-rectangle-overlapping lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1401] 圆和矩形是否有重叠
 *
 * https://leetcode.cn/problems/circle-and-rectangle-overlapping/description/
 *
 * algorithms
 * Medium (43.54%)
 * Total Accepted:    10.9K
 * Total Submissions: 22.2K
 * Testcase Example:  '1\n0\n0\n1\n-1\n3\n1'
 *
 * 给你一个以 (radius, xCenter, yCenter) 表示的圆和一个与坐标轴平行的矩形 (x1, y1, x2, y2) ，其中 (x1,
 * y1) 是矩形左下角的坐标，而 (x2, y2) 是右上角的坐标。
 * 
 * 如果圆和矩形有重叠的部分，请你返回 true ，否则返回 false 。
 * 
 * 换句话说，请你检测是否 存在 点 (xi, yi) ，它既在圆上也在矩形上（两者都包括点落在边界上的情况）。
 * 
 * 
 * 
 * 示例 1 ：
 * 
 * 
 * 输入：radius = 1, xCenter = 0, yCenter = 0, x1 = 1, y1 = -1, x2 = 3, y2 = 1
 * 输出：true
 * 解释：圆和矩形存在公共点 (1,0) 。
 * 
 * 
 * 示例 2 ：
 * 
 * 
 * 输入：radius = 1, xCenter = 1, yCenter = 1, x1 = 1, y1 = -3, x2 = 2, y2 = -1
 * 输出：false
 * 
 * 
 * 示例 3 ：
 * 
 * 
 * 输入：radius = 1, xCenter = 0, yCenter = 0, x1 = -1, y1 = 0, x2 = 0, y2 = 1
 * 输出：true
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= radius <= 2000
 * -10^4 <= xCenter, yCenter <= 10^4
 * -10^4 <= x1 < x2 <= 10^4
 * -10^4 <= y1 < y2 <= 10^4
 * 
 * 
 */
public class Solution
{
    private int Distance(int x1, int y1, int x2, int y2) =>
        (x1 - x2) * (x1 - x2) + (y1 - y2) * (y1 - y2);

    public bool CheckOverlap(int radius, int xCenter, int yCenter, int x1, int y1, int x2, int y2)
        => (x1 <= xCenter && xCenter <= x2,
            y1 <= yCenter && yCenter <= y2) switch
        {
            (true, true) => true,
            (true, _) => Math.Min(
                Math.Abs(yCenter - y1),
                Math.Abs(yCenter - y2)) <= radius,
            (_, true) => Math.Min(
                Math.Abs(xCenter - x1),
                Math.Abs(xCenter - x2)) <= radius,
            _ => Math.Min(
                Math.Min(
                    Distance(xCenter, yCenter, x1, y1),
                    Distance(xCenter, yCenter, x1, y2)),
                Math.Min(
                    Distance(xCenter, yCenter, x2, y1),
                    Distance(xCenter, yCenter, x2, y2))) <= radius * radius,
        };
}
