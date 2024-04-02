/*
 * @lc app=leetcode.cn id=circular-permutation-in-binary-representation lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1238] 循环码排列
 *
 * https://leetcode.cn/problems/circular-permutation-in-binary-representation/description/
 *
 * algorithms
 * Medium (67.66%)
 * Total Accepted:    4.8K
 * Total Submissions: 7.1K
 * Testcase Example:  '2\n3'
 *
 * 给你两个整数 n 和 start。你的任务是返回任意 (0,1,2,,...,2^n-1) 的排列 p，并且满足：
 * 
 * 
 * p[0] = start
 * p[i] 和 p[i+1] 的二进制表示形式只有一位不同
 * p[0] 和 p[2^n -1] 的二进制表示形式也只有一位不同
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：n = 2, start = 3
 * 输出：[3,2,0,1]
 * 解释：这个排列的二进制表示是 (11,10,00,01)
 * ⁠    所有的相邻元素都有一位是不同的，另一个有效的排列是 [3,1,0,2]
 * 
 * 
 * 示例 2：
 * 
 * 输出：n = 3, start = 2
 * 输出：[2,6,7,5,4,0,1,3]
 * 解释：这个排列的二进制表示是 (010,110,111,101,100,000,001,011)
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 16
 * 0 <= start < 2^n
 * 
 * 
 */
public class Solution
{
    private List<int> GrayCode(int n)
    {
        if (n == 0)
        {
            return new List<int> { 0 };
        }
        var last = GrayCode(n - 1);
        var ans = new List<int>();
        foreach (var v in last)
        {
            ans.Add(v);
        }
        for (int i = last.Count() - 1; i >= 0; i--)
        {
            ans.Add((1 << (n - 1)) | last[i]);
        }
        return ans;
    }

    public IList<int> CircularPermutation(int n, int start)
    {
        var original = GrayCode(n);
        var count = original.Count();
        var k = original.IndexOf(start);
        return Enumerable.Range(0, count)
            .Select(i => original[(i + k) % count])
            .ToList();
    }

    /*
     * public IList<int> CircularPermutation(int n, int start)
     * {
     *     var count = 1 << n;
     *     var visit = new bool[count];
     *     var ans = new List<int>();
     *     while (count-- > 0)
     *     {
     *         ans.Add(start);
     *         visit[start] = true;
     *         for (var i = 0; i < n; i++)
     *         {
     *             var next = (1 << i) ^ start;
     *             if (!visit[next])
     *             {
     *                 start = next;
     *                 break;
     *             }
     *         }
     *     }
     *     return ans;
     * }
     */

    /*
     * public IList<int> CircularPermutation(int n, int start)
     * {
     *     var visit = new bool[1 << n];
     *     return Enumerable.Range(0, 1 << n)
     *         .Select(_ =>
     *         {
     *             var prev = start;
     *             visit[prev] = true;
     *             for (var i = 0; i < n; i++)
     *             {
     *                 var next = start ^ (1 << i);
     *                 if (!visit[next])
     *                 {
     *                     start = next;
     *                     break;
     *                 }
     *             }
     *             return prev;
     *         }).ToList();
     * }
     */
}
