/*
 * @lc app=leetcode.cn id=find-the-kth-smallest-sum-of-a-matrix-with-sorted-rows lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1439] 有序矩阵中的第 k 个最小数组和
 *
 * https://leetcode.cn/problems/find-the-kth-smallest-sum-of-a-matrix-with-sorted-rows/description/
 *
 * algorithms
 * Hard (56.87%)
 * Total Accepted:    7.7K
 * Total Submissions: 13.4K
 * Testcase Example:  '[[1,3,11],[2,4,6]]\n5'
 *
 * 给你一个 m * n 的矩阵 mat，以及一个整数 k ，矩阵中的每一行都以非递减的顺序排列。
 * 
 * 你可以从每一行中选出 1 个元素形成一个数组。返回所有可能数组中的第 k 个 最小 数组和。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：mat = [[1,3,11],[2,4,6]], k = 5
 * 输出：7
 * 解释：从每一行中选出一个元素，前 k 个和最小的数组分别是：
 * [1,2], [1,4], [3,2], [3,4], [1,6]。其中第 5 个的和是 7 。  
 * 
 * 示例 2：
 * 
 * 输入：mat = [[1,3,11],[2,4,6]], k = 9
 * 输出：17
 * 
 * 
 * 示例 3：
 * 
 * 输入：mat = [[1,10,10],[1,4,5],[2,3,6]], k = 7
 * 输出：9
 * 解释：从每一行中选出一个元素，前 k 个和最小的数组分别是：
 * [1,1,2], [1,1,3], [1,4,2], [1,4,3], [1,1,6], [1,5,2], [1,5,3]。其中第 7 个的和是 9
 * 。 
 * 
 * 
 * 示例 4：
 * 
 * 输入：mat = [[1,1,10],[2,2,9]], k = 7
 * 输出：12
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * m == mat.length
 * n == mat.length[i]
 * 1 <= m, n <= 40
 * 1 <= k <= min(200, n ^ m)
 * 1 <= mat[i][j] <= 5000
 * mat[i] 是一个非递减数组
 * 
 * 
 */

/*
 * public class Solution
 * {
 *     private const long P = 13331;
 *     private static long GetHash(int[] arr) => arr.Select(v => (long)v).Aggregate((x, y) => x * P + y);
 * 
 *     public int KthSmallest(int[][] M, int k)
 *     {
 *         var (n, m) = (M.Length, M[0].Length);
 *         var Q = new PriorityQueue<int[], int>();
 *         var S = new HashSet<long>();
 *         var start = new int[n];
 *         Q.Enqueue(start, M.Zip(start, (row, col) => row[col]).Sum());
 *         S.Add(GetHash(start));
 *         for (var i = 1; true; i++)
 *         {
 *             Q.TryDequeue(out var arr, out var sum);
 *             if (i == k)
 *             {
 *                 return sum;
 *             }
 *             for (var r = 0; r < n; r++)
 *             {
 *                 if (arr[r] + 1 < m)
 *                 {
 *                     var next = arr.Select((v, i) => v + (i == r ? 1 : 0)).ToArray();
 *                     var hash = GetHash(next);
 *                     if (!S.Contains(hash))
 *                     {
 *                         Q.Enqueue(next, sum + (M[r][next[r]] - M[r][arr[r]]));
 *                         S.Add(hash);
 *                     }
 *                 }
 *             }
 *         }
 *     }
 * }
 */

public class Solution
{
    public int KthSmallest(int[][] M, int k)
    {
        var arr = new int[1];
        for (var (i, n) = (0, M.Length); i < n; i++)
        {
            arr = M[i]
                .SelectMany(v => arr.Select(a => v + a))
                .OrderBy(x => x)
                .Take(k)
                .ToArray();
        }
        return arr[arr.Length - 1];
    }
}
