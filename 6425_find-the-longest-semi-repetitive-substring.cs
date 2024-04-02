/*
 * @lc app=leetcode.cn id=find-the-longest-semi-repetitive-substring lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6425] 找到最长的半重复子字符串
 *
 * https://leetcode.cn/problems/find-the-longest-semi-repetitive-substring/description/
 *
 * algorithms
 * Medium (41.28%)
 * Total Accepted:    1.9K
 * Total Submissions: 4.5K
 * Testcase Example:  '"52233"'
 *
 * 给你一个下标从 0 开始的字符串 s ，这个字符串只包含 0 到 9 的数字字符。
 * 
 * 如果一个字符串 t 中至多有一对相邻字符是相等的，那么称这个字符串是 半重复的 。
 * 
 * 请你返回 s 中最长 半重复 子字符串的长度。
 * 
 * 一个 子字符串 是一个字符串中一段连续 非空 的字符。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s = "52233"
 * 输出：4
 * 解释：最长半重复子字符串是 "5223" ，子字符串从 i = 0 开始，在 j = 3 处结束。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：s = "5494"
 * 输出：4
 * 解释：s 就是一个半重复字符串，所以答案为 4 。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：s = "1111111"
 * 输出：2
 * 解释：最长半重复子字符串是 "11" ，子字符串从 i = 0 开始，在 j = 1 处结束。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= s.length <= 50
 * '0' <= s[i] <= '9'
 * 
 * 
 */
public class Solution
{
    public int LongestSemiRepetitiveSubstring(string s)
    {
        var ans = 1;
        for (var (p, n) = (0, s.Length); p < n; p++)
        {
            var has = false;
            for (var q = p + 1; q < n; q++)
            {
                if (s[q] == s[q - 1])
                {
                    if (has)
                    {
                        break;
                    }
                    has = true;
                }
                ans = Math.Max(ans, q - p + 1);
            }
        }
        return ans;
    }
}
