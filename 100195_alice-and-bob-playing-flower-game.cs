/*
 * @lc app=leetcode.cn id=alice-and-bob-playing-flower-game lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [100195] Alice 和 Bob 玩鲜花游戏
 *
 * https://leetcode.cn/problems/alice-and-bob-playing-flower-game/description/
 *
 * algorithms
 * Medium (50.88%)
 * Total Accepted:    2.3K
 * Total Submissions: 4.5K
 * Testcase Example:  '3\n2'
 *
 * Alice 和 Bob 在一个长满鲜花的环形草地玩一个回合制游戏。环形的草地上有一些鲜花，Alice 到 Bob 之间顺时针有 x 朵鲜花，逆时针有 y
 * 朵鲜花。
 * 
 * 游戏过程如下：
 * 
 * 
 * Alice 先行动。
 * 每一次行动中，当前玩家必须选择顺时针或者逆时针，然后在这个方向上摘一朵鲜花。
 * 一次行动结束后，如果所有鲜花都被摘完了，那么 当前 玩家抓住对手并赢得游戏的胜利。
 * 
 * 
 * 给你两个整数 n 和 m ，你的任务是求出满足以下条件的所有 (x, y) 对：
 * 
 * 
 * 按照上述规则，Alice 必须赢得游戏。
 * Alice 顺时针方向上的鲜花数目 x 必须在区间 [1,n] 之间。
 * Alice 逆时针方向上的鲜花数目 y 必须在区间 [1,m] 之间。
 * 
 * 
 * 请你返回满足题目描述的数对 (x, y) 的数目。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：n = 3, m = 2
 * 输出：3
 * 解释：以下数对满足题目要求：(1,2) ，(3,2) ，(2,1) 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：n = 1, m = 1
 * 输出：0
 * 解释：没有数对满足题目要求。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n, m <= 10^5
 * 
 * 
 */
public class Solution
{
    public long FlowerGame(int n, int m) => OddCount(n) * EvenCount(m)
        + EvenCount(n) * OddCount(m);
    private long OddCount(int n) => (long)(n + 1) >> 1;
    private long EvenCount(int n) => (long)n >> 1;
}
