/*
 * @lc app=leetcode.cn id=groups-of-special-equivalent-strings lang=cpp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [893] 特殊等价字符串组
 *
 * https://leetcode-cn.com/problems/groups-of-special-equivalent-strings/description/
 *
 * algorithms
 * Medium (72.63%)
 * Total Accepted:    11.5K
 * Total Submissions: 15.9K
 * Testcase Example:  '["abcd","cdab","cbad","xyzz","zzxy","zzyx"]'
 *
 * 给你一个字符串数组 words。
 * 
 * 一步操作中，你可以交换字符串 words[i] 的任意两个偶数下标对应的字符或任意两个奇数下标对应的字符。
 * 
 * 对两个字符串 words[i] 和 words[j] 而言，如果经过任意次数的操作，words[i] == words[j] ，那么这两个字符串是
 * 特殊等价 的。
 * 
 * 
 * 例如，words[i] = "zzxy" 和 words[j] = "xyzz" 是一对 特殊等价 字符串，因为可以按 "zzxy" -> "xzzy"
 * -> "xyzz" 的操作路径使 words[i] == words[j] 。
 * 
 * 
 * 现在规定，words 的 一组特殊等价字符串 就是 words 的一个同时满足下述条件的非空子集：
 * 
 * 
 * 该组中的每一对字符串都是 特殊等价 的
 * 该组字符串已经涵盖了该类别中的所有特殊等价字符串，容量达到理论上的最大值（也就是说，如果一个字符串不在该组中，那么这个字符串就 不会
 * 与该组内任何字符串特殊等价）
 * 
 * 
 * 返回 words 中 特殊等价字符串组 的数量。
 * 
 * 
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：words = ["abcd","cdab","cbad","xyzz","zzxy","zzyx"]
 * 输出：3
 * 解释：
 * 其中一组为 ["abcd", "cdab", "cbad"]，因为它们是成对的特殊等价字符串，且没有其他字符串与这些字符串特殊等价。
 * 另外两组分别是 ["xyzz", "zzxy"] 和 ["zzyx"]。特别需要注意的是，"zzxy" 不与 "zzyx" 特殊等价。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：words = ["abc","acb","bac","bca","cab","cba"]
 * 输出：3
 * 解释：3 组 ["abc","cba"]，["acb","bca"]，["bac","cab"]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= words.length <= 1000
 * 1 <= words[i].length <= 20
 * 所有 words[i] 都只由小写字母组成。
 * 所有 words[i] 都具有相同的长度。
 * 
 * 
 */

const int alpha = 26;

int find(vector<int>& uf, int k) {
    if (k != uf[k]) uf[k] = find(uf, uf[k]);
    return uf[k];
}

void join(vector<int>& uf, int p, int q) {
    uf[find(uf, p)] = find(uf, q);
}

vector<int> new_uf(int n) {
    vector<int> uf(n + 1, 0);
    for (int i = 0; i <= n; i++) uf[i] = i;
    return uf;
}

class Solution {
    bool specialEquivalent(string s, string t) {
        if (s.size() != t.size()) return false;
        auto n = s.size();
        vector<int> count(alpha, 0);
        for (auto i = 0; i < n; i += 2) count[s[i] - 'a']++;
        for (auto i = 0; i < n; i += 2) count[t[i] - 'a']--;
        for (auto c : count) if (c) return false;
        for (auto i = 1; i < n; i += 2) count[s[i] - 'a']++;
        for (auto i = 1; i < n; i += 2) count[t[i] - 'a']--;
        for (auto c : count) if (c) return false;
        return true;
    }

    string getSort(string w) {
        string odd, even;
        auto n = w.size();
        for (auto i = 0; i < n; i += 2) even.push_back(w[i]);
        for (auto i = 1; i < n; i += 2) odd.push_back(w[i]);
        sort(odd.begin(), odd.end());
        sort(even.begin(), even.end());
        return odd + " " + even;
    }

public:
    /*
     * int numSpecialEquivGroups(vector<string>& words) {
     *     auto n = words.size();
     *     if (n == 0) return 0;
     *     auto uf = new_uf(n);
     *     for (auto i = 0; i < n; i++)
     *         for (auto j = i + 1; j < n; j++)
     *             if (specialEquivalent(words[i], words[j]))
     *                 join(uf, i + 1, j + 1);
     *     set<int> s;
     *     for (int i = 1; i <= n; i++) if (find(uf, i)) s.insert(find(uf, i));
     *     return s.size();
     * }
     */

    /*
     * // 其实不需要用并查集...
     * int numSpecialEquivGroups(vector<string>& words) {
     *     auto n = words.size();
     *     if (n == 0) return 0;
     *     vector<string> equivalentWords;
     *     for (auto w : words) {
     *         bool seen = false;
     *         for (auto ew : equivalentWords)
     *             if (specialEquivalent(w, ew)) {
     *                 seen = true;
     *                 break;
     *             }
     *         if (!seen) equivalentWords.push_back(w);
     *     }
     *     return equivalentWords.size();
     * }
     */

    // 甚至可以直接算一个哈希出来...
    int numSpecialEquivGroups(vector<string>& words) {
        unordered_set<string> S(1 << 10);
        for (auto w : words) S.insert(getSort(w));
        return S.size();
    }
};
