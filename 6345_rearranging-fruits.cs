/*
 * @lc app=leetcode.cn id=rearranging-fruits lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6345] 重排水果
 *
 * https://leetcode.cn/problems/rearranging-fruits/description/
 *
 * algorithms
 * Hard (28.95%)
 * Total Accepted:    1.3K
 * Total Submissions: 4.3K
 * Testcase Example:  '[4,2,2,2]\n[1,4,1,2]'
 *
 * 你有两个果篮，每个果篮中有 n 个水果。给你两个下标从 0 开始的整数数组 basket1 和 basket2 ，用以表示两个果篮中每个水果的成本。
 * 
 * 你希望两个果篮相等。为此，可以根据需要多次执行下述操作：
 * 
 * 
 * 选中两个下标 i 和 j ，并交换 basket1 中的第 i 个水果和 basket2 中的第 j 个水果。
 * 交换的成本是 min(basket1i,basket2j) 。
 * 
 * 
 * 根据果篮中水果的成本进行排序，如果排序后结果完全相同，则认为两个果篮相等。
 * 
 * 返回使两个果篮相等的最小交换成本，如果无法使两个果篮相等，则返回 -1 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：basket1 = [4,2,2,2], basket2 = [1,4,1,2]
 * 输出：1
 * 解释：交换 basket1 中下标为 1 的水果和 basket2 中下标为 0 的水果，交换的成本为 1 。此时，basket1 =
 * [4,1,2,2] 且 basket2 = [2,4,1,2] 。重排两个数组，发现二者相等。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：basket1 = [2,3,4,1], basket2 = [3,2,5,1]
 * 输出：-1
 * 解释：可以证明无法使两个果篮相等。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * basket1.length == bakste2.length
 * 1 <= basket1.length <= 10^5
 * 1 <= basket1i,basket2i <= 10^9
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/rearranging-fruits/solution/tan-xin-gou-zao-by-endlesscheng-c2ui/
//
// 傻逼! 都想到贪心的点了还写不对! 还想用四指针! 羞耻不羞耻!
public class Solution
{
    public long MinCost(int[] basket1, int[] basket2)
    {
        var count = new Dictionary<int, int>();
        var count1 = new Dictionary<int, int>();
        foreach (var b in basket1)
        {
            count.TryGetValue(b, out var c);
            count[b] = c + 1;
            count1[b] = c + 1;
        }
        foreach (var b in basket2)
        {
            count.TryGetValue(b, out var c);
            count[b] = c + 1;
        }
        var minVal = count.Select(kv => kv.Key).Min();
        var fruits = new List<int>();
        foreach (var (k, c) in count)
        {
            if (c % 2 != 0)
            {
                return -1;
            }
            count1.TryGetValue(k, out var c1);
            fruits.AddRange(Enumerable.Range(1, Math.Abs(c / 2 - c1)).Select(_ => k));
        }
        fruits.Sort();
        return Enumerable.Range(0, fruits.Count / 2).Select(i => (long)Math.Min(fruits[i], minVal * 2)).Sum();
    }
}
