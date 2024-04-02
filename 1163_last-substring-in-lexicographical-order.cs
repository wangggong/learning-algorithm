/*
 * @lc app=leetcode.cn id=last-substring-in-lexicographical-order lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1163] 按字典序排在最后的子串
 *
 * https://leetcode.cn/problems/last-substring-in-lexicographical-order/description/
 *
 * algorithms
 * Hard (27.87%)
 * Total Accepted:    8.7K
 * Total Submissions: 30.4K
 * Testcase Example:  '"abab"'
 *
 * 给你一个字符串 s ，找出它的所有子串并按字典序排列，返回排在最后的那个子串。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s = "abab"
 * 输出："bab"
 * 解释：我们可以找出 7 个子串 ["a", "ab", "aba", "abab", "b", "ba", "bab"]。按字典序排在最后的子串是
 * "bab"。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：s = "leetcode"
 * 输出："tcode"
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= s.length <= 4 * 10^5
 * s 仅含有小写英文字符。
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/last-substring-in-lexicographical-order/solution/an-zi-dian-xu-pai-zai-zui-hou-de-zi-chua-31yl/
//
// 至少应该注意到这样一个事实: 最大的字串一定是后缀.
// PS: 想到了 trivial 的用最大 index 过滤的方式. 但 C# 的常数太大了.
public class Solution
{
    public string LastSubstring(string s)
    {
        var (i, j, n) = (0, 1, s.Length);
        while (j < n)
        {
            var k = 0;
            for (; j + k < n && s[i + k] == s[j + k]; k++) { }
            if (j + k < n && s[i + k] < s[j + k])
            {
                (i, j) = (j, Math.Max(j + 1, i + k + 1));
            }
            else
            {
                j += k + 1;
            }
        }
        return s.Substring(i);
    }
}
