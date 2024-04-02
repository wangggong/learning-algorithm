/*
 * @lc app=leetcode.cn id=minimum-number-of-operations-to-make-array-empty lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [100032] 使数组为空的最少操作次数
 *
 * https://leetcode.cn/problems/minimum-number-of-operations-to-make-array-empty/description/
 *
 * algorithms
 * Medium (60.73%)
 * Total Accepted:    2K
 * Total Submissions: 3.3K
 * Testcase Example:  '[2,3,3,2,2,4,2,3,4]'
 *
 * 给你一个下标从 0 开始的正整数数组 nums 。
 * 
 * 你可以对数组执行以下两种操作 任意次 ：
 * 
 * 
 * 从数组中选择 两个 值 相等 的元素，并将它们从数组中 删除 。
 * 从数组中选择 三个 值 相等 的元素，并将它们从数组中 删除 。
 * 
 * 
 * 请你返回使数组为空的 最少 操作次数，如果无法达成，请返回 -1 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [2,3,3,2,2,4,2,3,4]
 * 输出：4
 * 解释：我们可以执行以下操作使数组为空：
 * - 对下标为 0 和 3 的元素执行第一种操作，得到 nums = [3,3,2,4,2,3,4] 。
 * - 对下标为 2 和 4 的元素执行第一种操作，得到 nums = [3,3,4,3,4] 。
 * - 对下标为 0 ，1 和 3 的元素执行第二种操作，得到 nums = [4,4] 。
 * - 对下标为 0 和 1 的元素执行第一种操作，得到 nums = [] 。
 * 至少需要 4 步操作使数组为空。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [2,1,2,2,3,3]
 * 输出：-1
 * 解释：无法使数组为空。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 2 <= nums.length <= 10^5
 * 1 <= nums[i] <= 10^6
 * 
 * 
 */
public class Solution
{
    public int MinOperations(int[] nums)
    {
        var ans = 0;
        foreach (var c in nums
            .GroupBy(x => x)
            .Select(g => g.Count()))
        {
            if (c is 1) { return -1; }
            ans += (c % 3) switch
            {
                1 => (c - 1) / 3 + 1,
                2 => c / 3 + 1,
                _ => c / 3,
            };
        }
        return ans;
    }
}
