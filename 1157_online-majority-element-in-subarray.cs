/*
 * @lc app=leetcode.cn id=online-majority-element-in-subarray lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1157] 子数组中占绝大多数的元素
 *
 * https://leetcode.cn/problems/online-majority-element-in-subarray/description/
 *
 * algorithms
 * Hard (36.24%)
 * Total Accepted:    5K
 * Total Submissions: 13.6K
 * Testcase Example:  '["MajorityChecker","query","query","query"]\n' +
  '[[[1,1,2,2,1,1]],[0,5,4],[0,3,3],[2,3,2]]'
 *
 * 设计一个数据结构，有效地找到给定子数组的 多数元素 。
 * 
 * 子数组的 多数元素 是在子数组中出现 threshold 次数或次数以上的元素。
 * 
 * 实现 MajorityChecker 类:
 * 
 * 
 * MajorityChecker(int[] arr) 会用给定的数组 arr 对 MajorityChecker 初始化。
 * int query(int left, int right, int threshold) 返回子数组中的元素  arr[left...right]
 * 至少出现 threshold 次数，如果不存在这样的元素则返回 -1。
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入:
 * ["MajorityChecker", "query", "query", "query"]
 * [[[1, 1, 2, 2, 1, 1]], [0, 5, 4], [0, 3, 3], [2, 3, 2]]
 * 输出：
 * [null, 1, -1, 2]
 * 
 * 解释：
 * MajorityChecker majorityChecker = new MajorityChecker([1,1,2,2,1,1]);
 * majorityChecker.query(0,5,4); // 返回 1
 * majorityChecker.query(0,3,3); // 返回 -1
 * majorityChecker.query(2,3,2); // 返回 2
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= arr.length <= 2 * 10^4
 * 1 <= arr[i] <= 2 * 10^4
 * 0 <= left <= right < arr.length
 * threshold <= right - left + 1
 * 2 * threshold > right - left + 1
 * 调用 query 的次数最多为 10^4 
 * 
 * 
 */

/*
 * // 二分. 官解通过随机性确保性能, 不予置评.
 * public class MajorityChecker
 * {
 *     private const int Maxn = 20000;
 *     private List<int>[] CountInfo;
 *     private (int, int)[] SortedCount;
 * 
 *     public MajorityChecker(int[] arr)
 *     {
 *         var n = arr.Length;
 *         CountInfo = Enumerable.Range(0, Maxn + 1).Select(_ => new List<int>()).ToArray();
 *         var d = new Dictionary<int, int>();
 *         for (var i = 0; i < n; i++)
 *         {
 *             CountInfo[arr[i]].Add(i);
 *             d[arr[i]] = (d.ContainsKey(arr[i]) ? d[arr[i]] : 0) + 1;
 *         }
 *         SortedCount = d.OrderBy(kv => -kv.Value).Select(kv => (kv.Key, kv.Value)).ToArray();
 *     }
 *     
 *     public int Query(int left, int right, int threshold)
 *     {
 *         foreach (var (v, c) in SortedCount)
 *         {
 *             if (threshold > c)
 *             {
 *                 break;
 *             }
 *             var ci = CountInfo[v];
 *             int binarySearch(int k)
 *             {
 *                 var (p, q) = (0, ci.Count());
 *                 while (p < q)
 *                 {
 *                     var mid = (p + q) >> 1;
 *                     if (ci[mid] >= k)
 *                     {
 *                         q = mid;
 *                     }
 *                     else
 *                     {
 *                         p = mid + 1;
 *                     }
 *                 }
 *                 return p;
 *             }
 *             if (binarySearch(right + 1) - binarySearch(left) >= threshold)
 *             {
 *                 return v;
 *             }
 *         }
 *         return -1;
 *     }
 * }
 */

// 参考: https://leetcode.cn/problems/online-majority-element-in-subarray/solution/by-424479543-8wxq/
//
// 按位预处理, 避免 `O(nU)` 的预处理量, 可以获得 `O(nlogU)` 的性能. 就会好很多.
public class MajorityChecker
{
    private const int D = 15;
    private const int Maxn = 20000;
    private int[][] Counts;
    private List<int>[] Indexes;

    public MajorityChecker(int[] arr)
    {
        var n = arr.Length;
        Counts = Enumerable.Range(0, D).Select(_ => new int[n + 1]).ToArray();
        Indexes = Enumerable.Range(0, Maxn + 1).Select(_ => new List<int>()).ToArray();
        for (var i = 0; i < n; i++)
        {
            Indexes[arr[i]].Add(i);
            for (var d = 0; d < D; d++)
            {
                Counts[d][i + 1] = Counts[d][i] + ((arr[i] >> d) & 1);
            }
        }
    }
    
    public int Query(int left, int right, int threshold)
    {
        var ans = 0;
        right++;
        for (var d = 0; d < D; d++)
        {
            var ones = Counts[d][right] - Counts[d][left];
            if (ones >= threshold)
            {
                ans |= 1 << d;
            }
            else if (right - left - ones < threshold)
            {
                return -1;
            }
        }
        var index = Indexes[ans];
        int binarySearch(int k)
        {
            var (p, q) = (0, index.Count());
            while (p < q)
            {
                var mid = (p + q) >> 1;
                if (index[mid] >= k)
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
        return binarySearch(right) - binarySearch(left) >= threshold ? ans : -1;
    }
}

/**
 * Your MajorityChecker object will be instantiated and called as such:
 * MajorityChecker obj = new MajorityChecker(arr);
 * int param_1 = obj.Query(left,right,threshold);
 */
