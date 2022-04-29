/*
 * @lc app=leetcode.cn id=longest-common-prefix lang=cpp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [14] 最长公共前缀
 *
 * https://leetcode-cn.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (42.21%)
 * Total Accepted:    814K
 * Total Submissions: 1.9M
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * 编写一个函数来查找字符串数组中的最长公共前缀。
 * 
 * 如果不存在公共前缀，返回空字符串 ""。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：strs = ["flower","flow","flight"]
 * 输出："fl"
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：strs = ["dog","racecar","car"]
 * 输出：""
 * 解释：输入不存在公共前缀。
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= strs.length <= 200
 * 0 <= strs[i].length <= 200
 * strs[i] 仅由小写英文字母组成
 * 
 * 
 */
class Solution {
public:
    string longestCommonPrefix(vector<string>& strs) {
        auto n = strs.size();
        if (n == 0) return 0;
        string ans = strs[0];
        for (int i = 1; i < n; i++) {
            int j = 0;
            for (; j < ans.size() && j < strs[i].size() && strs[i][j] == ans[j]; j++);
            ans = ans.substr(0, j);
        }
        return ans;
    }
};
