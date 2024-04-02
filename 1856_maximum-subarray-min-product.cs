/*
 * Medium
 * 
 * 一个数组的 最小乘积 定义为这个数组中 最小值 乘以 数组的 和 。
 * 
 * 比方说，数组 [3,2,5] （最小值是 2）的最小乘积为 2 * (3+2+5) = 2 * 10 = 20 。
 * 给你一个正整数数组 nums ，请你返回 nums 任意 非空子数组 的最小乘积 的 最大值 。由于答案可能很大，请你返回答案对  109 + 7 取余 的结果。
 * 
 * 请注意，最小乘积的最大值考虑的是取余操作 之前 的结果。题目保证最小乘积的最大值在 不取余 的情况下可以用 64 位有符号整数 保存。
 * 
 * 子数组 定义为一个数组的 连续 部分。
 * 
 *  
 * 
 * 示例 1：
 * 
 * 输入：nums = [1,2,3,2]
 * 输出：14
 * 解释：最小乘积的最大值由子数组 [2,3,2] （最小值是 2）得到。
 * 2 * (2+3+2) = 2 * 7 = 14 。
 * 示例 2：
 * 
 * 输入：nums = [2,3,3,1,2]
 * 输出：18
 * 解释：最小乘积的最大值由子数组 [3,3] （最小值是 3）得到。
 * 3 * (3+3) = 3 * 6 = 18 。
 * 示例 3：
 * 
 * 输入：nums = [3,1,5,6,4,2]
 * 输出：60
 * 解释：最小乘积的最大值由子数组 [5,6,4] （最小值是 4）得到。
 * 4 * (5+6+4) = 4 * 15 = 60 。
 *  
 * 
 * 提示：
 * 
 * 1 <= nums.length <= 105
 * 1 <= nums[i] <= 107
 * 通过次数 10.3K
 * 提交次数 27.2K
 * 通过率 37.8%
 * 
 */

public class Solution
{
    public int MaxSumMinProduct(int[] nums)
    {
        const long Mod = (long)1e9 + 7;
        var n = nums.Length;
        var S = new long[n + 1];
        for (var i = 0; i < n; i++) { S[i + 1] = S[i] + (long)nums[i]; }
        var lefts = new int[n];
        var stk = new Stack<int>();
        stk.Push(-1);
        for (var i = 0; i < n; i++)
        {
            for (; stk.Peek() is not -1 && nums[stk.Peek()] >= nums[i]; stk.Pop()) { }
            lefts[i] = stk.Peek();
            stk.Push(i);
        }
        var rights = new int[n];
        stk = new Stack<int>();
        stk.Push(n);
        for (var i = n - 1; i >= 0; i--)
        {
            for (; stk.Peek() != n && nums[stk.Peek()] > nums[i]; stk.Pop()) { }
            rights[i] = stk.Peek();
            stk.Push(i);
        }
        return (int)(Enumerable
            .Range(0, n)
            .Select(i => (long)nums[i] * (S[rights[i]] - S[lefts[i] + 1]))
            .Max() % Mod);
    }
}
