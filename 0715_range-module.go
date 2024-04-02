/*
 * @lc app=leetcode.cn id=range-module lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [715] Range 模块
 *
 * https://leetcode.cn/problems/range-module/description/
 *
 * algorithms
 * Hard (44.12%)
 * Total Accepted:    8.6K
 * Total Submissions: 16.8K
 * Testcase Example:  '["RangeModule","addRange","removeRange","queryRange","queryRange","queryRange"]\n' +
  '[[],[10,20],[14,16],[10,14],[13,15],[16,17]]'
 *
 * Range模块是跟踪数字范围的模块。设计一个数据结构来跟踪表示为 半开区间 的范围并查询它们。
 *
 * 半开区间 [left, right) 表示所有 left <= x < right 的实数 x 。
 *
 * 实现 RangeModule 类:
 *
 *
 * RangeModule() 初始化数据结构的对象。
 * void addRange(int left, int right) 添加 半开区间 [left,
 * right)，跟踪该区间中的每个实数。添加与当前跟踪的数字部分重叠的区间时，应当添加在区间 [left, right)
 * 中尚未跟踪的任何数字到该区间中。
 * boolean queryRange(int left, int right) 只有在当前正在跟踪区间 [left, right)
 * 中的每一个实数时，才返回 true ，否则返回 false 。
 * void removeRange(int left, int right) 停止跟踪 半开区间 [left, right)
 * 中当前正在跟踪的每个实数。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入
 * ["RangeModule", "addRange", "removeRange", "queryRange", "queryRange",
 * "queryRange"]
 * [[], [10, 20], [14, 16], [10, 14], [13, 15], [16, 17]]
 * 输出
 * [null, null, null, true, false, true]
 *
 * 解释
 * RangeModule rangeModule = new RangeModule();
 * rangeModule.addRange(10, 20);
 * rangeModule.removeRange(14, 16);
 * rangeModule.queryRange(10, 14); 返回 true （区间 [10, 14) 中的每个数都正在被跟踪）
 * rangeModule.queryRange(13, 15); 返回 false（未跟踪区间 [13, 15) 中像 14, 14.03, 14.17
 * 这样的数字）
 * rangeModule.queryRange(16, 17); 返回 true （尽管执行了删除操作，区间 [16, 17) 中的数字 16
 * 仍然会被跟踪）
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= left < right <= 10^9
 * 在单个测试用例中，对 addRange 、  queryRange 和 removeRange 的调用总数不超过 10^4 次
 *
 *
*/

// 线段树模版++
// 希望啥时候能自己写出来吧...
const maxn int = 1e9

const (
	removeLazy = iota - 1
	noLazy
	addLazy
)

type SegTreeNode struct {
	L, R, Add, Cnt int
	Left, Right    *SegTreeNode
}

func NewSegTreeNode(l, r int) *SegTreeNode {
	return &SegTreeNode{
		L: l,
		R: r,
	}
}

type RangeModule = SegTreeNode

func Constructor() RangeModule {
	node := NewSegTreeNode(1, maxn+1)
	return *node
}

func (rm *RangeModule) pushdown() {
	mid := (rm.L + rm.R) >> 1
	if rm.Left == nil {
		rm.Left = NewSegTreeNode(rm.L, mid)
	}
	if rm.Right == nil {
		rm.Right = NewSegTreeNode(mid, rm.R)
	}
	if rm.Add == noLazy {
		return
	}
	if rm.Add == addLazy {
		rm.Left.Cnt = mid - rm.L
		rm.Right.Cnt = rm.R - mid
	} else {
		rm.Left.Cnt = 0
		rm.Right.Cnt = 0
	}
	rm.Left.Add, rm.Right.Add = rm.Add, rm.Add
	rm.Add = noLazy
	return
}

func (rm *RangeModule) update(l, r, add int) {
	if r <= rm.L || rm.R <= l {
		return
	}
	if l <= rm.L && rm.R <= r {
		if add == addLazy {
			rm.Cnt = rm.R - rm.L
		} else {
			rm.Cnt = 0
		}
		rm.Add = add
		return
	}
	rm.pushdown()
	rm.Left.update(l, r, add)
	rm.Right.update(l, r, add)
	// rm.pushup()
	rm.Cnt = rm.Left.Cnt + rm.Right.Cnt
	return
}

func (rm *RangeModule) query(l, r int) int {
	if r <= rm.L || rm.R <= l {
		return 0
	}
	if l <= rm.L && rm.R <= r {
		return rm.Cnt
	}
	rm.pushdown()
	return rm.Left.query(l, r) + rm.Right.query(l, r)
}

func (rm *RangeModule) AddRange(l, r int)        { rm.update(l, r, addLazy) }
func (rm *RangeModule) QueryRange(l, r int) bool { return rm.query(l, r) == r-l }
func (rm *RangeModule) RemoveRange(l, r int)     { rm.update(l, r, removeLazy) }

/**
 * Your RangeModule object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddRange(left,right);
 * param_2 := obj.QueryRange(left,right);
 * obj.RemoveRange(left,right);
 */
