/*
 * @lc app=leetcode.cn id=check-if-array-is-sorted-and-rotated lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1752] 检查数组是否经排序和轮转得到
 *
 * https://leetcode.cn/problems/check-if-array-is-sorted-and-rotated/description/
 *
 * algorithms
 * Easy (61.56%)
 * Total Accepted:    16.1K
 * Total Submissions: 26.4K
 * Testcase Example:  '[3,4,5,1,2]'
 *
 * 给你一个数组 nums 。nums 的源数组中，所有元素与 nums 相同，但按非递减顺序排列。
 * 
 * 如果 nums 能够由源数组轮转若干位置（包括 0 个位置）得到，则返回 true ；否则，返回 false 。
 * 
 * 源数组中可能存在 重复项 。
 * 
 * 注意：我们称数组 A 在轮转 x 个位置后得到长度相同的数组 B ，当它们满足 A[i] == B[(i+x) % A.length] ，其中 %
 * 为取余运算。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [3,4,5,1,2]
 * 输出：true
 * 解释：[1,2,3,4,5] 为有序的源数组。
 * 可以轮转 x = 3 个位置，使新数组从值为 3 的元素开始：[3,4,5,1,2] 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [2,1,3,4]
 * 输出：false
 * 解释：源数组无法经轮转得到 nums 。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [1,2,3]
 * 输出：true
 * 解释：[1,2,3] 为有序的源数组。
 * 可以轮转 x = 0 个位置（即不轮转）得到 nums 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 100
 * 1 <= nums[i] <= 100
 * 
 * 
 */

public class Solution
{
    public bool Check(int[] nums)
    {
        var duplicate = new List<int>();
        foreach (var num in nums) { duplicate.Add(num); }
        foreach (var num in nums) { duplicate.Add(num); }
        var n = nums.Length;
        Array.Sort(nums);
        for (int i = 1; i <= n; i++)
        {
            var valid = true;
            for (int j = 0; j < n; j++)
            {
                if (duplicate[i + j] != nums[j])
                {
                    valid = false;
                    break;
                }
            }
            if (valid) { return true; }
        }
        return false;
    }
}

/*
 * public class Solution
 * {
 *     private bool IsIncrease(int[] nums)
 *     {
 *         for (int i = 0, n = nums.Length; i + 1 < n; i++)
 *         {
 *             if (nums[i] > nums[i + 1]) { return false; }
 *         }
 *         return true;
 *     }
 * 
 *     public bool Check(int[] nums)
 *     {
 *         if (IsIncrease(nums)) { return true; }
 *         var ind = 0;
 *         var n = nums.Length;
 *         for (; ind + 1 < n; ind++)
 *         {
 *             if (nums[ind] > nums[ind + 1]) { break; }
 *         }
 *         for (int i = 1; i < n; i++)
 *         {
 *             if (nums[(ind + i) % n] > nums[(ind + 1 + i) % n]) { return false; }
 *         }
 *         return true;
 *     }
 * }
 */
