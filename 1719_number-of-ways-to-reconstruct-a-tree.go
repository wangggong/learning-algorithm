/*
 * @lc app=leetcode.cn id=number-of-ways-to-reconstruct-a-tree lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1719] 重构一棵树的方案数
 *
 * https://leetcode-cn.com/problems/number-of-ways-to-reconstruct-a-tree/description/
 *
 * algorithms
 * Hard (38.37%)
 * Total Accepted:    4.2K
 * Total Submissions: 6.2K
 * Testcase Example:  '[[1,2],[2,3]]'
 *
 * 给你一个数组 pairs ，其中 pairs[i] = [xi, yi] ，并且满足：
 *
 *
 * pairs 中没有重复元素
 * xi < yi
 *
 *
 * 令 ways 为满足下面条件的有根树的方案数：
 *
 *
 * 树所包含的所有节点值都在 pairs 中。
 * 一个数对 [xi, yi] 出现在 pairs 中 当且仅当 xi 是 yi 的祖先或者 yi 是 xi 的祖先。
 * 注意：构造出来的树不一定是二叉树。
 *
 *
 * 两棵树被视为不同的方案当存在至少一个节点在两棵树中有不同的父节点。
 *
 * 请你返回：
 *
 *
 * 如果 ways == 0 ，返回 0 。
 * 如果 ways == 1 ，返回 1 。
 * 如果 ways > 1 ，返回 2 。
 *
 *
 * 一棵 有根树 指的是只有一个根节点的树，所有边都是从根往外的方向。
 *
 * 我们称从根到一个节点路径上的任意一个节点（除去节点本身）都是该节点的 祖先 。根节点没有祖先。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：pairs = [[1,2],[2,3]]
 * 输出：1
 * 解释：如上图所示，有且只有一个符合规定的有根树。
 *
 *
 * 示例 2：
 *
 *
 * 输入：pairs = [[1,2],[2,3],[1,3]]
 * 输出：2
 * 解释：有多个符合规定的有根树，其中三个如上图所示。
 *
 *
 * 示例 3：
 *
 *
 * 输入：pairs = [[1,2],[2,3],[2,4],[1,5]]
 * 输出：0
 * 解释：没有符合规定的有根树。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= pairs.length <= 1e5
 * 1 <= xi < yi <= 500
 * pairs 中的元素互不相同。
 *
 *
 */
import "sort"

const MAXN = 500 + 10

type IndexCount struct {
	Index, Count int
}

var G [MAXN][MAXN]bool
var fa [MAXN]int

func checkWays(pairs [][]int) int {
	n := len(pairs)
	// assert n > 0 && n <= 1e5;
	m := make(map[int]int)
	for i := 0; i < MAXN; i++ {
		fa[i] = 0
		for j := 0; j < MAXN; j++ {
			G[i][j] = false
		}
	}
	// 统计出度并逆序, 寻找可能的根...
	for _, p := range pairs {
		s, t := p[0], p[1]
		m[s]++
		m[t]++
		G[s][t] = true
		G[t][s] = true
	}
	var indexCounts []IndexCount
	for index, count := range m {
		indexCounts = append(indexCounts, IndexCount{index, count})
	}
	sort.Slice(indexCounts, func(p, q int) bool { return indexCounts[p].Count > indexCounts[q].Count })
	// 1. 判断这是不是个森林...
	// 节点数为 n 的关系下限是 n-1 (一个根连下面所有其他节点), 比这个小的都是森林了.
	if n+1 < len(indexCounts) {
		//fmt.Printf("1. %v %v %v\n", indexCounts[0], n, 0)
		return 0
	}
	// 2. 试图去构造树
	for i, ic := range indexCounts {
		if i == 0 {
			// 根结点没 father
			fa[ic.Index] = -1
			continue
		}
		// 其余节点, 贪心地寻找度数最接近的 fa...
		var j int
		for j = i - 1; j >= 0; j-- {
			s, t := ic.Index, indexCounts[j].Index
			if G[s][t] {
				fa[s] = t
				break
			}
		}
		// ... 找不到?
		if j < 0 {
			//fmt.Printf("2. %v %v %v %v\n", indexCounts, i, j, 0)
			return 0
		}
	}
	// 3. 最后查一波对数.
	ans := 1
	pairsCount := 0
	for _, ic := range indexCounts {
		s := ic.Index
		for t := fa[s]; t != -1; t = fa[t] {
			pairsCount++
			if m[t] == m[s] {
				ans = 2
			}
			if !G[s][t] {
				//fmt.Printf("3. %v %v %v %v\n", indexCounts, s, t, 0)
				return 0
			}
		}
	}
	if pairsCount < n {
		//fmt.Printf("4. %v %v %v\n", pairsCount, n, 0)
		return 0
	}
	//fmt.Printf("5. %v\n", ans)
	return ans
}
