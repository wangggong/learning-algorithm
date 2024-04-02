/*
 * @lc app=leetcode.cn id=maximum-value-of-an-ordered-triplet-ii lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [100086] 有序三元组中的最大值 II
 *
 * https://leetcode.cn/problems/maximum-value-of-an-ordered-triplet-ii/description/
 *
 * algorithms
 * Medium (38.99%)
 * Total Accepted:    2.3K
 * Total Submissions: 5.9K
 * Testcase Example:  '[12,6,1,2,7]'
 *
 * 给你一个下标从 0 开始的整数数组 nums 。
 * 
 * 请你从所有满足 i < j < k 的下标三元组 (i, j, k) 中，找出并返回下标三元组的最大值。如果所有满足条件的三元组的值都是负数，则返回 0
 * 。
 * 
 * 下标三元组 (i, j, k) 的值等于 (nums[i] - nums[j]) * nums[k] 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [12,6,1,2,7]
 * 输出：77
 * 解释：下标三元组 (0, 2, 4) 的值是 (nums[0] - nums[2]) * nums[4] = 77 。
 * 可以证明不存在值大于 77 的有序下标三元组。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [1,10,3,4,19]
 * 输出：133
 * 解释：下标三元组 (1, 2, 4) 的值是 (nums[1] - nums[2]) * nums[4] = 133 。
 * 可以证明不存在值大于 133 的有序下标三元组。 
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [1,2,3]
 * 输出：0
 * 解释：唯一的下标三元组 (0, 1, 2) 的值是一个负数，(nums[0] - nums[1]) * nums[2] = -3 。因此，答案是 0
 * 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 3 <= nums.length <= 10^5
 * 1 <= nums[i] <= 10^6
 * 
 * 
 */
public class Solution
{
    public long MaximumTripletValue(int[] nums)
    {
        var ans = 0l;
        var n = nums.Length;
        var diffs = new long[n];
        var max = (long)nums.First();
        for (var i = 1; i < n; i++)
        {
            diffs[i] = max - (long)nums[i];
            max = Math.Max(max, (long)nums[i]);
        }
        max = (long)nums.Last();
        for (var i = n - 1; i >= 2; i--)
        {
            max = Math.Max(max, nums[i]);
            ans = Math.Max(ans, diffs[i - 1] * max);
        }
        return ans;
    }
}
