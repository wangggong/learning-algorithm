/*
 * @lc app=leetcode.cn id=count-palindromic-subsequences lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6251] 统计回文子序列数目
 *
 * https://leetcode.cn/problems/count-palindromic-subsequences/description/
 *
 * algorithms
 * Hard (36.49%)
 * Total Accepted:    1K
 * Total Submissions: 2.8K
 * Testcase Example:  '"103301"'
 *
 * 给你数字字符串 s ，请你返回 s 中长度为 5 的 回文子序列 数目。由于答案可能很大，请你将答案对 10^9 + 7 取余 后返回。
 * 
 * 提示：
 * 
 * 
 * 如果一个字符串从前往后和从后往前读相同，那么它是 回文字符串 。
 * 子序列是一个字符串中删除若干个字符后，不改变字符顺序，剩余字符构成的字符串。
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：s = "103301"
 * 输出：2
 * 解释：
 * 总共有 6 长度为 5 的子序列："10330" ，"10331" ，"10301" ，"10301" ，"13301" ，"03301" 。
 * 它们中有两个（都是 "10301"）是回文的。
 * 
 * 
 * 示例 2：
 * 
 * 输入：s = "0000000"
 * 输出：21
 * 解释：所有 21 个长度为 5 的子序列都是 "00000" ，都是回文的。
 * 
 * 
 * 示例 3：
 * 
 * 输入：s = "9999900000"
 * 输出：2
 * 解释：仅有的两个回文子序列是 "99999" 和 "00000" 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= s.length <= 10^4
 * s 只包含数字字符。
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/count-palindromic-subsequences/solution/by-tsreaper-ncct/
//
// 想到递推了, 没想到怎么维护到某个位置的出现次数.
public class Solution
{
    private const long Mod = (long) 1e9 + 7;
    private const int D = 10;

    public int CountPalindromes(string s)
    {
        var n = s.Count();
        var p1 = new long[n + 2][];
        var p2 = new long[n + 2][,];
        var s1 = new long[n + 2][];
        var s2 = new long[n + 2][,];
        for (int i = 0; i < n + 2; i++)
        {
            p1[i] = new long[D];
            p2[i] = new long[D, D];
            s1[i] = new long[D];
            s2[i] = new long[D, D];
        }
        for (int i = 1; i <= n; i++)
        {
            for (int j = 0; j < D; j++)
            {
                p1[i][j] = p1[i - 1][j];
                for (int k = 0; k < D; k++) { p2[i][j, k] = p2[i - 1][j, k]; }
            }
            var x = (int) (s[i - 1] - '0');
            for (int k = 0; k < D; k++)
            {
                p2[i][x, k] = (p2[i][x, k] + p1[i - 1][k]) % Mod;
            }
            p1[i][x]++;
        }
        for (int i = n; i >= 1; i--)
        {
            for (int j = 0; j < D; j++)
            {
                s1[i][j] = s1[i + 1][j];
                for (int k = 0; k < D; k++) { s2[i][j, k] = s2[i + 1][j, k]; }
            }
            var x = (int) (s[i - 1] - '0');
            for (int k = 0; k < D; k++)
            {
                s2[i][x, k] = (s2[i][x, k] + s1[i + 1][k]) % Mod;
            }
            s1[i][x]++;
        }
        long ans = 0;
        for (int i = 1; i <= n; i++)
        {
            for (int j = 0; j < D; j++)
            {
                for (int k = 0; k < D; k++)
                {
                    ans = (ans + p2[i - 1][j, k] * s2[i + 1][j, k]) % Mod;
                }
            }
        }
        return (int) ans;
    }
}
