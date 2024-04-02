/*
 * @lc app=leetcode.cn id=make-array-strictly-increasing lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1187] 使数组严格递增
 *
 * https://leetcode.cn/problems/make-array-strictly-increasing/description/
 *
 * algorithms
 * Hard (48.18%)
 * Total Accepted:    3.6K
 * Total Submissions: 7.1K
 * Testcase Example:  '[1,5,3,6,7]\n[1,3,2,4]'
 *
 * 给你两个整数数组 arr1 和 arr2，返回使 arr1 严格递增所需要的最小「操作」数（可能为 0）。
 * 
 * 每一步「操作」中，你可以分别从 arr1 和 arr2 中各选出一个索引，分别为 i 和 j，0 <= i < arr1.length 和 0 <= j
 * < arr2.length，然后进行赋值运算 arr1[i] = arr2[j]。
 * 
 * 如果无法让 arr1 严格递增，请返回 -1。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：arr1 = [1,5,3,6,7], arr2 = [1,3,2,4]
 * 输出：1
 * 解释：用 2 来替换 5，之后 arr1 = [1, 2, 3, 6, 7]。
 * 
 * 
 * 示例 2：
 * 
 * 输入：arr1 = [1,5,3,6,7], arr2 = [4,3,1]
 * 输出：2
 * 解释：用 3 来替换 5，然后用 4 来替换 3，得到 arr1 = [1, 3, 4, 6, 7]。
 * 
 * 
 * 示例 3：
 * 
 * 输入：arr1 = [1,5,3,6,7], arr2 = [1,6,3,3]
 * 输出：-1
 * 解释：无法使 arr1 严格递增。
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= arr1.length, arr2.length <= 2000
 * 0 <= arr1[i], arr2[i] <= 10^9
 * 
 * 
 * 
 * 
 */

/*
 * public class Solution
 * {
 *     public int MakeArrayIncreasing(int[] arr1, int[] arr2)
 *     {
 *         arr2 = arr2.ToHashSet().OrderBy(x => x).ToArray();
 *         var (n, m) = (arr1.Length, arr2.Length);
 *         var dp = new int[n][];
 *         for (var i = 0; i < n; i++)
 *         {
 *             dp[i] = new int[m + 1];
 *         }
 *         for (var i = 0; i <= m; i++)
 *         {
 *             dp[0][i] = (i == m || arr2[i] == arr1[0]) ? 0 : 1;
 *         }
 *         // int binarySearch(int k)
 *         // {
 *         //     var (p, q) = (0, m);
 *         //     while (p < q)
 *         //     {
 *         //         var mid = (p + q) >> 1;
 *         //         if (arr2[mid] > k)
 *         //         {
 *         //             q = mid;
 *         //         }
 *         //         else
 *         //         {
 *         //             p = mid + 1;
 *         //         }
 *         //     }
 *         //     return p;
 *         // }
 *         for (var i = 1; i < n; i++)
 *         {
 *             var k = 0;
 *             var min = n + 1;
 *             for (var j = 0; j < m; j++)
 *             {
 *                 for (; k < m && arr2[k] < arr2[j]; k++)
 *                 {
 *                     min = Math.Min(min, dp[i - 1][k]);
 *                 }
 *                 if (arr1[i - 1] < arr2[j])
 *                 {
 *                     min = Math.Min(min, dp[i - 1][m]);
 *                 }
 *                 dp[i][j] = min + (arr2[j] != arr1[i] ? 1 : 0);
 *             }
 *             for ((k, min) = (0, n + 1); k < m && arr2[k] < arr1[i]; k++)
 *             {
 *                 min = Math.Min(min, dp[i - 1][k]);
 *             }
 *             if (arr1[i - 1] < arr1[i])
 *             {
 *                 min = Math.Min(min, dp[i - 1][m]);
 *             }
 *             dp[i][m] = min;
 *         }
 *         var ans = dp[n - 1].Min();
 *         return ans <= n ? ans : -1;
 *     }
 * }
 */

// 参考: https://leetcode.cn/problems/make-array-strictly-increasing/solution/zui-chang-di-zeng-zi-xu-lie-de-bian-xing-jhgg/

/*
 * public class Solution
 * {
 *     private const int Maxn = (int)1e9;
 * 
 *     public int MakeArrayIncreasing(int[] arr1, int[] arr2)
 *     {
 *         Array.Sort(arr2);
 *         var memo = new Dictionary<(int, int), int>();
 *         int binarySearch(int k)
 *         {
 *             var (p, q) = (0, arr2.Length);
 *             while (p < q)
 *             {
 *                 var mid = (p + q) >> 1;
 *                 if (arr2[mid] >= k)
 *                 {
 *                     q = mid;
 *                 }
 *                 else
 *                 {
 *                     p = mid + 1;
 *                 }
 *             }
 *             return p - 1;
 *         }
 *         int dfs(int i, int pre)
 *         {
 *             var key = (i, pre);
 *             if (memo.ContainsKey(key))
 *             {
 *                 return memo[key];
 *             }
 *             if (i < 0)
 *             {
 *                 return (memo[key] = 0);
 *             }
 *             memo[key] = arr1[i] < pre ? dfs(i - 1, arr1[i]) : Maxn;
 *             var k = binarySearch(pre);
 *             if (k >= 0)
 *             {
 *                 memo[key] = Math.Min(memo[key], dfs(i - 1, arr2[k]) + 1);
 *             }
 *             return memo[key];
 *         }
 *         var ans = dfs(arr1.Length - 1, Maxn);
 *         return ans < Maxn ? ans : -1;
 *     }
 * }
 */

public class Solution
{
    private const int Maxn = (int)1e9;

    public int MakeArrayIncreasing(int[] arr1, int[] arr2)
    {
        var list = arr1.ToList();
        list.Add(Maxn);
        arr2 = arr2.ToHashSet().OrderBy(x => x).ToArray();
        int binarySearch(int k)
        {
            var (p, q) = (0, arr2.Length);
            while (p < q)
            {
                var mid = (p + q) >> 1;
                if (arr2[mid] >= k)
                {
                    q = mid;
                }
                else
                {
                    p = mid + 1;
                }
            }
            return p;
        }
        var n = list.Count();
        var memo = new int[n];
        int dfs(int i)
        {
            if (memo[i] != 0)
            {
                return memo[i];
            }
            var k = binarySearch(list[i]);
            if (k < i)
            {
                memo[i] = -Maxn;
            }
            if (i > 0 && list[i - 1] < list[i])
            {
                memo[i] = Math.Max(memo[i], dfs(i - 1));
            }
            for (var j = i - 2; j > i - k - 1 && j >= 0; j--)
            {
                if (arr2[k - (i - j - 1)] > list[j])
                {
                    memo[i] = Math.Max(memo[i], dfs(j));
                }
            }
            memo[i]++;
            return memo[i];
        }
        var ans = dfs(n - 1);
        return ans < 0 ? -1 : n - ans;
    }
}
