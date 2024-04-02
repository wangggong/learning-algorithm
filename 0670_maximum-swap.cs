/*
 * @lc app=leetcode.cn id=maximum-swap lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [670] 最大交换
 *
 * https://leetcode.cn/problems/maximum-swap/description/
 *
 * algorithms
 * Medium (48.03%)
 * Total Accepted:    70.9K
 * Total Submissions: 147.4K
 * Testcase Example:  '2736'
 *
 * 给定一个非负整数，你至多可以交换一次数字中的任意两位。返回你能得到的最大值。
 * 
 * 示例 1 :
 * 
 * 
 * 输入: 2736
 * 输出: 7236
 * 解释: 交换数字2和数字7。
 * 
 * 
 * 示例 2 :
 * 
 * 
 * 输入: 9973
 * 输出: 9973
 * 解释: 不需要交换。
 * 
 * 
 * 注意:
 * 
 * 
 * 给定数字的范围是 [0, 10^8]
 * 
 * 
 */
public class Solution
{
    public int MaximumSwap(int num)
    {
        var s = num.ToString().ToCharArray();
        var (i, n) = (1, s.Length);
        for (; i < n && s[i] <= s[i - 1]; i++) { }
        if (i < n)
        {
            for (var j = i; j < n; j++)
            {
                if (s[j] >= s[i]) { i = j; }
            }
            for (var j = 0; true; j++)
            {
                if (s[j] < s[i])
                {
                    (s[i], s[j]) = (s[j], s[i]);
                    break;
                }
            }
        }
        return int.Parse(new string(s));
    }
}
