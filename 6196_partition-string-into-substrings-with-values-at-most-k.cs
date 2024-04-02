/*
 * @lc app=leetcode.cn id=partition-string-into-substrings-with-values-at-most-k lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6196] 将字符串分割成值不超过 K 的子字符串
 *
 * https://leetcode.cn/problems/partition-string-into-substrings-with-values-at-most-k/description/
 *
 * algorithms
 * Medium (46.62%)
 * Total Accepted:    2.5K
 * Total Submissions: 5.4K
 * Testcase Example:  '"165462"\n60'
 *
 * 给你一个字符串 s ，它每一位都是 1 到 9 之间的数字组成，同时给你一个整数 k 。
 * 
 * 如果一个字符串 s 的分割满足以下条件，我们称它是一个 好 分割：
 * 
 * 
 * s 中每个数位 恰好 属于一个子字符串。
 * 每个子字符串的值都小于等于 k 。
 * 
 * 
 * 请你返回 s 所有的 好 分割中，子字符串的 最少 数目。如果不存在 s 的 好 分割，返回 -1 。
 * 
 * 注意：
 * 
 * 
 * 一个字符串的 值 是这个字符串对应的整数。比方说，"123" 的值为 123 ，"1" 的值是 1 。
 * 子字符串 是字符串中一段连续的字符序列。
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s = "165462", k = 60
 * 输出：4
 * 解释：我们将字符串分割成子字符串 "16" ，"54" ，"6" 和 "2" 。每个子字符串的值都小于等于 k = 60 。
 * 不存在小于 4 个子字符串的好分割。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：s = "238182", k = 5
 * 输出：-1
 * 解释：这个字符串不存在好分割。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= s.length <= 10^5
 * s[i] 是 '1' 到 '9' 之间的数字。
 * 1 <= k <= 10^9
 * 
 * 
 */
public class Solution
{
    public int MinimumPartition(string s, int k)
    {
        var n = s.Length;
        var dp = new int[n + 1];
        for (var i = 1; i <= n; i++)
        {
            dp[i] = n + 1;
        }
        for (var i = 1; i <= n; i++)
        {
            for (long j = 1, cur = 0, digit = 1; (int) j <= i; j++, digit *= 10)
            {
                cur = cur + (s[i - (int) j] - '0') * digit;
                if (cur > (long) k)
                {
                    break;
                }
                dp[i] = Math.Min(dp[i], dp[i - (int) j] + 1);
            }
        }
        return dp[n] > n ? -1 : dp[n];
    }
}
