/*
 * @lc app=leetcode.cn id=ways-to-split-array-into-good-subarrays lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6910] 将数组划分成若干好子数组的方式
 *
 * https://leetcode.cn/problems/ways-to-split-array-into-good-subarrays/description/
 *
 * algorithms
 * Medium (33.10%)
 * Total Accepted:    1.6K
 * Total Submissions: 4.9K
 * Testcase Example:  '[0,1,0,0,1]'
 *
 * 给你一个二元数组 nums 。
 * 
 * 如果数组中的某个子数组 恰好 只存在 一 个值为 1 的元素，则认为该子数组是一个 好子数组 。
 * 
 * 请你统计将数组 nums 划分成若干 好子数组 的方法数，并以整数形式返回。由于数字可能很大，返回其对 10^9 + 7 取余 之后的结果。
 * 
 * 子数组是数组中的一个连续 非空 元素序列。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：nums = [0,1,0,0,1]
 * 输出：3
 * 解释：存在 3 种可以将 nums 划分成若干好子数组的方式：
 * - [0,1] [0,0,1]
 * - [0,1,0] [0,1]
 * - [0,1,0,0] [1]
 * 
 * 
 * 示例 2：
 * 
 * 输入：nums = [0,1,0]
 * 输出：1
 * 解释：存在 1 种可以将 nums 划分成若干好子数组的方式：
 * - [0,1,0]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 10^5
 * 0 <= nums[i] <= 1
 * 
 * 
 */
public class Solution
{
    private const long Mod = (long)1e9 + 7;

    public int NumberOfGoodSubarraySplits(int[] nums)
    {
        var indexes = nums
            .Select((n, i) => (n, (long)i))
            .Where(x => x.n is 1)
            .Select(x => x.Item2)
            .ToList();
        return !indexes.Any()
            ? 0
            : (indexes.Count() is 1
                ? 1
                : (int)Enumerable
                    .Range(1, indexes.Count() - 1)
                    .Select(i => indexes[i] - indexes[i - 1])
                    .Aggregate((x, y) => x * y % Mod));
    }
}
