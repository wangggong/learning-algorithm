/*
 * @lc app=leetcode.cn id=maximum-sum-circular-subarray lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [918] 环形子数组的最大和
 *
 * https://leetcode.cn/problems/maximum-sum-circular-subarray/description/
 *
 * algorithms
 * Medium (37.84%)
 * Total Accepted:    60.1K
 * Total Submissions: 154.7K
 * Testcase Example:  '[1,-2,3,-2]'
 *
 * 给定一个长度为 n 的环形整数数组 nums ，返回 nums 的非空 子数组 的最大可能和 。
 * 
 * 环形数组 意味着数组的末端将会与开头相连呈环状。形式上， nums[i] 的下一个元素是 nums[(i + 1) % n] ， nums[i]
 * 的前一个元素是 nums[(i - 1 + n) % n] 。
 * 
 * 子数组 最多只能包含固定缓冲区 nums 中的每个元素一次。形式上，对于子数组 nums[i], nums[i + 1], ..., nums[j]
 * ，不存在 i <= k1, k2 <= j 其中 k1 % n == k2 % n 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [1,-2,3,-2]
 * 输出：3
 * 解释：从子数组 [3] 得到最大和 3
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [5,-3,5]
 * 输出：10
 * 解释：从子数组 [5,5] 得到最大和 5 + 5 = 10
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [3,-2,2,-3]
 * 输出：3
 * 解释：从子数组 [3] 和 [3,-2,2] 都可以得到最大和 3
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * n == nums.length
 * 1 <= n <= 3 * 10^4
 * -3 * 10^4 <= nums[i] <= 3 * 10^4​​​​​​​
 * 
 * 
 */
public class Solution
{
    public int MaxSubarraySumCircular(int[] nums)
    {
        var n = nums.Length;
        var dp = new int[n];
        dp[0] = nums[0];
        var ans = nums[0];
        for (var i = 1; i < n; i++)
        {
            dp[i] = Math.Max(dp[i - 1], 0) + nums[i];
            ans = Math.Max(ans, dp[i]);
        }
        var suffixs = new int[n + 1];
        for (var i = n - 1; i >= 0; i--)
        {
            suffixs[i] = suffixs[i + 1] + nums[i];
        }
        for (var i = n - 1; i >= 0; i--)
        {
            suffixs[i] = Math.Max(suffixs[i], suffixs[i + 1]);
        }
        var prefix = 0;
        for (var i = 0; i < n; i++)
        {
            prefix += nums[i];
            ans = Math.Max(ans, prefix + suffixs[i + 1]);
        }
        return ans;
    }
}
