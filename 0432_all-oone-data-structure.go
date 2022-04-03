import "fmt"

/*
 * @lc app=leetcode.cn id=all-oone-data-structure lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [432] 全 O(1) 的数据结构
 *
 * https://leetcode-cn.com/problems/all-oone-data-structure/description/
 *
 * algorithms
 * Hard (38.52%)
 * Total Accepted:    11.3K
 * Total Submissions: 26.7K
 * Testcase Example:  '["AllOne","inc","inc","getMaxKey","getMinKey","inc","getMaxKey","getMinKey"]\n' +
  '[[],["hello"],["hello"],[],[],["leet"],[],[]]'
 *
 * 请你设计一个用于存储字符串计数的数据结构，并能够返回计数最小和最大的字符串。
 *
 * 实现 AllOne 类：
 *
 *
 * AllOne() 初始化数据结构的对象。
 * inc(String key) 字符串 key 的计数增加 1 。如果数据结构中尚不存在 key ，那么插入计数为 1 的 key 。
 * dec(String key) 字符串 key 的计数减少 1 。如果 key 的计数在减少后为 0 ，那么需要将这个 key
 * 从数据结构中删除。测试用例保证：在减少计数前，key 存在于数据结构中。
 * getMaxKey() 返回任意一个计数最大的字符串。如果没有元素存在，返回一个空字符串 "" 。
 * getMinKey() 返回任意一个计数最小的字符串。如果没有元素存在，返回一个空字符串 "" 。
 *
 *
 *
 *
 * 示例：
 *
 *
 * 输入
 * ["AllOne", "inc", "inc", "getMaxKey", "getMinKey", "inc", "getMaxKey",
 * "getMinKey"]
 * [[], ["hello"], ["hello"], [], [], ["leet"], [], []]
 * 输出
 * [null, null, null, "hello", "hello", null, "hello", "leet"]
 *
 * 解释
 * AllOne allOne = new AllOne();
 * allOne.inc("hello");
 * allOne.inc("hello");
 * allOne.getMaxKey(); // 返回 "hello"
 * allOne.getMinKey(); // 返回 "hello"
 * allOne.inc("leet");
 * allOne.getMaxKey(); // 返回 "hello"
 * allOne.getMinKey(); // 返回 "leet"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= key.length <= 10
 * key 由小写英文字母组成
 * 测试用例保证：在每次调用 dec 时，数据结构中总存在 key
 * 最多调用 inc、dec、getMaxKey 和 getMinKey 方法 5 * 10^4 次
 *
 *
*/

/*
 * type AllOne struct {
 * 	Count          map[string]int
 * 	Index          map[int]map[string]struct{}
 * 	MaxCnt, MinCnt int
 * }
 *
 * func Constructor() AllOne {
 * 	return AllOne{
 * 		Count: make(map[string]int),
 * 		Index: make(map[int]map[string]struct{}),
 * 	}
 * }
 *
 * func (ao *AllOne) clearIndex(cnt int, key string) {
 * 	m := ao.Index[cnt]
 * 	if m == nil {
 * 		return
 * 	}
 * 	if _, ok := m[key]; ok {
 * 		delete(m, key)
 * 	}
 * 	if len(m) == 0 {
 * 		delete(ao.Index, cnt)
 * 	}
 * }
 *
 * func (ao *AllOne) getMaxCnt() {
 * 	ao.MaxCnt = ao.MinCnt
 * 	for k := range ao.Index {
 * 		if k > ao.MaxCnt {
 * 			ao.MaxCnt = k
 * 		}
 * 	}
 * }
 *
 * func (ao *AllOne) getMinCnt() {
 * 	ao.MinCnt = ao.MaxCnt
 * 	for k := range ao.Index {
 * 		if k < ao.MinCnt {
 * 			ao.MinCnt = k
 * 		}
 * 	}
 * }
 *
 * func (ao *AllOne) Inc(key string) {
 * 	cnt := ao.Count[key]
 * 	ao.clearIndex(cnt, key)
 * 	ao.Count[key]++
 * 	cnt++
 * 	if ao.Index[cnt] == nil {
 * 		ao.Index[cnt] = make(map[string]struct{})
 * 	}
 * 	ao.Index[cnt][key] = struct{}{}
 * 	if cnt > ao.MaxCnt {
 * 		ao.MaxCnt = cnt
 * 	}
 * 	if cnt < ao.MinCnt {
 * 		ao.MinCnt = cnt
 * 	}
 * 	if ao.Index[ao.MinCnt] == nil {
 * 		ao.getMinCnt()
 * 	}
 * 	// fmt.Println("inc", key, ao)
 * }
 *
 * func (ao *AllOne) Dec(key string) {
 * 	cnt := ao.Count[key]
 * 	ao.clearIndex(cnt, key)
 * 	ao.Count[key]--
 * 	cnt--
 * 	if cnt > 0 {
 * 		if ao.Index[cnt] == nil {
 * 			ao.Index[cnt] = make(map[string]struct{})
 * 		}
 * 		ao.Index[cnt][key] = struct{}{}
 * 	} else {
 * 		delete(ao.Count, key)
 * 	}
 * 	if cnt > ao.MaxCnt {
 * 		ao.MaxCnt = cnt
 * 	}
 * 	if cnt > 0 {
 * 		if cnt < ao.MinCnt {
 * 			ao.MinCnt = cnt
 * 		}
 * 	} else {
 * 		ao.getMinCnt()
 * 	}
 * 	if ao.Index[ao.MaxCnt] == nil {
 * 		ao.getMaxCnt()
 * 	}
 * 	// fmt.Println("dec", key, ao)
 * }
 *
 * func (ao *AllOne) GetMaxKey() string {
 * 	if len(ao.Count) == 0 || ao.Index[ao.MaxCnt] == nil {
 * 		return ""
 * 	}
 * 	for k := range ao.Index[ao.MaxCnt] {
 * 		return k
 * 	}
 * 	return ""
 * }
 *
 * func (ao *AllOne) GetMinKey() string {
 * 	if len(ao.Count) == 0 || ao.Index[ao.MinCnt] == nil {
 * 		return ""
 * 	}
 * 	for k := range ao.Index[ao.MinCnt] {
 * 		return k
 * 	}
 * 	return ""
 * }
 */

// 本意是一个 LFU:
type node struct {
	Cnt         int
	Set         map[string]struct{}
	Left, Right *node
}

type AllOne struct {
	Head, Tail *node
	Index      map[string]*node
}

func Constructor() AllOne {
	head, tail := &node{Cnt: -1}, &node{Cnt: -1}
	head.Right = tail
	tail.Left = head
	return AllOne{
		Head:  head,
		Tail:  tail,
		Index: make(map[string]*node),
	}
}

func (ao *AllOne) print() {
	fmt.Print(ao.Index, "\t")
	for c := ao.Head.Right; c != ao.Tail; c = c.Right {
		fmt.Print(c.Cnt, c.Set, "->")
	}
	fmt.Println()
}

func (ao *AllOne) clear(nd *node) {
	if len(nd.Set) == 0 {
		left, right := nd.Left, nd.Right
		left.Right, right.Left = right, left
	}
}

func (ao *AllOne) Inc(key string) {
	nd, ok := ao.Index[key]
	if ok {
		delete(nd.Set, key)
		left := nd.Left
		if left.Cnt != nd.Cnt+1 {
			newNd := &node{
				Cnt: nd.Cnt + 1,
				Set: make(map[string]struct{}),
			}
			newNd.Set[key] = struct{}{}
			ao.Index[key] = newNd
			newNd.Left = left
			newNd.Right = nd
			left.Right = newNd
			nd.Left = newNd
		} else {
			left.Set[key] = struct{}{}
			ao.Index[key] = left
		}
		ao.clear(nd)
		// ao.print()
		return
	}
	tail, left := ao.Tail, ao.Tail.Left
	if left.Cnt != 1 {
		nd = &node{
			Cnt: 1,
			Set: make(map[string]struct{}),
		}
		nd.Set[key] = struct{}{}
		ao.Index[key] = nd
		nd.Left = left
		nd.Right = tail
		tail.Left = nd
		left.Right = nd
	} else {
		left.Set[key] = struct{}{}
		ao.Index[key] = left
	}
	// ao.print()
}

func (ao *AllOne) Dec(key string) {
	nd, ok := ao.Index[key]
	if !ok {
		// ao.print()
		return
	}
	if nd.Cnt-1 == 0 {
		delete(ao.Index, key)
		delete(nd.Set, key)
		ao.clear(nd)
		// ao.print()
		return
	}
	right := nd.Right
	if right.Cnt != nd.Cnt-1 {
		newNd := &node{
			Cnt: nd.Cnt - 1,
			Set: make(map[string]struct{}),
		}
		newNd.Set[key] = struct{}{}
		ao.Index[key] = newNd
		newNd.Left = nd
		newNd.Right = right
		right.Left = newNd
		nd.Right = newNd
	} else {
		right.Set[key] = struct{}{}
		ao.Index[key] = right
	}
	delete(nd.Set, key)
	ao.clear(nd)
	// ao.print()
}

func (ao *AllOne) GetMaxKey() string {
	nd := ao.Head.Right
	if nd.Set != nil {
		for k := range nd.Set {
			return k
		}
	}
	return ""
}

func (ao *AllOne) GetMinKey() string {
	nd := ao.Tail.Left
	if nd.Set != nil {
		for k := range nd.Set {
			return k
		}
	}
	return ""
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
