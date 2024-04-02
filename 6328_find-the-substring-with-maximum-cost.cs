/*
 * @lc app=leetcode.cn id=find-the-substring-with-maximum-cost lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6328] 找到最大开销的子字符串
 *
 * https://leetcode.cn/problems/find-the-substring-with-maximum-cost/description/
 *
 * algorithms
 * Medium (50.03%)
 * Total Accepted:    3.4K
 * Total Submissions: 6.5K
 * Testcase Example:  '"adaa"\n"d"\n[-1000]'
 *
 * 给你一个字符串 s ，一个字符 互不相同 的字符串 chars 和一个长度与 chars 相同的整数数组 vals 。
 * 
 * 子字符串的开销 是一个子字符串中所有字符对应价值之和。空字符串的开销是 0 。
 * 
 * 字符的价值 定义如下：
 * 
 * 
 * 如果字符不在字符串 chars 中，那么它的价值是它在字母表中的位置（下标从 1 开始）。
 * 
 * 
 * 比方说，'a' 的价值为 1 ，'b' 的价值为 2 ，以此类推，'z' 的价值为 26 。
 * 
 * 
 * 否则，如果这个字符在 chars 中的位置为 i ，那么它的价值就是 vals[i] 。
 * 
 * 
 * 请你返回字符串 s 的所有子字符串中的最大开销。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：s = "adaa", chars = "d", vals = [-1000]
 * 输出：2
 * 解释：字符 "a" 和 "d" 的价值分别为 1 和 -1000 。
 * 最大开销子字符串是 "aa" ，它的开销为 1 + 1 = 2 。
 * 2 是最大开销。
 * 
 * 
 * 示例 2：
 * 
 * 输入：s = "abc", chars = "abc", vals = [-1,-1,-1]
 * 输出：0
 * 解释：字符 "a" ，"b" 和 "c" 的价值分别为 -1 ，-1 和 -1 。
 * 最大开销子字符串是 "" ，它的开销为 0 。
 * 0 是最大开销。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= s.length <= 10^5
 * s 只包含小写英文字母。
 * 1 <= chars.length <= 26
 * chars 只包含小写英文字母，且 互不相同 。
 * vals.length == chars.length
 * -1000 <= vals[i] <= 1000
 * 
 * 
 */

/*
 * public class Solution
 * {
 *     public int MaximumCostSubstring(string s, string chars, int[] vals)
 *     {
 *         var values = Enumerable.Range(0, 26).Select(i => i + 1).ToArray();
 *         var m = chars.Length;
 *         for (var i = 0; i < m; i++)
 *         {
 *             values[(int)(chars[i] - 'a')] = vals[i];
 *         }
 *         var n = s.Length;
 *         var dp = new int[n + 1];
 *         for (var i = 0; i < n; i++)
 *         {
 *             dp[i + 1] = Math.Max(dp[i], 0) + values[(int)(s[i] - 'a')];
 *         }
 *         return dp.Max();
 *     }
 * }
 */

public class Solution
{
    public int MaximumCostSubstring(string s, string chars, int[] vals)
    {
        var values = Enumerable.Range(0, 26).Select(i => i + 1).ToArray();
        for (var (i, m) = (0, chars.Length); i < m; i++)
        {
            values[(int)(chars[i] - 'a')] = vals[i];
        }
        var (ans, cur) = (0, 0);
        for (var (i, n) = (0, s.Length); i < n; i++)
        {
            cur = Math.Max(cur, 0) + values[(int)(s[i] - 'a')];
            ans = Math.Max(ans, cur);
        }
        return ans;
    }
}
