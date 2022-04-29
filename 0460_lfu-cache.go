/*
 * @lc app=leetcode.cn id=lfu-cache lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [460] LFU 缓存
 *
 * https://leetcode-cn.com/problems/lfu-cache/description/
 *
 * algorithms
 * Hard (43.93%)
 * Total Accepted:    46.5K
 * Total Submissions: 105.7K
 * Testcase Example:  '["LFUCache","put","put","get","put","get","get","put","get","get","get"]\n' +
  '[[2],[1,1],[2,2],[1],[3,3],[2],[3],[4,4],[1],[3],[4]]'
 *
 * 请你为 最不经常使用（LFU）缓存算法设计并实现数据结构。
 *
 * 实现 LFUCache 类：
 *
 *
 * LFUCache(int capacity) - 用数据结构的容量 capacity 初始化对象
 * int get(int key) - 如果键 key 存在于缓存中，则获取键的值，否则返回 -1 。
 * void put(int key, int value) - 如果键 key 已存在，则变更其值；如果键不存在，请插入键值对。当缓存达到其容量
 * capacity 时，则应该在插入新项之前，移除最不经常使用的项。在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，应该去除 最近最久未使用
 * 的键。
 *
 *
 * 为了确定最不常使用的键，可以为缓存中的每个键维护一个 使用计数器 。使用计数最小的键是最久未使用的键。
 *
 * 当一个键首次插入到缓存中时，它的使用计数器被设置为 1 (由于 put 操作)。对缓存中的键执行 get 或 put 操作，使用计数器的值将会递增。
 *
 * 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
 *
 *
 *
 * 示例：
 *
 *
 * 输入：
 * ["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get",
 * "get"]
 * [[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]
 * 输出：
 * [null, null, null, 1, null, -1, 3, null, -1, 3, 4]
 *
 * 解释：
 * // cnt(x) = 键 x 的使用计数
 * // cache=[] 将显示最后一次使用的顺序（最左边的元素是最近的）
 * LFUCache lfu = new LFUCache(2);
 * lfu.put(1, 1);   // cache=[1,_], cnt(1)=1
 * lfu.put(2, 2);   // cache=[2,1], cnt(2)=1, cnt(1)=1
 * lfu.get(1);      // 返回 1
 * ⁠                // cache=[1,2], cnt(2)=1, cnt(1)=2
 * lfu.put(3, 3);   // 去除键 2 ，因为 cnt(2)=1 ，使用计数最小
 * ⁠                // cache=[3,1], cnt(3)=1, cnt(1)=2
 * lfu.get(2);      // 返回 -1（未找到）
 * lfu.get(3);      // 返回 3
 * ⁠                // cache=[3,1], cnt(3)=2, cnt(1)=2
 * lfu.put(4, 4);   // 去除键 1 ，1 和 3 的 cnt 相同，但 1 最久未使用
 * ⁠                // cache=[4,3], cnt(4)=1, cnt(3)=2
 * lfu.get(1);      // 返回 -1（未找到）
 * lfu.get(3);      // 返回 3
 * ⁠                // cache=[3,4], cnt(4)=1, cnt(3)=3
 * lfu.get(4);      // 返回 4
 * ⁠                // cache=[3,4], cnt(4)=2, cnt(3)=3
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= capacity <= 10^4
 * 0 <= key <= 10^5
 * 0 <= value <= 10^9
 * 最多调用 2 * 10^5 次 get 和 put 方法
 *
 *
*/
type Node struct {
	Key, Val, Freq int
	Prev, Next     *Node
}

func NewNode(key, val int) *Node {
	return &Node{
		Key:  key,
		Val:  val,
		Freq: 1,
	}
}

type NodeList struct {
	Head, Tail *Node
	Prev, Next *NodeList
}

func (nl *NodeList) Empty() bool { return nl == nil || nl.Head == nil }

func (nl *NodeList) Insert(node *Node) {
	head := nl.Head
	nl.Head, node.Next, head.Prev = node, head, node
	return
}

func (nl *NodeList) Remove(node *Node) {
	defer func() { node.Prev, node.Next = nil, nil }()
	prev, next := node.Prev, node.Next
	if nl.Head == nl.Tail {
		nl.Head, nl.Tail = nil, nil
		return
	}
	if node == nl.Head {
		nl.Head, next.Prev = next, nil
		return
	}
	if node == nl.Tail {
		nl.Tail, prev.Next = prev, nil
		return
	}
	prev.Next, next.Prev = next, prev
	return
}

func NewNodeList(node *Node) *NodeList {
	return &NodeList{
		Head: node,
		Tail: node,
	}
}

type LFUCache struct {
	Len, Cap int
	Records  map[int]*Node
	Heads    map[*Node]*NodeList
	HeadList *NodeList
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		Cap:     capacity,
		Records: make(map[int]*Node),
		Heads:   make(map[*Node]*NodeList),
	}
}

func (lc *LFUCache) updateHeadList(nodeList *NodeList) bool {
	if !nodeList.Empty() {
		return false
	}
	prev, next := nodeList.Prev, nodeList.Next
	if lc.HeadList == nodeList {
		lc.HeadList = next
		if next != nil {
			next.Prev = nil
		}
		return true
	}
	prev.Next = next
	if next != nil {
		next.Prev = prev
	}
	return true
}

func (lc *LFUCache) move(node *Node, nodeList *NodeList) {
	nodeList.Remove(node)
	prev := nodeList
	if lc.updateHeadList(nodeList) {
		prev = nodeList.Prev
	}
	next := nodeList.Next
	if !next.Empty() && next.Head.Freq == node.Freq {
		next.Insert(node)
		lc.Heads[node] = next
		return
	}
	newList := NewNodeList(node)
	if prev != nil {
		prev.Next = newList
	}
	if next != nil {
		next.Prev = newList
	}
	newList.Prev, newList.Next = prev, next
	if lc.HeadList == nil || lc.HeadList == next {
		lc.HeadList = newList
	}
	lc.Heads[node] = newList
	return
}

func (lc *LFUCache) Get(key int) int {
	node, ok := lc.Records[key]
	if !ok {
		return -1
	}
	node.Freq++
	lc.move(node, lc.Heads[node])
	return node.Val
}

func (lc *LFUCache) Put(key int, val int) {
	if lc.Cap == 0 {
		return
	}
	node, ok := lc.Records[key]
	if ok {
		node.Freq++
		node.Val = val
		lc.move(node, lc.Heads[node])
		return
	}
	if lc.Len == lc.Cap {
		node := lc.HeadList.Tail
		lc.HeadList.Remove(node)
		lc.updateHeadList(lc.HeadList)
		delete(lc.Records, node.Key)
		delete(lc.Heads, node)
		lc.Len--
	}
	node = NewNode(key, val)
	if head := lc.HeadList; !head.Empty() && head.Head.Freq == node.Freq {
		head.Insert(node)
		lc.Records[key] = node
		lc.Heads[node] = head
		lc.Len++
		return
	}
	newList := NewNodeList(node)
	if head := lc.HeadList; head != nil {
		lc.HeadList, newList.Next, head.Prev = newList, head, newList
	} else {
		lc.HeadList = newList
	}
	lc.Records[key] = node
	lc.Heads[node] = lc.HeadList
	lc.Len++
	return
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
