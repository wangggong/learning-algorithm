/*
 * @lc app=leetcode.cn id=count-alternating-subarrays lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [100266] 交替子数组计数
 *
 * https://leetcode.cn/problems/count-alternating-subarrays/description/
 *
 * algorithms
 * Medium (47.75%)
 * Total Accepted:    3.5K
 * Total Submissions: 7.4K
 * Testcase Example:  '[0,1,1,1]'
 *
 * 给你一个二进制数组 nums 。
 * 
 * 如果一个子数组中 不存在 两个 相邻 元素的值 相同 的情况，我们称这样的子数组为 交替子数组 。
 * 
 * 返回数组 nums 中交替子数组的数量。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入： nums = [0,1,1,1]
 * 
 * 输出： 5
 * 
 * 解释：
 * 
 * 
 * 以下子数组是交替子数组：[0] 、[1] 、[1] 、[1] 以及 [0,1] 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入： nums = [1,0,1,0]
 * 
 * 输出： 10
 * 
 * 解释：
 * 
 * 
 * 数组的每个子数组都是交替子数组。可以统计在内的子数组共有 10 个。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 10^5
 * nums[i] 不是 0 就是 1 。
 * 
 * 
 */

// 基础不牢, 地动山摇
public class Solution
{
    public long CountAlternatingSubarrays(int[] nums)
    {
        var ans = 0l;
        for (var (p, q, n) = (0l, 0l, (long)nums.Length); p < n; p = q)
        {
            for (; q < n && Math.Abs(nums[q] - nums[p]) == (q - p) % 2; q++) { }
            var d = q - p;
            ans += d * (d + 1) / 2;
        }
        return ans;
    }
}
