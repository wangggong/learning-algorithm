/*
 * @lc app=leetcode.cn id=check-if-word-is-valid-after-substitutions lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1003] 检查替换后的词是否有效
 *
 * https://leetcode.cn/problems/check-if-word-is-valid-after-substitutions/description/
 *
 * algorithms
 * Medium (58.36%)
 * Total Accepted:    22.1K
 * Total Submissions: 36.2K
 * Testcase Example:  '"aabcbc"'
 *
 * 给你一个字符串 s ，请你判断它是否 有效 。
 * 字符串 s 有效 需要满足：假设开始有一个空字符串 t = "" ，你可以执行 任意次 下述操作将 t 转换为 s ：
 * 
 * 
 * 将字符串 "abc" 插入到 t 中的任意位置。形式上，t 变为 tleft + "abc" + tright，其中 t == tleft +
 * tright 。注意，tleft 和 tright 可能为 空 。
 * 
 * 
 * 如果字符串 s 有效，则返回 true；否则，返回 false。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s = "aabcbc"
 * 输出：true
 * 解释：
 * "" -> "abc" -> "aabcbc"
 * 因此，"aabcbc" 有效。
 * 
 * 示例 2：
 * 
 * 
 * 输入：s = "abcabcababcc"
 * 输出：true
 * 解释：
 * "" -> "abc" -> "abcabc" -> "abcabcabc" -> "abcabcababcc"
 * 因此，"abcabcababcc" 有效。
 * 
 * 示例 3：
 * 
 * 
 * 输入：s = "abccba"
 * 输出：false
 * 解释：执行操作无法得到 "abccba" 。
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= s.length <= 2 * 10^4
 * s 由字母 'a'、'b' 和 'c' 组成
 * 
 * 
 */

/*
 * public class Solution
 * {
 *     public bool IsValid(string s)
 *     {
 *         var S = new Stack<string>();
 *         foreach (var ch in s)
 *         {
 *             switch (ch)
 *             {
 *                 case 'a':
 *                     S.Push("a");
 *                     break;
 *                 case 'b':
 *                     if (S.Count == 0 || S.Peek() != "a")
 *                     {
 *                         return false;
 *                     }
 *                     S.Pop();
 *                     S.Push("ab");
 *                     break;
 *                 case 'c':
 *                     if (S.Count == 0 || S.Peek() != "ab")
 *                     {
 *                         return false;
 *                     }
 *                     S.Pop();
 *                     break;
 *             }
 *         }
 *         return S.Count == 0;
 *     }
 * }
 * public class Solution
 * {
 *     public bool IsValid(string s)
 *     {
 *         var S = new Stack<string>();
 *         foreach (var ch in s)
 *         {
 *             switch (ch)
 *             {
 *                 case 'a':
 *                     S.Push("a");
 *                     break;
 *                 case 'b':
 *                     if (S.Count == 0 || S.Peek() != "a")
 *                     {
 *                         return false;
 *                     }
 *                     S.Pop();
 *                     S.Push("ab");
 *                     break;
 *                 case 'c':
 *                     if (S.Count == 0 || S.Peek() != "ab")
 *                     {
 *                         return false;
 *                     }
 *                     S.Pop();
 *                     break;
 *             }
 *         }
 *         return S.Count == 0;
 *     }
 * }
 */

// 其实可以直接用 char 的.
public class Solution
{
    public bool IsValid(string s)
    {
        var S = new Stack<char>();
        foreach (var ch in s)
        {
            switch (ch)
            {
                case 'a':
                    S.Push('a');
                    break;
                case 'b':
                    if (S.Count == 0 || S.Peek() != 'a')
                    {
                        return false;
                    }
                    S.Pop();
                    S.Push('b');
                    break;
                case 'c':
                    if (S.Count == 0 || S.Peek() != 'b')
                    {
                        return false;
                    }
                    S.Pop();
                    break;
            }
        }
        return S.Count == 0;
    }
}
