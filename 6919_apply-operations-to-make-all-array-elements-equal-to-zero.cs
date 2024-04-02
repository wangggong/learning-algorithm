/*
 * @lc app=leetcode.cn id=apply-operations-to-make-all-array-elements-equal-to-zero lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6919] 使数组中的所有元素都等于零
 *
 * https://leetcode.cn/problems/apply-operations-to-make-all-array-elements-equal-to-zero/description/
 *
 * algorithms
 * Medium (29.67%)
 * Total Accepted:    2.1K
 * Total Submissions: 7K
 * Testcase Example:  '[2,2,3,1,1,0]\n3'
 *
 * 给你一个下标从 0 开始的整数数组 nums 和一个正整数 k 。
 * 
 * 你可以对数组执行下述操作 任意次 ：
 * 
 * 
 * 从数组中选出长度为 k 的 任一 子数组，并将子数组中每个元素都 减去 1 。
 * 
 * 
 * 如果你可以使数组中的所有元素都等于 0 ，返回  true ；否则，返回 false 。
 * 
 * 子数组 是数组中的一个非空连续元素序列。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：nums = [2,2,3,1,1,0], k = 3
 * 输出：true
 * 解释：可以执行下述操作：
 * - 选出子数组 [2,2,3] ，执行操作后，数组变为 nums = [1,1,2,1,1,0] 。
 * - 选出子数组 [2,1,1] ，执行操作后，数组变为 nums = [1,1,1,0,0,0] 。
 * - 选出子数组 [1,1,1] ，执行操作后，数组变为 nums = [0,0,0,0,0,0] 。
 * 
 * 
 * 示例 2：
 * 
 * 输入：nums = [1,3,1,1], k = 2
 * 输出：false
 * 解释：无法使数组中的所有元素等于 0 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= k <= nums.length <= 10^5
 * 0 <= nums[i] <= 10^6
 * 
 * 
 */
public class Solution
{
    /*
     * public bool CheckArray(int[] nums, int k)
     * {
     *     for (var (i, n) = (0, nums.Length); i < n; i++)
     *     {
     *         for (; i < n && nums[i] is 0; i++) { }
     *         if (i == n) { return true; }
     *         if (i + k > n) { return false; }
     *         for (var (j, d) = (0, nums[i]); j < k; j++)
     *         {
     *             nums[i + j] -= d;
     *             if (nums[i + j] < 0) { return false; }
     *         }
     *     }
     *     return true;
     * }
     */

    // 正解是差分.
    public bool CheckArray(int[] nums, int k)
    {
        var n = nums.Length;
        var d = new int[n + 1];
        d[0] = nums[0];
        for (var i = 1; i < n; i++)
        {
            d[i] = nums[i] - nums[i - 1];
        }
        for (var i = 0; i < n; i++)
        {
            if (d[i] is 0) { continue; }
            if (d[i] < 0) { return false; }
            if (i + k > n) { return false; }
            d[i + k] += d[i];
            d[i] -= d[i];
        }
        return true;
    }
}
