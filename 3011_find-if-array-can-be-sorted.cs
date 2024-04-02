/*
 * @lc app=leetcode.cn id=find-if-array-can-be-sorted lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [3011] 判断一个数组是否可以变为有序
 *
 * https://leetcode.cn/problems/find-if-array-can-be-sorted/description/
 *
 * algorithms
 * Medium (47.87%)
 * Total Accepted:    2.1K
 * Total Submissions: 4.3K
 * Testcase Example:  '[8,4,2,30,15]'
 *
 * 给你一个下标从 0 开始且全是 正 整数的数组 nums 。
 * 
 * 一次 操作 中，如果两个 相邻 元素在二进制下数位为 1 的数目 相同 ，那么你可以将这两个元素交换。你可以执行这个操作 任意次 （也可以 0 次）。
 * 
 * 如果你可以使数组变有序，请你返回 true ，否则返回 false 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [8,4,2,30,15]
 * 输出：true
 * 解释：我们先观察每个元素的二进制表示。 2 ，4 和 8 分别都只有一个数位为 1 ，分别为 "10" ，"100" 和 "1000" 。15 和 30
 * 分别有 4 个数位为 1 ："1111" 和 "11110" 。
 * 我们可以通过 4 个操作使数组有序：
 * - 交换 nums[0] 和 nums[1] 。8 和 4 分别只有 1 个数位为 1 。数组变为 [4,8,2,30,15] 。
 * - 交换 nums[1] 和 nums[2] 。8 和 2 分别只有 1 个数位为 1 。数组变为 [4,2,8,30,15] 。
 * - 交换 nums[0] 和 nums[1] 。4 和 2 分别只有 1 个数位为 1 。数组变为 [2,4,8,30,15] 。
 * - 交换 nums[3] 和 nums[4] 。30 和 15 分别有 4 个数位为 1 ，数组变为 [2,4,8,15,30] 。
 * 数组变成有序的，所以我们返回 true 。
 * 注意我们还可以通过其他的操作序列使数组变得有序。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [1,2,3,4,5]
 * 输出：true
 * 解释：数组已经是有序的，所以我们返回 true 。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [3,16,8,4,2]
 * 输出：false
 * 解释：无法通过操作使数组变为有序。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 100
 * 1 <= nums[i] <= 2^8
 * 
 * 
 */
public class Solution
{
    public bool CanSortArray(int[] nums)
    {
        const int B = 9;
        int bitCount(int n) => Enumerable
            .Range(0, B)
            .Count(b => ((n >> b) & 1) is not 0);
        var bits = nums
            .Select(bitCount)
            .ToArray();
        for (var (i, n) = (0, nums.Length); i < n; i++)
        {
            for (var j = 0; j + 1 < n; j++)
            {
                if (nums[j] > nums[j + 1] && bits[j] == bits[j + 1])
                { (nums[j], nums[j + 1]) = (nums[j + 1], nums[j]); }
            }
        }
        return Enumerable
            .Range(0, nums.Length - 1)
            .All(i => nums[i] <= nums[i + 1]);
    }
}
