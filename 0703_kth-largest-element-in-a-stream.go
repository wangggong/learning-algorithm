
import (
	"container/heap"
	"sort"
)

/*
 * @lc app=leetcode.cn id=kth-largest-element-in-a-stream lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [703] 数据流中的第 K 大元素
 *
 * https://leetcode-cn.com/problems/kth-largest-element-in-a-stream/description/
 *
 * algorithms
 * Easy (51.58%)
 * Total Accepted:    66.5K
 * Total Submissions: 129K
 * Testcase Example:  '["KthLargest","add","add","add","add","add"]\n' +
  '[[3,[4,5,8,2]],[3],[5],[10],[9],[4]]'
 *
 * 设计一个找到数据流中第 k 大元素的类（class）。注意是排序后的第 k 大元素，不是第 k 个不同的元素。
 *
 * 请实现 KthLargest 类：
 *
 *
 * KthLargest(int k, int[] nums) 使用整数 k 和整数流 nums 初始化对象。
 * int add(int val) 将 val 插入数据流 nums 后，返回当前数据流中第 k 大的元素。
 *
 *
 *
 *
 * 示例：
 *
 *
 * 输入：
 * ["KthLargest", "add", "add", "add", "add", "add"]
 * [[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
 * 输出：
 * [null, 4, 5, 5, 8, 8]
 *
 * 解释：
 * KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
 * kthLargest.add(3);   // return 4
 * kthLargest.add(5);   // return 5
 * kthLargest.add(10);  // return 5
 * kthLargest.add(9);   // return 8
 * kthLargest.add(4);   // return 8
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= k <= 1e4
 * 0 <= nums.length <= 1e4
 * -1e4 <= nums[i] <= 1e4
 * -1e4 <= val <= 1e4
 * 最多调用 add 方法 1e4 次
 * 题目数据保证，在查找第 k 大元素时，数组中至少有 k 个元素
 *
 *
*/
type Heap struct {
	sort.IntSlice
}

func (h *Heap) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *Heap) Pop() interface{} {
	l := len(h.IntSlice)
	v := h.IntSlice[l-1]
	h.IntSlice = h.IntSlice[:l-1]
	return v
}

type KthLargest struct {
	Heap
	Size int
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{Size: k, Heap: Heap{}}
	for _, n := range nums {
		_ = kl.Add(n)
	}
	return kl
}

func (kl *KthLargest) Add(val int) int {
	if kl.Len() < kl.Size || kl.IntSlice[0] < val {
		heap.Push(kl, val)
	}
	if kl.Len() > kl.Size {
		heap.Pop(kl)
	}
	if kl.Len() < 0 {
		return -1
	}
	return kl.IntSlice[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
