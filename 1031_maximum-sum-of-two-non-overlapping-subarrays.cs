/*
 * @lc app=leetcode.cn id=maximum-sum-of-two-non-overlapping-subarrays lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1031] 两个非重叠子数组的最大和
 *
 * https://leetcode.cn/problems/maximum-sum-of-two-non-overlapping-subarrays/description/
 *
 * algorithms
 * Medium (59.73%)
 * Total Accepted:    14.8K
 * Total Submissions: 23.3K
 * Testcase Example:  '[0,6,5,2,2,5,1,9,4]\n1\n2'
 *
 * 给你一个整数数组 nums 和两个整数 firstLen 和 secondLen，请你找出并返回两个非重叠 子数组 中元素的最大和，长度分别为
 * firstLen 和 secondLen 。
 * 
 * 长度为 firstLen 的子数组可以出现在长为 secondLen 的子数组之前或之后，但二者必须是不重叠的。
 * 
 * 子数组是数组的一个 连续 部分。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [0,6,5,2,2,5,1,9,4], firstLen = 1, secondLen = 2
 * 输出：20
 * 解释：子数组的一种选择中，[9] 长度为 1，[6,5] 长度为 2。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [3,8,1,3,2,1,8,9,0], firstLen = 3, secondLen = 2
 * 输出：29
 * 解释：子数组的一种选择中，[3,8,1] 长度为 3，[8,9] 长度为 2。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [2,1,5,6,0,9,5,0,3,8], firstLen = 4, secondLen = 3
 * 输出：31
 * 解释：子数组的一种选择中，[5,6,0,9] 长度为 4，[0,3,8] 长度为 3。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= firstLen, secondLen <= 1000
 * 2 <= firstLen + secondLen <= 1000
 * firstLen + secondLen <= nums.length <= 1000
 * 0 <= nums[i] <= 1000
 * 
 * 
 */
public class Solution
{
    public int MaxSumTwoNoOverlap(int[] nums, int firstLen, int secondLen)
    {
        var n = nums.Length;
        var S = new int[n + 1];
        for (var i = 0; i < n; i++)
        {
            S[i + 1] = S[i] + nums[i];
        }
        var ans = 0;
        for (var (i, max) = (firstLen, 0); i + secondLen <= n; i++)
        {
            max = Math.Max(max, S[i] - S[i - firstLen]);
            ans = Math.Max(ans, S[i + secondLen] - S[i] + max);
        }
        for (var (i, max) = (secondLen, 0); i + firstLen <= n; i++)
        {
            max = Math.Max(max, S[i] - S[i - secondLen]);
            ans = Math.Max(ans, S[i + firstLen] - S[i] + max);
        }
        return ans;
    }
}
