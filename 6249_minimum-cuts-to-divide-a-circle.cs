/*
 * @lc app=leetcode.cn id=minimum-cuts-to-divide-a-circle lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6249] 分割圆的最少切割次数
 *
 * https://leetcode.cn/problems/minimum-cuts-to-divide-a-circle/description/
 *
 * algorithms
 * Easy (49.47%)
 * Total Accepted:    2.8K
 * Total Submissions: 5.6K
 * Testcase Example:  '4'
 *
 * 圆内一个 有效切割 ，符合以下二者之一：
 * 
 * 
 * 该切割是两个端点在圆上的线段，且该线段经过圆心。
 * 该切割是一端在圆心另一端在圆上的线段。
 * 
 * 
 * 一些有效和无效的切割如下图所示。
 * 
 * 
 * 
 * 给你一个整数 n ，请你返回将圆切割成相等的 n 等分的 最少 切割次数。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 
 * 输入：n = 4
 * 输出：2
 * 解释：
 * 上图展示了切割圆 2 次，得到四等分。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 
 * 
 * 输入：n = 3
 * 输出：3
 * 解释：
 * 最少需要切割 3 次，将圆切成三等分。
 * 少于 3 次切割无法将圆切成大小相等面积相同的 3 等分。
 * 同时可以观察到，第一次切割无法将圆切割开。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 100
 * 
 * 
 */
public class Solution
{
    public int NumberOfCuts(int n) => n == 1 ? 0 : (n % 2 == 0 ? n / 2 : n);
}
