/*
 * @lc app=leetcode.cn id=design-hashmap lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [706] 设计哈希映射
 *
 * https://leetcode-cn.com/problems/design-hashmap/description/
 *
 * algorithms
 * Easy (64.04%)
 * Total Accepted:    68.9K
 * Total Submissions: 107.6K
 * Testcase Example:  '["MyHashMap","put","put","get","get","put","get","remove","get"]\n' +
  '[[],[1,1],[2,2],[1],[3],[2,1],[2],[2],[2]]'
 *
 * 不使用任何内建的哈希表库设计一个哈希映射（HashMap）。
 *
 * 实现 MyHashMap 类：
 *
 *
 * MyHashMap() 用空映射初始化对象
 * void put(int key, int value) 向 HashMap 插入一个键值对 (key, value) 。如果 key
 * 已经存在于映射中，则更新其对应的值 value 。
 * int get(int key) 返回特定的 key 所映射的 value ；如果映射中不包含 key 的映射，返回 -1 。
 * void remove(key) 如果映射中存在 key 的映射，则移除 key 和它所对应的 value 。
 *
 *
 *
 *
 * 示例：
 *
 *
 * 输入：
 * ["MyHashMap", "put", "put", "get", "get", "put", "get", "remove", "get"]
 * [[], [1, 1], [2, 2], [1], [3], [2, 1], [2], [2], [2]]
 * 输出：
 * [null, null, null, 1, -1, null, 1, null, -1]
 *
 * 解释：
 * MyHashMap myHashMap = new MyHashMap();
 * myHashMap.put(1, 1); // myHashMap 现在为 [[1,1]]
 * myHashMap.put(2, 2); // myHashMap 现在为 [[1,1], [2,2]]
 * myHashMap.get(1);    // 返回 1 ，myHashMap 现在为 [[1,1], [2,2]]
 * myHashMap.get(3);    // 返回 -1（未找到），myHashMap 现在为 [[1,1], [2,2]]
 * myHashMap.put(2, 1); // myHashMap 现在为 [[1,1], [2,1]]（更新已有的值）
 * myHashMap.get(2);    // 返回 1 ，myHashMap 现在为 [[1,1], [2,1]]
 * myHashMap.remove(2); // 删除键为 2 的数据，myHashMap 现在为 [[1,1]]
 * myHashMap.get(2);    // 返回 -1（未找到），myHashMap 现在为 [[1,1]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= key, value <= 10^6
 * 最多调用 10^4 次 put、get 和 remove 方法
 *
 *
*/

const (
	PRIME = 139
	N     = 114514
)

func myHash(key int) int {
	// return (key + N) % PRIME
	return key % PRIME
}

type Node struct {
	Key  int
	Val  int
	Next *Node
}

type MyHashMap struct {
	Arr [PRIME]*Node
}

func Constructor() MyHashMap {
	return MyHashMap{}
}

func (mhm *MyHashMap) Put(key int, value int) {
	k := myHash(key)

	/*
		defer func() {
			fmt.Printf("put %v ", k)
			for c := mhm.Arr[k]; c != nil; c = c.Next {
				fmt.Printf("%v ", c)
			}
			fmt.Println()
		}()
	*/

	var p, c *Node
	for c = mhm.Arr[k]; c != nil; p, c = c, c.Next {
		if c.Key == key {
			c.Val = value
			return
		}
	}
	node := Node{Key: key, Val: value}
	if p == nil {
		mhm.Arr[k] = &node
		return
	}
	p.Next = &node
}

func (mhm *MyHashMap) Get(key int) int {
	k := myHash(key)

	/*
		defer func() {
			fmt.Printf("get %v ", k)
			for c := mhm.Arr[k]; c != nil; c = c.Next {
				fmt.Printf("%v ", c)
			}
			fmt.Println()
		}()
	*/

	for c := mhm.Arr[k]; c != nil; c = c.Next {
		if c.Key == key {
			return c.Val
		}
	}
	return -1
}

func (mhm *MyHashMap) Remove(key int) {
	k := myHash(key)

	/*
		defer func() {
			fmt.Printf("get %v ", k)
			for c := mhm.Arr[k]; c != nil; c = c.Next {
				fmt.Printf("%v ", c)
			}
			fmt.Println()
		}()
	*/

	var p, c *Node
	for c = mhm.Arr[k]; c != nil; p, c = c, c.Next {
		if c.Key == key {
			if p == nil {
				mhm.Arr[k] = c.Next
				return
			}
			p.Next = c.Next
			return
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
