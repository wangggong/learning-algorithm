/*
 * @lc app=leetcode.cn id=count-subarrays-where-max-element-appears-at-least-k-times lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [100137] 统计最大元素出现至少 K 次的子数组
 *
 * https://leetcode.cn/problems/count-subarrays-where-max-element-appears-at-least-k-times/description/
 *
 * algorithms
 * Medium (41.51%)
 * Total Accepted:    3K
 * Total Submissions: 7.1K
 * Testcase Example:  '[1,3,2,3,3]\n2'
 *
 * 给你一个整数数组 nums 和一个 正整数 k 。
 * 
 * 请你统计有多少满足 「 nums 中的 最大 元素」至少出现 k 次的子数组，并返回满足这一条件的子数组的数目。
 * 
 * 子数组是数组中的一个连续元素序列。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [1,3,2,3,3], k = 2
 * 输出：6
 * 解释：包含元素 3 至少 2 次的子数组为：[1,3,2,3]、[1,3,2,3,3]、[3,2,3]、[3,2,3,3]、[2,3,3] 和
 * [3,3] 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [1,4,2,1], k = 3
 * 输出：0
 * 解释：没有子数组包含元素 4 至少 3 次。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 10^5
 * 1 <= nums[i] <= 10^6
 * 1 <= k <= 10^5
 * 
 * 
 */
public class Solution
{
    public long CountSubarrays(int[] nums, int k)
    {
        var ans = 0l;
        var max = nums.Max();
        for (var (p, q, n, c) = (0, 0, nums.Length, 0); q < n; q++)
        {
            c += nums[q] == max ? 1 : 0;
            for (; c >= k; p++) { c -= nums[p] == max ? 1 : 0; }
            ans += (long)p;
        }
        return ans;
    }
}
