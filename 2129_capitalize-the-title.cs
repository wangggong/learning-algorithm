/*
 * @lc app=leetcode.cn id=capitalize-the-title lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [2129] 将标题首字母大写
 *
 * https://leetcode.cn/problems/capitalize-the-title/description/
 *
 * algorithms
 * Easy (68.56%)
 * Total Accepted:    19.8K
 * Total Submissions: 28.8K
 * Testcase Example:  '"capiTalIze tHe titLe"'
 *
 * 给你一个字符串 title ，它由单个空格连接一个或多个单词组成，每个单词都只包含英文字母。请你按以下规则将每个单词的首字母 大写 ：
 * 
 * 
 * 如果单词的长度为 1 或者 2 ，所有字母变成小写。
 * 否则，将单词首字母大写，剩余字母变成小写。
 * 
 * 
 * 请你返回 大写后 的 title 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：title = "capiTalIze tHe titLe"
 * 输出："Capitalize The Title"
 * 解释：
 * 由于所有单词的长度都至少为 3 ，将每个单词首字母大写，剩余字母变为小写。
 * 
 * 
 * 示例 2：
 * 
 * 输入：title = "First leTTeR of EACH Word"
 * 输出："First Letter of Each Word"
 * 解释：
 * 单词 "of" 长度为 2 ，所以它保持完全小写。
 * 其他单词长度都至少为 3 ，所以其他单词首字母大写，剩余字母小写。
 * 
 * 
 * 示例 3：
 * 
 * 输入：title = "i lOve leetcode"
 * 输出："i Love Leetcode"
 * 解释：
 * 单词 "i" 长度为 1 ，所以它保留小写。
 * 其他单词长度都至少为 3 ，所以其他单词首字母大写，剩余字母小写。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= title.length <= 100
 * title 由单个空格隔开的单词组成，且不含有任何前导或后缀空格。
 * 每个单词由大写和小写英文字母组成，且都是 非空 的。
 * 
 * 
 */
public class Solution
{
    public string CapitalizeTitle(string title)
    {
        char toLower(char c) => 'a' <= c && c <= 'z'
            ? c
            : (char)(c - 'A' + 'a');
        char toUpper(char c) => 'a' <= c && c <= 'z'
            ? (char)(c - 'a' + 'A')
            : c;
        var n = title.Length;
        var ans = new StringBuilder();
        for (var (p, q) = (0, 0); p < n; p = q + 1)
        {
            var sb = new StringBuilder();
            for (q = p; q < n && title[q] != ' '; q++)
            {
                sb.Append(title[q]);
            }
            var s = sb.ToString();
            ans.Append(s.Length > 2 ? toUpper(s[0]) : toLower(s[0]));
            ans.Append(new string(s[1..].Select(c => toLower(c)).ToArray()));
            if (q < n)
            {
                ans.Append(' ');
            }
        }
        return ans.ToString();
    }
}
