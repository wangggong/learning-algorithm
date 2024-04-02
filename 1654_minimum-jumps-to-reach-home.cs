/*
 * @lc app=leetcode.cn id=minimum-jumps-to-reach-home lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1654] 到家的最少跳跃次数
 *
 * https://leetcode.cn/problems/minimum-jumps-to-reach-home/description/
 *
 * algorithms
 * Medium (31.19%)
 * Total Accepted:    15.2K
 * Total Submissions: 43.7K
 * Testcase Example:  '[14,4,18,1,15]\n3\n15\n9'
 *
 * 有一只跳蚤的家在数轴上的位置 x 处。请你帮助它从位置 0 出发，到达它的家。
 * 
 * 跳蚤跳跃的规则如下：
 * 
 * 
 * 它可以 往前 跳恰好 a 个位置（即往右跳）。
 * 它可以 往后 跳恰好 b 个位置（即往左跳）。
 * 它不能 连续 往后跳 2 次。
 * 它不能跳到任何 forbidden 数组中的位置。
 * 
 * 
 * 跳蚤可以往前跳 超过 它的家的位置，但是它 不能跳到负整数 的位置。
 * 
 * 给你一个整数数组 forbidden ，其中 forbidden[i] 是跳蚤不能跳到的位置，同时给你整数 a， b 和 x
 * ，请你返回跳蚤到家的最少跳跃次数。如果没有恰好到达 x 的可行方案，请你返回 -1 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：forbidden = [14,4,18,1,15], a = 3, b = 15, x = 9
 * 输出：3
 * 解释：往前跳 3 次（0 -> 3 -> 6 -> 9），跳蚤就到家了。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：forbidden = [8,3,16,6,12,20], a = 15, b = 13, x = 11
 * 输出：-1
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：forbidden = [1,6,2,14,5,17,4], a = 16, b = 9, x = 7
 * 输出：2
 * 解释：往前跳一次（0 -> 16），然后往回跳一次（16 -> 7），跳蚤就到家了。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * 1 
 * 0 
 * forbidden 中所有位置互不相同。
 * 位置 x 不在 forbidden 中。
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/minimum-jumps-to-reach-home/solutions/2414842/dao-jia-de-zui-shao-tiao-yue-ci-shu-by-l-sza1/
//
// 官解给出了一个更精确的上限.
public class Solution
{
    public int MinimumJumps(int[] forbidden, int a, int b, int x)
    {
        const int N = 2000;
        var S = forbidden.ToHashSet();
        bool valid(int x) => 0 <= x && x <= N * 4 && !S.Contains(x);
        var Q = new Queue<(int, bool)>();
        var visit = new HashSet<(int, bool)>();
        var start = (0, false);
        Q.Enqueue(start);
        visit.Add(start);
        for (var step = 0; Q.Count > 0; step++)
        {
            for (var c = Q.Count; c > 0; c--)
            {
                var (position, backed) = Q.Dequeue();
                if (position == x) { return step; }
                foreach (var next in new (int, bool)[]
                {
                    (position + a, false),
                    (position - b, true),
                })
                {
                    if ((backed && next.Item2)
                        || !valid(next.Item1)
                        || visit.Contains(next)) { continue; }
                    Q.Enqueue(next);
                    visit.Add(next);
                }
            }
        }
        return -1;
    }
}
