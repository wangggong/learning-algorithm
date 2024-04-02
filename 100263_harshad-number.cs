/*
 * @lc app=leetcode.cn id=harshad-number lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [100263] 哈沙德数
 *
 * https://leetcode.cn/problems/harshad-number/description/
 *
 * algorithms
 * Easy (83.21%)
 * Total Accepted:    3.9K
 * Total Submissions: 4.7K
 * Testcase Example:  '18'
 *
 * 如果一个整数能够被其各个数位上的数字之和整除，则称之为 哈沙德数（Harshad number）。给你一个整数 x 。如果 x 是 哈沙德数 ，则返回
 * x 各个数位上的数字之和，否则，返回 -1 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入： x = 18
 * 
 * 输出： 9
 * 
 * 解释：
 * 
 * x 各个数位上的数字之和为 9 。18 能被 9 整除。因此 18 是哈沙德数，答案是 9 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入： x = 23
 * 
 * 输出： -1
 * 
 * 解释：
 * 
 * x 各个数位上的数字之和为 5 。23 不能被 5 整除。因此 23 不是哈沙德数，答案是 -1 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= x <= 100
 * 
 * 
 */
public class Solution
{
    private int DigitsSum(int x) => x
        .ToString()
        .Select(c => (int)(c - '0'))
        .Sum();

    public int SumOfTheDigitsOfHarshadNumber(int x) => x % DigitsSum(x) == 0
        ? DigitsSum(x)
        : -1;
}
