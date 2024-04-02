/*
 * @lc app=leetcode.cn id=count-the-number-of-consistent-strings lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1684] 统计一致字符串的数目
 *
 * https://leetcode.cn/problems/count-the-number-of-consistent-strings/description/
 *
 * algorithms
 * Easy (82.07%)
 * Total Accepted:    42.3K
 * Total Submissions: 49.3K
 * Testcase Example:  '"ab"\n["ad","bd","aaab","baa","badab"]'
 *
 * 给你一个由不同字符组成的字符串 allowed 和一个字符串数组 words 。如果一个字符串的每一个字符都在 allowed 中，就称这个字符串是
 * 一致字符串 。
 * 
 * 请你返回 words 数组中 一致字符串 的数目。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：allowed = "ab", words = ["ad","bd","aaab","baa","badab"]
 * 输出：2
 * 解释：字符串 "aaab" 和 "baa" 都是一致字符串，因为它们只包含字符 'a' 和 'b' 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：allowed = "abc", words = ["a","b","c","ab","ac","bc","abc"]
 * 输出：7
 * 解释：所有字符串都是一致的。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：allowed = "cad", words = ["cc","acd","b","ba","bac","bad","ac","d"]
 * 输出：4
 * 解释：字符串 "cc"，"acd"，"ac" 和 "d" 是一致字符串。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= words.length <= 104
 * 1 <= allowed.length <= 26
 * 1 <= words[i].length <= 10
 * allowed 中的字符 互不相同 。
 * words[i] 和 allowed 只包含小写英文字母。
 * 
 * 
 */
public class Solution
{
    public int CountConsistentStrings(string allowed, string[] words)
    {
        int mask = 0, ans = 0;
        foreach (var a in allowed) { mask |= 1 << (a - 'a'); }
        foreach (var word in words)
        {
            var valid = true;
            foreach (var w in word)
            {
                if ((mask & (1 << (w - 'a'))) == 0)
                {
                    valid = false;
                    break;
                }
            }
            if (valid) { ans++; }
        }
        return ans;
    }
}
