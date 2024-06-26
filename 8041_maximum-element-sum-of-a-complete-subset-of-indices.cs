/*
 * @lc app=leetcode.cn id=maximum-element-sum-of-a-complete-subset-of-indices lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [8041] 完全子集的最大元素和
 *
 * https://leetcode.cn/problems/maximum-element-sum-of-a-complete-subset-of-indices/description/
 *
 * algorithms
 * Hard (31.05%)
 * Total Accepted:    918
 * Total Submissions: 3K
 * Testcase Example:  '[8,7,3,5,7,2,4,9]'
 *
 * 给你一个下标从 1 开始、由 n 个整数组成的数组。
 * 
 * 如果一组数字中每对元素的乘积都是一个完全平方数，则称这组数字是一个 完全集 。
 * 
 * 下标集 {1, 2, ..., n} 的子集可以表示为 {i1, i2, ..., ik}，我们定义对应该子集的 元素和 为 nums[i1] +
 * nums[i2] + ... + nums[ik] 。
 * 
 * 返回下标集 {1, 2, ..., n} 的 完全子集 所能取到的 最大元素和 。
 * 
 * 完全平方数是指可以表示为一个整数和其自身相乘的数。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [8,7,3,5,7,2,4,9]
 * 输出：16
 * 解释：除了由单个下标组成的子集之外，还有两个下标集的完全子集：{1,4} 和 {2,8} 。
 * 与下标 1 和 4 对应的元素和等于 nums[1] + nums[4] = 8 + 5 = 13 。
 * 与下标 2 和 8 对应的元素和等于 nums[2] + nums[8] = 7 + 9 = 16 。
 * 因此，下标集的完全子集可以取到的最大元素和为 16 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [5,10,3,10,1,13,7,9,4]
 * 输出：19
 * 解释：除了由单个下标组成的子集之外，还有四个下标集的完全子集：{1,4}、{1,9}、{2,8}、{4,9} 和 {1,4,9} 。 
 * 与下标 1 和 4 对应的元素和等于 nums[1] + nums[4] = 5 + 10 = 15 。 
 * 与下标 1 和 9 对应的元素和等于 nums[1] + nums[9] = 5 + 4 = 9 。 
 * 与下标 2 和 8 对应的元素和等于 nums[2] + nums[8] = 10 + 9 = 19 。
 * 与下标 4 和 9 对应的元素和等于 nums[4] + nums[9] = 10 + 4 = 14 。 
 * 与下标 1、4 和 9 对应的元素和等于 nums[1] + nums[4] + nums[9] = 5 + 10 + 4 = 19 。 
 * 因此，下标集的完全子集可以取到的最大元素和为 19 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n == nums.length <= 10^4
 * 1 <= nums[i] <= 10^9
 * 
 * 
 */
public class Solution
{
    public long MaximumSum(IList<int> nums)
    {
        var n = nums.Count();
        var visits = new bool[n];
        var ans = 0l;
        for (var i = 1; i <= n; i++)
        {
            if (visits[i - 1]) { continue; }
            var total = 0l;
            for (var j = 1; i * j * j <= n; j++)
            {
                total += (long)nums[i * j * j - 1];
                visits[i * j * j - 1] = true;
            }
            ans = Math.Max(ans, total);
        }
        return ans;
    }
}
