/*
 * @lc app=leetcode.cn id=left-and-right-sum-differences lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6369] 左右元素和的差值
 *
 * https://leetcode.cn/problems/left-and-right-sum-differences/description/
 *
 * algorithms
 * Easy (90.83%)
 * Total Accepted:    5.1K
 * Total Submissions: 5.7K
 * Testcase Example:  '[10,4,8,3]'
 *
 * 给你一个下标从 0 开始的整数数组 nums ，请你找出一个下标从 0 开始的整数数组 answer ，其中：
 * 
 * 
 * answer.length == nums.length
 * answer[i] = |leftSum[i] - rightSum[i]|
 * 
 * 
 * 其中：
 * 
 * 
 * leftSum[i] 是数组 nums 中下标 i 左侧元素之和。如果不存在对应的元素，leftSum[i] = 0 。
 * rightSum[i] 是数组 nums 中下标 i 右侧元素之和。如果不存在对应的元素，rightSum[i] = 0 。
 * 
 * 
 * 返回数组 answer 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：nums = [10,4,8,3]
 * 输出：[15,1,11,22]
 * 解释：数组 leftSum 为 [0,10,14,22] 且数组 rightSum 为 [15,11,3,0] 。
 * 数组 answer 为 [|0 - 15|,|10 - 11|,|14 - 3|,|22 - 0|] = [15,1,11,22] 。
 * 
 * 
 * 示例 2：
 * 
 * 输入：nums = [1]
 * 输出：[0]
 * 解释：数组 leftSum 为 [0] 且数组 rightSum 为 [0] 。
 * 数组 answer 为 [|0 - 0|] = [0] 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 1000
 * 1 <= nums[i] <= 10^5
 * 
 * 
 */
public class Solution
{
    public int[] LeftRigthDifference(int[] nums)
    {
        var n = nums.Length;
        var S = new int[n + 1];
        for (var i = 0; i < n; i++)
        {
            S[i + 1] = S[i] + nums[i];
        }
        var ans = new int[n];
        for (var i = 0; i < n; i++)
        {
            ans[i] = Math.Abs((S[i] - S[0]) - (S[n] - S[i + 1]));
        }
        return ans;
    }
}
