/*
 * @lc app=leetcode.cn id=peeking-iterator lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [284] 顶端迭代器
 *
 * https://leetcode-cn.com/problems/peeking-iterator/description/
 *
 * algorithms
 * Medium (76.98%)
 * Total Accepted:    26.9K
 * Total Submissions: 34.9K
 * Testcase Example:  '["PeekingIterator","next","peek","next","next","hasNext"]\n' +
  '[[[1,2,3]],[],[],[],[],[]]'
 *
 * 请你在设计一个迭代器，在集成现有迭代器拥有的 hasNext 和 next 操作的基础上，还额外支持 peek 操作。
 *
 * 实现 PeekingIterator 类：
 *
 *
 * PeekingIterator(Iterator<int> nums) 使用指定整数迭代器 nums 初始化迭代器。
 * int next() 返回数组中的下一个元素，并将指针移动到下个元素处。
 * bool hasNext() 如果数组中存在下一个元素，返回 true ；否则，返回 false 。
 * int peek() 返回数组中的下一个元素，但 不 移动指针。
 *
 *
 * 注意：每种语言可能有不同的构造函数和迭代器 Iterator，但均支持 int next() 和 boolean hasNext() 函数。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：
 * ["PeekingIterator", "next", "peek", "next", "next", "hasNext"]
 * [[[1, 2, 3]], [], [], [], [], []]
 * 输出：
 * [null, 1, 2, 2, 3, false]
 *
 * 解释：
 * PeekingIterator peekingIterator = new PeekingIterator([1, 2, 3]); // [1,2,3]
 * peekingIterator.next();    // 返回 1 ，指针移动到下一个元素 [1,2,3]
 * peekingIterator.peek();    // 返回 2 ，指针未发生移动 [1,2,3]
 * peekingIterator.next();    // 返回 2 ，指针移动到下一个元素 [1,2,3]
 * peekingIterator.next();    // 返回 3 ，指针移动到下一个元素 [1,2,3]
 * peekingIterator.hasNext(); // 返回 False
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 1000
 * 1 <= nums[i] <= 1000
 * 对 next 和 peek 的调用均有效
 * next、hasNext 和 peek 最多调用  1000 次
 *
 *
 *
 *
 * 进阶：你将如何拓展你的设计？使之变得通用化，从而适应所有的类型，而不只是整数型？
 *
*/
/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

/*
 * type PeekingIterator struct {
 * 	Iter   *Iterator
 * 	Val    int
 * 	Peeked bool
 * }
 *
 * func Constructor(iter *Iterator) *PeekingIterator {
 * 	return &PeekingIterator{Iter: iter, Peeked: true}
 * }
 *
 * func (pi *PeekingIterator) hasNext() bool {
 * 	if !pi.Peeked {
 * 		return true
 * 	}
 * 	return pi.Iter.hasNext()
 * }
 *
 * func (pi *PeekingIterator) next() int {
 * 	val := pi.peek()
 * 	pi.Peeked = true
 * 	return val
 * }
 *
 * func (pi *PeekingIterator) peek() int {
 * 	if !pi.hasNext() {
 * 		return -1
 * 	}
 * 	if pi.Peeked {
 * 		pi.Val = pi.Iter.next()
 * 		pi.Peeked = false
 * 	}
 * 	return pi.Val
 * }
 */

type PeekingIterator struct {
	*Iterator
	Next *int
}

func Constructor(iter *Iterator) *PeekingIterator {
	pi := PeekingIterator{Iterator: iter}
	if pi.Iterator.hasNext() {
		v := pi.Iterator.next()
		pi.Next = &v
	}
	return &pi
}

func (pi *PeekingIterator) hasNext() bool {
	return pi.Next != nil
}

func (pi *PeekingIterator) next() int {
	val := *pi.Next
	if pi.Iterator.hasNext() {
		v := pi.Iterator.next()
		pi.Next = &v
	} else {
		pi.Next = nil
	}
	return val
}

func (pi *PeekingIterator) peek() int {
	return *pi.Next
}
