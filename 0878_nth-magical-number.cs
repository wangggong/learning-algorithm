/*
 * @lc app=leetcode.cn id=nth-magical-number lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [878] 第 N 个神奇数字
 *
 * https://leetcode.cn/problems/nth-magical-number/description/
 *
 * algorithms
 * Hard (30.37%)
 * Total Accepted:    10.6K
 * Total Submissions: 29.8K
 * Testcase Example:  '1\n2\n3'
 *
 * 一个正整数如果能被 a 或 b 整除，那么它是神奇的。
 * 
 * 给定三个整数 n , a , b ，返回第 n 个神奇的数字。因为答案可能很大，所以返回答案 对 10^9 + 7 取模 后的值。
 * 
 * 
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：n = 1, a = 2, b = 3
 * 输出：2
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：n = 4, a = 2, b = 3
 * 输出：6
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 10^9
 * 2 <= a, b <= 4 * 10^4
 * 
 * 
 * 
 * 
 */
public class Solution
{
    private const long Mod = (long) 1e9 + 7;

    private long Gcd(long x, long y) => y == 0 ? x : Gcd(y, x % y);

    private long Lcm(long x, long y) => x * y / Gcd(x, y);

    public int NthMagicalNumber(int n, int a, int b)
    {
        var lcm = Lcm((long) a, (long) b);
        var k = lcm / a + lcm / b - 1;
        var ans = lcm % Mod * (n / k) % Mod;
        if ((long) n % k == 0) { return (int) ans; }
        long na = a;
        long nb = b;
        for (var r = (long) n % k - 1; r > 0; r--)
        {
            // assert na < lcm && nb < lcm && na != nb;
            if (na < nb) { na += a; }
            else { nb += b; }
        }
        ans = (ans + Math.Min(na, nb)) % Mod;
        return (int) ans;
    }
}
