/*
 * @lc app=leetcode.cn id=minimum-deletions-to-make-string-balanced lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1653] 使字符串平衡的最少删除次数
 *
 * https://leetcode.cn/problems/minimum-deletions-to-make-string-balanced/description/
 *
 * algorithms
 * Medium (55.60%)
 * Total Accepted:    6.9K
 * Total Submissions: 12.4K
 * Testcase Example:  '"aababbab"'
 *
 * 给你一个字符串 s ，它仅包含字符 'a' 和 'b'​​​​ 。
 * 
 * 你可以删除 s 中任意数目的字符，使得 s 平衡 。当不存在下标对 (i,j) 满足 i < j ，且 s[i] = 'b' 的同时 s[j]= 'a'
 * ，此时认为 s 是 平衡 的。
 * 
 * 请你返回使 s 平衡 的 最少 删除次数。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s = "aababbab"
 * 输出：2
 * 解释：你可以选择以下任意一种方案：
 * 下标从 0 开始，删除第 2 和第 6 个字符（"aababbab" -> "aaabbb"），
 * 下标从 0 开始，删除第 3 和第 6 个字符（"aababbab" -> "aabbbb"）。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：s = "bbaaaaabb"
 * 输出：2
 * 解释：唯一的最优解是删除最前面两个字符。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= s.length <= 10^5
 * s[i] 要么是 'a' 要么是 'b'​ 。​
 * 
 * 
 */
public class Solution
{
    public int MinimumDeletions(string s)
    {
        var n = s.Length;
        var S = new int[n + 1][];
        for (var i = 0; i <= n; i++)
        {
            S[i] = new int[2];
        }
        for (var i = 0; i < n; i++)
        {
            (S[i + 1][0], S[i + 1][1]) = (S[i][0], S[i][1]);
            S[i + 1][s[i] - 'a']++;
        }
        return Enumerable
            .Range(0, n + 1)
            .Select(i => (S[i][1] - S[0][1]) + (S[n][0] - S[i][0]))
            .Min();
    }
}
