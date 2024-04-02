/*
 * @lc app=leetcode.cn id=minimum-additions-to-make-valid-string lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6375] 构造有效字符串的最少插入数
 *
 * https://leetcode.cn/problems/minimum-additions-to-make-valid-string/description/
 *
 * algorithms
 * Medium (50.37%)
 * Total Accepted:    3.8K
 * Total Submissions: 7.4K
 * Testcase Example:  '"b"'
 *
 * 给你一个字符串 word ，你可以向其中任何位置插入 "a"、"b" 或 "c" 任意次，返回使 word 有效 需要插入的最少字母数。
 * 
 * 如果字符串可以由 "abc" 串联多次得到，则认为该字符串 有效 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：word = "b"
 * 输出：2
 * 解释：在 "b" 之前插入 "a" ，在 "b" 之后插入 "c" 可以得到有效字符串 "abc" 。
 * 
 * 
 * 示例 2：
 * 
 * 输入：word = "aaa"
 * 输出：6
 * 解释：在每个 "a" 之后依次插入 "b" 和 "c" 可以得到有效字符串 "abcabcabc" 。
 * 
 * 
 * 示例 3：
 * 
 * 输入：word = "abc"
 * 输出：0
 * 解释：word 已经是有效字符串，不需要进行修改。 
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= word.length <= 50
 * word 仅由字母 "a"、"b" 和 "c" 组成。
 * 
 * 
 */
public class Solution
{
    public int AddMinimum(string word)
    {
        var statuses = word.Select(ch => (int)(ch - 'a'));
        var (ans, expect) = (0, 0);
        foreach (var s in statuses)
        {
            for (; expect % 3 != s; expect++)
            {
                ans++;
            }
            expect++;
        }
        for (; expect % 3 != 0; expect++)
        {
            ans++;
        }
        return ans;
    }
}
