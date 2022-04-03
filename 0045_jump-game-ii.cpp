/*
 * @lc app=leetcode.cn id=jump-game-ii lang=cpp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [45] 跳跃游戏 II
 *
 * https://leetcode-cn.com/problems/jump-game-ii/description/
 *
 * algorithms
 * Medium (43.70%)
 * Total Accepted:    259.1K
 * Total Submissions: 590.7K
 * Testcase Example:  '[2,3,1,1,4]'
 *
 * 给你一个非负整数数组 nums ，你最初位于数组的第一个位置。
 * 
 * 数组中的每个元素代表你在该位置可以跳跃的最大长度。
 * 
 * 你的目标是使用最少的跳跃次数到达数组的最后一个位置。
 * 
 * 假设你总是可以到达数组的最后一个位置。
 * 
 * 
 * 
 * 示例 1:
 * 
 * 
 * 输入: nums = [2,3,1,1,4]
 * 输出: 2
 * 解释: 跳到最后一个位置的最小跳跃数是 2。
 * 从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
 * 
 * 
 * 示例 2:
 * 
 * 
 * 输入: nums = [2,3,0,1,4]
 * 输出: 2
 * 
 * 
 * 
 * 
 * 提示:
 * 
 * 
 * 1 <= nums.length <= 1e4
 * 0 <= nums[i] <= 1000
 * 
 * 
 */
class Solution {
public:
    // int jump(vector<int>& nums) {
    //     auto n = nums.size();
    //     queue<pair<int, int>> Q;
    //     Q.push(make_pair(0, 0));
    //     vector<bool> visit(n, false);
    //     while (!Q.empty()) {
    //         auto q = Q.front(); Q.pop();
    //         if (q.first >= n-1) return q.second;
    //         for (auto i = 1; i <= nums[q.first]; i++) {
    //             auto v = q.first + i;
    //             if (v >= n || visit[v]) continue;
    //             Q.push(make_pair(v, q.second + 1));
    //             visit[v] = true;
    //         }
    //     }
    //     return n-1;
    // }

	int jump(vector<int>& nums) {
		auto n = nums.size();
		if (n <= 1) return 0;
		auto ans = 0;
		auto maxPos = 0;
		auto i = 0;
		while (i < n) {
			ans++;
			auto next = maxPos;
			for (; i <= maxPos; i++) {
				next = max(next, i + nums[i]);
			}
			maxPos = next;
			if (maxPos >= n-1) break;
		}
		return ans;
	}
};
