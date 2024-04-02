/*
 * @lc app=leetcode.cn id=minimum-time-to-revert-word-to-initial-state-ii lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [100203] 将单词恢复初始状态所需的最短时间 II
 *
 * https://leetcode.cn/problems/minimum-time-to-revert-word-to-initial-state-ii/description/
 *
 * algorithms
 * Hard (32.18%)
 * Total Accepted:    1.5K
 * Total Submissions: 4.2K
 * Testcase Example:  '"abacaba"\n3'
 *
 * 给你一个下标从 0 开始的字符串 word 和一个整数 k 。
 * 
 * 在每一秒，你必须执行以下操作：
 * 
 * 
 * 移除 word 的前 k 个字符。
 * 在 word 的末尾添加 k 个任意字符。
 * 
 * 
 * 注意 添加的字符不必和移除的字符相同。但是，必须在每一秒钟都执行 两种 操作。
 * 
 * 返回将 word 恢复到其 初始 状态所需的 最短 时间（该时间必须大于零）。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：word = "abacaba", k = 3
 * 输出：2
 * 解释：
 * 第 1 秒，移除 word 的前缀 "aba"，并在末尾添加 "bac" 。因此，word 变为 "cababac"。
 * 第 2 秒，移除 word 的前缀 "cab"，并在末尾添加 "aba" 。因此，word 变为 "abacaba" 并恢复到始状态。
 * 可以证明，2 秒是 word 恢复到其初始状态所需的最短时间。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：word = "abacaba", k = 4
 * 输出：1
 * 解释：
 * 第 1 秒，移除 word 的前缀 "abac"，并在末尾添加 "caba" 。因此，word 变为 "abacaba" 并恢复到初始状态。
 * 可以证明，1 秒是 word 恢复到其初始状态所需的最短时间。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：word = "abcbabcd", k = 2
 * 输出：4
 * 解释：
 * 每一秒，我们都移除 word 的前 2 个字符，并在 word 末尾添加相同的字符。
 * 4 秒后，word 变为 "abcbabcd" 并恢复到初始状态。
 * 可以证明，4 秒是 word 恢复到其初始状态所需的最短时间。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= word.length <= 10^5 
 * 1 <= k <= word.length
 * word仅由小写英文字母组成。
 * 
 * 
 */
public class Solution
{
    public int MinimumTimeToInitialState(string word, int k)
    {
        const long P = 13331;
        var n = word.Length;
        var hashs = new long[n + 1];
        var multis = new long[n + 1];
        multis[0] = 1;
        for (var i = 0; i < n; i++)
        {
            hashs[i + 1] = hashs[i] * P + (long)(word[i] - 'a');
            multis[i + 1] = multis[i] * P;
        }
        return Enumerable.Range(1, n + 1)
            .First(i => n <= i * k
                || hashs[n - i * k] == hashs[n] - hashs[i * k] * multis[n - i * k]);
    }
}
