/*
 * @lc app=leetcode.cn id=count-integers-in-intervals lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [2276] 统计区间中的整数数目
 *
 * https://leetcode-cn.com/problems/count-integers-in-intervals/description/
 *
 * algorithms
 * Hard (26.68%)
 * Total Accepted:    2.7K
 * Total Submissions: 10.3K
 * Testcase Example:  '["CountIntervals","add","add","count","add","count"]\n' +
  '[[],[2,3],[7,10],[],[5,8],[]]'
 *
 * 给你区间的 空 集，请你设计并实现满足要求的数据结构：
 *
 *
 * 新增：添加一个区间到这个区间集合中。
 * 统计：计算出现在 至少一个 区间中的整数个数。
 *
 *
 * 实现 CountIntervals 类：
 *
 *
 * CountIntervals() 使用区间的空集初始化对象
 * void add(int left, int right) 添加区间 [left, right] 到区间集合之中。
 * int count() 返回出现在 至少一个 区间中的整数个数。
 *
 *
 * 注意：区间 [left, right] 表示满足 left <= x <= right 的所有整数 x 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入
 * ["CountIntervals", "add", "add", "count", "add", "count"]
 * [[], [2, 3], [7, 10], [], [5, 8], []]
 * 输出
 * [null, null, null, 6, null, 8]
 *
 * 解释
 * CountIntervals countIntervals = new CountIntervals(); // 用一个区间空集初始化对象
 * countIntervals.add(2, 3);  // 将 [2, 3] 添加到区间集合中
 * countIntervals.add(7, 10); // 将 [7, 10] 添加到区间集合中
 * countIntervals.count();    // 返回 6
 * ⁠                          // 整数 2 和 3 出现在区间 [2, 3] 中
 * ⁠                          // 整数 7、8、9、10 出现在区间 [7, 10] 中
 * countIntervals.add(5, 8);  // 将 [5, 8] 添加到区间集合中
 * countIntervals.count();    // 返回 8
 * ⁠                          // 整数 2 和 3 出现在区间 [2, 3] 中
 * ⁠                          // 整数 5 和 6 出现在区间 [5, 8] 中
 * ⁠                          // 整数 7 和 8 出现在区间 [5, 8] 和区间 [7, 10] 中
 * ⁠                          // 整数 9 和 10 出现在区间 [7, 10] 中
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= left <= right <= 10^9
 * 最多调用  add 和 count 方法 总计 10^5 次
 * 调用 count 方法至少一次
 *
 *
*/

// 参考: https://leetcode.cn/problems/count-integers-in-intervals/solution/by-endlesscheng-clk2/
//
// 线段树模版题, 背吧...
const maxn int = 1e9

type CountIntervals struct {
	Left, Right *CountIntervals
	L, R, Cnt   int
}

func Constructor() CountIntervals {
	ci := NewCountIntervals(1, maxn+1)
	return *ci
}

func NewCountIntervals(left, right int) *CountIntervals {
	return &CountIntervals{
		L: left,
		R: right,
	}
}

func (ci *CountIntervals) Add(left int, right int) {
	if ci.Cnt == ci.R-ci.L {
		return
	}
	if left <= ci.L && ci.R <= right+1 {
		ci.Cnt = ci.R - ci.L
		return
	}
	mid := (ci.L + ci.R) >> 1
	if ci.Left == nil {
		ci.Left = NewCountIntervals(ci.L, mid)
	}
	if ci.Right == nil {
		ci.Right = NewCountIntervals(mid, ci.R)
	}
	if left < mid {
		ci.Left.Add(left, right)
	}
	if mid <= right {
		ci.Right.Add(left, right)
	}
	ci.Cnt = ci.Left.Cnt + ci.Right.Cnt
	return
}

func (ci *CountIntervals) Count() int {
	return ci.Cnt
}

/**
 * Your CountIntervals object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(left,right);
 * param_2 := obj.Count();
 */
