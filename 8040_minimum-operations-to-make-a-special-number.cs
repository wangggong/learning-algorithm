/*
 * @lc app=leetcode.cn id=minimum-operations-to-make-a-special-number lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [8040] 生成特殊数字的最少操作
 *
 * https://leetcode.cn/problems/minimum-operations-to-make-a-special-number/description/
 *
 * algorithms
 * Medium (36.76%)
 * Total Accepted:    3.2K
 * Total Submissions: 8.6K
 * Testcase Example:  '"2245047"'
 *
 * 给你一个下标从 0 开始的字符串 num ，表示一个非负整数。
 * 
 * 在一次操作中，您可以选择 num 的任意一位数字并将其删除。请注意，如果你删除 num 中的所有数字，则 num 变为 0。
 * 
 * 返回最少需要多少次操作可以使 num 变成特殊数字。
 * 
 * 如果整数 x 能被 25 整除，则该整数 x 被认为是特殊数字。
 * 
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：num = "2245047"
 * 输出：2
 * 解释：删除数字 num[5] 和 num[6] ，得到数字 "22450" ，可以被 25 整除。
 * 可以证明要使数字变成特殊数字，最少需要删除 2 位数字。
 * 
 * 示例 2：
 * 
 * 
 * 输入：num = "2908305"
 * 输出：3
 * 解释：删除 num[3]、num[4] 和 num[6] ，得到数字 "2900" ，可以被 25 整除。
 * 可以证明要使数字变成特殊数字，最少需要删除 3 位数字。
 * 
 * 示例 3：
 * 
 * 
 * 输入：num = "10"
 * 输出：1
 * 解释：删除 num[0] ，得到数字 "0" ，可以被 25 整除。
 * 可以证明要使数字变成特殊数字，最少需要删除 1 位数字。
 * 
 * 
 * 
 * 
 * 提示
 * 
 * 
 * 1 <= num.length <= 100
 * num 仅由数字 '0' 到 '9' 组成
 * num 不含任何前导零
 * 
 * 
 */
public class Solution
{
    public int MinimumOperations(string num)
    {
        var n = num.Length;
        var ans = n - (num.Any(x => x is '0') ? 1 : 0);
        void find(int k, Predicate<int> p)
        {
            for (var i = k; i >= 0; i--)
            {
                if (p(num[i])) { ans = Math.Min(ans, n - i - 2); }
            }
        }
        for (var i = n - 1; i >= 0; i--)
        {
            if (num[i] is '0') { find(i - 1, x => x is '0' or '5'); }
            if (num[i] is '5') { find(i - 1, x => x is '2' or '7'); }
        }
        return ans;
    }
}
