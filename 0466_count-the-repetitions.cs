/*
 * @lc app=leetcode.cn id=count-the-repetitions lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [466] 统计重复个数
 *
 * https://leetcode.cn/problems/count-the-repetitions/description/
 *
 * algorithms
 * Hard (37.77%)
 * Total Accepted:    17.2K
 * Total Submissions: 43.5K
 * Testcase Example:  '"acb"\n4\n"ab"\n2'
 *
 * 定义 str = [s, n] 表示 str 由 n 个字符串 s 连接构成。
 * 
 * 
 * 例如，str == ["abc", 3] =="abcabcabc" 。
 * 
 * 
 * 如果可以从 s2 中删除某些字符使其变为 s1，则称字符串 s1 可以从字符串 s2 获得。
 * 
 * 
 * 例如，根据定义，s1 = "abc" 可以从 s2 = "abdbec" 获得，仅需要删除加粗且用斜体标识的字符。
 * 
 * 
 * 现在给你两个字符串 s1 和 s2 和两个整数 n1 和 n2 。由此构造得到两个字符串，其中 str1 = [s1, n1]、str2 = [s2,
 * n2] 。
 * 
 * 请你找出一个最大整数 m ，以满足 str = [str2, m] 可以从 str1 获得。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s1 = "acb", n1 = 4, s2 = "ab", n2 = 2
 * 输出：2
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：s1 = "acb", n1 = 1, s2 = "acb", n2 = 1
 * 输出：1
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * s1 和 s2 由小写英文字母组成
 * 1 
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/count-the-repetitions/solutions/2069646/by-sln178716-rufd/?envType=daily-question&envId=2024-01-02
//
// 本质上是找循环节, 但这个方案找起来相对简洁, 代价就是效率会差很多.
public class Solution
{
    public int GetMaxRepetitions(string s1, int n1, string s2, int n2)
    {
        var (l1, l2, f1, f2) = (s1.Length, s2.Length, 0, 0);
        for (var l = l1 * n1; f1 < l; )
        {
            if (s1[f1 % l1] == s2[f2 % l2]) { f2++; }
            f1++;
            if ((f1 % l1, f2 % l2) is (0, 0))
            {
                var t = l / f1;
                (f1, f2) = (f1 * t, f2 * t);
            }
        }
        return f2 / l2 / n2;
    }
}
