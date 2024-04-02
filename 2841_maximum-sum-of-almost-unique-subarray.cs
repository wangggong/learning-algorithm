/*
 * @lc app=leetcode.cn id=maximum-sum-of-almost-unique-subarray lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [2841] 几乎唯一子数组的最大和
 *
 * https://leetcode.cn/problems/maximum-sum-of-almost-unique-subarray/description/
 *
 * algorithms
 * Medium (39.28%)
 * Total Accepted:    2.3K
 * Total Submissions: 5.9K
 * Testcase Example:  '[2,6,7,3,1,7]\n3\n4'
 *
 * 给你一个整数数组 nums 和两个正整数 m 和 k 。
 * 
 * 请你返回 nums 中长度为 k 的 几乎唯一 子数组的 最大和 ，如果不存在几乎唯一子数组，请你返回 0 。
 * 
 * 如果 nums 的一个子数组有至少 m 个互不相同的元素，我们称它是 几乎唯一 子数组。
 * 
 * 子数组指的是一个数组中一段连续 非空 的元素序列。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [2,6,7,3,1,7], m = 3, k = 4
 * 输出：18
 * 解释：总共有 3 个长度为 k = 4 的几乎唯一子数组。分别为 [2, 6, 7, 3] ，[6, 7, 3, 1] 和 [7, 3, 1, 7]
 * 。这些子数组中，和最大的是 [2, 6, 7, 3] ，和为 18 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [5,9,9,2,4,5,4], m = 1, k = 3
 * 输出：23
 * 解释：总共有 5 个长度为 k = 3 的几乎唯一子数组。分别为 [5, 9, 9] ，[9, 9, 2] ，[9, 2, 4] ，[2, 4, 5]
 * 和 [4, 5, 4] 。这些子数组中，和最大的是 [5, 9, 9] ，和为 23 。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [1,2,1,2,1,2,1], m = 3, k = 3
 * 输出：0
 * 解释：输入数组中不存在长度为 k = 3 的子数组含有至少  m = 3 个互不相同元素的子数组。所以不存在几乎唯一子数组，最大和为 0 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 2 * 10^4
 * 1 <= m <= k <= nums.length
 * 1 <= nums[i] <= 10^9
 * 
 * 
 */
public class Solution
{
    public long MaxSum(IList<int> nums, int m, int k)
    {
        var (ans, total) = (0L, 0L);
        var d = new Dictionary<int, int>();
        for (var i = 0; i < k; i++)
        {
            total += (long)nums[i];
            d.TryGetValue(nums[i], out var c);
            d[nums[i]] = c + 1;
        }
        if (d.Count() >= m) { ans = Math.Max(ans, total); }
        for (var (i, n) = (k, nums.Count()); i < n; i++)
        {
            total -= (long)nums[i - k];
            d[nums[i - k]]--;
            if (d[nums[i - k]] is 0) { d.Remove(nums[i - k]); }
            total += (long)nums[i];
            d.TryGetValue(nums[i], out var c);
            d[nums[i]] = c + 1;
            if (d.Count() >= m) { ans = Math.Max(ans, total); }
        }
        return ans;
    }
}
