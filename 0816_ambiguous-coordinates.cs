/*
 * @lc app=leetcode.cn id=ambiguous-coordinates lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [816] 模糊坐标
 *
 * https://leetcode.cn/problems/ambiguous-coordinates/description/
 *
 * algorithms
 * Medium (50.79%)
 * Total Accepted:    19.5K
 * Total Submissions: 31.7K
 * Testcase Example:  '"(123)"'
 *
 * 我们有一些二维坐标，如 "(1, 3)" 或 "(2,
 * 0.5)"，然后我们移除所有逗号，小数点和空格，得到一个字符串S。返回所有可能的原始字符串到一个列表中。
 * 
 * 原始的坐标表示法不会存在多余的零，所以不会出现类似于"00", "0.0", "0.00", "1.0", "001",
 * "00.01"或一些其他更小的数来表示坐标。此外，一个小数点前至少存在一个数，所以也不会出现“.1”形式的数字。
 * 
 * 最后返回的列表可以是任意顺序的。而且注意返回的两个数字中间（逗号之后）都有一个空格。
 * 
 * 
 * 
 * 
 * 示例 1:
 * 输入: "(123)"
 * 输出: ["(1, 23)", "(12, 3)", "(1.2, 3)", "(1, 2.3)"]
 * 
 * 
 * 
 * 示例 2:
 * 输入: "(00011)"
 * 输出:  ["(0.001, 1)", "(0, 0.011)"]
 * 解释: 
 * 0.0, 00, 0001 或 00.01 是不被允许的。
 * 
 * 
 * 
 * 示例 3:
 * 输入: "(0123)"
 * 输出: ["(0, 123)", "(0, 12.3)", "(0, 1.23)", "(0.1, 23)", "(0.1, 2.3)",
 * "(0.12, 3)"]
 * 
 * 
 * 
 * 示例 4:
 * 输入: "(100)"
 * 输出: [(10, 0)]
 * 解释: 
 * 1.0 是不被允许的。
 * 
 * 
 * 
 * 
 * 提示: 
 * 
 * 
 * 4 <= S.length <= 12.
 * S[0] = "(", S[S.length - 1] = ")", 且字符串 S 中的其他元素都是数字。
 * 
 * 
 * 
 * 
 */
public class Solution
{
    private IList<string> ans = new List<string>();

    public IList<string> AmbiguousCoordinates(string s)
    {
        int n = s.Count();
        s = s.Substring(1, n - 2);
        n -= 2;
        for (int i = 1; i < n; i++)
        {
            string x = s.Substring(0, i);
            string y = s.Substring(i);
            IList<string> xs = GenValid(x);
            IList<string> ys = GenValid(y);
            foreach (var cx in xs)
            {
                foreach (var cy in ys)
                {
                    ans.Add($"({cx}, {cy})");
                }
            }
        }
        return ans;
    }

    private IList<string> GenValid(string s)
    {
        IList<string> floats = new List<string>();
        if (Valid(s)) { floats.Add(s); }
        int n = s.Count();
        for (int i = 1; i < n; i++)
        {
            string ns = s.Substring(0, i) + "." + s.Substring(i);
            if (Valid(ns)) { floats.Add(ns); }
        }
        return floats;
    }

    private bool Valid(string s)
    {
        int n = s.Count();
        int idx = s.IndexOf(".");
        if (idx < 0) { return ValidInt(s); }
        if (!ValidInt(s.Substring(0, idx))) { return false; }
        return s[n - 1] != '.' && s[n - 1] != '0';
    }

    private bool ValidInt(string s)
        => s.Count() == 1 || s.Count() > 1 && s[0] != '0';
}
