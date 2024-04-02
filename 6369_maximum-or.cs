/*
 * @lc app=leetcode.cn id=maximum-or lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6369] 最大或值
 *
 * https://leetcode.cn/problems/maximum-or/description/
 *
 * algorithms
 * Medium (37.67%)
 * Total Accepted:    1.4K
 * Total Submissions: 3.8K
 * Testcase Example:  '[12,9]\n1'
 *
 * 给你一个下标从 0 开始长度为 n 的整数数组 nums 和一个整数 k 。每一次操作中，你可以选择一个数并将它乘 2 。
 * 
 * 你最多可以进行 k 次操作，请你返回 nums[0] | nums[1] | ... | nums[n - 1] 的最大值。
 * 
 * a | b 表示两个整数 a 和 b 的 按位或 运算。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [12,9], k = 1
 * 输出：30
 * 解释：如果我们对下标为 1 的元素进行操作，新的数组为 [12,18] 。此时得到最优答案为 12 和 18 的按位或运算的结果，也就是 30 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [8,1,2], k = 2
 * 输出：35
 * 解释：如果我们对下标 0 处的元素进行操作，得到新数组 [32,1,2] 。此时得到最优答案为 32|1|2 = 35 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 10^5
 * 1 <= nums[i] <= 10^9
 * 1 <= k <= 15
 * 
 * 
 */

/*
 * public class Solution
 * {
 *     public long MaximumOr(int[] nums, int k)
 *     {
 *         var n = nums.Length;
 *         var arr = nums.Select(n => (long)n).ToArray();
 *         var dp = Enumerable.Range(0, n + 1).Select(_ => new long[k + 1]).ToArray();
 *         for (var i = 0; i < n; i++)
 *         {
 *             dp[i + 1][0] = dp[i][0] | arr[i];
 *             for (var j = 1; j <= k; j++)
 *             {
 *                 var v = arr[i];
 *                 for (var p = j; p >= 0; p--)
 *                 {
 *                     dp[i + 1][j] = Math.Max(dp[i + 1][j], dp[i][p] | v);
 *                     v <<= 1;
 *                 }
 *             }
 *         }
 *         // for (var i = 0; i < n; i++)
 *         // {
 *         //     Console.WriteLine(string.Join(',', dp[i]));
 *         // }
 *         return dp[n][k];
 *     }
 * }
 */

public class Solution
{
    public long MaximumOr(int[] nums, int k)
    {
        var n = nums.Length;
        var suffix = new long[n + 1];
        for (var i = n - 1; i >= 0; i--)
        {
            suffix[i] = suffix[i + 1] | (long)nums[i];
        }
        var (ans, prefix) = ((long)0, (long)0);
        for (var i = 0; i < n; i++)
        {
            ans = Math.Max(ans, prefix | ((long)(nums[i]) << k) | suffix[i + 1]);
            prefix |= (long)nums[i];
        }
        return ans;
    }
}
