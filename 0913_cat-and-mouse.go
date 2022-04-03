/*
 * @lc app=leetcode.cn id=cat-and-mouse lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [913] 猫和老鼠
 *
 * https://leetcode-cn.com/problems/cat-and-mouse/description/
 *
 * algorithms
 * Hard (62.53%)
 * Total Accepted:    15.7K
 * Total Submissions: 25.3K
 * Testcase Example:  '[[2,5],[3],[0,4,5],[1,4,5],[2,3],[0,2,3]]'
 *
 * 两位玩家分别扮演猫和老鼠，在一张 无向 图上进行游戏，两人轮流行动。
 *
 * 图的形式是：graph[a] 是一个列表，由满足 ab 是图中的一条边的所有节点 b 组成。
 *
 * 老鼠从节点 1 开始，第一个出发；猫从节点 2 开始，第二个出发。在节点 0 处有一个洞。
 *
 * 在每个玩家的行动中，他们 必须 沿着图中与所在当前位置连通的一条边移动。例如，如果老鼠在节点 1 ，那么它必须移动到 graph[1] 中的任一节点。
 *
 * 此外，猫无法移动到洞中（节点 0）。
 *
 * 然后，游戏在出现以下三种情形之一时结束：
 *
 *
 * 如果猫和老鼠出现在同一个节点，猫获胜。
 * 如果老鼠到达洞中，老鼠获胜。
 * 如果某一位置重复出现（即，玩家的位置和移动顺序都与上一次行动相同），游戏平局。
 *
 *
 * 给你一张图 graph ，并假设两位玩家都都以最佳状态参与游戏：
 *
 *
 * 如果老鼠获胜，则返回 1；
 * 如果猫获胜，则返回 2；
 * 如果平局，则返回 0 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：graph = [[2,5],[3],[0,4,5],[1,4,5],[2,3],[0,2,3]]
 * 输出：0
 *
 *
 * 示例 2：
 *
 *
 * 输入：graph = [[1,3],[0],[3],[0,2]]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 3 <= graph.length <= 50
 * 1 <= graph[i].length < graph.length
 * 0 <= graph[i][j] < graph.length
 * graph[i][j] != i
 * graph[i] 互不相同
 * 猫和老鼠在游戏中总是移动
 *
 *
 */
const MAXN = 50 + 5

const (
	defaultState = iota - 1
	draw
	mouseWin
	catWin
)

const (
	minMaxMouseWin = iota - 1
	minMaxDraw
	minMaxCatWin
)

const (
	mouseTurn = iota
	catTurn
)

var n int

var dp [2 * MAXN * MAXN][MAXN][MAXN]int
var visit [2 * MAXN * MAXN][MAXN][MAXN]bool

/*
 * // 记忆化搜索
 *
 * func mouseMove(G [][]int, t, m, c int) int {
 * 	w, d := false, false
 * 	for _, n := range G[m] {
 * 		ret := dfs(G, t+1, n, c)
 * 		switch ret {
 * 		case mouseWin:
 * 			w = true
 * 		case draw:
 * 			d = true
 * 		}
 * 		if w {
 * 			break
 * 		}
 * 	}
 * 	if w {
 * 		return mouseWin
 * 	}
 * 	if d {
 * 		return draw
 * 	}
 * 	return catWin
 * }
 *
 * func catMove(G [][]int, t, m, c int) int {
 * 	w, d := false, false
 * 	for _, n := range G[c] {
 * 		// NOTE: 猫不能钻洞
 * 		if n == 0 {
 * 			continue
 * 		}
 * 		ret := dfs(G, t+1, m, n)
 * 		switch ret {
 * 		case catWin:
 * 			w = true
 * 		case draw:
 * 			d = true
 * 		}
 * 		if w {
 * 			break
 * 		}
 * 	}
 * 	if w {
 * 		return catWin
 * 	}
 * 	if d {
 * 		return draw
 * 	}
 * 	return mouseWin
 * }
 *
 * func dfs(G [][]int, t, m, c int) (v int) {
 * 	if v = dp[t][m][c]; v > defaultState {
 * 		return
 * 	}
 * 	defer func() {
 * 		// if t < 5 {
 * 		// 	fmt.Printf("%v %v %v %v\n", t, m, c, v)
 * 		// }
 * 		dp[t][m][c] = v
 * 	}()
 * 	if m == 0 {
 * 		v = mouseWin
 * 		return
 * 	}
 * 	if m == c {
 * 		v = catWin
 * 		return
 * 	}
 * 	// NOTE: 抽屉原理, 行进到 2n^2 步时一定重复了
 * 	if t >= 2*n*n {
 * 		v = draw
 * 		return
 * 	}
 * 	if t%2 == 0 {
 * 		v = mouseMove(G, t, m, c)
 * 	} else {
 * 		v = catMove(G, t, m, c)
 * 	}
 * 	return
 * }
 *
 * func catMouseGame(G [][]int) int {
 * 	n = len(G)
 * 	for k := 0; k < 2*MAXN*MAXN; k++ {
 * 		for i := 0; i < MAXN; i++ {
 * 			for j := 0; j < MAXN; j++ {
 * 				dp[k][i][j] = defaultState
 * 			}
 * 		}
 * 	}
 *
 * 	v := dfs(G, 0, 1, 2)
 * 	return v
 * }
 */

/*
 * // min-max 极大极小博弈
 *
 * func min(x, y int) int {
 * 	if x < y {
 * 		return x
 * 	}
 * 	return y
 * }
 *
 * func max(x, y int) int {
 * 	if x > y {
 * 		return x
 * 	}
 * 	return y
 * }
 *
 * func minMax(G [][]int, t, m, c int) (v int) {
 * 	if visit[t][m][c] {
 * 		return dp[t][m][c]
 * 	}
 * 	defer func() { dp[t][m][c], visit[t][m][c] = v, true }()
 * 	if t >= 2*n*n {
 * 		v = minMaxDraw
 * 		return
 * 	}
 * 	if m == 0 {
 * 		v = minMaxMouseWin
 * 		return
 * 	}
 * 	if m == c {
 * 		v = minMaxCatWin
 * 		return
 * 	}
 * 	if t%2 == 0 {
 * 		// mouseMove, 初始默认 v == catwin
 * 		v = minMaxCatWin
 * 		for _, n := range G[m] {
 * 			v = min(v, minMax(G, t+1, n, c))
 * 			if v == minMaxMouseWin {
 * 				break
 * 			}
 * 		}
 * 	} else {
 * 		// catMove, 初始默认 v == mousewin
 * 		v = minMaxMouseWin
 * 		for _, n := range G[c] {
 * 			if n == 0 {
 * 				continue
 * 			}
 * 			v = max(v, minMax(G, t+1, m, n))
 * 			if v == minMaxCatWin {
 * 				break
 * 			}
 * 		}
 * 	}
 * 	return
 * }
 *
 * func catMouseGame(G [][]int) int {
 * 	n = len(G)
 * 	for k := 0; k < 2*MAXN*MAXN; k++ {
 * 		for i := 0; i < MAXN; i++ {
 * 			for j := 0; j < MAXN; j++ {
 * 				visit[k][i][j] = false
 * 			}
 * 		}
 * 	}
 * 	switch minMax(G, 0, 1, 2) {
 * 	case minMaxMouseWin:
 * 		return mouseWin
 * 	case minMaxCatWin:
 * 		return catWin
 * 	default:
 * 		return draw
 * 	}
 * }
 */

type State struct{ m, c, t int }

func prevStatus(G [][]int, s State) []State {
	var ans []State
	switch s.t {
	case mouseTurn:
		for _, n := range G[s.c] {
			// NOTE: 猫不能钻洞
			if n == 0 {
				continue
			}
			ans = append(ans, State{s.m, n, catTurn})
		}
	case catTurn:
		for _, n := range G[s.m] {
			ans = append(ans, State{n, s.c, mouseTurn})
		}
	}
	return ans
}

func catMouseGame(G [][]int) int {
	n = len(G)
	for i := 0; i < MAXN; i++ {
		for j := 0; j < MAXN; j++ {
			dp[mouseTurn][i][j] = 0
			dp[catTurn][i][j] = 0
		}
	}
	var Q []State
	for i := 1; i < n; i++ {
		dp[mouseTurn][0][i] = mouseWin
		dp[catTurn][0][i] = mouseWin
		Q = append(Q, State{0, i, mouseTurn})
		Q = append(Q, State{0, i, catTurn})
	}
	for i := 1; i < n; i++ {
		dp[mouseTurn][i][i] = catWin
		dp[catTurn][i][i] = catWin
		Q = append(Q, State{i, i, mouseTurn})
		Q = append(Q, State{i, i, catTurn})
	}
	for len(Q) > 0 {
		q := Q[0]
		Q = Q[1:]
		ps := prevStatus(G, q)
		for _, p := range ps {
			if dp[p.t][p.m][p.c] != draw {
				continue
			}
			if (dp[q.t][q.m][q.c] == catWin && p.t == catTurn) || (dp[q.t][q.m][q.c] == mouseWin && p.t == mouseTurn) {
				s := catWin
				if p.t == mouseTurn {
					s = mouseWin
				}
				dp[p.t][p.m][p.c] = s
				Q = append(Q, p)
				continue
			}
			allLose := true
			var s int
			switch p.t {
			case mouseTurn:
				s = catWin
				for _, n := range G[p.m] {
					if dp[q.t][n][p.c] != s {
						allLose = false
						break
					}
				}
			case catTurn:
				s = mouseWin
				for _, n := range G[p.c] {
					// NOTE: 猫不能钻洞
					if n == 0 {
						continue
					}
					if dp[q.t][p.m][n] != s {
						allLose = false
						break
					}
				}
			}
			if allLose {
				dp[p.t][p.m][p.c] = s
				Q = append(Q, p)
			}
		}
	}
	return dp[0][1][2]
}
