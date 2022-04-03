/*
 * @lc app=leetcode.cn id=largest-number lang=cpp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [179] 最大数
 *
 * https://leetcode-cn.com/problems/largest-number/description/
 *
 * algorithms
 * Medium (41.02%)
 * Total Accepted:    134.5K
 * Total Submissions: 327.9K
 * Testcase Example:  '[10,2]'
 *
 * 给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。
 * 
 * 注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [10,2]
 * 输出："210"
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [3,30,34,5,9]
 * 输出："9534330"
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [1]
 * 输出："1"
 * 
 * 
 * 示例 4：
 * 
 * 
 * 输入：nums = [10]
 * 输出："10"
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 100
 * 10 <= nums[i] <= 1e9
 * 
 * 
 */
class Solution {
public:
    string largestNumber(vector<int>& N) {
        sort(N.begin(), N.end(), [](const int &x, const int &y) {
            auto sx = to_string(x), sy = to_string(y);
            return sx + sy > sy + sx;
        });
        string ans;
        for (auto x : N) {
            ans = ans + to_string(x);
        }
        // 去除前导 '0'. 如果去空了还得还一个回来.
        ans.erase(0, ans.find_first_not_of("0"));
        if (ans.size() == 0) {
            ans = "0";
        }
        return ans;
    }
};
