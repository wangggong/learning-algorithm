/*
 * @lc app=leetcode.cn id=sign-of-the-product-of-an-array lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1822] 数组元素积的符号
 *
 * https://leetcode.cn/problems/sign-of-the-product-of-an-array/description/
 *
 * algorithms
 * Easy (72.42%)
 * Total Accepted:    46.2K
 * Total Submissions: 62.9K
 * Testcase Example:  '[-1,-2,-3,-4,3,2,1]'
 *
 * 已知函数 signFunc(x) 将会根据 x 的正负返回特定值：
 * 
 * 
 * 如果 x 是正数，返回 1 。
 * 如果 x 是负数，返回 -1 。
 * 如果 x 是等于 0 ，返回 0 。
 * 
 * 
 * 给你一个整数数组 nums 。令 product 为数组 nums 中所有元素值的乘积。
 * 
 * 返回 signFunc(product) 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [-1,-2,-3,-4,3,2,1]
 * 输出：1
 * 解释：数组中所有值的乘积是 144 ，且 signFunc(144) = 1
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [1,5,0,2,-3]
 * 输出：0
 * 解释：数组中所有值的乘积是 0 ，且 signFunc(0) = 0
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [-1,1,-1,1,-1]
 * 输出：-1
 * 解释：数组中所有值的乘积是 -1 ，且 signFunc(-1) = -1
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * -100 
 * 
 * 
 */
public class Solution
{
    public int ArraySign(int[] nums)
    {
        int zeroCnt = 0, negCnt = 0;
        foreach (var n in nums)
        {
            if (n == 0)
            {
                zeroCnt++;
            }
            else if (n < 0)
            {
                negCnt++;
            }
        }
        return zeroCnt > 0 ? 0 : (negCnt % 2 == 0 ? 1 : -1);
    }
}
