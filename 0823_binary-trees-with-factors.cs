/*
 * @lc app=leetcode.cn id=binary-trees-with-factors lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [823] 带因子的二叉树
 *
 * https://leetcode.cn/problems/binary-trees-with-factors/description/
 *
 * algorithms
 * Medium (43.56%)
 * Total Accepted:    7.7K
 * Total Submissions: 16.8K
 * Testcase Example:  '[2,4]'
 *
 * 给出一个含有不重复整数元素的数组 arr ，每个整数 arr[i] 均大于 1。
 * 
 * 用这些整数来构建二叉树，每个整数可以使用任意次数。其中：每个非叶结点的值应等于它的两个子结点的值的乘积。
 * 
 * 满足条件的二叉树一共有多少个？答案可能很大，返回 对 10^9 + 7 取余 的结果。
 * 
 * 
 * 
 * 示例 1:
 * 
 * 
 * 输入: arr = [2, 4]
 * 输出: 3
 * 解释: 可以得到这些二叉树: [2], [4], [4, 2, 2]
 * 
 * 示例 2:
 * 
 * 
 * 输入: arr = [2, 4, 5, 10]
 * 输出: 7
 * 解释: 可以得到这些二叉树: [2], [4], [5], [10], [4, 2, 2], [10, 2, 5], [10, 5, 2].
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= arr.length <= 1000
 * 2 <= arr[i] <= 10^9
 * arr 中的所有值 互不相同
 * 
 * 
 */
public class Solution
{
    public int NumFactoredBinaryTrees(int[] arr)
    {
        const long Mod = (long)1e9 + 7;
        var n = arr.Length;
        Array.Sort(arr);
        var d = new Dictionary<int, int>();
        var dp = new long[n];
        Array.Fill(dp, 1);
        for (var i = 0; i < n; i++)
        {
            for (var j = 0; j < i; j++)
            {
                if ((arr[i] % arr[j]) is not 0) { continue; }
                var k = arr[i] / arr[j];
                if (k > arr[j] || !d.ContainsKey(k)) { continue; }
                var t = arr[j] == k ? 1 : 2;
                dp[i] = (dp[i] + dp[j] * dp[d[k]] % Mod * t) % Mod;
            }
            d[arr[i]] = i;
        }
        return (int)dp.Aggregate((x, y) => (x + y) % Mod);
    }
}
