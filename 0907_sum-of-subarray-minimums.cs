/*
 * @lc app=leetcode.cn id=sum-of-subarray-minimums lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [907] 子数组的最小值之和
 *
 * https://leetcode.cn/problems/sum-of-subarray-minimums/description/
 *
 * algorithms
 * Medium (38.50%)
 * Total Accepted:    49K
 * Total Submissions: 127.3K
 * Testcase Example:  '[3,1,2,4]'
 *
 * 给定一个整数数组 arr，找到 min(b) 的总和，其中 b 的范围为 arr 的每个（连续）子数组。
 * 
 * 由于答案可能很大，因此 返回答案模 10^9 + 7 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：arr = [3,1,2,4]
 * 输出：17
 * 解释：
 * 子数组为 [3]，[1]，[2]，[4]，[3,1]，[1,2]，[2,4]，[3,1,2]，[1,2,4]，[3,1,2,4]。 
 * 最小值为 3，1，2，4，1，1，2，1，1，1，和为 17。
 * 
 * 示例 2：
 * 
 * 
 * 输入：arr = [11,81,94,43,3]
 * 输出：444
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * 1 
 * 
 * 
 * 
 * 
 */
public class Solution
{
    public int SumSubarrayMins(int[] arr)
    {
        const long Mod = (long)1e9 + 7;
        var n = arr.Length;
        var lefts = new int[n];
        var rights = new int[n];
        var S = new Stack<int>();
        S.Push(-1);
        for (var i = 0; i < n; i++)
        {
            for (; S.Peek() is not -1 && arr[S.Peek()] >= arr[i]; S.Pop()) { }
            lefts[i] = S.Peek();
            S.Push(i);
        }
        S = new Stack<int>();
        S.Push(n);
        for (var i = n - 1; i >= 0; i--)
        {
            for (; S.Peek() != n && arr[S.Peek()] > arr[i]; S.Pop()) { }
            rights[i] = S.Peek();
            S.Push(i);
        }
        var ans = 0l;
        for (var i = 0; i < n; i++)
        {
            ans = (ans + (long)arr[i] * (long)(i - lefts[i]) % Mod * (long)(rights[i] - i) % Mod) % Mod;
        }
        return (int)ans;
    }
}
