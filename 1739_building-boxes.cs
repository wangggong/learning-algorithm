/*
 * @lc app=leetcode.cn id=building-boxes lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1739] 放置盒子
 *
 * https://leetcode.cn/problems/building-boxes/description/
 *
 * algorithms
 * Hard (49.74%)
 * Total Accepted:    5.3K
 * Total Submissions: 9.6K
 * Testcase Example:  '3'
 *
 * 有一个立方体房间，其长度、宽度和高度都等于 n 个单位。请你在房间里放置 n 个盒子，每个盒子都是一个单位边长的立方体。放置规则如下：
 * 
 * 
 * 你可以把盒子放在地板上的任何地方。
 * 如果盒子 x 需要放置在盒子 y 的顶部，那么盒子 y 竖直的四个侧面都 必须 与另一个盒子或墙相邻。
 * 
 * 
 * 给你一个整数 n ，返回接触地面的盒子的 最少 可能数量。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 
 * 输入：n = 3
 * 输出：3
 * 解释：上图是 3 个盒子的摆放位置。
 * 这些盒子放在房间的一角，对应左侧位置。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 
 * 
 * 输入：n = 4
 * 输出：3
 * 解释：上图是 3 个盒子的摆放位置。
 * 这些盒子放在房间的一角，对应左侧位置。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 
 * 
 * 输入：n = 10
 * 输出：6
 * 解释：上图是 10 个盒子的摆放位置。
 * 这些盒子放在房间的一角，对应后方位置。
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * 
 * 
 */

// 自己做出来了. 具体的可以看: https://leetcode.cn/problems/building-boxes/solution/mei-xiang-ming-bai-yi-ge-dong-hua-miao-d-8vbe/
public class Solution
{
    public int MinimumBoxes(int n)
    {
        var tot = 0;
        var lastLevel = 0;
        var levels = 0;
        var additional = 1;
        for (; tot + lastLevel + levels <= n; levels++)
        {
            lastLevel += levels;
            tot += lastLevel;
        }
        for (; tot < n; additional++)
        {
            lastLevel++;
            tot += additional;
        }
        return lastLevel;
    }
}
