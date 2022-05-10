/*
 * @lc app=leetcode.cn id=minimum-genetic-mutation lang=cpp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [433] 最小基因变化
 *
 * https://leetcode-cn.com/problems/minimum-genetic-mutation/description/
 *
 * algorithms
 * Medium (53.49%)
 * Total Accepted:    25.6K
 * Total Submissions: 47.2K
 * Testcase Example:  '"AACCGGTT"\n"AACCGGTA"\n["AACCGGTA"]'
 *
 * 基因序列可以表示为一条由 8 个字符组成的字符串，其中每个字符都是 'A'、'C'、'G' 和 'T' 之一。
 * 
 * 假设我们需要调查从基因序列 start 变为 end 所发生的基因变化。一次基因变化就意味着这个基因序列中的一个字符发生了变化。
 * 
 * 
 * 例如，"AACCGGTT" --> "AACCGGTA" 就是一次基因变化。
 * 
 * 
 * 另有一个基因库 bank 记录了所有有效的基因变化，只有基因库中的基因才是有效的基因序列。
 * 
 * 给你两个基因序列 start 和 end ，以及一个基因库 bank ，请你找出并返回能够使 start 变化为 end
 * 所需的最少变化次数。如果无法完成此基因变化，返回 -1 。
 * 
 * 注意：起始基因序列 start 默认是有效的，但是它并不一定会出现在基因库中。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：start = "AACCGGTT", end = "AACCGGTA", bank = ["AACCGGTA"]
 * 输出：1
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：start = "AACCGGTT", end = "AAACGGTA", bank =
 * ["AACCGGTA","AACCGCTA","AAACGGTA"]
 * 输出：2
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：start = "AAAAACCC", end = "AACCCCCC", bank =
 * ["AAAACCCC","AAACCCCC","AACCCCCC"]
 * 输出：3
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * start.length == 8
 * end.length == 8
 * 0 <= bank.length <= 10
 * bank[i].length == 8
 * start、end 和 bank[i] 仅由字符 ['A', 'C', 'G', 'T'] 组成
 * 
 * 
 */

const string genes = "AGCT";

class Solution {
public:
    int minMutation(string start, string end, vector<string>& bank) {
        int ans = 0;
        unordered_set<string> S, visit;
        for (string b : bank) S.insert(b);
        queue<string> Q;
        Q.push(start);
        visit.insert(start);
        bool found = false;
        while (!Q.empty()) {
            for (int size = Q.size(); size; size--) {
                auto curr = Q.front(); Q.pop();
                if (curr == end) {
                    found = true;
                    break;
                }
                auto next = curr;
                for (int i = 0; i < next.size(); i++) {
                    for (auto g : genes) if (g != curr[i]) {
                        next[i] = g;
                        if (!S.count(next) || visit.count(next)) continue;
                        Q.push(next);
                        visit.insert(next);
                    }
                    next[i] = curr[i];
                }
            }
            if (found) break;
            ans++;
        }
        return found ? ans : -1;
    }
};
