/*
 * @lc app=leetcode.cn id=max-pair-sum-in-an-array lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6939] 数组中的最大数对和
 *
 * https://leetcode.cn/problems/max-pair-sum-in-an-array/description/
 *
 * algorithms
 * Easy (68.54%)
 * Total Accepted:    4.6K
 * Total Submissions: 6.7K
 * Testcase Example:  '[51,71,17,24,42]'
 *
 * 给你一个下标从 0 开始的整数数组 nums 。请你从 nums 中找出和 最大 的一对数，且这两个数数位上最大的数字相等。
 * 
 * 返回最大和，如果不存在满足题意的数字对，返回 -1 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：nums = [51,71,17,24,42]
 * 输出：88
 * 解释：
 * i = 1 和 j = 2 ，nums[i] 和 nums[j] 数位上最大的数字相等，且这一对的总和 71 + 17 = 88 。 
 * i = 3 和 j = 4 ，nums[i] 和 nums[j] 数位上最大的数字相等，且这一对的总和 24 + 42 = 66 。
 * 可以证明不存在其他数对满足数位上最大的数字相等，所以答案是 88 。
 * 
 * 示例 2：
 * 
 * 输入：nums = [1,2,3,4]
 * 输出：-1
 * 解释：不存在数对满足数位上最大的数字相等。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 2 <= nums.length <= 100
 * 1 <= nums[i] <= 10^4
 * 
 * 
 */
public class Solution
{
    public int MaxSum(int[] nums) => nums
        .Select((n, i) => (n, i))
        .SelectMany(x => nums.Select((n, j) => (n: x.n, m: n, i: x.i, j: j)))
        .Where(x => x.i != x.j && x.n.ToString().Max() == x.m.ToString().Max())
        .Select(x => x.n + x.m)
        .OrderByDescending(x => x)
        .FirstOrDefault(-1);
}
