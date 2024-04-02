/*
 * @lc app=leetcode.cn id=split-with-minimum-sum lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6312] 最小和分割
 *
 * https://leetcode.cn/problems/split-with-minimum-sum/description/
 *
 * algorithms
 * Easy (76.24%)
 * Total Accepted:    3.4K
 * Total Submissions: 4.5K
 * Testcase Example:  '4325'
 *
 * 给你一个正整数 num ，请你将它分割成两个非负整数 num1 和 num2 ，满足：
 * 
 * 
 * num1 和 num2 直接连起来，得到 num 各数位的一个排列。
 * 
 * 
 * 换句话说，num1 和 num2 中所有数字出现的次数之和等于 num 中所有数字出现的次数。
 * 
 * 
 * num1 和 num2 可以包含前导 0 。
 * 
 * 
 * 请你返回 num1 和 num2 可以得到的和的 最小 值。
 * 
 * 注意：
 * 
 * 
 * num 保证没有前导 0 。
 * num1 和 num2 中数位顺序可以与 num 中数位顺序不同。
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：num = 4325
 * 输出：59
 * 解释：我们可以将 4325 分割成 num1 = 24 和 num2 = 35 ，和为 59 ，59 是最小和。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：num = 687
 * 输出：75
 * 解释：我们可以将 687 分割成 num1 = 68 和 num2 = 7 ，和为最优值 75 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 10 <= num <= 10^9
 * 
 * 
 */
public class Solution
{
    public int SplitNum(int num)
    {
        var s = num.ToString();
        s = new string(s.OrderBy(ch => ch).ToArray());
        var sbs = new StringBuilder[2];
        for (var i = 0; i < 2; i++)
        {
            sbs[i] = new();
        }
        for (int i = 0, n = s.Length; i < n; i++)
        {
            sbs[i % 2].Append(s[i]);
        }
        return sbs.Select(sb => int.Parse(sb.ToString())).Sum();
    }
}
