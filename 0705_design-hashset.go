/*
 * @lc app=leetcode.cn id=design-hashset lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [705] 设计哈希集合
 *
 * https://leetcode-cn.com/problems/design-hashset/description/
 *
 * algorithms
 * Easy (64.03%)
 * Total Accepted:    76.9K
 * Total Submissions: 120.1K
 * Testcase Example:  '["MyHashSet","add","add","contains","contains","add","contains","remove","contains"]\n' +
  '[[],[1],[2],[1],[3],[2],[2],[2],[2]]'
 *
 * 不使用任何内建的哈希表库设计一个哈希集合（HashSet）。
 *
 * 实现 MyHashSet 类：
 *
 *
 * void add(key) 向哈希集合中插入值 key 。
 * bool contains(key) 返回哈希集合中是否存在这个值 key 。
 * void remove(key) 将给定值 key 从哈希集合中删除。如果哈希集合中没有这个值，什么也不做。
 *
 *
 *
 * 示例：
 *
 *
 * 输入：
 * ["MyHashSet", "add", "add", "contains", "contains", "add", "contains",
 * "remove", "contains"]
 * [[], [1], [2], [1], [3], [2], [2], [2], [2]]
 * 输出：
 * [null, null, null, true, false, null, true, null, false]
 *
 * 解释：
 * MyHashSet myHashSet = new MyHashSet();
 * myHashSet.add(1);      // set = [1]
 * myHashSet.add(2);      // set = [1, 2]
 * myHashSet.contains(1); // 返回 True
 * myHashSet.contains(3); // 返回 False ，（未找到）
 * myHashSet.add(2);      // set = [1, 2]
 * myHashSet.contains(2); // 返回 True
 * myHashSet.remove(2);   // set = [1]
 * myHashSet.contains(2); // 返回 False ，（已移除）
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= key <= 10^6
 * 最多调用 10^4 次 add、remove 和 contains
 *
 *
*/

const PRIME = 139

func myHash(n int) int {
	return n % PRIME
}

type Node struct {
	Key  int
	Next *Node
}

type MyHashSet struct {
	Arr [PRIME]*Node
}

func Constructor() MyHashSet {
	return MyHashSet{}
}

func (mhs *MyHashSet) Add(key int) {
	k := myHash(key)
	if mhs.Contains(key) {
		return
	}
	if mhs.Arr[k] == nil {
		mhs.Arr[k] = &Node{Key: -1}
	}
	p, c := mhs.Arr[k], mhs.Arr[k].Next
	node := Node{Key: key}
	p.Next, node.Next = &node, c
}

func (mhs *MyHashSet) Remove(key int) {
	k := myHash(key)
	if !mhs.Contains(key) {
		return
	}
	p, c := mhs.Arr[k], mhs.Arr[k].Next
	for ; c != nil; p, c = c, c.Next {
		if c.Key == key {
			p.Next = c.Next
			return
		}
	}
}

func (mhs *MyHashSet) Contains(key int) bool {
	k := myHash(key)
	for c := mhs.Arr[k]; c != nil; c = c.Next {
		if c.Key == key {
			return true
		}
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
