/*
 * @lc app=leetcode.cn id=longest-non-decreasing-subarray-from-two-arrays lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6912] 构造最长非递减子数组
 *
 * https://leetcode.cn/problems/longest-non-decreasing-subarray-from-two-arrays/description/
 *
 * algorithms
 * Medium (28.18%)
 * Total Accepted:    2.4K
 * Total Submissions: 8.5K
 * Testcase Example:  '[2,3,1]\n[1,2,1]'
 *
 * 给你两个下标从 0 开始的整数数组 nums1 和 nums2 ，长度均为 n 。
 * 
 * 让我们定义另一个下标从 0 开始、长度为 n 的整数数组，nums3 。对于范围 [0, n - 1] 的每个下标 i ，你可以将 nums1[i] 或
 * nums2[i] 的值赋给 nums3[i] 。
 * 
 * 你的任务是使用最优策略为 nums3 赋值，以最大化 nums3 中 最长非递减子数组 的长度。
 * 
 * 以整数形式表示并返回 nums3 中 最长非递减 子数组的长度。
 * 
 * 注意：子数组 是数组中的一个连续非空元素序列。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：nums1 = [2,3,1], nums2 = [1,2,1]
 * 输出：2
 * 解释：构造 nums3 的方法之一是： 
 * nums3 = [nums1[0], nums2[1], nums2[2]] => [2,2,1]
 * 从下标 0 开始到下标 1 结束，形成了一个长度为 2 的非递减子数组 [2,2] 。 
 * 可以证明 2 是可达到的最大长度。
 * 
 * 示例 2：
 * 
 * 输入：nums1 = [1,3,2,1], nums2 = [2,2,3,4]
 * 输出：4
 * 解释：构造 nums3 的方法之一是： 
 * nums3 = [nums1[0], nums2[1], nums2[2], nums2[3]] => [1,2,3,4]
 * 整个数组形成了一个长度为 4 的非递减子数组，并且是可达到的最大长度。
 * 
 * 
 * 示例 3：
 * 
 * 输入：nums1 = [1,1], nums2 = [2,2]
 * 输出：2
 * 解释：构造 nums3 的方法之一是： 
 * nums3 = [nums1[0], nums1[1]] => [1,1] 
 * 整个数组形成了一个长度为 2 的非递减子数组，并且是可达到的最大长度。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums1.length == nums2.length == n <= 10^5
 * 1 <= nums1[i], nums2[i] <= 10^9
 * 
 * 
 */
public class Solution
{
    public int MaxNonDecreasingLength(int[] nums1, int[] nums2)
    {
        var n = nums1.Length;
        var dp = Enumerable.Range(0, n).Select(_ => new int[2]).ToArray();
        (dp[0][0], dp[0][1]) = (1, 1);
        for (var i = 1; i < n; i++)
        {
            dp[i][0] = Math.Max(
                (nums1[i - 1] <= nums1[i] ? dp[i - 1][0] : 0) + 1,
                (nums2[i - 1] <= nums1[i] ? dp[i - 1][1] : 0) + 1);
            dp[i][1] = Math.Max(
                (nums1[i - 1] <= nums2[i] ? dp[i - 1][0] : 0) + 1,
                (nums2[i - 1] <= nums2[i] ? dp[i - 1][1] : 0) + 1);
        }
        return dp.SelectMany(row => row).Max();
    }
}
