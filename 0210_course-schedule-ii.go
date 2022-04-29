/*
 * @lc app=leetcode.cn id=course-schedule-ii lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [210] 课程表 II
 *
 * https://leetcode-cn.com/problems/course-schedule-ii/description/
 *
 * algorithms
 * Medium (54.96%)
 * Total Accepted:    108.9K
 * Total Submissions: 198.1K
 * Testcase Example:  '2\n[[1,0]]'
 *
 * 现在你总共有 numCourses 门课需要选，记为 0 到 numCourses - 1。给你一个数组 prerequisites ，其中
 * prerequisites[i] = [ai, bi] ，表示在选修课程 ai 前 必须 先选修 bi 。
 *
 *
 * 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示：[0,1] 。
 *
 *
 * 返回你为了学完所有课程所安排的学习顺序。可能会有多个正确的顺序，你只要返回 任意一种 就可以了。如果不可能完成所有课程，返回 一个空数组 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：numCourses = 2, prerequisites = [[1,0]]
 * 输出：[0,1]
 * 解释：总共有 2 门课程。要学习课程 1，你需要先完成课程 0。因此，正确的课程顺序为 [0,1] 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
 * 输出：[0,2,1,3]
 * 解释：总共有 4 门课程。要学习课程 3，你应该先完成课程 1 和课程 2。并且课程 1 和课程 2 都应该排在课程 0 之后。
 * 因此，一个正确的课程顺序是 [0,1,2,3] 。另一个正确的排序是 [0,2,1,3] 。
 *
 * 示例 3：
 *
 *
 * 输入：numCourses = 1, prerequisites = []
 * 输出：[0]
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= numCourses <= 2000
 * 0 <= prerequisites.length <= numCourses * (numCourses - 1)
 * prerequisites[i].length == 2
 * 0 <= ai, bi < numCourses
 * ai != bi
 * 所有[ai, bi] 互不相同
 *
 *
 */

var G [][]int
var ans []int
var visit []bool

func findOrder(n int, P [][]int) []int {
	G = make([][]int, n)
	visit = make([]bool, n)
	ans = nil
	for _, p := range P {
		G[p[1]] = append(G[p[1]], p[0])
	}
	for i := 0; i < n; i++ {
		if hasRing(i) {
			return ans
		}
	}
	for i := 0; i < n; i++ {
		dfs(i)
	}
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return ans
}

func hasRing(k int) bool {
	if visit[k] {
		return false
	}
	visit[k] = true
	for _, g := range G[k] {
		if visit[g] || hasRing(g) {
			return true
		}
	}
	visit[k] = false
	return false
}

func dfs(k int) {
	if visit[k] {
		return
	}
	visit[k] = true
	for _, g := range G[k] {
		dfs(g)
	}
	ans = append(ans, k)
	return
}

/*
 * // BFS
 * const MAXN = 2000 + 10
 *
 * var visit [MAXN]bool
 * var S, Q, ans []int
 * var G [MAXN][]int
 * var indegree [MAXN]int
 * var valid bool
 *
 * func findOrder(n int, prerequisites [][]int) []int {
 * 	ans, Q = nil, nil
 * 	for i := 0; i < MAXN; i++ {
 * 		G[i] = nil
 * 		indegree[i] = 0
 * 	}
 * 	for _, p := range prerequisites {
 * 		G[p[0]] = append(G[p[0]], p[1])
 * 		indegree[p[1]]++
 * 	}
 * 	for i := 0; i < n; i++ {
 * 		if indegree[i] > 0 {
 * 			continue
 * 		}
 * 		Q = append(Q, i)
 * 	}
 * 	for len(Q) > 0 {
 * 		q := Q[0]
 * 		Q = Q[1:]
 * 		ans = append(ans, q)
 * 		for _, v := range G[q] {
 * 			indegree[v]--
 * 			if indegree[v] == 0 {
 * 				Q = append(Q, v)
 * 			}
 * 		}
 * 	}
 * 	reverse(ans)
 * 	if len(ans) < n {
 * 		ans = ans[:0]
 * 	}
 * 	return ans
 * }
 */
