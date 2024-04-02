/*
 * @lc app=leetcode.cn id=count-substrings-starting-and-ending-with-given-character lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [100236] 统计以给定字符开头和结尾的子字符串总数
 *
 * https://leetcode.cn/problems/count-substrings-starting-and-ending-with-given-character/description/
 *
 * algorithms
 * Medium (50.90%)
 * Total Accepted:    4.8K
 * Total Submissions: 9.2K
 * Testcase Example:  '"abada"\n"a"'
 *
 * 给你一个字符串 s 和一个字符 c 。返回在字符串 s 中并且以 c 字符开头和结尾的非空子字符串的总数。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s = "abada", c = "a"
 * 
 * 输出：6
 * 
 * 解释：以 "a" 开头和结尾的子字符串有： "abada"、"abada"、"abada"、"abada"、"abada"、"abada"。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：s = "zzz", c = "z"
 * 
 * 输出：6
 * 
 * 解释：字符串 s 中总共有 6 个子字符串，并且它们都以 "z" 开头和结尾。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= s.length <= 10^5
 * s 和 c 均由小写英文字母组成。
 * 
 * 
 */
public class Solution
{
    public long CountSubstrings(string s, char ch)
    {
        var count = (long)s.Count(c => c == ch);
        return count * (count + 1) / 2;
    }
}
