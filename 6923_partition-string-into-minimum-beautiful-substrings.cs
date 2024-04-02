/*
 * @lc app=leetcode.cn id=partition-string-into-minimum-beautiful-substrings lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6923] 将字符串分割为最少的美丽子字符串
 *
 * https://leetcode.cn/problems/partition-string-into-minimum-beautiful-substrings/description/
 *
 * algorithms
 * Medium (49.47%)
 * Total Accepted:    1.4K
 * Total Submissions: 2.8K
 * Testcase Example:  '"1011"'
 *
 * 给你一个二进制字符串 s ，你需要将字符串分割成一个或者多个 子字符串  ，使每个子字符串都是 美丽 的。
 * 
 * 如果一个字符串满足以下条件，我们称它是 美丽 的：
 * 
 * 
 * 它不包含前导 0 。
 * 它是 5 的幂的 二进制 表示。
 * 
 * 
 * 请你返回分割后的子字符串的 最少 数目。如果无法将字符串 s 分割成美丽子字符串，请你返回 -1 。
 * 
 * 子字符串是一个字符串中一段连续的字符序列。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：s = "1011"
 * 输出：2
 * 解释：我们可以将输入字符串分成 ["101", "1"] 。
 * - 字符串 "101" 不包含前导 0 ，且它是整数 5^1 = 5 的二进制表示。
 * - 字符串 "1" 不包含前导 0 ，且它是整数 5^0 = 1 的二进制表示。
 * 最少可以将 s 分成 2 个美丽子字符串。
 * 
 * 
 * 示例 2：
 * 
 * 输入：s = "111"
 * 输出：3
 * 解释：我们可以将输入字符串分成 ["1", "1", "1"] 。
 * - 字符串 "1" 不包含前导 0 ，且它是整数 5^0 = 1 的二进制表示。
 * 最少可以将 s 分成 3 个美丽子字符串。
 * 
 * 
 * 示例 3：
 * 
 * 输入：s = "0"
 * 输出：-1
 * 解释：无法将给定字符串分成任何美丽子字符串。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= s.length <= 15
 * s[i] 要么是 '0' 要么是 '1' 。
 * 
 * 
 */
public class Solution
{
    private static List<string> primes;
    private const int K = 5;
    private const int N = 15;

    static Solution()
    {
        primes = new();
        for (var i = 1; true; i *= K)
        {
            var sb = new StringBuilder();
            for (var k = i; k is not 0; k >>= 1)
            {
                sb.Append((char)('0' + (k % 2)));
            }
            var s = new string(sb.ToString().Reverse().ToArray());
            if (s.Length > N) { break; }
            primes.Add(s);
        }
    }

    public int MinimumBeautifulSubstrings(string s)
    {
        var n = s.Length;
        var dp = new int[n + 1];
        Array.Fill(dp, n + 1);
        dp[n] = 0;
        for (var i = n - 1; i >= 0; i--)
        {
            foreach (var p in primes)
            {
                var m = p.Length;
                if (i + m <= n && s.Substring(i, m) == p)
                {
                    dp[i] = Math.Min(dp[i], 1 + dp[i + m]);
                }
            }
        }
        return dp[0] > n ? -1 : dp[0];
    }
}
