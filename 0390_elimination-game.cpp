/*
 * @lc app=leetcode.cn id=elimination-game lang=cpp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [390] 消除游戏
 *
 * https://leetcode-cn.com/problems/elimination-game/description/
 *
 * algorithms
 * Medium (60.88%)
 * Total Accepted:    30.2K
 * Total Submissions: 49.6K
 * Testcase Example:  '9'
 *
 * 列表 arr 由在范围 [1, n] 中的所有整数组成，并按严格递增排序。请你对 arr 应用下述算法：
 * 
 * 
 * 
 * 
 * 从左到右，删除第一个数字，然后每隔一个数字删除一个，直到到达列表末尾。
 * 重复上面的步骤，但这次是从右到左。也就是，删除最右侧的数字，然后剩下的数字每隔一个删除一个。
 * 不断重复这两步，从左到右和从右到左交替进行，直到只剩下一个数字。
 * 
 * 
 * 给你整数 n ，返回 arr 最后剩下的数字。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：n = 9
 * 输出：6
 * 解释：
 * arr = [1, 2, 3, 4, 5, 6, 7, 8, 9]
 * arr = [2, 4, 6, 8]
 * arr = [2, 6]
 * arr = [6]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：n = 1
 * 输出：1
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 10^9
 * 
 * 
 * 
 * 
 */
class Solution {
    int remainingLeft(int n) {
        return n <= 1 ? n : remainingRight(n >> 1) << 1;
    }

    int remainingRight(int n) {
        return n + 1 - remainingLeft(n);
    }

    int f(int n, bool dir) {
        int v = 0;
        if (n == 1) {
            v = 1;
            // cout << n << " " << dir << " " << v << endl;
            return v;
        }
        if (dir) {
            v = f(n >> 1, !dir) << 1;
            // cout << n << " " << dir << " " << v << endl;
            return v;
        }
        if (n & 1) {
            v = f(n >> 1, !dir) << 1;
            // cout << n << " " << dir << " " << v << endl;
            return v;
        }
        v = ((f((n >> 1), !dir) - 1) << 1) | 1;
        // cout << n << " " << dir << " " << v << endl;
        return v;
    }
public:
    /*
     * int lastRemaining(int n) {
     *     // return f(n, true);
     *     return n == 1 ? 1 : 2 * (((n >> 1) + 1) - lastRemaining(n >> 1));
     * }
     */

    int lastRemaining(int n) {
        return remainingLeft(n);
    }
};
