/*
 * @lc app=leetcode.cn id=maximum-sum-of-distinct-subarrays-with-length-k lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6230] 长度为 K 子数组中的最大和
 *
 * https://leetcode.cn/problems/maximum-sum-of-distinct-subarrays-with-length-k/description/
 *
 * algorithms
 * Medium (26.13%)
 * Total Accepted:    3.9K
 * Total Submissions: 14.7K
 * Testcase Example:  '[1,5,4,2,9,9,9]\n3'
 *
 * 给你一个整数数组 nums 和一个整数 k 。请你从 nums 中满足下述条件的全部子数组中找出最大子数组和：
 * 
 * 
 * 子数组的长度是 k，且
 * 子数组中的所有元素 各不相同 。
 * 
 * 
 * 返回满足题面要求的最大子数组和。如果不存在子数组满足这些条件，返回 0 。
 * 
 * 子数组 是数组中一段连续非空的元素序列。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：nums = [1,5,4,2,9,9,9], k = 3
 * 输出：15
 * 解释：nums 中长度为 3 的子数组是：
 * - [1,5,4] 满足全部条件，和为 10 。
 * - [5,4,2] 满足全部条件，和为 11 。
 * - [4,2,9] 满足全部条件，和为 15 。
 * - [2,9,9] 不满足全部条件，因为元素 9 出现重复。
 * - [9,9,9] 不满足全部条件，因为元素 9 出现重复。
 * 因为 15 是满足全部条件的所有子数组中的最大子数组和，所以返回 15 。
 * 
 * 
 * 示例 2：
 * 
 * 输入：nums = [4,4,4], k = 3
 * 输出：0
 * 解释：nums 中长度为 3 的子数组是：
 * - [4,4,4] 不满足全部条件，因为元素 4 出现重复。
 * 因为不存在满足全部条件的子数组，所以返回 0 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= k <= nums.length <= 10^5
 * 1 <= nums[i] <= 10^5
 * 
 * 
 */
public class Solution
{
    public long MaximumSubarraySum(int[] nums, int k)
    {
        long ans = 0;
        long tot = 0;
        int n = nums.Length;
        Dictionary<int, int> d = new();
        for (int i = 0; i < k; i++)
        {
            d[nums[i]] = d.ContainsKey(nums[i]) ? d[nums[i]] + 1 : 1;
            tot += (long) nums[i];
        }
        if (d.Count() == k)
        {
            ans = tot;
        }
        for (int i = k; i < n; i++)
        {
            tot -= (long) nums[i - k];
            d[nums[i - k]]--;
            if (d[nums[i - k]] == 0)
            {
                d.Remove(nums[i - k]);
            }
            d[nums[i]] = d.ContainsKey(nums[i]) ? d[nums[i]] + 1 : 1;
            tot += (long) nums[i];
            if (d.Count() == k)
            {
                ans = Math.Max(ans, tot);
            }
        }
        return ans;
    }
}
