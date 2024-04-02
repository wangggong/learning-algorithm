/*
 * @lc app=leetcode.cn id=make-number-of-distinct-characters-equal lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6284] 使字符串总不同字符的数目相等
 *
 * https://leetcode.cn/problems/make-number-of-distinct-characters-equal/description/
 *
 * algorithms
 * Medium (22.88%)
 * Total Accepted:    2.5K
 * Total Submissions: 11K
 * Testcase Example:  '"ac"\n"b"'
 *
 * 给你两个下标从 0 开始的字符串 word1 和 word2 。
 * 
 * 一次 移动 由以下两个步骤组成：
 * 
 * 
 * 选中两个下标 i 和 j ，分别满足 0 <= i < word1.length 和 0 <= j < word2.length ，
 * 交换 word1[i] 和 word2[j] 。
 * 
 * 
 * 如果可以通过 恰好一次 移动，使 word1 和 word2 中不同字符的数目相等，则返回 true ；否则，返回 false 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：word1 = "ac", word2 = "b"
 * 输出：false
 * 解释：交换任何一组下标都会导致第一个字符串中有 2 个不同的字符，而在第二个字符串中只有 1 个不同字符。
 * 
 * 
 * 示例 2：
 * 
 * 输入：word1 = "abcc", word2 = "aab"
 * 输出：true
 * 解释：交换第一个字符串的下标 2 和第二个字符串的下标 0 。之后得到 word1 = "abac" 和 word2 = "cab" ，各有 3
 * 个不同字符。
 * 
 * 
 * 示例 3：
 * 
 * 输入：word1 = "abcde", word2 = "fghij"
 * 输出：true
 * 解释：无论交换哪一组下标，两个字符串中都会有 5 个不同字符。
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= word1.length, word2.length <= 10^5
 * word1 和 word2 仅由小写英文字母组成。
 * 
 * 
 */
public class Solution
{
    private const int AlphaCnt = 26;

    public bool IsItPossible(string word1, string word2)
    {
        var count1 = new int[AlphaCnt];
        var count2 = new int[AlphaCnt];
        foreach (var ch in word1)
        {
            count1[ch - 'a']++;
        }
        foreach (var ch in word2)
        {
            count2[ch - 'a']++;
        }
        for (var i = 0; i < AlphaCnt; i++)
        {
            if (count1[i] > 0)
            {
                for (var j = 0; j < AlphaCnt; j++)
                {
                    if (count2[j] > 0)
                    {
                        count1[i]--; count2[i]++; count1[j]++; count2[j]--;
                        if (count1.Where(x => x > 0).Count() == count2.Where(x => x > 0).Count())
                        {
                            return true;
                        }
                        count1[i]++; count2[i]--; count1[j]--; count2[j]++;
                    }
                }
            }
        }
        return false;
    }
}
