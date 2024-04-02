/*
 * @lc app=leetcode.cn id=swap-for-longest-repeated-character-substring lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1156] 单字符重复子串的最大长度
 *
 * https://leetcode.cn/problems/swap-for-longest-repeated-character-substring/description/
 *
 * algorithms
 * Medium (43.25%)
 * Total Accepted:    9.1K
 * Total Submissions: 20.6K
 * Testcase Example:  '"ababa"'
 *
 * 如果字符串中的所有字符都相同，那么这个字符串是单字符重复的字符串。
 * 
 * 给你一个字符串 text，你只能交换其中两个字符一次或者什么都不做，然后得到一些单字符重复的子串。返回其中最长的子串的长度。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：text = "ababa"
 * 输出：3
 * 
 * 
 * 示例 2：
 * 
 * 输入：text = "aaabaaa"
 * 输出：6
 * 
 * 
 * 示例 3：
 * 
 * 输入：text = "aaabbaaa"
 * 输出：4
 * 
 * 
 * 示例 4：
 * 
 * 输入：text = "aaaaa"
 * 输出：5
 * 
 * 
 * 示例 5：
 * 
 * 输入：text = "abcdef"
 * 输出：1
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= text.length <= 20000
 * text 仅由小写英文字母组成。
 * 
 * 
 */
public class Solution
{
    public int MaxRepOpt1(string text)
    {
        var ans = 0;
        for (var ch = 'a'; ch <= 'z'; ch++)
        {
            var parts = new List<(int, int)>();
            for (var (p, q, n) = (0, 0, text.Length); p < n; p = q)
            {
                for (; p < n && text[p] != ch; p++) { }
                if (p == n) { break; }
                for (q = p; q < n && text[q] == ch; q++) { }
                parts.Add((p, q));
            }
            for (var (i, n) = (0, parts.Count()); i < n; i++)
            {
                var (p, q) = parts[i];
                ans = Math.Max(ans, q - p + (n > 1 ? 1 : 0));
                if (i + 1 < n && parts[i + 1].Item1 == q + 1)
                {
                    ans = Math.Max(ans, parts[i + 1].Item2 - p - (n == 2 ? 1 : 0));
                }
            }
        }
        return ans;
    }
}
