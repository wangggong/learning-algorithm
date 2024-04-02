/*
 * @lc app=leetcode.cn id=find-longest-special-substring-that-occurs-thrice-ii lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [100184] 找出出现至少三次的最长特殊子字符串 II
 *
 * https://leetcode.cn/problems/find-longest-special-substring-that-occurs-thrice-ii/description/
 *
 * algorithms
 * Medium (27.20%)
 * Total Accepted:    1.6K
 * Total Submissions: 6K
 * Testcase Example:  '"aaaa"'
 *
 * 给你一个仅由小写英文字母组成的字符串 s 。
 * 
 * 如果一个字符串仅由单一字符组成，那么它被称为 特殊 字符串。例如，字符串 "abc" 不是特殊字符串，而字符串 "ddd"、"zz" 和 "f"
 * 是特殊字符串。
 * 
 * 返回在 s 中出现 至少三次 的 最长特殊子字符串 的长度，如果不存在出现至少三次的特殊子字符串，则返回 -1 。
 * 
 * 子字符串 是字符串中的一个连续 非空 字符序列。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s = "aaaa"
 * 输出：2
 * 解释：出现三次的最长特殊子字符串是 "aa" ：子字符串 "aaaa"、"aaaa" 和 "aaaa"。
 * 可以证明最大长度是 2 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：s = "abcdef"
 * 输出：-1
 * 解释：不存在出现至少三次的特殊子字符串。因此返回 -1 。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：s = "abcaba"
 * 输出：1
 * 解释：出现三次的最长特殊子字符串是 "a" ：子字符串 "abcaba"、"abcaba" 和 "abcaba"。
 * 可以证明最大长度是 1 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 3 <= s.length <= 5 * 10^5
 * s 仅由小写英文字母组成。
 * 
 * 
 */
public class Solution
{
    public int MaximumLength(string s)
    {
        IEnumerable<int> getLengths(char c)
        {
            for (var (p, q, n) = (0, 0, s.Length); p < n; p = q)
            {
                for (; p < n && s[p] != c; p++) { }
                if (p == n) { yield break; }
                for (q = p; q < n && s[q] == c; q++) { }
                yield return q - p;
            }
        }
        int getLength(char c)
        {
            var list = getLengths(c)
                .OrderByDescending(x => x)
                .ToList();
            for (var k = list.FirstOrDefault(); k > 0; k--)
            {
                if (list.Select(l => Math.Max(l - k + 1, 0)).Sum() >= 3)
                { return k; }
            }
            return -1;
        }
        var ans = -1;
        for (var c = 'a'; c <= 'z'; c++) { ans = Math.Max(ans, getLength(c)); }
        return ans;
    }
}
