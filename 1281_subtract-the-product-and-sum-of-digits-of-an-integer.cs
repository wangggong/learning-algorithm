/*
 * @lc app=leetcode.cn id=subtract-the-product-and-sum-of-digits-of-an-integer lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1281] 整数的各位积和之差
 *
 * https://leetcode.cn/problems/subtract-the-product-and-sum-of-digits-of-an-integer/description/
 *
 * algorithms
 * Easy (82.96%)
 * Total Accepted:    76.1K
 * Total Submissions: 91.5K
 * Testcase Example:  '234'
 *
 * 给你一个整数 n，请你帮忙计算并返回该整数「各位数字之积」与「各位数字之和」的差。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：n = 234
 * 输出：15 
 * 解释：
 * 各位数之积 = 2 * 3 * 4 = 24 
 * 各位数之和 = 2 + 3 + 4 = 9 
 * 结果 = 24 - 9 = 15
 * 
 * 
 * 示例 2：
 * 
 * 输入：n = 4421
 * 输出：21
 * 解释： 
 * 各位数之积 = 4 * 4 * 2 * 1 = 32 
 * 各位数之和 = 4 + 4 + 2 + 1 = 11 
 * 结果 = 32 - 11 = 21
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 10^5
 * 
 * 
 */
public class Solution
{
    public int SubtractProductAndSum(int n)
    {
        var (mul, sum) = (1, 0);
        for (; n > 0; n /= 10)
        {
            mul *= n % 10;
            sum += n % 10;
        }
        return mul - sum;
    }
}
