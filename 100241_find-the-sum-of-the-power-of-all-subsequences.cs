/*
 * @lc app=leetcode.cn id=find-the-sum-of-the-power-of-all-subsequences lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [100241] 求出所有子序列的能量和
 *
 * https://leetcode.cn/problems/find-the-sum-of-the-power-of-all-subsequences/description/
 *
 * algorithms
 * Hard (43.52%)
 * Total Accepted:    1K
 * Total Submissions: 2.3K
 * Testcase Example:  '[1,2,3]\n3'
 *
 * 给你一个长度为 n 的整数数组 nums 和一个 正 整数 k 。
 * 
 * 一个整数数组的 能量 定义为和 等于 k 的子序列的数目。
 * 
 * 请你返回 nums 中所有子序列的 能量和 。
 * 
 * 由于答案可能很大，请你将它对 10^9 + 7 取余 后返回。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：  nums = [1,2,3], k = 3 
 * 
 * 输出：  6 
 * 
 * 解释：
 * 
 * 总共有 5 个能量不为 0 的子序列：
 * 
 * 
 * 子序列 [1,2,3] 有 2 个和为 3 的子序列：[1,2,3] 和 [1,2,3] 。
 * 子序列 [1,2,3] 有 1 个和为 3 的子序列：[1,2,3] 。
 * 子序列 [1,2,3] 有 1 个和为 3 的子序列：[1,2,3] 。
 * 子序列 [1,2,3] 有 1 个和为 3 的子序列：[1,2,3] 。
 * 子序列 [1,2,3] 有 1 个和为 3 的子序列：[1,2,3] 。
 * 
 * 
 * 所以答案为 2 + 1 + 1 + 1 + 1 = 6 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：  nums = [2,3,3], k = 5 
 * 
 * 输出：  4 
 * 
 * 解释：
 * 
 * 总共有 3 个能量不为 0 的子序列：
 * 
 * 
 * 子序列 [2,3,3] 有 2 个子序列和为 5 ：[2,3,3] 和 [2,3,3] 。
 * 子序列 [2,3,3] 有 1 个子序列和为 5 ：[2,3,3] 。
 * 子序列 [2,3,3] 有 1 个子序列和为 5 ：[2,3,3] 。
 * 
 * 
 * 所以答案为 2 + 1 + 1 = 4 。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：  nums = [1,2,3], k = 7 
 * 
 * 输出：  0 
 * 
 * 解释：不存在和为 7 的子序列，所以 nums 的能量和为 0 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 100
 * 1 <= nums[i] <= 10^4
 * 1 <= k <= 100
 * 
 * 
 */

/*
 * public class Solution
 * {
 *     private const long Mod = (long)1e9 + 7;
 *     private const int N = 100;
 *     private static long[] powers = new long[N + 1];
 * 
 *     static Solution()
 *     {
 *         powers[0] = 1;
 *         for (var i = 1; i <= N; i++)
 *         {
 *             powers[i] = powers[i - 1] * 2 % Mod;
 *         }
 *     }
 * 
 *     public int SumOfPower(int[] nums, int k)
 *     {
 *         var n = nums.Length;
 *         var dp = new long[n + 1][];
 *         for (var i = 0; i <= n; i++)
 *         {
 *             dp[i] = new long[k + 1];
 *         }
 *         dp[0][0] = 1;
 *         foreach (var num in nums)
 *         {
 *             for (var i = n; i > 0; i--)
 *             {
 *                 for (var j = num; j <= k; j++)
 *                 {
 *                     dp[i][j] = (dp[i][j] + dp[i - 1][j - num]) % Mod;
 *                 }
 *             }
 *         }
 *         var ans = 0l;
 *         for (var i = 0; i <= n; i++)
 *         {
 *             ans = (ans + dp[i][k] * powers[n - i]) % Mod;
 *         }
 *         return (int)ans;
 *     }
 * }
 */

// 参考: https://leetcode.cn/problems/find-the-sum-of-the-power-of-all-subsequences/solutions/2691661/liang-chong-fang-fa-er-wei-yi-wei-0-1-be-2e47/
public class Solution
{
    public int SumOfPower(int[] nums, int k)
    {
        const long Mod = (long)1e9 + 7;
        var dp = new long[k + 1];
        dp[0] = 1;
        foreach (var n in nums)
        {
            for (var i = k; i >= 0; i--)
            {
                dp[i] = (dp[i] * 2 + (i >= n ? dp[i - n] : 0)) % Mod;
            }
        }
        return (int)dp.Last();
    }
}
