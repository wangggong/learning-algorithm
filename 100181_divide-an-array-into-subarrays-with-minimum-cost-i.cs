/*
 * @lc app=leetcode.cn id=divide-an-array-into-subarrays-with-minimum-cost-i lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [100181] 将数组分成最小总代价的子数组 I
 *
 * https://leetcode.cn/problems/divide-an-array-into-subarrays-with-minimum-cost-i/description/
 *
 * algorithms
 * Easy (69.77%)
 * Total Accepted:    2.3K
 * Total Submissions: 3.3K
 * Testcase Example:  '[1,2,3,12]'
 *
 * 给你一个长度为 n 的整数数组 nums 。
 * 
 * 一个数组的 代价 是它的 第一个 元素。比方说，[1,2,3] 的代价是 1 ，[3,4,1] 的代价是 3 。
 * 
 * 你需要将 nums 分成 3 个 连续且没有交集 的子数组。
 * 
 * 请你返回这些子数组的 最小 代价 总和 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [1,2,3,12]
 * 输出：6
 * 解释：最佳分割成 3 个子数组的方案是：[1] ，[2] 和 [3,12] ，总代价为 1 + 2 + 3 = 6 。
 * 其他得到 3 个子数组的方案是：
 * - [1] ，[2,3] 和 [12] ，总代价是 1 + 2 + 12 = 15 。
 * - [1,2] ，[3] 和 [12] ，总代价是 1 + 3 + 12 = 16 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [5,4,3]
 * 输出：12
 * 解释：最佳分割成 3 个子数组的方案是：[5] ，[4] 和 [3] ，总代价为 5 + 4 + 3 = 12 。
 * 12 是所有分割方案里的最小总代价。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [10,3,1,1]
 * 输出：12
 * 解释：最佳分割成 3 个子数组的方案是：[10,3] ，[1] 和 [1] ，总代价为 10 + 1 + 1 = 12 。
 * 12 是所有分割方案里的最小总代价。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 3 <= n <= 50
 * 1 <= nums[i] <= 50
 * 
 * 
 */
public class Solution
{
    public int MinimumCost(int[] nums) => Enumerable
        .Range(1, nums.Length - 1)
        .SelectMany(i => Enumerable
            .Range(i + 1, nums.Length - (i + 1))
            .Select(j => nums[i] + nums[j] + nums[0]))
        .Min();
}
