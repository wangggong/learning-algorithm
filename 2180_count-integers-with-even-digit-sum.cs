/*
 * @lc app=leetcode.cn id=count-integers-with-even-digit-sum lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [2180] 统计各位数字之和为偶数的整数个数
 *
 * https://leetcode.cn/problems/count-integers-with-even-digit-sum/description/
 *
 * algorithms
 * Easy (64.74%)
 * Total Accepted:    13.8K
 * Total Submissions: 20.6K
 * Testcase Example:  '4'
 *
 * 给你一个正整数 num ，请你统计并返回 小于或等于 num 且各位数字之和为 偶数 的正整数的数目。
 * 
 * 正整数的 各位数字之和 是其所有位上的对应数字相加的结果。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：num = 4
 * 输出：2
 * 解释：
 * 只有 2 和 4 满足小于等于 4 且各位数字之和为偶数。    
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：num = 30
 * 输出：14
 * 解释：
 * 只有 14 个整数满足小于等于 30 且各位数字之和为偶数，分别是： 
 * 2、4、6、8、11、13、15、17、19、20、22、24、26 和 28 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= num <= 1000
 * 
 * 
 */
public class Solution
{
    public int CountEven(int num)
        => Enumerable
            .Range(1, num)
            .Where(x =>
                {
                    var tot = 0;
                    for (; x > 0; x /= 10)
                    {
                        tot += x % 10;
                    }
                    return tot % 2 == 0;
                }).Count();
}
