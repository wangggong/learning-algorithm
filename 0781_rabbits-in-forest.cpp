/*
 * @lc app=leetcode.cn id=rabbits-in-forest lang=cpp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [781] 森林中的兔子
 *
 * https://leetcode-cn.com/problems/rabbits-in-forest/description/
 *
 * algorithms
 * Medium (60.58%)
 * Total Accepted:    47.6K
 * Total Submissions: 78.6K
 * Testcase Example:  '[1,1,2]'
 *
 * 森林中有未知数量的兔子。提问其中若干只兔子 "还有多少只兔子与你（指被提问的兔子）颜色相同?" ，将答案收集到一个整数数组 answers 中，其中
 * answers[i] 是第 i 只兔子的回答。
 * 
 * 给你数组 answers ，返回森林中兔子的最少数量。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：answers = [1,1,2]
 * 输出：5
 * 解释：
 * 两只回答了 "1" 的兔子可能有相同的颜色，设为红色。 
 * 之后回答了 "2" 的兔子不会是红色，否则他们的回答会相互矛盾。
 * 设回答了 "2" 的兔子为蓝色。 
 * 此外，森林中还应有另外 2 只蓝色兔子的回答没有包含在数组中。 
 * 因此森林中兔子的最少数量是 5 只：3 只回答的和 2 只没有回答的。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：answers = [10,10,10]
 * 输出：11
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= answers.length <= 1000
 * 0 <= answers[i] < 1000
 * 
 * 
 */
class Solution {
public:
    /*
     * int numRabbits(vector<int>& answers) {
     *     int n = answers.size();
     *     sort(answers.begin(), answers.end());
     *     int ans = 0, ind = 0;
     *     while (ind < n) {
     *         auto i = ind;
     *         for (; i < min(n, ind + answers[ind] + 1) && answers[i] == answers[ind]; i++);
     *         ans += answers[ind] + 1;
     *         ind = i;
     *     }
     *     return ans;
     * }
     */

    // 是的, 官解又是哈希...
    int numRabbits(vector<int>& answers) {
        unordered_map<int, int> count;
        for (auto a : answers) count[a]++;
        int ans = 0;
        for (auto c : count) {
            auto k = c.first, v = c.second;
            ans += (k + 1) * ((v + k) / (k + 1));
        }
        return ans;
    }
};
