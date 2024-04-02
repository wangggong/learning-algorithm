/*
 * @lc app=leetcode.cn id=apply-operations-to-make-two-strings-equal lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [8028] 执行操作使两个字符串相等
 *
 * https://leetcode.cn/problems/apply-operations-to-make-two-strings-equal/description/
 *
 * algorithms
 * Medium (16.65%)
 * Total Accepted:    939
 * Total Submissions: 5.6K
 * Testcase Example:  '"1100011000"\n"0101001010"\n2'
 *
 * 给你两个下标从 0 开始的二进制字符串 s1 和 s2 ，两个字符串的长度都是 n ，再给你一个正整数 x 。
 * 
 * 你可以对字符串 s1 执行以下操作 任意次 ：
 * 
 * 
 * 选择两个下标 i 和 j ，将 s1[i] 和 s1[j] 都反转，操作的代价为 x 。
 * 选择满足 i < n - 1 的下标 i ，反转 s1[i] 和 s1[i + 1] ，操作的代价为 1 。
 * 
 * 
 * 请你返回使字符串 s1 和 s2 相等的 最小 操作代价之和，如果无法让二者相等，返回 -1 。
 * 
 * 注意 ，反转字符的意思是将 0 变成 1 ，或者 1 变成 0 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s1 = "1100011000", s2 = "0101001010", x = 2
 * 输出：4
 * 解释：我们可以执行以下操作：
 * - 选择 i = 3 执行第二个操作。结果字符串是 s1 = "1101111000" 。
 * - 选择 i = 4 执行第二个操作。结果字符串是 s1 = "1101001000" 。
 * - 选择 i = 0 和 j = 8 ，执行第一个操作。结果字符串是 s1 = "0101001010" = s2 。
 * 总代价是 1 + 1 + 2 = 4 。这是最小代价和。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：s1 = "10110", s2 = "00011", x = 4
 * 输出：-1
 * 解释：无法使两个字符串相等。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * n == s1.length == s2.length
 * 1 <= n, x <= 500
 * s1 和 s2 只包含字符 '0' 和 '1' 。
 * 
 * 
 */

// 参考:
// https://leetcode.cn/problems/apply-operations-to-make-two-strings-equal/solutions/2472122/ji-yi-hua-sou-suo-by-endlesscheng-64vq/
// https://leetcode.cn/problems/apply-operations-to-make-two-strings-equal/solutions/2472137/xiao-yang-xiao-en-o1-kong-jian-yu-on-shi-wg6j/
public class Solution
{
    
    /*
     * // `O(n^2)` DP
     * public int MinOperations(string s1, string s2, int x)
     * {
     *     const int Maxn = 0x3f3f3f3f;
     *     if (s1.Count(x => x is '0') % 2 != s2.Count(x => x is '0') % 2) { return -1; }
     *     var n = s1.Length;
     *     var memo = new Dictionary<(int, int, bool), int>();
     *     int dfs(int i, int j, bool pre)
     *     {
     *         var key = (i, j, pre);
     *         if (memo.ContainsKey(key)) { return memo[key]; }
     *         if (i == n) { return memo[key] = j > 0 || pre ? Maxn : 0; }
     *         if ((s1[i] == s2[i]) != pre) { return memo[key] = dfs(i + 1, j, false); }
     *         var ans = Math.Min( dfs(i + 1, j, true) + 1, dfs(i + 1, j + 1, false) + x);
     *         if (j > 0) { ans = Math.Min(ans, dfs(i + 1, j - 1, false)); }
     *         return memo[key] = ans;
     *     }
     *     return dfs(0, 0, false);
     * }
     */

    /*
     * // `O(n)` DP
     * public int MinOperations(string s1, string s2, int x)
     * {
     *     const int Maxn = 0x3f3f3f3f;
     *     var arr = s1.Zip(s2, (s, t) => (s, t))
     *         .Select((x, i) => (s: x.s, t: x.t, i: i))
     *         .Where(x => x.s != x.t)
     *         .Select(x => x.i)
     *         .ToArray();
     *     var n = arr.Length;
     *     if ((n % 2) is not 0) { return -1; }
     *     var dp = new int[n + 1];
     *     Array.Fill(dp, Maxn);
     *     dp[0] = 0;
     *     for (int i = 0; i < n; i++)
     *     {
     *         dp[i + 1] = dp[i] + ((i % 2) is 0 ? 0 : x);
     *         if (i > 0) { dp[i + 1] = Math.Min(dp[i + 1], dp[i - 1] + arr[i] - arr[i - 1]); }
     *     }
     *     return dp.Last();
     * }
     */

    // `O(n)` DP
    // 把 "两次操作总共花费 `x`" 转化为 "一次操作花费 `x/2`"
    public int MinOperations(string s1, string s2, int x)
    {
        const int Maxn = 0x3f3f3f3f;
        var arr = s1.Zip(s2, (s, t) => (s, t))
            .Select((x, i) => (s: x.s, t: x.t, i: i))
            .Where(x => x.s != x.t)
            .Select(x => x.i)
            .ToArray();
        var n = arr.Length;
        if ((n % 2) is not 0) { return -1; }
        var dp = new int[n + 2];
        dp[0] = Maxn;
        for (var i = 0; i < n; i++)
        {
            dp[i + 2] = dp[i + 1] + x;
            if (i > 0)
            {
                dp[i + 2] = Math.Min(
                    dp[i + 2],
                    dp[i] + 2 * (arr[i] - arr[i - 1]));
            }
        }
        return dp.Last() / 2;
    }
}
