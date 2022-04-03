/*
 * @lc app=leetcode.cn id=reconstruct-itinerary lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [332] 重新安排行程
 *
 * https://leetcode-cn.com/problems/reconstruct-itinerary/description/
 *
 * algorithms
 * Hard (45.39%)
 * Total Accepted:    49.3K
 * Total Submissions: 108.3K
 * Testcase Example:  '[["MUC","LHR"],["JFK","MUC"],["SFO","SJC"],["LHR","SFO"]]'
 *
 * 给你一份航线列表 tickets ，其中 tickets[i] = [fromi, toi]
 * 表示飞机出发和降落的机场地点。请你对该行程进行重新规划排序。
 *
 * 所有这些机票都属于一个从 JFK（肯尼迪国际机场）出发的先生，所以该行程必须从 JFK
 * 开始。如果存在多种有效的行程，请你按字典排序返回最小的行程组合。
 *
 *
 * 例如，行程 ["JFK", "LGA"] 与 ["JFK", "LGB"] 相比就更小，排序更靠前。
 *
 *
 * 假定所有机票至少存在一种合理的行程。且所有的机票 必须都用一次 且 只能用一次。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：tickets = [["MUC","LHR"],["JFK","MUC"],["SFO","SJC"],["LHR","SFO"]]
 * 输出：["JFK","MUC","LHR","SFO","SJC"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：tickets =
 * [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
 * 输出：["JFK","ATL","JFK","SFO","ATL","SFO"]
 * 解释：另一种有效的行程是 ["JFK","SFO","ATL","JFK","ATL","SFO"] ，但是它字典排序更大更靠后。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * tickets[i].length == 2
 * fromi.length == 3
 * toi.length == 3
 * fromi 和 toi 由大写英文字母组成
 * fromi != toi
 *
 *
 */

import "sort"

/*
 * const JFK = "JFK"
 *
 * var ans []string
 * var cur []string
 * var found bool
 *
 * var G map[string]map[string]int
 *
 * func findItinerary(tickets [][]string) []string {
 * 	n := len(tickets)
 * 	ans, cur, found = nil, nil, false
 * 	G = make(map[string]map[string]int)
 * 	for _, t := range tickets {
 * 		from, to := t[0], t[1]
 * 		if _, ok := G[from]; !ok {
 * 			G[from] = make(map[string]int)
 * 		}
 * 		G[from][to]++
 * 	}
 * 	backtrack(tickets, JFK, 0, n)
 * 	ans = append([]string{JFK}, ans...)
 * 	return ans
 * }
 *
 * func backtrack(arr [][]string, from string, k, n int) {
 * 	if k == n {
 * 		ans = make([]string, n)
 * 		copy(ans, cur)
 * 		found = true
 * 		return
 * 	}
 * 	var tos []string
 * 	for to := range G[from] {
 * 		tos = append(tos, to)
 * 	}
 * 	// NOTE: 重点是把当前结果的目的地按照字典序排序, 这样我们直接查第一个结果必然符合条件.
 * 	sort.Strings(tos)
 * 	for _, to := range tos {
 * 		if G[from][to] == 0 {
 * 			continue
 * 		}
 * 		cur = append(cur, to)
 * 		G[from][to]--
 * 		backtrack(arr, to, k+1, n)
 * 		if found {
 * 			return
 * 		}
 * 		cur = cur[:len(cur)-1]
 * 		G[from][to]++
 * 	}
 * }
 */

var ans []string
var cur []string
var G map[string][]int
var visit []bool

func findItinerary(tickets [][]string) []string {
	// NOTE: 在哪都是排序, 直接一开始把目的地按照字典序排序不香吗?
	sort.Slice(tickets, func(p, q int) bool { return tickets[p][1] < tickets[q][1] })
	n := len(tickets)
	visit = make([]bool, n)
	ans, cur = nil, nil
	G = make(map[string][]int)
	for i, t := range tickets {
		from := t[0]
		G[from] = append(G[from], i)
	}
	backtrack(tickets, "JFK", n)
	return ans
}

func backtrack(arr [][]string, from string, k int) bool {
	if k == 0 {
		cur = append(cur, from)
		if len(ans) == 0 || less(cur, ans) {
			ans = copyOf(cur)
		}
		cur = cur[:len(cur)-1]
		return true
	}
	cur = append(cur, from)
	for _, i := range G[from] {
		if visit[i] {
			continue
		}
		visit[i] = true
		to := arr[i][1]
		if backtrack(arr, to, k-1) {
			return true
		}
		visit[i] = false
	}
	cur = cur[:len(cur)-1]
	return false
}

func less(p, q []string) bool {
	n := len(p)
	// assert len(p) == len(q);
	for i := 0; i < n; i++ {
		if p[i] < q[i] {
			return true
		} else if p[i] > q[i] {
			return false
		}
	}
	return false
}

func copyOf(arr []string) []string {
	ans := make([]string, len(arr))
	copy(ans, arr)
	return ans
}
