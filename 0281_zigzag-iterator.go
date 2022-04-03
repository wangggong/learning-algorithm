/*
 * Medium
 *
 * 给出两个一维的向量，请你实现一个迭代器，交替返回它们中间的元素。
 *
 * 示例:
 *
 * 输入:
 * v1 = [1,2]
 * v2 = [3,4,5,6]
 *
 * 输出: [1,3,2,4,5,6]
 *
 * 解析: 通过连续调用 next 函数直到 hasNext 函数返回 false，
 *      next 函数返回值的次序应依次为: [1,3,2,4,5,6]。
 * 拓展：假如给你 k 个一维向量呢？你的代码在这种情况下的扩展性又会如何呢?
 *
 * 拓展声明：
 *  “锯齿” 顺序对于 k > 2 的情况定义可能会有些歧义。所以，假如你觉得 “锯齿” 这个表述不妥，也可以认为这是一种 “循环”。例如：
 *
 * 输入:
 * [1,2,3]
 * [4,5,6,7]
 * [8,9]
 *
 * 输出: [1,4,8,2,5,9,3,6,7].
 * 通过次数3,307
 * 提交次数4,344
 */

type ZigzagIterator struct {
	Arr1, Arr2 []int
	Pos1, Pos2 int
}

func Constructor(v1, v2 []int) *ZigzagIterator {
	return &ZigzagIterator{
		Arr1: v1,
		Arr2: v2,
	}
}

func (it *ZigzagIterator) next() int {
	if it.Pos1 == len(it.Arr1) {
		v := it.Arr2[it.Pos2]
		it.Pos2++
		return v
	}
	if it.Pos2 == len(it.Arr2) {
		v := it.Arr1[it.Pos1]
		it.Pos1++
		return v
	}
	if it.Pos1 == it.Pos2 {
		v := it.Arr1[it.Pos1]
		it.Pos1++
		return v
	}
	if it.Pos1 > it.Pos2 {
		v := it.Arr2[it.Pos2]
		it.Pos2++
		return v
	}
	// unreachable
	return -1
}

func (it *ZigzagIterator) hasNext() bool {
	return it.Pos1 < len(it.Arr1) || it.Pos2 < len(it.Arr2)
}

/**
 * Your ZigzagIterator object will be instantiated and called as such:
 * obj := Constructor(param_1, param_2);
 * for obj.hasNext() {
 *	 ans = append(ans, obj.next())
 * }
 */
