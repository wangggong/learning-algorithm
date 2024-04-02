/*
 * @lc app=leetcode.cn id=finding-mk-average lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1825] 求出 MK 平均值
 *
 * https://leetcode.cn/problems/finding-mk-average/description/
 *
 * algorithms
 * Hard (29.82%)
 * Total Accepted:    9.6K
 * Total Submissions: 24.3K
 * Testcase Example:  '["MKAverage","addElement","addElement","calculateMKAverage","addElement","calculateMKAverage","addElement","addElement","addElement","calculateMKAverage"]\n' +
  '[[3,1],[3],[1],[],[10],[],[5],[5],[5],[]]'
 *
 * 给你两个整数 m 和 k ，以及数据流形式的若干整数。你需要实现一个数据结构，计算这个数据流的 MK 平均值 。
 * 
 * MK 平均值 按照如下步骤计算：
 * 
 * 
 * 如果数据流中的整数少于 m 个，MK 平均值 为 -1 ，否则将数据流中最后 m 个元素拷贝到一个独立的容器中。
 * 从这个容器中删除最小的 k 个数和最大的 k 个数。
 * 计算剩余元素的平均值，并 向下取整到最近的整数 。
 * 
 * 
 * 请你实现 MKAverage 类：
 * 
 * 
 * MKAverage(int m, int k) 用一个空的数据流和两个整数 m 和 k 初始化 MKAverage 对象。
 * void addElement(int num) 往数据流中插入一个新的元素 num 。
 * int calculateMKAverage() 对当前的数据流计算并返回 MK 平均数 ，结果需 向下取整到最近的整数 。
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：
 * ["MKAverage", "addElement", "addElement", "calculateMKAverage",
 * "addElement", "calculateMKAverage", "addElement", "addElement",
 * "addElement", "calculateMKAverage"]
 * [[3, 1], [3], [1], [], [10], [], [5], [5], [5], []]
 * 输出：
 * [null, null, null, -1, null, 3, null, null, null, 5]
 * 
 * 解释：
 * MKAverage obj = new MKAverage(3, 1); 
 * obj.addElement(3);        // 当前元素为 [3]
 * obj.addElement(1);        // 当前元素为 [3,1]
 * obj.calculateMKAverage(); // 返回 -1 ，因为 m = 3 ，但数据流中只有 2 个元素
 * obj.addElement(10);       // 当前元素为 [3,1,10]
 * obj.calculateMKAverage(); // 最后 3 个元素为 [3,1,10]
 * ⁠                         // 删除最小以及最大的 1 个元素后，容器为 [3]
 * ⁠                         // [3] 的平均值等于 3/1 = 3 ，故返回 3
 * obj.addElement(5);        // 当前元素为 [3,1,10,5]
 * obj.addElement(5);        // 当前元素为 [3,1,10,5,5]
 * obj.addElement(5);        // 当前元素为 [3,1,10,5,5,5]
 * obj.calculateMKAverage(); // 最后 3 个元素为 [5,5,5]
 * ⁠                         // 删除最小以及最大的 1 个元素后，容器为 [5]
 * ⁠                         // [5] 的平均值等于 5/1 = 5 ，故返回 5
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 3 <= m <= 10^5
 * 1 <= k*2 < m
 * 1 <= num <= 10^5
 * addElement 与 calculateMKAverage 总操作次数不超过 10^5 次。
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/finding-mk-average/solution/python-shu-zhuang-shu-zu-by-daxin-2021-9h61/
//
// BIT + 二分
public class MKAverage
{
    class BIT
    {
        private int[] Tr;

        private int LowBit(int n) => n & (-n);
        public BIT(int n) => Tr = new int[n];

        public int Query(int n)
        {
            var ans = 0;
            for (; n > 0; n -= LowBit(n))
            {
                ans += Tr[n];
            }
            return ans;
        }

        public void Add(int n, int v)
        {
            for (; n < Tr.Length; n += LowBit(n))
            {
                Tr[n] += v;
            }
        }
    }

    private const int N = (int)2e5;
    private int M;
    private int K;
    private Queue<int> Q;
    private BIT Count;
    private BIT PrefixSum;

    public MKAverage(int m, int k)
    {
        M = m;
        K = k;
        Q = new();
        Count = new(N + 1);
        PrefixSum = new(N + 1);
    }

    private int GetTop(int n)
    {
        int p = 0, q = N;
        while (p < q)
        {
            var mid = (p + q) >> 1;
            if (Count.Query(mid) >= n)
            {
                q = mid;
            }
            else
            {
                p = mid + 1;
            }
        }
        var d = Count.Query(p) - n;
        return PrefixSum.Query(p) - d * p;
    }

    public void AddElement(int n)
    {
        Q.Enqueue(n);
        Count.Add(n, 1);
        PrefixSum.Add(n, n);
        while (Q.Count > M)
        {
            var k = Q.Dequeue();
            Count.Add(k, -1);
            PrefixSum.Add(k, -k);
        }
    }

    public int CalculateMKAverage() => Q.Count < M ? -1 : (GetTop(M - K) - GetTop(K)) / (M - 2 * K);
}

/**
 * Your MKAverage object will be instantiated and called as such:
 * MKAverage obj = new MKAverage(m, k);
 * obj.AddElement(num);
 * int param_2 = obj.CalculateMKAverage();
 */
