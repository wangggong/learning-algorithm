/*
 * @lc app=leetcode.cn id=letter-case-permutation lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [784] 字母大小写全排列
 *
 * https://leetcode.cn/problems/letter-case-permutation/description/
 *
 * algorithms
 * Medium (70.35%)
 * Total Accepted:    83.6K
 * Total Submissions: 117K
 * Testcase Example:  '"a1b2"'
 *
 * 给定一个字符串 s ，通过将字符串 s 中的每个字母转变大小写，我们可以获得一个新的字符串。
 * 
 * 返回 所有可能得到的字符串集合 。以 任意顺序 返回输出。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s = "a1b2"
 * 输出：["a1b2", "a1B2", "A1b2", "A1B2"]
 * 
 * 
 * 示例 2:
 * 
 * 
 * 输入: s = "3z4"
 * 输出: ["3z4","3Z4"]
 * 
 * 
 * 
 * 
 * 提示:
 * 
 * 
 * 1 <= s.length <= 12
 * s 由小写英文字母、大写英文字母和数字组成
 * 
 * 
 */
public class Solution
{
    private IList<string> ans = new List<string>();

    public IList<string> LetterCasePermutation(string s)
    {
        traversal(s, "", 0);
        return ans;
    }

    private void traversal(string s, string c, int k)
    {
        if (k == s.Count())
        {
            ans.Add(c);
            return;
        }
        traversal(s, c + s[k], k + 1);
        if ('0' <= s[k] && s[k] <= '9')
        {
            return;
        }
        traversal(s, c + (char) ('a' <= s[k] && s[k] <= 'z' ? s[k] + 'A' - 'a' : s[k] + 'a' - 'A'), k + 1);
        return;
    }
}
