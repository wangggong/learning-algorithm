/*
 * @lc app=leetcode.cn id=add-strings lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [415] 字符串相加
 *
 * https://leetcode.cn/problems/add-strings/description/
 *
 * algorithms
 * Easy (54.59%)
 * Total Accepted:    270.2K
 * Total Submissions: 494.9K
 * Testcase Example:  '"11"\n"123"'
 *
 * 给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和并同样以字符串形式返回。
 * 
 * 你不能使用任何內建的用于处理大整数的库（比如 BigInteger）， 也不能直接将输入的字符串转换为整数形式。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：num1 = "11", num2 = "123"
 * 输出："134"
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：num1 = "456", num2 = "77"
 * 输出："533"
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：num1 = "0", num2 = "0"
 * 输出："0"
 * 
 * 
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= num1.length, num2.length <= 10^4
 * num1 和num2 都只包含数字 0-9
 * num1 和num2 都不包含任何前导零
 * 
 * 
 */
public class Solution
{
    public string AddStrings(string num1, string num2) =>
        new string(AddStrings(num1.ToCharArray(), num2.ToCharArray()));

    private char[] AddStrings(char[] num1, char[] num2)
    {
        var ans = new List<char>();
        num1 = num1.Reverse().ToArray();
        num2 = num2.Reverse().ToArray();
        var C = 0;
        for (var (i, m, n) = (0, num1.Length, num2.Length); i < m || i < n; i++)
        {
            var p = i < m ? (int)(num1[i] - '0') : 0;
            var q = i < n ? (int)(num2[i] - '0') : 0;
            ans.Add((char)((p + q + C) % 10 + '0'));
            C = (p + q + C) / 10;
        }
        if (C is not 0) { ans.Add((char)(C + '0')); }
        ans.Reverse();
        return ans.ToArray();
    }
}
