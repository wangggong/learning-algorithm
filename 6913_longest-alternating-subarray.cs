/*
 * @lc app=leetcode.cn id=longest-alternating-subarray lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6913] 最长交替子序列
 *
 * https://leetcode.cn/problems/longest-alternating-subarray/description/
 *
 * algorithms
 * Easy (40.39%)
 * Total Accepted:    1.9K
 * Total Submissions: 4.8K
 * Testcase Example:  '[2,3,4,3,4]'
 *
 * 给你一个下标从 0 开始的整数数组 nums 。如果 nums 中长度为 m 的子数组 s 满足以下条件，我们称它是一个交替子序列：
 * 
 * 
 * m 大于 1 。
 * s1 = s0 + 1 。
 * 下标从 0 开始的子数组 s 与数组 [s0, s1, s0, s1,...,s(m-1) % 2] 一样。也就是说，s1 - s0 = 1 ，s2 -
 * s1 = -1 ，s3 - s2 = 1 ，s4 - s3 = -1 ，以此类推，直到 s[m - 1] - s[m - 2] = (-1)^m 。
 * 
 * 
 * 请你返回 nums 中所有 交替 子数组中，最长的长度，如果不存在交替子数组，请你返回 -1 。
 * 
 * 子数组是一个数组中一段连续 非空 的元素序列。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：nums = [2,3,4,3,4]
 * 输出：4
 * 解释：交替子数组有 [3,4] ，[3,4,3] 和 [3,4,3,4] 。最长的子数组为 [3,4,3,4] ，长度为4 。
 * 
 * 
 * 示例 2：
 * 
 * 输入：nums = [4,5,6]
 * 输出：2
 * 解释：[4,5] 和 [5,6] 是仅有的两个交替子数组。它们长度都为 2 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 2 <= nums.length <= 100
 * 1 <= nums[i] <= 10^4
 * 
 * 
 */
public class Solution
{
    public int AlternatingSubarray(int[] nums)
    {
        var ans = -1;
        for (var (p, n) = (0, nums.Length); p < n; p++)
        {
            var q = p + 1;
            for (; q < n && nums[q] - nums[p] == (q - p) % 2; q++) { }
            if (q - p >= 2) { ans = Math.Max(ans, q - p); }
        }
        return ans;
    }
}
