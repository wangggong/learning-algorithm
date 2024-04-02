/*
 * @lc app=leetcode.cn id=maximum-number-of-non-overlapping-palindrome-substrings lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6236] 不重叠回文子字符串的最大数目
 *
 * https://leetcode.cn/problems/maximum-number-of-non-overlapping-palindrome-substrings/description/
 *
 * algorithms
 * Hard (29.95%)
 * Total Accepted:    2.1K
 * Total Submissions: 6.1K
 * Testcase Example:  '"abaccdbbd"\n3'
 *
 * 给你一个字符串 s 和一个 正 整数 k 。
 * 
 * 从字符串 s 中选出一组满足下述条件且 不重叠 的子字符串：
 * 
 * 
 * 每个子字符串的长度 至少 为 k 。
 * 每个子字符串是一个 回文串 。
 * 
 * 
 * 返回最优方案中能选择的子字符串的 最大 数目。
 * 
 * 子字符串 是字符串中一个连续的字符序列。
 * 
 * 
 * 
 * 示例 1 ：
 * 
 * 
 * 输入：s = "abaccdbbd", k = 3
 * 输出：2
 * 解释：可以选择 s = "abaccdbbd" 中斜体加粗的子字符串。"aba" 和 "dbbd" 都是回文，且长度至少为 k = 3 。
 * 可以证明，无法选出两个以上的有效子字符串。
 * 
 * 
 * 示例 2 ：
 * 
 * 
 * 输入：s = "adbcda", k = 2
 * 输出：0
 * 解释：字符串中不存在长度至少为 2 的回文子字符串。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= k <= s.length <= 2000
 * s 仅由小写英文字母组成
 * 
 * 
 */
public class Solution
{
    public int MaxPalindromes(string s, int k)
    {
        int n = s.Count();
        int[] dp = new int[n + 1];
        for (int i = 0; i <= n; i++)
        {
            if (i > 0) { dp[i] = Math.Max(dp[i], dp[i - 1]); }
            for (int l = i, r = i; l >= 0 && r < n && s[l] == s[r]; l--, r++)
            {
                if (r - l + 1 >= k) { dp[r + 1] = Math.Max(dp[r + 1], dp[l] + 1); }
            }
            for (int l = i, r = i + 1; l >= 0 && r < n && s[l] == s[r]; l--, r++)
            {
                if (r - l + 1 >= k) { dp[r + 1] = Math.Max(dp[r + 1], dp[l] + 1); }
            }
        }
        return dp[n];
    }
}
