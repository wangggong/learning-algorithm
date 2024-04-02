/*
 * @lc app=leetcode.cn id=count-substrings-that-differ-by-one-character lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1638] 统计只差一个字符的子串数目
 *
 * https://leetcode.cn/problems/count-substrings-that-differ-by-one-character/description/
 *
 * algorithms
 * Medium (72.67%)
 * Total Accepted:    7.1K
 * Total Submissions: 9.4K
 * Testcase Example:  '"aba"\n"baba"'
 *
 * 给你两个字符串 s 和 t ，请你找出 s 中的非空子串的数目，这些子串满足替换 一个不同字符 以后，是 t 串的子串。换言之，请你找到 s 和 t
 * 串中 恰好 只有一个字符不同的子字符串对的数目。
 * 
 * 比方说， "computer" and "computation" 只有一个字符不同： 'e'/'a' ，所以这一对子字符串会给答案加 1 。
 * 
 * 请你返回满足上述条件的不同子字符串对数目。
 * 
 * 一个 子字符串 是一个字符串中连续的字符。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s = "aba", t = "baba"
 * 输出：6
 * 解释：以下为只相差 1 个字符的 s 和 t 串的子字符串对：
 * ("aba", "baba")
 * ("aba", "baba")
 * ("aba", "baba")
 * ("aba", "baba")
 * ("aba", "baba")
 * ("aba", "baba")
 * 加粗部分分别表示 s 和 t 串选出来的子字符串。
 * 
 * 示例 2：
 * 
 * 
 * 输入：s = "ab", t = "bb"
 * 输出：3
 * 解释：以下为只相差 1 个字符的 s 和 t 串的子字符串对：
 * ("ab", "bb")
 * ("ab", "bb")
 * ("ab", "bb")
 * 加粗部分分别表示 s 和 t 串选出来的子字符串。
 * 
 * 示例 3：
 * 
 * 
 * 输入：s = "a", t = "a"
 * 输出：0
 * 
 * 
 * 示例 4：
 * 
 * 
 * 输入：s = "abe", t = "bbc"
 * 输出：10
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= s.length, t.length <= 100
 * s 和 t 都只包含小写英文字母。
 * 
 * 
 */

/*
 * // 枚举, 大力出奇迹
 * public class Solution
 * {
 *     public int CountSubstrings(string s, string t)
 *     {
 *         var count = new Dictionary<string, int>();
 *         var m = t.Length;
 *         for (var i = 0; i < m; i++)
 *         {
 *             for (var j = i + 1; j <= m; j++)
 *             {
 *                 var cur = t.Substring(i, j - i);
 *                 count.TryGetValue(cur, out var c);
 *                 count[cur] = c + 1;
 *             }
 *         }
 *         var n = s.Length;
 *         var ans = 0;
 *         for (var i = 0; i < n; i++)
 *         {
 *             for (var j = i + 1; j <= n; j++)
 *             {
 *                 var cur = s.Substring(i, j - i);
 *                 var charArray = cur.ToCharArray();
 *                 for (var k = 0; k < j - i; k++)
 *                 {
 *                     var curChar = charArray[k];
 *                     for (var ch = 'a'; ch <= 'z'; ch++)
 *                     {
 *                         if (ch != curChar)
 *                         {
 *                             charArray[k] = ch;
 *                             count.TryGetValue(new string(charArray), out var c);
 *                             ans += c;
 *                         }
 *                     }
 *                     charArray[k] = curChar;
 *                 }
 *             }
 *         }
 *         return ans;
 *     }
 * }
 */

// 参考: https://leetcode.cn/problems/count-substrings-that-differ-by-one-character/solution/tu-jie-fei-bao-li-onm-suan-fa-pythonjava-k5og/
public class Solution
{
    public int CountSubstrings(string s, string t)
    {
        var n = s.Length;
        var m = t.Length;
        var ans = 0;
        for (var d = 1 - m; d < n; d++)
        {
            var i = Math.Max(d, 0);
            for (var (k0, k1) = (i - 1, i - 1); i < n && i - d < m; i++)
            {
                if (s[i] != t[i - d])
                {
                    k0 = k1;
                    k1 = i;
                }
                ans += k1 - k0;
            }
        }
        return ans;
    }
}
