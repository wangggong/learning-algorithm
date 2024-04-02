/*
 * @lc app=leetcode.cn id=smallest-value-after-replacing-with-sum-of-prime-factors lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6266] 使用质因数之和替换后可以取到的最小值
 *
 * https://leetcode.cn/problems/smallest-value-after-replacing-with-sum-of-prime-factors/description/
 *
 * algorithms
 * Medium (43.13%)
 * Total Accepted:    2.8K
 * Total Submissions: 6.5K
 * Testcase Example:  '15'
 *
 * 给你一个正整数 n 。
 * 
 * 请你将 n 的值替换为 n 的 质因数 之和，重复这一过程。
 * 
 * 
 * 注意，如果 n 能够被某个质因数多次整除，则在求和时，应当包含这个质因数同样次数。
 * 
 * 
 * 返回 n 可以取到的最小值。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：n = 15
 * 输出：5
 * 解释：最开始，n = 15 。
 * 15 = 3 * 5 ，所以 n 替换为 3 + 5 = 8 。
 * 8 = 2 * 2 * 2 ，所以 n 替换为 2 + 2 + 2 = 6 。
 * 6 = 2 * 3 ，所以 n 替换为 2 + 3 = 5 。
 * 5 是 n 可以取到的最小值。
 * 
 * 
 * 示例 2：
 * 
 * 输入：n = 3
 * 输出：3
 * 解释：最开始，n = 3 。
 * 3 是 n 可以取到的最小值。
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 2 <= n <= 10^5
 * 
 * 
 */

/*
 * public class Solution
 * {
 *     private const int N = 100000;
 * 
 *     public int SmallestValue(int n)
 *     {
 *         var primes = new bool[N + 1];
 *         for (var i = 2; i <= N; i++)
 *         {
 *             primes[i] = true;
 *         }
 *         for (var i = 2; i <= N; i++)
 *         {
 *             if (primes[i])
 *             {
 *                 for (var j = 2; i * j <= N; j++)
 *                 {
 *                     primes[i * j] = false;
 *                 }
 *             }
 *         }
 *         int get(int n)
 *         {
 *             var ans = 0;
 *             for (var i = 2; i <= N && n > 1; i++)
 *             {
 *                 if (primes[i])
 *                 {
 *                     for (; n % i == 0; n /= i)
 *                     {
 *                         ans += i;
 *                     }
 *                 }
 *             }
 *             return ans;
 *         }
 *         var ans = n;
 *         for (var nx = get(n); nx != n; nx = get(n))
 *         {
 *             ans = Math.Min(ans, nx);
 *             n = nx;
 *         }
 *         return ans;
 *     }
 * }
 */

public class Solution
{
    public int SmallestValue(int n)
    {
        int get(int n)
        {
            var ans = 0;
            for (var i = 2; i * i <= n; i++)
            {
                while (n % i == 0)
                {
                    ans += i;
                    n /= i;
                }
            }
            if (n > 1)
            {
                ans += n;
            }
            return ans;
        }
        while (true)
        {
            var ne = get(n);
            if (n == ne)
            {
                return n;
            }
            n = ne;
        }
    }
}
