/*
 * @lc app=leetcode.cn id=sum-of-number-and-its-reverse lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6219] 反转之后的数字和
 *
 * https://leetcode.cn/problems/sum-of-number-and-its-reverse/description/
 *
 * algorithms
 * Medium (45.34%)
 * Total Accepted:    5.7K
 * Total Submissions: 12.5K
 * Testcase Example:  '443'
 *
 * 给你一个 非负 整数 num 。如果存在某个 非负 整数 k 满足 k + reverse(k) = num  ，则返回 true ；否则，返回
 * false 。
 * 
 * reverse(k) 表示 k 反转每个数位后得到的数字。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：num = 443
 * 输出：true
 * 解释：172 + 271 = 443 ，所以返回 true 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：num = 63
 * 输出：false
 * 解释：63 不能表示为非负整数及其反转后数字之和，返回 false 。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：num = 181
 * 输出：true
 * 解释：140 + 041 = 181 ，所以返回 true 。注意，反转后的数字可能包含前导零。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 0 <= num <= 10^5
 * 
 * 
 */
public class Solution
{
    public bool SumOfNumberAndReverse(int num)
    {
        for (int i = 0; i <= num; i++)
        {
            if (GetSum(i) == num)
            {
                return true;
            }
        }
        return false;
    }

    private int GetSum(int n)
    {
        int ans = 0;
        for (int k = n; k != 0; k /= 10)
        {
            ans = ans * 10 + k % 10;
        }
        return ans + n;
    }
}
