/*
 * @lc app=leetcode.cn id=booking-concert-tickets-in-groups lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [10011] 以组为单位订音乐会的门票
 *
 * https://leetcode.cn/problems/booking-concert-tickets-in-groups/description/
 *
 * algorithms
 * Hard (12.86%)
 * Total Accepted:    1.6K
 * Total Submissions: 8.9K
 * Testcase Example:  '["BookMyShow","gather","gather","scatter","scatter"]\n' +
  '[[2,5],[4,0],[2,0],[5,1],[5,1]]'
 *
 * 一个音乐会总共有 n 排座位，编号从 0 到 n - 1 ，每一排有 m 个座椅，编号为 0 到 m - 1
 * 。你需要设计一个买票系统，针对以下情况进行座位安排：
 *
 *
 * 同一组的 k 位观众坐在 同一排座位，且座位连续 。
 * k 位观众中 每一位 都有座位坐，但他们 不一定 坐在一起。
 *
 *
 * 由于观众非常挑剔，所以：
 *
 *
 * 只有当一个组里所有成员座位的排数都 小于等于 maxRow ，这个组才能订座位。每一组的 maxRow 可能 不同 。
 * 如果有多排座位可以选择，优先选择 最小 的排数。如果同一排中有多个座位可以坐，优先选择号码 最小 的。
 *
 *
 * 请你实现 BookMyShow 类：
 *
 *
 * BookMyShow(int n, int m) ，初始化对象，n 是排数，m 是每一排的座位数。
 * int[] gather(int k, int maxRow) 返回长度为 2 的数组，表示 k 个成员中 第一个座位 的排数和座位编号，这 k
 * 位成员必须坐在 同一排座位，且座位连续 。换言之，返回最小可能的 r 和 c 满足第 r 排中 [c, c + k - 1] 的座位都是空的，且 r
 * <= maxRow 。如果 无法 安排座位，返回 [] 。
 * boolean scatter(int k, int maxRow) 如果组里所有 k 个成员 不一定 要坐在一起的前提下，都能在第 0 排到第
 * maxRow 排之间找到座位，那么请返回 true 。这种情况下，每个成员都优先找排数 最小 ，然后是座位编号最小的座位。如果不能安排所有 k
 * 个成员的座位，请返回 false 。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：
 * ["BookMyShow", "gather", "gather", "scatter", "scatter"]
 * [[2, 5], [4, 0], [2, 0], [5, 1], [5, 1]]
 * 输出：
 * [null, [0, 0], [], true, false]
 *
 * 解释：
 * BookMyShow bms = new BookMyShow(2, 5); // 总共有 2 排，每排 5 个座位。
 * bms.gather(4, 0); // 返回 [0, 0]
 * ⁠                 // 这一组安排第 0 排 [0, 3] 的座位。
 * bms.gather(2, 0); // 返回 []
 * ⁠                 // 第 0 排只剩下 1 个座位。
 * ⁠                 // 所以无法安排 2 个连续座位。
 * bms.scatter(5, 1); // 返回 True
 * ⁠                  // 这一组安排第 0 排第 4 个座位和第 1 排 [0, 3] 的座位。
 * bms.scatter(5, 1); // 返回 False
 * ⁠                  // 总共只剩下 2 个座位。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 5 * 10^4
 * 1 <= m, k <= 10^9
 * 0 <= maxRow <= n - 1
 * gather 和 scatter 总 调用次数不超过 5 * 10^4 次。
 *
 *
*/

/*
 * // 参考:
 * // - https://leetcode.cn/problems/booking-concert-tickets-in-groups/solution/by-endlesscheng-okcu/
 * // - https://www.bilibili.com/video/BV18t4y1p736
 * //
 * // 线段树二分, 动态开点
 * const maxn int = 1e5
 *
 * type SegTreeNode struct {
 * 	L, R        int
 * 	Sum, Min    int
 * 	Left, Right *SegTreeNode
 * }
 *
 * func (s *SegTreeNode) buildChildren() {
 * 	mid := (s.L + s.R) >> 1
 * 	if s.Left == nil {
 * 		s.Left = NewSegTreeNode(s.L, mid)
 * 	}
 * 	if s.Right == nil {
 * 		s.Right = NewSegTreeNode(mid, s.R)
 * 	}
 * 	return
 * }
 *
 * func (s *SegTreeNode) Add(k, val int) {
 * 	// defer func() { fmt.Println(k, val, s) }()
 * 	if k < s.L || s.R <= k {
 * 		return
 * 	}
 * 	if s.L+1 == s.R {
 * 		s.Sum += val
 * 		s.Min += val
 * 		return
 * 	}
 * 	s.buildChildren()
 * 	s.Left.Add(k, val)
 * 	s.Right.Add(k, val)
 * 	s.Sum = s.Left.Sum + s.Right.Sum
 * 	s.Min = min(s.Left.Min, s.Right.Min)
 * 	return
 * }
 *
 * func (s *SegTreeNode) Query(l, r int) int {
 * 	if r <= s.L || s.R <= l {
 * 		return 0
 * 	}
 * 	if l <= s.L && s.R <= r {
 * 		return s.Sum
 * 	}
 * 	ans := 0
 * 	if s.Left != nil {
 * 		ans += s.Left.Query(l, r)
 * 	}
 * 	if s.Right != nil {
 * 		ans += s.Right.Query(l, r)
 * 	}
 * 	return ans
 * }
 *
 * func (s *SegTreeNode) Index(l, r, k int) int {
 * 	if r <= s.L || s.R <= l {
 * 		return 0
 * 	}
 * 	if s.L+1 == s.R {
 * 		return s.L
 * 	}
 * 	s.buildChildren()
 * 	if s.Left.Min <= k {
 * 		return s.Left.Index(l, r, k)
 * 	}
 * 	if s.Right.Min <= k {
 * 		return s.Right.Index(l, r, k)
 * 	}
 * 	return 0
 * }
 *
 * func NewSegTreeNode(l, r int) *SegTreeNode {
 * 	return &SegTreeNode{
 * 		L: l,
 * 		R: r,
 * 	}
 * }
 *
 * type BookMyShow struct {
 * 	*SegTreeNode
 * 	N, M int
 * }
 *
 * func Constructor(n int, m int) BookMyShow {
 * 	return BookMyShow{
 * 		SegTreeNode: NewSegTreeNode(1, maxn+1),
 * 		N:           n,
 * 		M:           m,
 * 	}
 * }
 */

// 使用数组模拟线段树.
const (
	_sum = iota
	_min
)

type SegTree [][2]int

// add 单点更新
func (s SegTree) add(o, l, r, k, val int) {
	// assert l < r
	if l+1 == r {
		s[o][_sum] += val
		s[o][_min] += val
		return
	}
	mid := (l + r) >> 1
	left, right := o<<1, (o<<1)|1
	if k < mid {
		s.add(left, l, mid, k, val)
	} else {
		s.add(right, mid, r, k, val)
	}
	s[o][_sum] = s[left][_sum] + s[right][_sum]
	s[o][_min] = min(s[left][_min], s[right][_min])
	return
}

// query 查询区间和
func (s SegTree) query(o, l, r, L, R int) int {
	// assert l < r && L < R
	if L <= l && r <= R {
		return s[o][_sum]
	}
	mid := (l + r) >> 1
	left, right := o<<1, (o<<1)|1
	ans := 0
	if L < mid {
		ans += s.query(left, l, mid, L, R)
	}
	if mid < R {
		ans += s.query(right, mid, r, L, R)
	}
	return ans
}

// index 查询区间最小值, 返回 `[L, R)` 区间中 `sum <= val` 的最小 index.
func (s SegTree) index(o, l, r, L, R, val int) int {
	// assert l < r && L < R
	if l+1 == r {
		if s[o][_min] <= val {
			return l
		}
		return 0
	}
	mid := (l + r) >> 1
	left, right := o<<1, (o<<1)|1
	if L < mid && s[left][_min] <= val {
		return s.index(left, l, mid, L, R, val)
	}
	if mid < R && s[right][_min] <= val {
		return s.index(right, mid, r, L, R, val)
	}
	return 0
}

func (s SegTree) Add(k, val int) {
	n := len(s) >> 2
	s.add(1, 1, n+1, k, val)
	return
}

func (s SegTree) Query(l, r int) int {
	n := len(s) >> 2
	return s.query(1, 1, n+1, l, r)
}

func (s SegTree) Index(l, r, val int) int {
	n := len(s) >> 2
	return s.index(1, 1, n+1, l, r, val)
}

func NewSegTree(n int) *SegTree {
	arr := make([][2]int, n<<2)
	tr := SegTree(arr)
	return &tr
}

type BookMyShow struct {
	*SegTree
	N, M int
}

func Constructor(n int, m int) BookMyShow {
	return BookMyShow{
		SegTree: NewSegTree(n),
		N:       n,
		M:       m,
	}
}

func (b *BookMyShow) Gather(k int, maxRow int) []int {
	maxRow++
	i := b.Index(1, maxRow+1, b.M-k)
	if i == 0 {
		return nil
	}
	seats := b.Query(i, i+1)
	b.Add(i, k)
	return []int{i - 1, seats}
}

func (b *BookMyShow) Scatter(k int, maxRow int) bool {
	maxRow++
	tot := b.M*(maxRow) - b.Query(1, maxRow+1)
	if tot < k {
		return false
	}
	for i := b.Index(1, maxRow+1, b.M-1); i <= maxRow && k > 0; i++ {
		seats := b.Query(i, i+1)
		add := min(k, b.M-seats)
		b.Add(i, add)
		k -= add
	}
	return true
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/**
 * Your BookMyShow object will be instantiated and called as such:
 * obj := Constructor(n, m);
 * param_1 := obj.Gather(k,maxRow);
 * param_2 := obj.Scatter(k,maxRow);
 */
