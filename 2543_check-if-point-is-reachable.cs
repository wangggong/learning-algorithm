/*
 * @lc app=leetcode.cn id=check-if-point-is-reachable lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [2543] 判断一个点是否可以到达
 *
 * https://leetcode.cn/problems/check-if-point-is-reachable/description/
 *
 * algorithms
 * Hard (37.42%)
 * Total Accepted:    889
 * Total Submissions: 2.4K
 * Testcase Example:  '6\n9'
 *
 * 给你一个无穷大的网格图。一开始你在 (1, 1) ，你需要通过有限步移动到达点 (targetX, targetY) 。
 * 
 * 每一步 ，你可以从点 (x, y) 移动到以下点之一：
 * 
 * 
 * (x, y - x)
 * (x - y, y)
 * (2 * x, y)
 * (x, 2 * y)
 * 
 * 
 * 给你两个整数 targetX 和 targetY ，分别表示你最后需要到达点的 X 和 Y 坐标。如果你可以从 (1, 1)
 * 出发到达这个点，请你返回true ，否则返回 false 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：targetX = 6, targetY = 9
 * 输出：false
 * 解释：没法从 (1,1) 出发到达 (6,9) ，所以返回 false 。
 * 
 * 
 * 示例 2：
 * 
 * 输入：targetX = 4, targetY = 7
 * 输出：true
 * 解释：你可以按照以下路径到达：(1,1) -> (1,2) -> (1,4) -> (1,8) -> (1,7) -> (2,7) -> (4,7)
 * 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= targetX, targetY <= 10^9
 * 
 * 
 */
public class Solution
{
    public bool IsReachable(int targetX, int targetY) => Gcd(RemoveEven(targetX), RemoveEven(targetY)) == 1;

    private int Gcd(int x, int y) => y == 0 ? x : Gcd(y, x % y);
    private int RemoveEven(int n) => n % 2 != 0 ? n : RemoveEven(n / 2);
}
