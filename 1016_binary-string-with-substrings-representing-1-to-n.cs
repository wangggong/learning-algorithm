/*
 * @lc app=leetcode.cn id=binary-string-with-substrings-representing-1-to-n lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1016] 子串能表示从 1 到 N 数字的二进制串
 *
 * https://leetcode.cn/problems/binary-string-with-substrings-representing-1-to-n/description/
 *
 * algorithms
 * Medium (58.50%)
 * Total Accepted:    9.2K
 * Total Submissions: 15.3K
 * Testcase Example:  '"0110"\n3'
 *
 * 给定一个二进制字符串 s 和一个正整数 n，如果对于 [1, n] 范围内的每个整数，其二进制表示都是 s 的 子字符串 ，就返回 true，否则返回
 * false 。
 * 
 * 子字符串 是字符串中连续的字符序列。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s = "0110", n = 3
 * 输出：true
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：s = "0110", n = 4
 * 输出：false
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= s.length <= 1000
 * s[i] 不是 '0' 就是 '1'
 * 1 <= n <= 10^9
 * 
 * 
 */

/*
 * public class Solution
 * {
 *     public bool QueryString(string s, int n)
 *     {
 *         var sb = new StringBuilder();
 *         for (var i = s.Length - 1; i >= 0; i--)
 *         {
 *             sb.Append(s[i]);
 *         }
 *         s = sb.ToString();
 *         var S = new HashSet<string>();
 *         for (var (p, m) = (0, s.Length); p < m; p++)
 *         {
 *             for (var q = p + 1; q <= m; q++)
 *             {
 *                 S.Add(s.Substring(p, q - p));
 *             }
 *         }
 *         if (S.Count < n)
 *         {
 *             return false;
 *         }
 *         for (var i = 1; i <= n; i++)
 *         {
 *             sb = new StringBuilder();
 *             for (var k = i; k > 0; k >>= 1)
 *             {
 *                 sb.Append((char)((k & 1) + '0'));
 *             }
 *             var cur = sb.ToString();
 *             if (!S.Contains(cur))
 *             {
 *                 return false;
 *             }
 *         }
 *         return true;
 *     }
 * }
 */

// 参考: https://leetcode.cn/problems/binary-string-with-substrings-representing-1-to-n/solution/san-chong-suan-fa-cong-bao-li-dao-you-hu-nmtq/
public class Solution
{
    public bool QueryString(string s, int n)
    {
        var S = new HashSet<int>();
        for (var (i, m) = (0, s.Length); i < m; i++)
        {
            var cur = 0;
            for (var j = i; j < m; j++)
            {
                cur = (cur << 1) | (int)(s[j] - '0');
                if (cur > n)
                {
                    break;
                }
                S.Add(cur);
            }
        }
        return S.Count == n + (S.Contains(0) ? 1 : 0);
    }
}
