/*
 * @lc app=leetcode.cn id=reverse-subarray-to-maximize-array-value lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1330] 翻转子数组得到最大的数组值
 *
 * https://leetcode.cn/problems/reverse-subarray-to-maximize-array-value/description/
 *
 * algorithms
 * Hard (42.36%)
 * Total Accepted:    5.6K
 * Total Submissions: 10.3K
 * Testcase Example:  '[2,3,1,5,4]'
 *
 * 给你一个整数数组 nums 。「数组值」定义为所有满足 0 <= i < nums.length-1 的 |nums[i]-nums[i+1]|
 * 的和。
 * 
 * 你可以选择给定数组的任意子数组，并将该子数组翻转。但你只能执行这个操作 一次 。
 * 
 * 请你找到可行的最大 数组值 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：nums = [2,3,1,5,4]
 * 输出：10
 * 解释：通过翻转子数组 [3,1,5] ，数组变成 [2,5,1,3,4] ，数组值为 10 。
 * 
 * 
 * 示例 2：
 * 
 * 输入：nums = [2,4,9,24,2,1,10]
 * 输出：68
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 3*10^4
 * -10^5 <= nums[i] <= 10^5
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/reverse-subarray-to-maximize-array-value/solution/bu-hui-hua-jian-qing-kan-zhe-pythonjavac-c2s6/
//
// 带讨论, 确实来不了...
public class Solution
{
    public int MaxValueAfterReverse(int[] nums)
    {
        var (_base, d) = (0, 0);
        var (max, min) = (0, int.MaxValue);
        for (var (i, n) = (0, nums.Length); i + 1 < n; i++)
        {
            _base += Math.Abs(nums[i + 1] - nums[i]);
            max = Math.Max(max, Math.Min(nums[i], nums[i + 1]));
            min = Math.Min(min, Math.Max(nums[i], nums[i + 1]));
            d = Math.Max(d, Math.Max(
                Math.Abs(nums[0] - nums[i + 1]) - Math.Abs(nums[i] - nums[i + 1]),
                Math.Abs(nums[n - 1] - nums[i]) - Math.Abs(nums[i + 1] - nums[i])));
        }
        return _base + Math.Max(d, 2 * (max - min));
    }
}
