/*
 * @lc app=leetcode.cn id=minimum-elements-to-add-to-form-a-given-sum lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1785] 构成特定和需要添加的最少元素
 *
 * https://leetcode.cn/problems/minimum-elements-to-add-to-form-a-given-sum/description/
 *
 * algorithms
 * Medium (36.60%)
 * Total Accepted:    12.9K
 * Total Submissions: 31.9K
 * Testcase Example:  '[1,-1,1]\n3\n-4'
 *
 * 给你一个整数数组 nums ，和两个整数 limit 与 goal 。数组 nums 有一条重要属性：abs(nums[i])  。
 * 
 * 返回使数组元素总和等于 goal 所需要向数组中添加的 最少元素数量 ，添加元素 不应改变 数组中 abs(nums[i])  这一属性。
 * 
 * 注意，如果 x >= 0 ，那么 abs(x) 等于 x ；否则，等于 -x 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [1,-1,1], limit = 3, goal = -4
 * 输出：2
 * 解释：可以将 -2 和 -3 添加到数组中，数组的元素总和变为 1 - 1 + 1 - 2 - 3 = -4 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [1,-10,9,1], limit = 100, goal = 0
 * 输出：1
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * 1 
 * -limit 
 * -10^9 
 * 
 * 
 */
public class Solution
{
    public int MinElements(int[] nums, int limit, int goal)
        => (int) ((Math.Abs((long) goal - nums.Select(x => (long) x).Sum()) + (long) limit - 1) / (long) limit);
}
