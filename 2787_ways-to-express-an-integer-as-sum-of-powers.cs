/*
 * @lc app=leetcode.cn id=ways-to-express-an-integer-as-sum-of-powers lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [2787] 将一个数字表示成幂的和的方案数
 *
 * https://leetcode.cn/problems/ways-to-express-an-integer-as-sum-of-powers/description/
 *
 * algorithms
 * Medium (40.94%)
 * Total Accepted:    1.9K
 * Total Submissions: 4.7K
 * Testcase Example:  '10\n2'
 *
 * 给你两个 正 整数 n 和 x 。
 * 
 * 请你返回将 n 表示成一些 互不相同 正整数的 x 次幂之和的方案数。换句话说，你需要返回互不相同整数 [n1, n2, ..., nk]
 * 的集合数目，满足 n = n1^x + n2^x + ... + nk^x 。
 * 
 * 由于答案可能非常大，请你将它对 10^9 + 7 取余后返回。
 * 
 * 比方说，n = 160 且 x = 3 ，一个表示 n 的方法是 n = 2^3 + 3^3 + 5^3^ 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：n = 10, x = 2
 * 输出：1
 * 解释：我们可以将 n 表示为：n = 3^2 + 1^2 = 10 。
 * 这是唯一将 10 表达成不同整数 2 次方之和的方案。
 * 
 * 
 * 示例 2：
 * 
 * 输入：n = 4, x = 1
 * 输出：2
 * 解释：我们可以将 n 按以下方案表示：
 * - n = 4^1 = 4 。
 * - n = 3^1 + 1^1 = 4 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 300
 * 1 <= x <= 5
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/ways-to-express-an-integer-as-sum-of-powers/solution/0-1-bei-bao-mo-ban-ti-by-endlesscheng-ap09/
//
// 0-1 背包模板题, 这都不会写了?!
public class Solution
{
    private const int N = 300;
    private const int K = 5;
    private const long Mod = (long)1e9 + 7;
    private static long[][] dp;

    private static int Pow(int n, int k)
    {
        var ans = 1;
        for (; k > 0; k--)
        {
            ans *= n;
        }
        return ans;
    }

    static Solution()
    {
        dp = Enumerable.Range(0, K + 1)
            .Select(_ => new long[N + 1])
            .ToArray();
        for (var k = 1; k <= K; k++)
        {
            dp[k][0] = 1;
            for (var i = 1; true; i++)
            {
                var v = Pow(i, k);
                if (v > N)
                {
                    break;
                }
                for (var s = N; s >= v; s--)
                {
                    dp[k][s] = (dp[k][s] + dp[k][s - v]) % Mod;
                }
            }
        }
    }

    public int NumberOfWays(int n, int x) => (int)dp[x][n];
}
