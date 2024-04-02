/*
 * @lc app=leetcode.cn id=maximum-length-substring-with-two-occurrences lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [100245] 每个字符最多出现两次的最长子字符串
 *
 * https://leetcode.cn/problems/maximum-length-substring-with-two-occurrences/description/
 *
 * algorithms
 * Easy (64.87%)
 * Total Accepted:    3.7K
 * Total Submissions: 5.7K
 * Testcase Example:  '"bcbbbcba"'
 *
 * 给你一个字符串 s ，请找出满足每个字符最多出现两次的最长子字符串，并返回该子字符串的 最大 长度。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入： s = "bcbbbcba"
 * 
 * 输出： 4
 * 
 * 解释：
 * 
 * 以下子字符串长度为 4，并且每个字符最多出现两次："bcbbbcba"。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入： s = "aaaa"
 * 
 * 输出： 2
 * 
 * 解释：
 * 
 * 以下子字符串长度为 2，并且每个字符最多出现两次："aaaa"。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 2 <= s.length <= 100
 * 
 * s 仅由小写英文字母组成。
 * 
 * 
 */
public class Solution
{
    public int MaximumLengthSubstring(string s) => Enumerable.Range(0, s.Length)
        .SelectMany(i => Enumerable.Range(i + 1, s.Length - i)
            .Where(j => s[i..j].GroupBy(c => c)
                .Select(g => g.Count())
                .Max() <= 2)
            .Select(j => j - i))
        .Max();
}
