/*
 * @lc app=leetcode.cn id=shortest-common-supersequence lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1092] 最短公共超序列
 *
 * https://leetcode.cn/problems/shortest-common-supersequence/description/
 *
 * algorithms
 * Hard (54.33%)
 * Total Accepted:    6.4K
 * Total Submissions: 11.4K
 * Testcase Example:  '"abac"\n"cab"'
 *
 * 给出两个字符串 str1 和 str2，返回同时以 str1 和 str2
 * 作为子序列的最短字符串。如果答案不止一个，则可以返回满足条件的任意一个答案。
 * 
 * （如果从字符串 T 中删除一些字符（也可能不删除，并且选出的这些字符可以位于 T 中的 任意位置），可以得到字符串 S，那么 S 就是 T
 * 的子序列）
 * 
 * 
 * 
 * 示例：
 * 
 * 输入：str1 = "abac", str2 = "cab"
 * 输出："cabac"
 * 解释：
 * str1 = "abac" 是 "cabac" 的一个子串，因为我们可以删去 "cabac" 的第一个 "c"得到 "abac"。 
 * str2 = "cab" 是 "cabac" 的一个子串，因为我们可以删去 "cabac" 末尾的 "ac" 得到 "cab"。
 * 最终我们给出的答案是满足上述属性的最短字符串。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= str1.length, str2.length <= 1000
 * str1 和 str2 都由小写英文字母组成。
 * 
 * 
 */

/*
 * public class Solution
 * {
 *     public string ShortestCommonSupersequence(string str1, string str2)
 *     {
 *         var n = str1.Length;
 *         var m = str2.Length;
 *         var dp = new int[n + 1][];
 *         for (var i = 0; i <= n; i++)
 *         {
 *             dp[i] = new int[m + 1];
 *         }
 *         for (var i = 0; i < n; i++)
 *         {
 *             for (var j = 0; j < m; j++)
 *             {
 *                 dp[i + 1][j + 1] = str1[i] == str2[j]
 *                     ? dp[i][j] + 1
 *                     : Math.Max(dp[i + 1][j], dp[i][j + 1]);
 *             }
 *         }
 *         var sb = new StringBuilder();
 *         for (int i = n, j = m; i > 0 && j > 0; )
 *         {
 *             if (dp[i][j] == dp[i - 1][j - 1] + 1 && str1[i - 1] == str2[j - 1])
 *             {
 *                 sb.Append(str1[i - 1]);
 *                 i--;
 *                 j--;
 *             }
 *             else if (dp[i - 1][j] == dp[i][j])
 *             {
 *                 i--;
 *             }
 *             else
 *             {
 *                 j--;
 *             }
 *         }
 *         var charArray = sb.ToString().ToCharArray();
 *         for (int i = 0, j = charArray.Length - 1; i < j; i++, j--)
 *         {
 *             (charArray[i], charArray[j]) = (charArray[j], charArray[i]);
 *         }
 *         sb = new StringBuilder();
 *         var p = 0;
 *         var q = 0;
 *         foreach (var c in charArray)
 *         {
 *             for (; str1[p] != c; p++)
 *             {
 *                 sb.Append(str1[p]);
 *             }
 *             for (; str2[q] != c; q++)
 *             {
 *                 sb.Append(str2[q]);
 *             }
 *             sb.Append(c);
 *             p++;
 *             q++;
 *         }
 *         for (; p < n; p++)
 *         {
 *             sb.Append(str1[p]);
 *         }
 *         for (; q < m; q++)
 *         {
 *             sb.Append(str2[q]);
 *         }
 *         return sb.ToString();
 *     }
 * }
 */

// 参考: https://leetcode.cn/problems/shortest-common-supersequence/solution/cong-di-gui-dao-di-tui-jiao-ni-yi-bu-bu-auy8z/
public class Solution
{
    public string ShortestCommonSupersequence(string s, string t)
    {
        var n = s.Length;
        var m = t.Length;
        var dp = new int[n + 1][];
        for (var i = 0; i <= n; i++)
        {
            dp[i] = new int[m + 1];
        }
        for (var i = 0; i <= n; i++)
        {
            dp[i][0] = i;
        }
        for (var j = 0; j <= m; j++)
        {
            dp[0][j] = j;
        }
        for (var i = 0; i < n; i++)
        {
            for (var j = 0; j < m; j++)
            {
                dp[i + 1][j + 1] = (s[i] == t[j]
                    ? dp[i][j]
                    : Math.Min(dp[i + 1][j], dp[i][j + 1])) + 1;
            }
        }
        var sb = new StringBuilder();
        var p = n;
        var q = m;
        while (p > 0 && q > 0)
        {
            if (s[p - 1] == t[q - 1])
            {
                sb.Append(s[p - 1]);
                p--;
                q--;
            }
            else if (dp[p][q] == dp[p - 1][q] + 1)
            {
                sb.Append(s[p - 1]);
                p--;
            }
            else
            {
                sb.Append(t[q - 1]);
                q--;
            }
        }
        var charArray = sb.ToString().ToCharArray();
        for (int i = 0, j = charArray.Length - 1; i < j; i++, j--)
        {
            (charArray[i], charArray[j]) = (charArray[j], charArray[i]);
        }
        return s.Substring(0, p) + t.Substring(0, q) + new string(charArray);
    }
}
