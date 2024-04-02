/*
 * @lc app=leetcode.cn id=my-calendar-iii lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [732] 我的日程安排表 III
 *
 * https://leetcode.cn/problems/my-calendar-iii/description/
 *
 * algorithms
 * Hard (63.41%)
 * Total Accepted:    8.4K
 * Total Submissions: 12.1K
 * Testcase Example:  '["MyCalendarThree","book","book","book","book","book","book"]\n' +
  '[[],[10,20],[50,60],[10,40],[5,15],[5,10],[25,55]]'
 *
 * 当 k 个日程安排有一些时间上的交叉时（例如 k 个日程安排都在同一时间内），就会产生 k 次预订。
 *
 * 给你一些日程安排 [start, end) ，请你在每个日程安排添加后，返回一个整数 k ，表示所有先前日程安排会产生的最大 k 次预订。
 *
 * 实现一个 MyCalendarThree 类来存放你的日程安排，你可以一直添加新的日程安排。
 *
 *
 * MyCalendarThree() 初始化对象。
 * int book(int start, int end) 返回一个整数 k ，表示日历中存在的 k 次预订的最大值。
 *
 *
 *
 *
 * 示例：
 *
 *
 * 输入：
 * ["MyCalendarThree", "book", "book", "book", "book", "book", "book"]
 * [[], [10, 20], [50, 60], [10, 40], [5, 15], [5, 10], [25, 55]]
 * 输出：
 * [null, 1, 1, 2, 3, 3, 3]
 *
 * 解释：
 * MyCalendarThree myCalendarThree = new MyCalendarThree();
 * myCalendarThree.book(10, 20); // 返回 1 ，第一个日程安排可以预订并且不存在相交，所以最大 k 次预订是 1 次预订。
 * myCalendarThree.book(50, 60); // 返回 1 ，第二个日程安排可以预订并且不存在相交，所以最大 k 次预订是 1 次预订。
 * myCalendarThree.book(10, 40); // 返回 2 ，第三个日程安排 [10, 40) 与第一个日程安排相交，所以最大 k
 * 次预订是 2 次预订。
 * myCalendarThree.book(5, 15); // 返回 3 ，剩下的日程安排的最大 k 次预订是 3 次预订。
 * myCalendarThree.book(5, 10); // 返回 3
 * myCalendarThree.book(25, 55); // 返回 3
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= start < end <= 1e9
 * 每个测试用例，调用 book 函数最多不超过 400次
 *
 *
*/

const maxn int = 1e9

type SegTreeNode struct {
	L, R, Max   int
	Left, Right *SegTreeNode
}

func (s *SegTreeNode) Add(l, r int) {
	if r <= s.L || s.R <= l {
		return
	}
	if l <= s.L && s.R <= r {
		s.Max++
		// 如果要动态开点的话别忘了把子区间也更新了...
		if s.Left != nil {
			s.Left.Add(l, r)
		}
		if s.Right != nil {
			s.Right.Add(l, r)
		}
		return
	}
	mid := (s.L + s.R) >> 1
	if s.Left == nil {
		s.Left = NewSegTreeNode(s.L, mid, s.Max)
	}
	if s.Right == nil {
		s.Right = NewSegTreeNode(mid, s.R, s.Max)
	}
	s.Left.Add(l, r)
	s.Right.Add(l, r)
	s.Max = max(s.Left.Max, s.Right.Max)
	return
}

func NewSegTreeNode(l, r, m int) *SegTreeNode {
	return &SegTreeNode{
		L:   l,
		R:   r,
		Max: m,
	}
}

type MyCalendarThree struct{ *SegTreeNode }

func Constructor() MyCalendarThree {
	return MyCalendarThree{SegTreeNode: NewSegTreeNode(0, maxn, 0)}
}

func (m *MyCalendarThree) Book(start int, end int) int {
	m.Add(start, end)
	return m.Max
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
 * // 差分数组的做法, 用红黑树模拟差分数组.
 * import "github.com/emirpasic/gods/trees/redblacktree"
 * 
 * type MyCalendarThree struct {
 * 	*redblacktree.Tree
 * }
 * 
 * func Constructor() MyCalendarThree { return MyCalendarThree{Tree: redblacktree.NewWithIntComparator()} }
 * 
 * func (m *MyCalendarThree) Add(x, delta int) {
 * 	if v, ok := m.Get(x); ok {
 * 		delta += v.(int)
 * 	}
 * 	m.Put(x, delta)
 * 	return
 * }
 * 
 * func (m *MyCalendarThree) Book(start int, end int) int {
 * 	m.Add(start, 1)
 * 	m.Add(end, -1)
 * 	ans := 0
 * 	cur := 0
 * 	for it := m.Iterator(); it.Next(); {
 * 		cur += it.Value().(int)
 * 		ans = max(ans, cur)
 * 	}
 * 	return ans
 * }
 */

/**
 * Your MyCalendarThree object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
