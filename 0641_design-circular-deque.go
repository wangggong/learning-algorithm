/*
 * @lc app=leetcode.cn id=design-circular-deque lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [641] 设计循环双端队列
 *
 * https://leetcode.cn/problems/design-circular-deque/description/
 *
 * algorithms
 * Medium (53.16%)
 * Total Accepted:    46K
 * Total Submissions: 81.1K
 * Testcase Example:  '["MyCircularDeque","insertLast","insertLast","insertFront","insertFront","getRear","isFull","deleteLast","insertFront","getFront"]\n' +
  '[[3],[1],[2],[3],[4],[],[],[],[4],[]]'
 *
 * 设计实现双端队列。
 *
 * 实现 MyCircularDeque 类:
 *
 *
 * MyCircularDeque(int k) ：构造函数,双端队列最大为 k 。
 * boolean insertFront()：将一个元素添加到双端队列头部。 如果操作成功返回 true ，否则返回 false 。
 * boolean insertLast() ：将一个元素添加到双端队列尾部。如果操作成功返回 true ，否则返回 false 。
 * boolean deleteFront() ：从双端队列头部删除一个元素。 如果操作成功返回 true ，否则返回 false 。
 * boolean deleteLast() ：从双端队列尾部删除一个元素。如果操作成功返回 true ，否则返回 false 。
 * int getFront() )：从双端队列头部获得一个元素。如果双端队列为空，返回 -1 。
 * int getRear() ：获得双端队列的最后一个元素。 如果双端队列为空，返回 -1 。
 * boolean isEmpty() ：若双端队列为空，则返回 true ，否则返回 false  。
 * boolean isFull() ：若双端队列满了，则返回 true ，否则返回 false 。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入
 * ["MyCircularDeque", "insertLast", "insertLast", "insertFront",
 * "insertFront", "getRear", "isFull", "deleteLast", "insertFront", "getFront"]
 * [[3], [1], [2], [3], [4], [], [], [], [4], []]
 * 输出
 * [null, true, true, true, false, 2, true, true, true, 4]
 *
 * 解释
 * MyCircularDeque circularDeque = new MycircularDeque(3); // 设置容量大小为3
 * circularDeque.insertLast(1);                    // 返回 true
 * circularDeque.insertLast(2);                    // 返回 true
 * circularDeque.insertFront(3);                    // 返回 true
 * circularDeque.insertFront(4);                    // 已经满了，返回 false
 * circularDeque.getRear();                  // 返回 2
 * circularDeque.isFull();                        // 返回 true
 * circularDeque.deleteLast();                    // 返回 true
 * circularDeque.insertFront(4);                    // 返回 true
 * circularDeque.getFront();                // 返回 4
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= k <= 1000
 * 0 <= value <= 1000
 * insertFront, insertLast, deleteFront, deleteLast, getFront, getRear,
 * isEmpty, isFull  调用次数不大于 2000 次
 *
 *
*/

type MyCircularDeque struct {
	Arr        []int
	N          int
	Head, Tail int
}

func Constructor(n int) MyCircularDeque {
	return MyCircularDeque{
		N:    n,
		Head: n,
		Tail: n,
		Arr:  make([]int, 3*n),
	}
}

func (mcd *MyCircularDeque) realloc() {
	arr := make([]int, 3*mcd.N)
	newHead, newTail := mcd.N, mcd.N+(mcd.Tail-mcd.Head)
	copy(arr[newHead:newTail], mcd.Arr[mcd.Head:mcd.Tail])
	mcd.Arr = arr
	mcd.Head, mcd.Tail = newHead, newTail
	return
}

func (mcd *MyCircularDeque) InsertFront(val int) bool {
	if mcd.IsFull() {
		return false
	}
	if mcd.Head == 0 {
		mcd.realloc()
	}
	mcd.Arr[mcd.Head-1] = val
	mcd.Head--
	return true
}

func (mcd *MyCircularDeque) InsertLast(val int) bool {
	if mcd.IsFull() {
		return false
	}
	if mcd.Tail == len(mcd.Arr) {
		mcd.realloc()
	}
	mcd.Arr[mcd.Tail] = val
	mcd.Tail++
	return true

}

func (mcd *MyCircularDeque) DeleteFront() bool {
	if mcd.IsEmpty() {
		return false
	}
	mcd.Head++
	return true
}

func (mcd *MyCircularDeque) DeleteLast() bool {
	if mcd.IsEmpty() {
		return false
	}
	mcd.Tail--
	return true
}

func (mcd *MyCircularDeque) GetFront() int {
	if mcd.IsEmpty() {
		return -1
	}
	return mcd.Arr[mcd.Head]
}

func (mcd *MyCircularDeque) GetRear() int {
	if mcd.IsEmpty() {
		return -1
	}
	return mcd.Arr[mcd.Tail-1]
}

func (mcd *MyCircularDeque) IsEmpty() bool {
	return mcd.Head == mcd.Tail
}

func (mcd *MyCircularDeque) IsFull() bool {
	return mcd.Tail-mcd.Head >= mcd.N
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
