/*
 * @lc app=leetcode.cn id=string-transformation lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [8020] 字符串转换
 *
 * https://leetcode.cn/problems/string-transformation/description/
 *
 * algorithms
 * Hard (30.36%)
 * Total Accepted:    487
 * Total Submissions: 1.6K
 * Testcase Example:  '"abcd"\n"cdab"\n2'
 *
 * 给你两个长度都为 n 的字符串 s 和 t 。你可以对字符串 s 执行以下操作：
 * 
 * 
 * 将 s 长度为 l （0 < l < n）的 后缀字符串 删除，并将它添加在 s 的开头。
 * 比方说，s = 'abcd' ，那么一次操作中，你可以删除后缀 'cd' ，并将它添加到 s 的开头，得到 s = 'cdab' 。
 * 
 * 
 * 给你一个整数 k ，请你返回 恰好 k 次操作将 s 变为 t 的方案数。
 * 
 * 由于答案可能很大，返回答案对 10^9 + 7 取余 后的结果。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s = "abcd", t = "cdab", k = 2
 * 输出：2
 * 解释：
 * 第一种方案：
 * 第一次操作，选择 index = 3 开始的后缀，得到 s = "dabc" 。
 * 第二次操作，选择 index = 3 开始的后缀，得到 s = "cdab" 。
 * 
 * 第二种方案：
 * 第一次操作，选择 index = 1 开始的后缀，得到 s = "bcda" 。
 * 第二次操作，选择 index = 1 开始的后缀，得到 s = "cdab" 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：s = "ababab", t = "ababab", k = 1
 * 输出：2
 * 解释：
 * 第一种方案：
 * 选择 index = 2 开始的后缀，得到 s = "ababab" 。
 * 
 * 第二种方案：
 * 选择 index = 4 开始的后缀，得到 s = "ababab" 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 2 <= s.length <= 5 * 10^5
 * 1 <= k <= 10^15
 * s.length == t.length
 * s 和 t 都只包含小写英文字母。
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/string-transformation/solutions/2435348/kmp-ju-zhen-kuai-su-mi-you-hua-dp-by-end-vypf/
//
// 超大融合怪. 状态机 DP + 快速幂优化.
public class Solution
{
    public int NumberOfWays(string s, string t, long k)
    {
        var n = s.Length;
        var c = Kmp(s + s[.. ^1], t);
        return (int)Pow(new long[][]
        {
            new long[] { c - 1, c, },
            new long[] { n - c, n - c - 1, },
        }, k)[0][s == t ? 0 : 1];
    }

    private int Kmp(string haystack, string needle)
    {
        var (n, m) = (haystack.Length, needle.Length);
        var fails = new int[m];
        Array.Fill(fails, -1);
        for (var i = 1; i < m; i++)
        {
            var j = fails[i - 1];
            for (; j != -1 && needle[j + 1] != needle[i]; j = fails[j]) { }
            if (needle[j + 1] == needle[i]) { fails[i] = j + 1; }
        }
        var (ans, match) = (0, -1);
        for (var i = 0; i < n; i++)
        {
            for (; match != -1 && needle[match + 1] != haystack[i]; match = fails[match]) { }
            if (needle[match + 1] == haystack[i])
            {
                match++;
                if (match == m - 1)
                {
                    ans++;
                    match = fails[match];
                }
            }
        }
        return ans;
    }

    private long[][] Pow(long[][] M, long k)
    {
        var n = M.Length;
        var ans = new long[n][];
        for (var i = 0; i < n; i++)
        {
            ans[i] = new long[n];
            ans[i][i] = 1;
        }
        for (; k > 0; k >>= 1)
        {
            if ((k & 1) is not 0) { ans = Mul(ans, M); }
            M = Mul(M, M);
        }
        return ans;
    }

    private long[][] Mul(long[][] A, long[][] B)
    {
        const long Mod = (long)1e9 + 7;
        var n = A.Length;
        var ans = new long[n][];
        for (var i = 0; i < n; i++) { ans[i] = new long[n]; }
        for (var i = 0; i < n; i++)
        {
            for (var j = 0; j < n; j++)
            {
                for (var k = 0; k < n; k++)
                {
                    ans[i][j] = (ans[i][j] + A[i][k] * B[k][j]) % Mod;
                }
            }
        }
        return ans;
    }
}
