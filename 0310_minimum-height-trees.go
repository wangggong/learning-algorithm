/*
 * @lc app=leetcode.cn id=minimum-height-trees lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [310] 最小高度树
 *
 * https://leetcode-cn.com/problems/minimum-height-trees/description/
 *
 * algorithms
 * Medium (38.35%)
 * Total Accepted:    33K
 * Total Submissions: 82.2K
 * Testcase Example:  '4\n[[1,0],[1,2],[1,3]]'
 *
 * 树是一个无向图，其中任何两个顶点只通过一条路径连接。 换句话说，一个任何没有简单环路的连通图都是一棵树。
 *
 * 给你一棵包含 n 个节点的树，标记为 0 到 n - 1 。给定数字 n 和一个有 n - 1 条无向边的 edges
 * 列表（每一个边都是一对标签），其中 edges[i] = [ai, bi] 表示树中节点 ai 和 bi 之间存在一条无向边。
 *
 * 可选择树中任何一个节点作为根。当选择节点 x 作为根节点时，设结果树的高度为 h 。在所有可能的树中，具有最小高度的树（即，min(h)）被称为
 * 最小高度树 。
 *
 * 请你找到所有的 最小高度树 并按 任意顺序 返回它们的根节点标签列表。
 * 树的 高度 是指根节点和叶子节点之间最长向下路径上边的数量。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 4, edges = [[1,0],[1,2],[1,3]]
 * 输出：[1]
 * 解释：如图所示，当根是标签为 1 的节点时，树的高度是 1 ，这是唯一的最小高度树。
 *
 * 示例 2：
 *
 *
 * 输入：n = 6, edges = [[3,0],[3,1],[3,2],[3,4],[5,4]]
 * 输出：[3,4]
 *
 *
 *
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 2 * 10^4
 * edges.length == n - 1
 * 0 <= ai, bi < n
 * ai != bi
 * 所有 (ai, bi) 互不相同
 * 给定的输入 保证 是一棵树，并且 不会有重复的边
 *
 *
 */

// 确实没刷过这玩意, 太难了...
// 参考:
// - https://leetcode-cn.com/problems/minimum-height-trees/solution/zui-xiao-gao-du-shu-by-leetcode-solution-6v6f/
// - https://leetcode-cn.com/problems/minimum-height-trees/solution/by-ac_oier-7xio/
//
// PS: 提示很有意思. `How many MHTs can a graph have at most?` 答案是 2.

/*
 * // BFS, 类似树的直径.
 * func findMinHeightTrees(n int, edges [][]int) []int {
 * 	var ans []int
 * 	if n == 0 {
 * 		return ans
 * 	}
 * 	visit := make([]bool, n)
 * 	G := make([][]int, n)
 * 	for _, e := range edges {
 * 		G[e[0]] = append(G[e[0]], e[1])
 * 		G[e[1]] = append(G[e[1]], e[0])
 * 	}
 * 	var Q []int
 * 	curr := 0
 * 	Q = append(Q, curr)
 * 	visit[curr] = true
 * 	for len(Q) > 0 {
 * 		curr = Q[0]
 * 		Q = Q[1:]
 * 		for _, i := range G[curr] {
 * 			if visit[i] {
 * 				continue
 * 			}
 * 			Q = append(Q, i)
 * 			visit[i] = true
 * 		}
 * 	}
 * 	level := 0
 * 	for i := range visit {
 * 		visit[i] = false
 * 	}
 * 	// NOTE:
 * 	// 具体过程如下:
 * 	// 1. `0` 节点走一波 BFS, 找到最远节点 `x`;
 * 	// 2. `x` 节点走一波 BFS, 找到最远节点 `y`;
 * 	// 3. 找 `x -> y` 路径上的中间节点, 为答案;
 * 	// 这里只需要找 `x -> y` 路径上的中间节点即可, 因此只记录这个路径.
 * 	fa := make([]int, n)
 * 	Q = append(Q, curr)
 * 	visit[curr] = true
 * 	fa[curr] = -1
 * 	for len(Q) > 0 {
 * 		level++
 * 		for size := len(Q); size > 0; size-- {
 * 			curr = Q[0]
 * 			Q = Q[1:]
 * 			for _, i := range G[curr] {
 * 				if visit[i] {
 * 					continue
 * 				}
 * 				Q = append(Q, i)
 * 				visit[i] = true
 * 				fa[i] = curr
 * 			}
 * 		}
 * 	}
 * 	for curr != -1 {
 * 		ans = append(ans, curr)
 * 		curr = fa[curr]
 * 	}
 * 	m := len(ans)
 * 	if m%2 == 0 {
 * 		return ans[m/2-1 : m/2+1]
 * 	}
 * 	return ans[m/2 : m/2+1]
 * }
 */

/*
 * // DFS, 大差不差. 也是刚才那个结论.
 * func findMinHeightTrees(n int, edges [][]int) []int {
 * 	var ans []int
 * 	if n == 0 {
 * 		return ans
 * 	}
 * 	G := make([][]int, n)
 * 	for _, e := range edges {
 * 		G[e[0]] = append(G[e[0]], e[1])
 * 		G[e[1]] = append(G[e[1]], e[0])
 * 	}
 * 	visit := make([]bool, n)
 * 	fa := make([]int, n)
 * 	fa[0] = -1
 * 	x, _ := dfs(G, 0, visit, fa)
 * 	visit = make([]bool, n)
 * 	fa[x] = -1
 * 	y, _ := dfs(G, x, visit, fa)
 * 	for curr := y; curr != -1; curr = fa[curr] {
 * 		ans = append(ans, curr)
 * 	}
 * 	switch m := len(ans); m % 2 {
 * 	case 0:
 * 		return ans[m/2-1 : m/2+1]
 * 	case 1:
 * 		return ans[m/2 : m/2+1]
 * 	default:
 * 		// unreachable
 * 	}
 * 	return ans
 * }
 */

func dfs(G [][]int, root int, visit []bool, fa []int) (int, int) {
	visit[root] = true
	node, maxDepth := root, 0
	for _, child := range G[root] {
		if visit[child] {
			continue
		}
		fa[child] = root
		n, m := dfs(G, child, visit, fa)
		if m+1 > maxDepth {
			node, maxDepth = n, m+1
		}
	}
	return node, maxDepth
}

/*
 * // 拓扑排序 (OOM 了, 就说吧)
 * func findMinHeightTrees(n int, edges [][]int) []int {
 * 	if n == 0 {
 * 		return nil
 * 	}
 * 	if n == 1 {
 * 		return []int{0}
 * 	}
 * 	G := make([]map[int]struct{}, n)
 * 	for i := range G {
 * 		G[i] = make(map[int]struct{}, n)
 * 	}
 * 	for _, e := range edges {
 * 		G[e[0]][e[1]] = struct{}{}
 * 		G[e[1]][e[0]] = struct{}{}
 * 	}
 * 	var Q []int
 * 	visit := make([]bool, n)
 * 	rest := n
 * 	for i, m := range G {
 * 		if len(m) == 1 {
 * 			Q = append(Q, i)
 * 			visit[i] = true
 * 		}
 * 	}
 * 	for rest > 2 {
 * 		for size := len(Q); size > 0; size-- {
 * 			q := Q[0]
 * 			Q = Q[1:]
 * 			rest--
 * 			for k := range G[q] {
 * 				if visit[k] {
 * 					continue
 * 				}
 * 				delete(G[k], q)
 * 				if len(G[k]) == 1 {
 * 					Q = append(Q, k)
 * 					visit[k] = true
 * 				}
 * 			}
 * 			G[q] = nil
 * 		}
 * 	}
 * 	return Q
 * }
 */

/*
 * // 拓扑排序 (维护度数的 BFS). 差不多还是那个结论, 但具体实现还是有点坑的.
 * func findMinHeightTrees(n int, edges [][]int) []int {
 * 	if n == 0 {
 * 		return nil
 * 	}
 * 	if n == 1 {
 * 		return []int{0}
 * 	}
 * 	G := make([][]int, n)
 * 	deg := make([]int, n)
 * 	for _, e := range edges {
 * 		G[e[0]] = append(G[e[0]], e[1])
 * 		G[e[1]] = append(G[e[1]], e[0])
 * 		deg[e[0]]++
 * 		deg[e[1]]++
 * 	}
 * 	visit := make([]bool, n)
 * 	var Q []int
 * 	for i, d := range deg {
 * 		if d == 1 {
 * 			Q = append(Q, i)
 * 			visit[i] = true
 * 		}
 * 	}
 * 	rest := n
 * 	for rest > 2 {
 * 		for size := len(Q); size > 0; size-- {
 * 			q := Q[0]
 * 			Q = Q[1:]
 * 			rest--
 * 			for _, k := range G[q] {
 * 				if visit[k] {
 * 					continue
 * 				}
 * 				deg[k]--
 * 				if deg[k] == 1 {
 * 					Q = append(Q, k)
 * 					visit[k] = true
 * 				}
 * 			}
 * 		}
 * 	}
 * 	return Q
 * }
 */

const (
	maxInd = iota
	subMaxInd
)

var up []int
var down [][2]int
var nodes []int

// 树形 DP
// 树形 DP 通常根据 `方向` 划分. 先预处理出 `up[u]` & `down[u]`, 分别代表以 `0` 为根节点的树中 "向下" 和 "向上" 的最大高度.
// 这样以当前节点为根的树的深度为 `max(up[u], down[u])`.
func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 0 {
		return nil
	}
	G := make([][]int, n)
	for _, e := range edges {
		s, t := e[0], e[1]
		G[s] = append(G[s], t)
		G[t] = append(G[t], s)
	}
	up = make([]int, n)
	down = make([][2]int, n)
	nodes = make([]int, n)
	fillDown(G, 0, -1)
	fillUp(G, 0, -1)
	var ans []int
	minDepth := n
	for i := 0; i < n; i++ {
		if dep := max(up[i], down[i][maxInd]); minDepth > dep {
			minDepth = dep
			ans = []int{i}
		} else if minDepth == dep {
			ans = append(ans, i)
		} else {
			// do nothing
		}
	}
	return ans
}

// fillDown 填充 "以 0 号节点作为根节点时向下的最大 / 次大深度" 数组.
//
// 就... dfs 就完事了. 与一般 dfs 不一样的是需要把次大深度也取出来, 下面要用.
func fillDown(G [][]int, root, fa int) int {
	for _, child := range G[root] {
		if child == fa {
			continue
		}
		dep := fillDown(G, child, root) + 1
		if dep > down[root][maxInd] {
			down[root][subMaxInd] = down[root][maxInd]
			down[root][maxInd] = dep
			nodes[root] = child
			continue
		}
		if dep > down[root][subMaxInd] {
			down[root][subMaxInd] = dep
		}
	}
	return down[root][maxInd]
}

// fillUp 填充 "以 0 号节点作为根节点时向上的最大深度" 数组.
//
// 首先, 当前节点 `root` 向上的路径必经父节点 `fa`. 在此基础上分几种情况:
// 1. 父节点向上的路径 `up[fa]+1` (+1 是把父节点本身算上了, 下同);
// 2. 父节点向下的路径:
// 2.1 父节点向下的最长路径如果途径当前节点, 就得选次长路径: `down[fa][subMaxInd]+1` (没得选? 无所谓, 取最小值.)
// 2.2 否则, 直接选最长路径: `down[fa][maxInd]+1`
// 以上两种情况取较大值.
func fillUp(G [][]int, root, fa int) {
	for _, child := range G[root] {
		if child == fa {
			continue
		}
		up[child] = up[root] + 1
		if nodes[root] == child {
			up[child] = max(up[child], down[root][subMaxInd]+1)
		} else {
			up[child] = max(up[child], down[root][maxInd]+1)
		}
		fillUp(G, child, root)
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

