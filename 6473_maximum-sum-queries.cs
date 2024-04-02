/*
 * @lc app=leetcode.cn id=maximum-sum-queries lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6473] 最大和查询
 *
 * https://leetcode.cn/problems/maximum-sum-queries/description/
 *
 * algorithms
 * Hard (16.36%)
 * Total Accepted:    429
 * Total Submissions: 2.3K
 * Testcase Example:  '[4,3,1,2]\n[2,4,9,5]\n[[4,1],[1,3],[2,5]]'
 *
 * 给你两个长度为 n 、下标从 0 开始的整数数组 nums1 和 nums2 ，另给你一个下标从 1 开始的二维数组 queries ，其中
 * queries[i] = [xi, yi] 。
 * 
 * 对于第 i 个查询，在所有满足 nums1[j] >= xi 且 nums2[j] >= yi 的下标 j (0 <= j < n) 中，找出
 * nums1[j] + nums2[j] 的 最大值 ，如果不存在满足条件的 j 则返回 -1 。
 * 
 * 返回数组 answer ，其中 answer[i] 是第 i 个查询的答案。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：nums1 = [4,3,1,2], nums2 = [2,4,9,5], queries = [[4,1],[1,3],[2,5]]
 * 输出：[6,10,7]
 * 解释：
 * 对于第 1 个查询：xi = 4 且 yi = 1 ，可以选择下标 j = 0 ，此时 nums1[j] >= 4 且 nums2[j] >= 1
 * 。nums1[j] + nums2[j] 等于 6 ，可以证明 6 是可以获得的最大值。
 * 对于第 2 个查询：xi = 1 且 yi = 3 ，可以选择下标 j = 2 ，此时 nums1[j] >= 1 且 nums2[j] >= 3
 * 。nums1[j] + nums2[j] 等于 10 ，可以证明 10 是可以获得的最大值。
 * 对于第 3 个查询：xi = 2 且 yi = 5 ，可以选择下标 j = 3 ，此时 nums1[j] >= 2 且 nums2[j] >= 5
 * 。nums1[j] + nums2[j] 等于 7 ，可以证明 7 是可以获得的最大值。
 * 因此，我们返回 [6,10,7] 。
 * 
 * 
 * 示例 2：
 * 
 * 输入：nums1 = [3,2,5], nums2 = [2,3,4], queries = [[4,4],[3,2],[1,1]]
 * 输出：[9,9,9]
 * 解释：对于这个示例，我们可以选择下标 j = 2 ，该下标可以满足每个查询的限制。
 * 
 * 
 * 示例 3：
 * 
 * 输入：nums1 = [2,1], nums2 = [2,3], queries = [[3,3]]
 * 输出：[-1]
 * 解释：示例中的查询 xi = 3 且 yi = 3 。对于每个下标 j ，都只满足 nums1[j] < xi 或者 nums2[j] < yi
 * 。因此，不存在答案。 
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * nums1.length == nums2.length 
 * n == nums1.length 
 * 1 <= n <= 10^5
 * 1 <= nums1[i], nums2[i] <= 10^9 
 * 1 <= queries.length <= 10^5
 * queries[i].length == 2
 * xi == queries[i][1]
 * yi == queries[i][2]
 * 1 <= xi, yi <= 10^9
 * 
 * 
 */

/*
 * // 不知道哪里来的优先队列的做法, 很神奇.
 * public class Solution
 * {
 *     public int[] MaximumSumQueries(int[] nums1, int[] nums2, int[][] queries)
 *     {
 *         var Q = new PriorityQueue<(int, int), int>();
 *         foreach (var (n1, n2) in nums1
 *             .Zip(nums2))
 *         {
 *             Q.Enqueue((n1, n2), -(n1 + n2));
 *         }
 *         var ans = new int[queries.Length];
 *         Array.Fill(ans, -1);
 *         var last = -1;
 *         var candidates = new List<(int, int)>();
 *         foreach (var (x, y, i) in queries
 *             .Select((q, i) => (q[0], q[1], i))
 *             .OrderBy(x => x.Item1)
 *             .ThenBy(x => x.Item2))
 *         {
 *             if (last != x)
 *             {
 *                 foreach (var c in candidates)
 *                 {
 *                     Q.Enqueue(c, -(c.Item1 + c.Item2));
 *                 }
 *                 candidates = new();
 *                 last = x;
 *             }
 *             while (Q.Count > 0)
 *             {
 *                 Q.TryPeek(out var ns, out var k);
 *                 var (n1, n2) = ns;
 *                 if (n1 >= x && n2 >= y)
 *                 {
 *                     break;
 *                 }
 *                 Q.Dequeue();
 *                 if (n1 >= x)
 *                 {
 *                     candidates.Add(ns);
 *                 }
 *             }
 *             if (Q.Count > 0)
 *             {
 *                 Q.TryPeek(out var _, out var k);
 *                 ans[i] = -k;
 *             }
 *         }
 *         return ans;
 *     }
 * }
 */

// 参考: https://leetcode.cn/problems/maximum-sum-queries/solutions/2305395/pai-xu-dan-diao-zhan-shang-er-fen-by-end-of9h/?envType=daily-question&envId=2023-11-17
//
// 补充了单调栈上二分的做法.
public class Solution
{
    public int[] MaximumSumQueries(int[] nums1, int[] nums2, int[][] queries)
    {
        var pairs = nums1
            .Zip(nums2, (n1, n2) => (x: n1, y: n2, z: n1 + n2))
            .OrderByDescending(x => x.x)
            .ToArray();
        var ans = new int[queries.Length];
        Array.Fill(ans, -1);
        var n = nums1.Length;
        var S = new (int x, int y, int z)[n];
        var (i, j) = (0, 0);
        foreach (var query in queries
            .Select((q, i) => (x: q[0], y: q[1], i: i))
            .OrderByDescending(x => x.x))
        {
            for (; j < n && pairs[j].x >= query.x; j++)
            {
                for (; i > 0 && S[i - 1].z <= pairs[j].z; i--) { }
                if (i is 0 || S[i - 1].y < pairs[j].y) { S[i++] = pairs[j]; }
            }
            var (p, q) = (0, i);
            while (p < q)
            {
                var mid = (p + q) >> 1;
                if (S[mid].y >= query.y) { q = mid; }
                else { p = mid + 1; }
            }
            if (p < i) { ans[query.i] = S[p].z; }
        }
        return ans;
    }
}
