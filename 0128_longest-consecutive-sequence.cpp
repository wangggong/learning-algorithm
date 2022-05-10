/*
 * @lc app=leetcode.cn id=longest-consecutive-sequence lang=cpp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [128] 最长连续序列
 *
 * https://leetcode-cn.com/problems/longest-consecutive-sequence/description/
 *
 * algorithms
 * Medium (54.84%)
 * Total Accepted:    257.9K
 * Total Submissions: 470.1K
 * Testcase Example:  '[100,4,200,1,3,2]'
 *
 * 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
 * 
 * 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [100,4,200,1,3,2]
 * 输出：4
 * 解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [0,3,7,2,5,8,4,6,0,1]
 * 输出：9
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 0 <= nums.length <= 1e5
 * -1e9 <= nums[i] <= 1e9
 * 
 * 
 */
class Solution {
    int find(vector<int> &uf, int p) {
        int k = p;
        for (; uf[k] != k; k = uf[k]);
        uf[p] = k;
        return k;
    }

    void merge(vector<int> &uf, int p, int q) {
        uf[find(uf, p)] = uf[find(uf, q)];
        return;
    }
public:
    /*
     * int longestConsecutive(vector<int>& nums) {
     *     vector<int> unique;
     *     unordered_map<int, int> m;
     *     for (int n : nums)
     *         if (!m.count(n))
     *             m[n] = 0, unique.push_back(n);
     *     int n = unique.size();
     *     for (int i = 0; i < n; i++) m[unique[i]] = i + 1;
     *     vector<int> uf(n + 1, 0);
     *     for (int i = 1; i <= n; i++) uf[i] = i;
     *     for (int u : unique) {
     *         if (m.count(u - 1))
     *             merge(uf, m[u], m[u - 1]);
     *         if (m.count(u + 1))
     *             merge(uf, m[u], m[u + 1]);
     *     }
     *     m.clear();
     *     for (int i = 1; i < uf.size(); i++) m[find(uf, i)]++;
     *     int ans = 0;
     *     for (auto c : m) ans = max(ans, c.second);
     *     return ans;
     * }
     */
    int longestConsecutive(vector<int>& nums) {
        unordered_map<int, int> m;
        int ans = 0;
        for (int n : nums)
            if (!m.count(n)) {
                int left = m.count(n - 1) ? m[n - 1] : 0;
                int right = m.count(n + 1) ? m[n + 1] : 0;
                int cur = left + right + 1;
                ans = max(ans, cur);
                m[n] = m[n - left] = m[n + right] = cur;

            }
        return ans;
    }
};
