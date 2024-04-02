/*
 * @lc app=leetcode.cn id=make-array-empty lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6404] 将数组清空
 *
 * https://leetcode.cn/problems/make-array-empty/description/
 *
 * algorithms
 * Hard (26.97%)
 * Total Accepted:    866
 * Total Submissions: 3.2K
 * Testcase Example:  '[3,4,-1]'
 *
 * 给你一个包含若干 互不相同 整数的数组 nums ，你需要执行以下操作 直到数组为空 ：
 * 
 * 
 * 如果数组中第一个元素是当前数组中的 最小值 ，则删除它。
 * 否则，将第一个元素移动到数组的 末尾 。
 * 
 * 
 * 请你返回需要多少个操作使 nums 为空。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [3,4,-1]
 * 输出：5
 * 
 * 
 * 
 * 
 * 
 * Operation
 * Array
 * 
 * 
 * 
 * 
 * 1
 * [4, -1, 3]
 * 
 * 
 * 2
 * [-1, 3, 4]
 * 
 * 
 * 3
 * [3, 4]
 * 
 * 
 * 4
 * [4]
 * 
 * 
 * 5
 * []
 * 
 * 
 * 
 * 
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [1,2,4,3]
 * 输出：5
 * 
 * 
 * 
 * 
 * 
 * Operation
 * Array
 * 
 * 
 * 
 * 
 * 1
 * [2, 4, 3]
 * 
 * 
 * 2
 * [4, 3]
 * 
 * 
 * 3
 * [3, 4]
 * 
 * 
 * 4
 * [4]
 * 
 * 
 * 5
 * []
 * 
 * 
 * 
 * 
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [1,2,3]
 * 输出：3
 * 
 * 
 * 
 * 
 * 
 * Operation
 * Array
 * 
 * 
 * 
 * 
 * 1
 * [2, 3]
 * 
 * 
 * 2
 * [3]
 * 
 * 
 * 3
 * []
 * 
 * 
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 10^5
 * -10^9 <= nums[i] <= 10^9
 * nums 中的元素 互不相同 。
 * 
 * 
 */
// 参考: https://leetcode.cn/problems/make-array-empty/solution/shu-zhuang-shu-zu-mo-ni-pythonjavacgo-by-ygvb/
public class Solution
{
    public long CountOperationsToEmptyArray(int[] nums)
    {
        var n = nums.Length;
        var indexes = Enumerable
            .Range(0, n)
            .Select(i => (i, nums[i]))
            .OrderBy(k => k.Item2)
            .Select(k => k.Item1)
            .ToList();
        var total = (long)n;
        var cur = 1;
        var tr = new Tree(n + 1);
        for (var i = 0; i < n; i++)
        {
            var k = indexes[i] + 1;
            total += cur <= k
                ? k - cur - tr.Query(cur, k)
                : k - cur + (n - i) + tr.Query(k, cur - 1);
            cur = k;
            tr.Add(k);
        }
        return total;
    }

    public class Tree
    {
        private int[] Tr;
        private int LowBit(int x) => x & (-x);
        public Tree(int n) => Tr = new int[n];

        public int Query(int l, int r) => Query(r) - Query(l - 1);
        public int Query(int x)
        {
            var ans = 0;
            for (; x > 0; x -= LowBit(x))
            {
                ans += Tr[x];
            }
            return ans;
        }

        public void Add(int x)
        {
            for (var n = Tr.Length; x < n; x += LowBit(x))
            {
                Tr[x]++;
            }
        }

    }
}
