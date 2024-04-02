/*
 * @lc app=leetcode.cn id=number-of-valid-clock-times lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [2437] 有效时间的数目
 *
 * https://leetcode.cn/problems/number-of-valid-clock-times/description/
 *
 * algorithms
 * Easy (44.48%)
 * Total Accepted:    10.9K
 * Total Submissions: 22.2K
 * Testcase Example:  '"?5:00"'
 *
 * 给你一个长度为 5 的字符串 time ，表示一个电子时钟当前的时间，格式为 "hh:mm" 。最早 可能的时间是 "00:00" ，最晚 可能的时间是
 * "23:59" 。
 * 
 * 在字符串 time 中，被字符 ? 替换掉的数位是 未知的 ，被替换的数字可能是 0 到 9 中的任何一个。
 * 
 * 请你返回一个整数 answer ，将每一个 ? 都用 0 到 9 中一个数字替换后，可以得到的有效时间的数目。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：time = "?5:00"
 * 输出：2
 * 解释：我们可以将 ? 替换成 0 或 1 ，得到 "05:00" 或者 "15:00" 。注意我们不能替换成 2 ，因为时间 "25:00"
 * 是无效时间。所以我们有两个选择。
 * 
 * 
 * 示例 2：
 * 
 * 输入：time = "0?:0?"
 * 输出：100
 * 解释：两个 ? 都可以被 0 到 9 之间的任意数字替换，所以我们总共有 100 种选择。
 * 
 * 
 * 示例 3：
 * 
 * 输入：time = "??:??"
 * 输出：1440
 * 解释：小时总共有 24 种选择，分钟总共有 60 种选择。所以总共有 24 * 60 = 1440 种选择。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * time 是一个长度为 5 的有效字符串，格式为 "hh:mm" 。
 * "00" <= hh <= "23"
 * "00" <= mm <= "59"
 * 字符串中有的数位是 '?' ，需要用 0 到 9 之间的数字替换。
 * 
 * 
 */
public class Solution
{
    public int CountTime(string time) => Enumerable
        .Range(0, 24)
        .Select(h => Enumerable
            .Range(0, 60)
            .Select(m => $"{h.ToString("D2")}:{m.ToString("D2")}")
            .Where(s => Enumerable
                .Range(0, time.Length)
                .All(i => time[i] == '?' || time[i] == s[i]))
            .Count())
        .Sum();
}
