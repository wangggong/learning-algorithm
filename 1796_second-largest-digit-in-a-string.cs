/*
 * @lc app=leetcode.cn id=second-largest-digit-in-a-string lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1796] 字符串中第二大的数字
 *
 * https://leetcode.cn/problems/second-largest-digit-in-a-string/description/
 *
 * algorithms
 * Easy (48.91%)
 * Total Accepted:    13.5K
 * Total Submissions: 27K
 * Testcase Example:  '"dfa12321afd"'
 *
 * 给你一个混合字符串 s ，请你返回 s 中 第二大 的数字，如果不存在第二大的数字，请你返回 -1 。
 * 
 * 混合字符串 由小写英文字母和数字组成。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s = "dfa12321afd"
 * 输出：2
 * 解释：出现在 s 中的数字包括 [1, 2, 3] 。第二大的数字是 2 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：s = "abc1111"
 * 输出：-1
 * 解释：出现在 s 中的数字只包含 [1] 。没有第二大的数字。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * s 只包含小写英文字母和（或）数字。
 * 
 * 
 */
public class Solution
{
    public int SecondHighest(string s)
        => s.Where(ch => '0' <= ch && ch <= '9')
            .Select(ch => ch - '0')
            .Distinct()
            .OrderBy(ch => -ch)
            .Skip(1)
            .FirstOrDefault(-1);
}
