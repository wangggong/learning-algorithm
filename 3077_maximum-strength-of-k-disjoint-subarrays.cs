/*
 * @lc app=leetcode.cn id=maximum-strength-of-k-disjoint-subarrays lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [3077] K 个不相交子数组的最大能量值
 *
 * https://leetcode.cn/problems/maximum-strength-of-k-disjoint-subarrays/description/
 *
 * algorithms
 * Hard (37.04%)
 * Total Accepted:    1.3K
 * Total Submissions: 3.5K
 * Testcase Example:  '[1,2,3,-1,2]\n3'
 *
 * 给你一个长度为 n 下标从 0 开始的整数数组 nums 和一个 正奇数 整数 k 。
 * 
 * x 个子数组的能量值定义为 strength = sum[1] * x - sum[2] * (x - 1) + sum[3] * (x - 2) -
 * sum[4] * (x - 3) + ... + sum[x] * 1 ，其中 sum[i] 是第 i 个子数组的和。更正式的，能量值是满足 1 <=
 * i <= x 的所有 i 对应的 (-1)^i+1 * sum[i] * (x - i + 1) 之和。
 * 
 * 你需要在 nums 中选择 k 个 不相交子数组 ，使得 能量值最大 。
 * 
 * 请你返回可以得到的 最大能量值 。
 * 
 * 注意，选出来的所有子数组 不 需要覆盖整个数组。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [1,2,3,-1,2], k = 3
 * 输出：22
 * 解释：选择 3 个子数组的最好方式是选择：nums[0..2] ，nums[3..3] 和 nums[4..4] 。能量值为 (1 + 2 + 3) *
 * 3 - (-1) * 2 + 2 * 1 = 22 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [12,-2,-2,-2,-2], k = 5
 * 输出：64
 * 解释：唯一一种选 5 个不相交子数组的方案是：nums[0..0] ，nums[1..1] ，nums[2..2] ，nums[3..3] 和
 * nums[4..4] 。能量值为 12 * 5 - (-2) * 4 + (-2) * 3 - (-2) * 2 + (-2) * 1 = 64 。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [-1,-2,-3], k = 1
 * 输出：-1
 * 解释：选择 1 个子数组的最优方案是：nums[0..0] 。能量值为 -1 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 10^4
 * -10^9 <= nums[i] <= 10^9
 * 1 <= k <= n
 * 1 <= n * k <= 10^6
 * k 是奇数。
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/maximum-strength-of-k-disjoint-subarrays/solutions/2678061/qian-zhui-he-hua-fen-xing-dpshi-zi-bian-ap5z5/
//
// 划分型 DP, 思路很明确, 但最后 `O(n^2*k)` 到 `O(n^2)` 的化简很值得学习.
public class Solution
{
    public long MaximumStrength(int[] nums, int k)
    {
        var n = nums.Length;
        var S = new long[n + 1];
        for (var i = 0; i < n; i++)
        {
            S[i + 1] = S[i] + (long)nums[i];
        }
        var dp = new long[2][]
        {
            new long[n + 1],
            new long[n + 1],
        };
        for (var i = 1; i <= k; i++)
        {
            dp[i % 2][i - 1] = long.MinValue;
            var max = long.MinValue;
            var w = (k - i + 1) * (i % 2 == 1 ? 1 : -1);
            for (var j = i; j <= n - k + i; j++)
            {
                max = Math.Max(max, dp[(i - 1) % 2][j - 1] - S[j - 1] * w);
                dp[i % 2][j] = Math.Max(dp[i % 2][j - 1], S[j] * w + max);
            }
        }
        return dp[k % 2].Last();
    }
}
