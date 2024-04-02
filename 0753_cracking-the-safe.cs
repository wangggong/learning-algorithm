/*
 * @lc app=leetcode.cn id=cracking-the-safe lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [753] 破解保险箱
 *
 * https://leetcode.cn/problems/cracking-the-safe/description/
 *
 * algorithms
 * Hard (61.63%)
 * Total Accepted:    11.2K
 * Total Submissions: 15.6K
 * Testcase Example:  '1\n2'
 *
 * 有一个需要密码才能打开的保险箱。密码是 n 位数, 密码的每一位是 k 位序列 0, 1, ..., k-1 中的一个 。
 * 
 * 你可以随意输入密码，保险箱会自动记住最后 n 位输入，如果匹配，则能够打开保险箱。
 * 
 * 举个例子，假设密码是 "345"，你可以输入 "012345" 来打开它，只是你输入了 6 个字符.
 * 
 * 请返回一个能打开保险箱的最短字符串。
 * 
 * 
 * 
 * 示例1:
 * 
 * 输入: n = 1, k = 2
 * 输出: "01"
 * 说明: "10"也可以打开保险箱。
 * 
 * 
 * 
 * 
 * 示例2:
 * 
 * 输入: n = 2, k = 2
 * 输出: "00110"
 * 说明: "01100", "10011", "11001" 也能打开保险箱。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * n 的范围是 [1, 4]。
 * k 的范围是 [1, 10]。
 * k^n 最大可能为 4096。
 * 
 * 
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/cracking-the-safe/solution/po-jie-bao-xian-xiang-by-leetcode-solution/
// Hierholzer - 欧拉图中找欧拉回路的算法.

/*
 * public class Solution
 * {
 *     public string CrackSafe(int n, int k)
 *     {
 *         var sb = new StringBuilder();
 *         var visit = new HashSet<int>();
 *         var maxDigit = (int)Math.Pow(10, n - 1);
 *         void dfs(int node)
 *         {
 *             for (var i = 0; i < k; i++)
 *             {
 *                 var edge = node * 10 + i;
 *                 if (!visit.Contains(edge))
 *                 {
 *                     visit.Add(edge);
 *                     dfs(edge % maxDigit);
 *                     sb.Append((char)('0' + i));
 *                 }
 *             }
 *         }
 *         dfs(0);
 *         sb.Append('0', n - 1);
 *         return sb.ToString();
 *     }
 * }
 */

public class Solution
{
    public string CrackSafe(int n, int k)
    {
        StringBuilder sb = new();
        HashSet<string> S = new();
        void dfs(string s)
        {
            for (var i = 0; i < k; i++)
            {
                var ns = s + new string(new char[] { (char)('0' + i), });
                if (!S.Contains(ns))
                {
                    S.Add(ns);
                    dfs(ns.Substring(1));
                    sb.Append((char)('0' + i));
                }
            }
        }
        var start = new string('0', n - 1);
        dfs(start);
        sb.Append(start);
        return sb.ToString();
    }
}
