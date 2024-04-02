/*
 * @lc app=leetcode.cn id=sum-of-beauty-of-all-substrings lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1781] 所有子字符串美丽值之和
 *
 * https://leetcode.cn/problems/sum-of-beauty-of-all-substrings/description/
 *
 * algorithms
 * Medium (54.46%)
 * Total Accepted:    9.8K
 * Total Submissions: 15.9K
 * Testcase Example:  '"aabcb"'
 *
 * 一个字符串的 美丽值 定义为：出现频率最高字符与出现频率最低字符的出现次数之差。
 * 
 * 
 * 比方说，"abaacc" 的美丽值为 3 - 1 = 2 。
 * 
 * 
 * 给你一个字符串 s ，请你返回它所有子字符串的 美丽值 之和。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s = "aabcb"
 * 输出：5
 * 解释：美丽值不为零的字符串包括 ["aab","aabc","aabcb","abcb","bcb"] ，每一个字符串的美丽值都为 1 。
 * 
 * 示例 2：
 * 
 * 
 * 输入：s = "aabcbaa"
 * 输出：17
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * s 只包含小写英文字母。
 * 
 * 
 */
public class Solution
{
    private const int AlphaBetaCnt = 26;

    public int BeautySum(string s)
    {
        var n = s.Count();
        var ps = new int[n + 1][];
        for (var i = 0; i <= n; i++)
        {
            ps[i] = new int[AlphaBetaCnt];
        }
        for (var i = 0; i < n; i++)
        {
            for (var a = 0; a < AlphaBetaCnt; a++)
            {
                ps[i + 1][a] = ps[i][a] + (a == s[i] - 'a' ? 1 : 0);
            }
        }
        var ans = 0;
        for (var i = 0; i < n; i++)
        {
            for (var j = i + 1; j <= n; j++)
            {
                var max = 0;
                var min = n + 1;
                for (var a = 0; a < AlphaBetaCnt; a++)
                {
                    var cur = ps[j][a] - ps[i][a];
                    if (cur == 0)
                    {
                        continue;
                    }
                    max = Math.Max(max, cur);
                    min = Math.Min(min, cur);
                }
                ans += max - min;
            }
        }
        return ans;
    }
}
