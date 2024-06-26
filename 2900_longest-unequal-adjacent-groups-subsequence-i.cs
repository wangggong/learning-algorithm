/*
 * Medium
 * 
 * 给你一个整数 n 和一个下标从 0 开始的字符串数组 words ，和一个下标从 0 开始的 二进制 数组 groups ，两个数组长度都是 n 。
 * 
 * 你需要从下标 [0, 1, ..., n - 1] 中选出一个 最长子序列 ，将这个子序列记作长度为 k 的 [i0, i1, ..., ik - 1] ，对于所有满足 0 < j + 1 < k 的 j 都有 groups[ij] != groups[ij + 1] 。
 * 
 * 请你返回一个字符串数组，它是下标子序列 依次 对应 words 数组中的字符串连接形成的字符串数组。如果有多个答案，返回任意一个。
 * 
 * 子序列 指的是从原数组中删掉一些（也可能一个也不删掉）元素，剩余元素不改变相对位置得到的新的数组。
 * 
 * 注意：words 中的字符串长度可能 不相等 。
 * 
 *  
 * 
 * 示例 1：
 * 
 * 输入：n = 3, words = ["e","a","b"], groups = [0,0,1]
 * 输出：["e","b"]
 * 解释：一个可行的子序列是 [0,2] ，因为 groups[0] != groups[2] 。
 * 所以一个可行的答案是 [words[0],words[2]] = ["e","b"] 。
 * 另一个可行的子序列是 [1,2] ，因为 groups[1] != groups[2] 。
 * 得到答案为 [words[1],words[2]] = ["a","b"] 。
 * 这也是一个可行的答案。
 * 符合题意的最长子序列的长度为 2 。
 * 示例 2：
 * 
 * 输入：n = 4, words = ["a","b","c","d"], groups = [1,0,1,1]
 * 输出：["a","b","c"]
 * 解释：一个可行的子序列为 [0,1,2] 因为 groups[0] != groups[1] 且 groups[1] != groups[2] 。
 * 所以一个可行的答案是 [words[0],words[1],words[2]] = ["a","b","c"] 。
 * 另一个可行的子序列为 [0,1,3] 因为 groups[0] != groups[1] 且 groups[1] != groups[3] 。
 * 得到答案为 [words[0],words[1],words[3]] = ["a","b","d"] 。
 * 这也是一个可行的答案。
 * 符合题意的最长子序列的长度为 3 。
 *  
 * 
 * 提示：
 * 
 * 1 <= n == words.length == groups.length <= 100
 * 1 <= words[i].length <= 10
 * 0 <= groups[i] < 2
 * words 中的字符串 互不相同 。
 * words[i] 只包含小写英文字母。
 *
 * 通过次数 2.4K
 * 提交次数 3.3K
 * 通过率 72.9%
 * 
 */
public class Solution
{
    public IList<string> GetWordsInLongestSubsequence(int n, string[] words, int[] groups)
    {
        var ans = new List<string>();
        var last = -1;
        foreach (var (w, g) in words.Zip(groups))
        {
            if (last != g)
            {
                ans.Add(w);
                last = g;
            }
        }
        return ans;
    }
}
