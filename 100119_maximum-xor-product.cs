/*
 * @lc app=leetcode.cn id=maximum-xor-product lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [100119] 最大异或乘积
 *
 * https://leetcode.cn/problems/maximum-xor-product/description/
 *
 * algorithms
 * Medium (19.54%)
 * Total Accepted:    1.2K
 * Total Submissions: 6.2K
 * Testcase Example:  '12\n5\n4'
 *
 * 给你三个整数 a ，b 和 n ，请你返回 (a XOR x) * (b XOR x) 的 最大值 且 x 需要满足 0 <= x < 2^n。
 * 
 * 由于答案可能会很大，返回它对 10^9 + 7 取余 后的结果。
 * 
 * 注意，XOR 是按位异或操作。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：a = 12, b = 5, n = 4
 * 输出：98
 * 解释：当 x = 2 时，(a XOR x) = 14 且 (b XOR x) = 7 。所以，(a XOR x) * (b XOR x) = 98 。
 * 98 是所有满足 0 <= x < 2^n 中 (a XOR x) * (b XOR x) 的最大值。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：a = 6, b = 7 , n = 5
 * 输出：930
 * 解释：当 x = 25 时，(a XOR x) = 31 且 (b XOR x) = 30 。所以，(a XOR x) * (b XOR x) =
 * 930 。
 * 930 是所有满足 0 <= x < 2^n 中 (a XOR x) * (b XOR x) 的最大值。
 * 
 * 示例 3：
 * 
 * 
 * 输入：a = 1, b = 6, n = 3
 * 输出：12
 * 解释： 当 x = 5 时，(a XOR x) = 4 且 (b XOR x) = 3 。所以，(a XOR x) * (b XOR x) = 12 。
 * 12 是所有满足 0 <= x < 2^n 中 (a XOR x) * (b XOR x) 的最大值。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 0 <= a, b < 2^50
 * 0 <= n <= 50
 * 
 * 
 */
public class Solution
{
    public int MaximumXorProduct(long a, long b, int n)
    {
        const long Mod = (long)1e9 + 7;
        for (var d = 0; d < n; d++)
        {
            var (p, q, mask) = (a >> d & 1l, b >> d & 1l, 1l << d);
            if ((p, q) is (0, 0))
            {
                a |= mask;
                b |= mask;
                continue;
            }
            if ((p, q) is (0, 1) or (1, 0))
            {
                var (s, t) = q is 0 ? (a, b) : (b, a);
                if (s - t - mask > 0)
                {
                    a ^= mask;
                    b ^= mask;
                }
            }
        }
        return (int)((a % Mod + Mod) % Mod * (b % Mod + Mod) % Mod);
    }
}
