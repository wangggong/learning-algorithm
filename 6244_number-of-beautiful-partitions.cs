/*
 * @lc app=leetcode.cn id=number-of-beautiful-partitions lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6244] 完美分割的方案数
 *
 * https://leetcode.cn/problems/number-of-beautiful-partitions/description/
 *
 * algorithms
 * Hard (34.30%)
 * Total Accepted:    1.7K
 * Total Submissions: 4.9K
 * Testcase Example:  '"23542185131"\n3\n2'
 *
 * 给你一个字符串 s ，每个字符是数字 '1' 到 '9' ，再给你两个整数 k 和 minLength 。
 * 
 * 如果对 s 的分割满足以下条件，那么我们认为它是一个 完美 分割：
 * 
 * 
 * s 被分成 k 段互不相交的子字符串。
 * 每个子字符串长度都 至少 为 minLength 。
 * 每个子字符串的第一个字符都是一个 质数 数字，最后一个字符都是一个 非质数 数字。质数数字为 '2' ，'3' ，'5' 和 '7'
 * ，剩下的都是非质数数字。
 * 
 * 
 * 请你返回 s 的 完美 分割数目。由于答案可能很大，请返回答案对 10^9 + 7 取余 后的结果。
 * 
 * 一个 子字符串 是字符串中一段连续字符串序列。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s = "23542185131", k = 3, minLength = 2
 * 输出：3
 * 解释：存在 3 种完美分割方案：
 * "2354 | 218 | 5131"
 * "2354 | 21851 | 31"
 * "2354218 | 51 | 31"
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：s = "23542185131", k = 3, minLength = 3
 * 输出：1
 * 解释：存在一种完美分割方案："2354 | 218 | 5131" 。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：s = "3312958", k = 3, minLength = 1
 * 输出：1
 * 解释：存在一种完美分割方案："331 | 29 | 58" 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= k, minLength <= s.length <= 1000
 * s 每个字符都为数字 '1' 到 '9' 之一。
 * 
 * 
 */
// 参考: https://leetcode.cn/problems/number-of-beautiful-partitions/solutions/1981346/dong-tai-gui-hua-jian-ji-xie-fa-xun-huan-xyw3/
// 小优化太多了, 想不到...
public class Solution
{
    private const long Mod = (long) 1e9 + 7;

    public int BeautifulPartitions(string s, int k, int l)
    {
        var n = s.Length;
        var prime = new HashSet<char>{ '2', '3', '5', '7' };
        if (k * l > n || !prime.Contains(s[0]) || prime.Contains(s[n - 1])) { return 0; }
        bool canPartition(int k) => k == 0 || k == n || !prime.Contains(s[k - 1]) && prime.Contains(s[k]);
        var dp = new long[k + 1, n + 1];
        dp[0, 0] = 1;
        for (int i = 1; i <= k; i++)
        {
            long tot = 0;
            for (int j = l * i; j + l * (k - i) <= n; j++)
            {
                if (canPartition(j - l)) { tot = (tot + dp[i - 1, j - l]) % Mod; }
                if (canPartition(j)) { dp[i, j] = tot; }
            }
        }
        return (int) dp[k, n];
    }
}
