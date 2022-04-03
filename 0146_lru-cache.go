/*
 * @lc app=leetcode.cn id=lru-cache lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [146] LRU 缓存
 *
 * https://leetcode-cn.com/problems/lru-cache/description/
 *
 * algorithms
 * Medium (52.46%)
 * Total Accepted:    272.6K
 * Total Submissions: 519.6K
 * Testcase Example:  '["LRUCache","put","put","get","put","get","put","get","get","get"]\n[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]'
 *
 * 请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
 *
 * 实现 LRUCache 类：
 *
 *
 *
 *
 * LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
 * int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
 * void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组
 * key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
 *
 *
 * 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
 *
 *
 *
 *
 *
 * 示例：
 *
 *
 * 输入
 * ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
 * [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
 * 输出
 * [null, null, null, 1, null, -1, null, -1, 3, 4]
 *
 * 解释
 * LRUCache lRUCache = new LRUCache(2);
 * lRUCache.put(1, 1); // 缓存是 {1=1}
 * lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
 * lRUCache.get(1);    // 返回 1
 * lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
 * lRUCache.get(2);    // 返回 -1 (未找到)
 * lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
 * lRUCache.get(1);    // 返回 -1 (未找到)
 * lRUCache.get(3);    // 返回 3
 * lRUCache.get(4);    // 返回 4
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= capacity <= 3000
 * 0 <= key <= 10000
 * 0 <= value <= 10^5
 * 最多调用 2 * 10^5 次 get 和 put
 *
 *
 */

type ListNode struct {
	Key, Val   int
	Prev, Next *ListNode
}

type LRUCache struct {
	Len, Cap   int
	m          map[int]*ListNode
	Head, Tail *ListNode
}

func Constructor(c int) LRUCache {
	C := LRUCache{
		Cap:  c,
		m:    make(map[int]*ListNode),
		Head: &ListNode{Key: -1, Val: -1},
		Tail: &ListNode{Key: -1, Val: -1},
	}
	C.Head.Next, C.Tail.Prev = C.Tail, C.Head
	return C
}

func (C *LRUCache) Get(k int) int {
	v, ok := C.m[k]
	if !ok {
		return -1
	}
	C.gotoHead(v)
	return v.Val
}

func (C *LRUCache) Put(k int, x int) {
	if v, ok := C.m[k]; ok {
		v.Val = x
		C.gotoHead(v)
		return
	}

	v := ListNode{Key: k, Val: x}
	C.insert(&v)
	if C.Len > C.Cap {
		C.removeTail()
	}
}

func (C *LRUCache) gotoHead(v *ListNode) {
	C.remove(v)
	C.insert(v)
}

func (C *LRUCache) removeTail() {
	C.remove(C.Tail.Prev)
}

func (C *LRUCache) remove(v *ListNode) {
	v.Prev.Next, v.Next.Prev = v.Next, v.Prev
	delete(C.m, v.Key)
	C.Len--
}

func (C *LRUCache) insert(v *ListNode) {
	C.Head.Next, C.Head.Next.Prev, v.Prev, v.Next = v, v, C.Head, C.Head.Next
	C.m[v.Key] = v
	C.Len++
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
