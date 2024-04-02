/*
 * @lc app=leetcode.cn id=append-characters-to-string-to-make-subsequence lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6246] 追加字符以获得子序列
 *
 * https://leetcode.cn/problems/append-characters-to-string-to-make-subsequence/description/
 *
 * algorithms
 * Medium (62.53%)
 * Total Accepted:    4K
 * Total Submissions: 6.3K
 * Testcase Example:  '"coaching"\n"coding"'
 *
 * 给你两个仅由小写英文字母组成的字符串 s 和 t 。
 * 
 * 现在需要通过向 s 末尾追加字符的方式使 t 变成 s 的一个 子序列 ，返回需要追加的最少字符数。
 * 
 * 子序列是一个可以由其他字符串删除部分（或不删除）字符但不改变剩下字符顺序得到的字符串。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s = "coaching", t = "coding"
 * 输出：4
 * 解释：向 s 末尾追加字符串 "ding" ，s = "coachingding" 。
 * 现在，t 是 s ("coachingding") 的一个子序列。
 * 可以证明向 s 末尾追加任何 3 个字符都无法使 t 成为 s 的一个子序列。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：s = "abcde", t = "a"
 * 输出：0
 * 解释：t 已经是 s ("abcde") 的一个子序列。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：s = "z", t = "abcde"
 * 输出：5
 * 解释：向 s 末尾追加字符串 "abcde" ，s = "zabcde" 。
 * 现在，t 是 s ("zabcde") 的一个子序列。 
 * 可以证明向 s 末尾追加任何 4 个字符都无法使 t 成为 s 的一个子序列。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= s.length, t.length <= 10^5
 * s 和 t 仅由小写英文字母组成
 * 
 * 
 */
public class Solution
{
    public int AppendCharacters(string s, string t)
    {
        int n = t.Count();
        int m = s.Count();
        int k = 0;
        for (int j = 0; k < n && j < m; j++)
        {
            if (s[j] == t[k]) { k++; }
        }
        return n - k;
    }
}
