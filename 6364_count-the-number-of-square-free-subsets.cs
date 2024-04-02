/*
 * @lc app=leetcode.cn id=count-the-number-of-square-free-subsets lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6364] 无平方子集计数
 *
 * https://leetcode.cn/problems/count-the-number-of-square-free-subsets/description/
 *
 * algorithms
 * Medium (10.83%)
 * Total Accepted:    919
 * Total Submissions: 5.8K
 * Testcase Example:  '[3,4,4,5]'
 *
 * 给你一个正整数数组 nums 。
 * 
 * 如果数组 nums 的子集中的元素乘积是一个 无平方因子数 ，则认为该子集是一个 无平方 子集。
 * 
 * 无平方因子数 是无法被除 1 之外任何平方数整除的数字。
 * 
 * 返回数组 nums 中 无平方 且 非空 的子集数目。因为答案可能很大，返回对 10^9 + 7 取余的结果。
 * 
 * nums 的 非空子集 是可以由删除 nums
 * 中一些元素（可以不删除，但不能全部删除）得到的一个数组。如果构成两个子集时选择删除的下标不同，则认为这两个子集不同。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [3,4,4,5]
 * 输出：3
 * 解释：示例中有 3 个无平方子集：
 * - 由第 0 个元素 [3] 组成的子集。其元素的乘积是 3 ，这是一个无平方因子数。
 * - 由第 3 个元素 [5] 组成的子集。其元素的乘积是 5 ，这是一个无平方因子数。
 * - 由第 0 个和第 3 个元素 [3,5] 组成的子集。其元素的乘积是 15 ，这是一个无平方因子数。
 * 可以证明给定数组中不存在超过 3 个无平方子集。
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [1]
 * 输出：1
 * 解释：示例中有 1 个无平方子集：
 * - 由第 0 个元素 [1] 组成的子集。其元素的乘积是 1 ，这是一个无平方因子数。
 * 可以证明给定数组中不存在超过 1 个无平方子集。
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 1000
 * 1 <= nums[i] <= 30
 * 
 * 
 */
public class Solution
{
    private const long Mod = (long)1e9 + 7;
    private const int N = 30;
    private const int M = 10;
    private const int S = 1 << M;
    private int[] Primes = new int[M] { 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, };
    private int[] Masks = new int[N + 1];

    public Solution()
    {
        for (var n = 2; n <= N; n++)
        {
            for (var j = 0; j < M; j++)
            {
                if (n % Primes[j] == 0)
                {
                    if (n % (Primes[j] * Primes[j]) == 0)
                    {
                        Masks[n] = -1;
                        break;
                    }
                    Masks[n] |= 1 << j;
                }
            }
        }
    }

    /*
     * // 背包方案
     * public int SquareFreeSubsets(int[] nums)
     * {
     *     var dp = new long[S];
     *     dp[0] = 1;
     *     foreach (var n in nums)
     *     {
     *         var mask = Masks[n];
     *         if (mask >= 0)
     *         {
     *             for (var s = S - 1; s >= mask; s--)
     *             {
     *                 if ((s | mask) == s)
     *                 {
     *                     dp[s] = (dp[s] + dp[s ^ mask]) % Mod;
     *                 }
     *             }
     *         }
     *     }
     *     return (int)((dp.Sum() - 1 + Mod) % Mod);
     * }
     */

    // 状态压缩 DP 方案
        public int SquareFreeSubsets(int[] nums)
    {
        long c1 = 1;
        var dp = new long[S];
        dp[0] = 1;
        foreach (var (k, c) in nums
            .GroupBy(x => x)
            .ToDictionary(g => g.Key, g => g.Count()))
        {
            if (k == 1)
            {
                for (var i = 0; i < c; i++)
                {
                    c1 = (c1 << 1) % Mod;
                }
                continue;
            }
            var mask = Masks[k];
            if (mask > 0)
            {
                for (var s = (S - 1) ^ mask; s >= 0; s--)
                {
                    if ((s & mask) == 0)
                    {
                        dp[s | mask] = (dp[s | mask] + dp[s] * c) % Mod;
                    }
                }
            }
        }
        return (int)((dp.Sum() % Mod * c1 + Mod - 1) % Mod);
    }
}
