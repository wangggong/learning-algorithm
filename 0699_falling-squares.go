/*
 * @lc app=leetcode.cn id=falling-squares lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [699] 掉落的方块
 *
 * https://leetcode-cn.com/problems/falling-squares/description/
 *
 * algorithms
 * Hard (49.74%)
 * Total Accepted:    9.5K
 * Total Submissions: 17.5K
 * Testcase Example:  '[[1,2],[2,3],[6,1]]'
 *
 * 在二维平面上的 x 轴上，放置着一些方块。
 *
 * 给你一个二维整数数组 positions ，其中 positions[i] = [lefti, sideLengthi] 表示：第 i 个方块边长为
 * sideLengthi ，其左侧边与 x 轴上坐标点 lefti 对齐。
 *
 * 每个方块都从一个比目前所有的落地方块更高的高度掉落而下。方块沿 y 轴负方向下落，直到着陆到 另一个正方形的顶边 或者是 x 轴上
 * 。一个方块仅仅是擦过另一个方块的左侧边或右侧边不算着陆。一旦着陆，它就会固定在原地，无法移动。
 *
 * 在每个方块掉落后，你必须记录目前所有已经落稳的 方块堆叠的最高高度 。
 *
 * 返回一个整数数组 ans ，其中 ans[i] 表示在第 i 块方块掉落后堆叠的最高高度。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：positions = [[1,2],[2,3],[6,1]]
 * 输出：[2,5,5]
 * 解释：
 * 第 1 个方块掉落后，最高的堆叠由方块 1 组成，堆叠的最高高度为 2 。
 * 第 2 个方块掉落后，最高的堆叠由方块 1 和 2 组成，堆叠的最高高度为 5 。
 * 第 3 个方块掉落后，最高的堆叠仍然由方块 1 和 2 组成，堆叠的最高高度为 5 。
 * 因此，返回 [2, 5, 5] 作为答案。
 *
 *
 * 示例 2：
 *
 *
 * 输入：positions = [[100,100],[200,100]]
 * 输出：[100,100]
 * 解释：
 * 第 1 个方块掉落后，最高的堆叠由方块 1 组成，堆叠的最高高度为 100 。
 * 第 2 个方块掉落后，最高的堆叠可以由方块 1 组成也可以由方块 2 组成，堆叠的最高高度为 100 。
 * 因此，返回 [100, 100] 作为答案。
 * 注意，方块 2 擦过方块 1 的右侧边，但不会算作在方块 1 上着陆。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= positions.length <= 1000
 * 1 <= lefti <= 10^8
 * 1 <= sideLengthi <= 10^6
 *
 *
 */

// 线段树模版题+1
//
// 参考: https://leetcode.cn/problems/falling-squares/solution/-by-yu-niang-niang-14yb/
const maxn int = 1e8

var maxVal int

type SegTreeNode struct {
	L, R, Val   int
	Left, Right *SegTreeNode
}

func NewSegTreeNode(l, r int) *SegTreeNode {
	return &SegTreeNode{L: l, R: r}
}

func (s *SegTreeNode) Add(l, r, h int) {
	if s.R <= l || r <= s.L {
		return
	}
	cur := s.GetMax(l, r) + h
	maxVal = max(maxVal, cur)
	s.Update(l, r, cur)
	return
}

func (s *SegTreeNode) GetMax(l, r int) int {
	if s.R <= l || r <= s.L {
		return 0
	}
	ans := 0
	if s.L <= l && r <= s.R {
		ans = s.Val
	}
	mid := (s.L + s.R) >> 1
	if s.Left != nil && l < mid {
		ans = max(ans, s.Left.GetMax(l, min(mid, r)))
	}
	if s.Right != nil && mid < r {
		ans = max(ans, s.Right.GetMax(max(l, mid), r))
	}
	return ans
}

func (s *SegTreeNode) Update(l, r, h int) {
	if s.R <= l || r <= s.L {
		return
	}
	if l <= s.L && s.R <= r {
		s.Val = h
		return
	}
	mid := (s.L + s.R) >> 1
	if s.Left == nil {
		s.Left = NewSegTreeNode(s.L, mid)
	}
	if s.Right == nil {
		s.Right = NewSegTreeNode(mid, s.R)
	}
	if l < mid {
		s.Left.Update(l, min(mid, r), h)
	}
	if mid < r {
		s.Right.Update(max(l, mid), r, h)
	}
	return
}

func fallingSquares(P [][]int) (ans []int) {
	maxVal = 0
	root := NewSegTreeNode(1, 1e8+1)
	for _, p := range P {
		root.Add(p[0], p[0]+p[1], p[1])
		ans = append(ans, maxVal)
	}
	return
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
