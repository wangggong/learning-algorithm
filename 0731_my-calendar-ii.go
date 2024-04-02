/*
 * @lc app=leetcode.cn id=my-calendar-ii lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [731] 我的日程安排表 II
 *
 * https://leetcode.cn/problems/my-calendar-ii/description/
 *
 * algorithms
 * Medium (54.67%)
 * Total Accepted:    23.7K
 * Total Submissions: 37.4K
 * Testcase Example:  '["MyCalendarTwo","book","book","book","book","book","book"]\n' +
  '[[],[10,20],[50,60],[10,40],[5,15],[5,10],[25,55]]'
 *
 * 实现一个 MyCalendar 类来存放你的日程安排。如果要添加的时间内不会导致三重预订时，则可以存储这个新的日程安排。
 *
 * MyCalendar 有一个 book(int start, int end)方法。它意味着在 start 到 end
 * 时间内增加一个日程安排，注意，这里的时间是半开区间，即 [start, end), 实数 x 的范围为，  start <= x < end。
 *
 * 当三个日程安排有一些时间上的交叉时（例如三个日程安排都在同一时间内），就会产生三重预订。
 *
 * 每次调用 MyCalendar.book方法时，如果可以将日程安排成功添加到日历中而不会导致三重预订，返回 true。否则，返回 false
 * 并且不要将该日程安排添加到日历中。
 *
 * 请按照以下步骤调用MyCalendar 类: MyCalendar cal = new MyCalendar();
 * MyCalendar.book(start, end)
 *
 *
 *
 * 示例：
 *
 * MyCalendar();
 * MyCalendar.book(10, 20); // returns true
 * MyCalendar.book(50, 60); // returns true
 * MyCalendar.book(10, 40); // returns true
 * MyCalendar.book(5, 15); // returns false
 * MyCalendar.book(5, 10); // returns true
 * MyCalendar.book(25, 55); // returns true
 * 解释：
 * 前两个日程安排可以添加至日历中。 第三个日程安排会导致双重预订，但可以添加至日历中。
 * 第四个日程安排活动（5,15）不能添加至日历中，因为它会导致三重预订。
 * 第五个日程安排（5,10）可以添加至日历中，因为它未使用已经双重预订的时间10。
 * 第六个日程安排（25,55）可以添加至日历中，因为时间 [25,40] 将和第三个日程安排双重预订；
 * 时间 [40,50] 将单独预订，时间 [50,55）将和第二个日程安排双重预订。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 每个测试用例，调用 MyCalendar.book 函数最多不超过 1000次。
 * 调用函数 MyCalendar.book(start, end)时， start 和 end 的取值范围为 [0, 10^9]。
 *
 *
*/

// 参考: https://leetcode.cn/problems/my-calendar-ii/solution/by-ac_oier-a1b3/
//
// 所以自己就是写不出来呗...
const maxn int = 1e9

type SegTree struct {
	L, R        int
	Max, Add    int
	Left, Right *SegTree
}

func (s *SegTree) pushdown() {
	mid := (s.L + s.R) >> 1
	if s.Left == nil {
		s.Left = NewSegTree(s.L, mid)
	}
	if s.Right == nil {
		s.Right = NewSegTree(mid, s.R)
	}
	// 注意这里不能直接按 `s.Left.Max = s.Add` 初始化呀傻逼!
	s.Left.Max += s.Add
	s.Right.Max += s.Add
	s.Left.Add += s.Add
	s.Right.Add += s.Add
	s.Add = 0
	return
}

func (s *SegTree) query(l, r int) (ans int) {
	if s.R <= l || r <= s.L {
		return
	}
	if l <= s.L && s.R <= r {
		ans = s.Max
		return
	}
	s.pushdown()
	ans = max(s.Left.query(l, r), s.Right.query(l, r))
	return
}

func (s *SegTree) add(l, r, v int) {
	if s.R <= l || r <= s.L {
		return
	}
	if l <= s.L && s.R <= r {
		// 这里 `s.Add += v` 是为了向子区间传播, 必不可少啊!
		s.Max += v
		s.Add += v
		return
	}
	s.pushdown()
	s.Left.add(l, r, v)
	s.Right.add(l, r, v)
	s.Max = max(s.Left.Max, s.Right.Max)
	return
}

func NewSegTree(l, r int) *SegTree {
	return &SegTree{
		L: l,
		R: r,
	}
}

type MyCalendarTwo struct{ *SegTree }

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{SegTree: NewSegTree(0, maxn+1)}
}

func (m *MyCalendarTwo) Book(l, r int) bool {
	if m.query(l, r) >= 2 {
		return false
	}
	m.add(l, r, 1)
	return true
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
