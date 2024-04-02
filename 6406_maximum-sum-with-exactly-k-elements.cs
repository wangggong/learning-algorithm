/*
 * @lc app=leetcode.cn id=maximum-sum-with-exactly-k-elements lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6406] K 个元素的最大和
 *
 * https://leetcode.cn/problems/maximum-sum-with-exactly-k-elements/description/
 *
 * algorithms
 * Easy (89.59%)
 * Total Accepted:    2.1K
 * Total Submissions: 2.3K
 * Testcase Example:  '[1,2,3,4,5]\n3'
 *
 * 给你一个下标从 0 开始的整数数组 nums 和一个整数 k 。你需要执行以下操作 恰好 k 次，最大化你的得分：
 * 
 * 
 * 从 nums 中选择一个元素 m 。
 * 将选中的元素 m 从数组中删除。
 * 将新元素 m + 1 添加到数组中。
 * 你的得分增加 m 。
 * 
 * 
 * 请你返回执行以上操作恰好 k 次后的最大得分。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [1,2,3,4,5], k = 3
 * 输出：18
 * 解释：我们需要从 nums 中恰好选择 3 个元素并最大化得分。
 * 第一次选择 5 。和为 5 ，nums = [1,2,3,4,6] 。
 * 第二次选择 6 。和为 6 ，nums = [1,2,3,4,7] 。
 * 第三次选择 7 。和为 5 + 6 + 7 = 18 ，nums = [1,2,3,4,8] 。
 * 所以我们返回 18 。
 * 18 是可以得到的最大答案。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [5,5,5], k = 2
 * 输出：11
 * 解释：我们需要从 nums 中恰好选择 2 个元素并最大化得分。
 * 第一次选择 5 。和为 5 ，nums = [5,5,6] 。
 * 第二次选择 6 。和为 6 ，nums = [5,5,7] 。
 * 所以我们返回 11 。
 * 11 是可以得到的最大答案。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 100
 * 1 <= nums[i] <= 100
 * 1 <= k <= 100
 * 
 * 
 */
public class Solution
{
    public int MaximizeSum(int[] nums, int k)
    {
        var total = 0;
        var n = nums.Length;
        while (k-- > 0)
        {
            Array.Sort(nums);
            total += nums[n - 1];
            nums[n - 1]++;
        }
        return total;
    }
}
