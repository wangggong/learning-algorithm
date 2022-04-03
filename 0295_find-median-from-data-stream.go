import "container/heap"

/*
 * @lc app=leetcode.cn id=find-median-from-data-stream lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [295] 数据流的中位数
 *
 * https://leetcode-cn.com/problems/find-median-from-data-stream/description/
 *
 * algorithms
 * Hard (52.36%)
 * Total Accepted:    73.2K
 * Total Submissions: 139.8K
 * Testcase Example:  '["MedianFinder","addNum","addNum","findMedian","addNum","findMedian"]\n' +
  '[[],[1],[2],[],[3],[]]'
 *
 * 中位数是有序列表中间的数。如果列表长度是偶数，中位数则是中间两个数的平均值。
 *
 * 例如，
 *
 * [2,3,4] 的中位数是 3
 *
 * [2,3] 的中位数是 (2 + 3) / 2 = 2.5
 *
 * 设计一个支持以下两种操作的数据结构：
 *
 *
 * void addNum(int num) - 从数据流中添加一个整数到数据结构中。
 * double findMedian() - 返回目前所有元素的中位数。
 *
 *
 * 示例：
 *
 * addNum(1)
 * addNum(2)
 * findMedian() -> 1.5
 * addNum(3)
 * findMedian() -> 2
 *
 * 进阶:
 *
 *
 * 如果数据流中所有整数都在 0 到 100 范围内，你将如何优化你的算法？
 * 如果数据流中 99% 的整数都在 0 到 100 范围内，你将如何优化你的算法？
 *
 *
*/

type Heap struct {
	Nums []int
}

func (h Heap) Len() int            { return len(h.Nums) }
func (h Heap) Less(p, q int) bool  { return h.Nums[p] > h.Nums[q] }
func (h Heap) Swap(p, q int)       { h.Nums[p], h.Nums[q] = h.Nums[q], h.Nums[p] }
func (h *Heap) Push(v interface{}) { h.Nums = append(h.Nums, v.(int)) }
func (h *Heap) Pop() interface{}   { l := len(h.Nums); v := h.Nums[l-1]; h.Nums = h.Nums[:l-1]; return v }

type MedianFinder struct {
	Size               int
	BigRoot, SmallRoot Heap
}

func Constructor() MedianFinder {
	return MedianFinder{
		BigRoot:   Heap{},
		SmallRoot: Heap{},
	}
}

func (mf *MedianFinder) AddNum(num int) {
	mf.Size++
	br, sr := &mf.BigRoot, &mf.SmallRoot
	heap.Push(br, num)
	for br.Len() > 0 && sr.Len() > 0 && br.Nums[0] > -sr.Nums[0] {
		b, s := heap.Pop(br).(int), heap.Pop(sr).(int)
		heap.Push(br, -s)
		heap.Push(sr, -b)
	}
	if br.Len() > sr.Len()+1 {
		heap.Push(sr, -heap.Pop(br).(int))
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	br, sr := &mf.BigRoot, &mf.SmallRoot
	// fmt.Printf("%v %v\n", br.Nums, sr.Nums)
	if mf.Size%2 > 0 {
		return float64(br.Nums[0])
	}
	return float64(br.Nums[0]-sr.Nums[0]) / 2
}

/**
 * NOTE:
 * 1. heap.Push & heap.Pop, 类型本身的 push 和 pop 方法不合理!
 * 2. heap 是 "堆", 本质是二叉树; 实际实现是按照优先队列处理, 不按栈模型考虑;
 * 3. container/heap 里面的堆优先级最高的是 0 号元素! 坑了一宿了!
 */

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
