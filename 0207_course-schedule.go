/*
 * @lc app=leetcode.cn id=course-schedule lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [207] 课程表
 *
 * https://leetcode-cn.com/problems/course-schedule/description/
 *
 * algorithms
 * Medium (54.07%)
 * Total Accepted:    174.2K
 * Total Submissions: 322.1K
 * Testcase Example:  '2\n[[1,0]]'
 *
 * 你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
 *
 * 在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi]
 * ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。
 *
 *
 * 例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
 *
 *
 * 请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：numCourses = 2, prerequisites = [[1,0]]
 * 输出：true
 * 解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。
 *
 * 示例 2：
 *
 *
 * 输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
 * 输出：false
 * 解释：总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= numCourses <= 1e5
 * 0 <= prerequisites.length <= 5000
 * prerequisites[i].length == 2
 * 0 <= ai, bi < numCourses
 * prerequisites[i] 中的所有课程对 互不相同
 *
 *
 */

const MAXN = 1e5 + 10

var visit [MAXN]int
var vs [MAXN]bool
var cs [MAXN]bool
var ans bool
var G [MAXN][]int
var indegree [MAXN]int

func canFinish(n int, prerequisites [][]int) bool {
	for i := 0; i < MAXN; i++ {
		vs[i] = false
		cs[i] = false
		G[i] = nil
	}
	ans = false
	for _, p := range prerequisites {
		G[p[0]] = append(G[p[0]], p[1])
	}
	for i := 0; i < n; i++ {
		hasRing2(i)
		if ans {
			break
		}
	}
	return !ans
}

func hasRing2(k int) {
	if vs[k] {
		return
	}
	vs[k] = true
	cs[k] = true
	for _, v := range G[k] {
		if cs[v] {
			ans = true
			return
		}
		hasRing2(v)
		if ans {
			return
		}
	}
	cs[k] = false
}

/*
 * func canFinish(numCourses int, prerequisites [][]int) bool {
 * 	ans = false
 * 	for i := 0; i < MAXN; i++ {
 * 		visit[i] = 0
 * 	}
 * 	for i := 0; i < MAXN; i++ {
 * 		G[i] = nil
 * 	}
 * 	for _, p := range prerequisites {
 * 		G[p[0]] = append(G[p[0]], p[1])
 * 	}
 * 	return !hasRing(numCourses)
 * }
 */

/*
 * func canFinish(n int, prerequisites [][]int) bool {
 * 	for i := 0; i < MAXN; i++ {
 * 		G[i] = nil
 * 	}
 * 	visited := 0
 * 	for i := 0; i < MAXN; i++ {
 * 		indegree[i] = 0
 * 	}
 * 	for _, p := range prerequisites {
 * 		G[p[0]] = append(G[p[0]], p[1])
 * 		indegree[p[1]]++
 * 	}
 * 	var Q []int
 * 	for i := 0; i < n; i++ {
 * 		if indegree[i] == 0 {
 * 			Q = append(Q, i)
 * 		}
 * 	}
 * 	for len(Q) > 0 {
 * 		q := Q[0]
 * 		Q = Q[1:]
 * 		visited++
 * 		for _, v := range G[q] {
 * 			indegree[v]--
 * 			if indegree[v] == 0 {
 * 				Q = append(Q, v)
 * 			}
 * 		}
 * 	}
 * 	return visited == n
 * }
 */

func hasRing(n int) bool {
	for i := 0; i < n; i++ {
		dfs(i)
		if ans {
			return true
		}
	}
	return false
}

func dfs(k int) {
	visit[k] = 1
	for _, v := range G[k] {
		switch visit[v] {
		case 0:
			dfs(v)
			if ans {
				return
			}
		case 1:
			ans = true
			return
		case 2:
			// do nothing
		}
	}
	visit[k] = 2
}

