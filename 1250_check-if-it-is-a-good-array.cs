/*
 * @lc app=leetcode.cn id=check-if-it-is-a-good-array lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1250] 检查「好数组」
 *
 * https://leetcode.cn/problems/check-if-it-is-a-good-array/description/
 *
 * algorithms
 * Hard (59.78%)
 * Total Accepted:    5K
 * Total Submissions: 8.1K
 * Testcase Example:  '[12,5,7,23]'
 *
 * 给你一个正整数数组 nums，你需要从中任选一些子集，然后将子集中每一个数乘以一个 任意整数，并求出他们的和。
 * 
 * 假如该和结果为 1，那么原数组就是一个「好数组」，则返回 True；否则请返回 False。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：nums = [12,5,7,23]
 * 输出：true
 * 解释：挑选数字 5 和 7。
 * 5*3 + 7*(-2) = 1
 * 
 * 
 * 示例 2：
 * 
 * 输入：nums = [29,6,10]
 * 输出：true
 * 解释：挑选数字 29, 6 和 10。
 * 29*1 + 6*(-3) + 10*(-1) = 1
 * 
 * 
 * 示例 3：
 * 
 * 输入：nums = [3,6]
 * 输出：false
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 10^5
 * 1 <= nums[i] <= 10^9
 * 
 * 
 */
public class Solution
{
    private int ExGcd(int x, int y, out int a, out int b)
    {
        if (y == 0)
        {
            (a, b) = (1, 0);
            return x;
        }
        var gcd = ExGcd(y, x % y, out var d, out var c);
        (a, b) = (c, d - x / y * c);
        return gcd;
    }

    public bool IsGoodArray(int[] nums) => nums
        .Aggregate((x, y) => ExGcd(x, y, out _, out _)) == 1;
}
