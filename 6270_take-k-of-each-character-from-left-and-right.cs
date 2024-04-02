/*
 * @lc app=leetcode.cn id=take-k-of-each-character-from-left-and-right lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6270] 每种字符至少取 K 个
 *
 * https://leetcode.cn/problems/take-k-of-each-character-from-left-and-right/description/
 *
 * algorithms
 * Medium (22.11%)
 * Total Accepted:    1.3K
 * Total Submissions: 5.6K
 * Testcase Example:  '"aabaaaacaabc"\n2'
 *
 * 给你一个由字符 'a'、'b'、'c' 组成的字符串 s 和一个非负整数 k 。每分钟，你可以选择取走 s 最左侧 还是 最右侧 的那个字符。
 * 
 * 你必须取走每种字符 至少 k 个，返回需要的 最少 分钟数；如果无法取到，则返回 -1 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s = "aabaaaacaabc", k = 2
 * 输出：8
 * 解释：
 * 从 s 的左侧取三个字符，现在共取到两个字符 'a' 、一个字符 'b' 。
 * 从 s 的右侧取五个字符，现在共取到四个字符 'a' 、两个字符 'b' 和两个字符 'c' 。
 * 共需要 3 + 5 = 8 分钟。
 * 可以证明需要的最少分钟数是 8 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：s = "a", k = 1
 * 输出：-1
 * 解释：无法取到一个字符 'b' 或者 'c'，所以返回 -1 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= s.length <= 10^5
 * s 仅由字母 'a'、'b'、'c' 组成
 * 0 <= k <= s.length
 * 
 * 
 */
public class Solution
{
    public int TakeCharacters(string s, int k)
    {
        var n = s.Length;
        var cnt = new int[n + 1][];
        for (var i = 0; i <= n; i++)
        {
            cnt[i] = new int[3];
        }
        for (var i = 0; i < n; i++)
        {
            for (var j = 0; j < 3; j++)
            {
                cnt[i + 1][j] = cnt[i][j] + (s[i] - 'a' == j ? 1 : 0);
            }
        }
        for (var i = 0; i < 3; i++)
        {
            if (cnt[n][i] < k)
            {
                return -1;
            }
        }
        bool check(int start, int v)
        {
            for (var i = 0; i < 3; i++)
            {
                if (cnt[start][i] + (cnt[n][i] - cnt[v][i]) < k)
                {
                    return false;
                }
            }
            return true;
        }
        var ans = n + 1;
        for (var i = 0; i <= n; i++)
        {
            var p = 0;
            var q = n;
            while (p < q)
            {
                var mid = (p + q) >> 1;
                if (!check(i, mid))
                {
                    q = mid;
                }
                else
                {
                    p = mid + 1;
                }
            }
            p--;
            if (cnt[i][0] >= k && cnt[i][1] >= k && cnt[i][2] >= k)
            {
                ans = Math.Min(ans, i);
            }
            if (cnt[n][0] - cnt[i][0] >= k && cnt[n][1] - cnt[i][1] >= k && cnt[n][2] - cnt[i][2] >= k)
            {
                ans = Math.Min(ans, n - i);
            }
            ans = Math.Min(ans, i + (n - p));
        }
        return ans;
    }
}
