/*
 * @lc app=leetcode.cn id=maximum-product-of-word-lengths lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [318] 最大单词长度乘积
 *
 * https://leetcode.cn/problems/maximum-product-of-word-lengths/description/
 *
 * algorithms
 * Medium (72.42%)
 * Total Accepted:    71.5K
 * Total Submissions: 98.6K
 * Testcase Example:  '["abcw","baz","foo","bar","xtfn","abcdef"]'
 *
 * 给你一个字符串数组 words ，找出并返回 length(words[i]) * length(words[j])
 * 的最大值，并且这两个单词不含有公共字母。如果不存在这样的两个单词，返回 0 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：words = ["abcw","baz","foo","bar","xtfn","abcdef"]
 * 输出：16 
 * 解释：这两个单词为 "abcw", "xtfn"。
 * 
 * 示例 2：
 * 
 * 
 * 输入：words = ["a","ab","abc","d","cd","bcd","abcd"]
 * 输出：4 
 * 解释：这两个单词为 "ab", "cd"。
 * 
 * 示例 3：
 * 
 * 
 * 输入：words = ["a","aa","aaa","aaaa"]
 * 输出：0 
 * 解释：不存在这样的两个单词。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 2 <= words.length <= 1000
 * 1 <= words[i].length <= 1000
 * words[i] 仅包含小写字母
 * 
 * 
 */
public class Solution
{
    public int MaxProduct(string[] words)
    {
        var ans = 0;
        var masks = words
            .Select(w => w
                .Select(c => 1 << (int)(c - 'a'))
                .Aggregate((x, y) => x | y))
            .ToArray();
        var counts = words
            .Select(w => w
                .Count())
            .ToArray();
        for (var (i, n) = (0, words.Length); i < n; i++)
        {
            for (var j = i + 1; j < n; j++)
            {
                if ((masks[i] & masks[j]) is not 0) { continue; }
                ans = Math.Max(ans, counts[i] * counts[j]);
            }
        }
        return ans;
    }
}
