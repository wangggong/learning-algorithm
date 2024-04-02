/*
 * @lc app=leetcode.cn id=count-symmetric-integers lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [7020] 统计对称整数的数目
 *
 * https://leetcode.cn/problems/count-symmetric-integers/description/
 *
 * algorithms
 * Easy (69.00%)
 * Total Accepted:    4K
 * Total Submissions: 5.8K
 * Testcase Example:  '1\n100'
 *
 * 给你两个正整数 low 和 high 。
 * 
 * 对于一个由 2 * n 位数字组成的整数 x ，如果其前 n 位数字之和与后 n 位数字之和相等，则认为这个数字是一个对称整数。
 * 
 * 返回在 [low, high] 范围内的 对称整数的数目 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：low = 1, high = 100
 * 输出：9
 * 解释：在 1 到 100 范围内共有 9 个对称整数：11、22、33、44、55、66、77、88 和 99 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：low = 1200, high = 1230
 * 输出：4
 * 解释：在 1200 到 1230 范围内共有 4 个对称整数：1203、1212、1221 和 1230 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= low <= high <= 10^4
 * 
 * 
 */
public class Solution
{
    public int CountSymmetricIntegers(int low, int high) => Enumerable
        .Range(low, high - low + 1)
        .Where(i =>
        {
            var s = i.ToString();
            return (s.Length & 1) is 0
                && s[.. (s.Length >> 1)].Select(c => (int)(c - '0')).Sum()
                    == s[(s.Length >> 1) ..].Select(c => (int)(c - '0')).Sum();
        })
        .Count();
}
