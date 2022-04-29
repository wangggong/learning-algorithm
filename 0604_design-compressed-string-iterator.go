/*
 * Easy
 *
 * 设计并实现一个迭代压缩字符串的数据结构。给定的压缩字符串的形式是，每个字母后面紧跟一个正整数，表示该字母在原始未压缩字符串中出现的次数。
 *
 * 设计一个数据结构，它支持如下两种操作： next 和 hasNext。
 *
 * next() - 如果原始字符串中仍有未压缩字符，则返回下一个字符，否则返回空格。
 * hasNext() - 如果原始字符串中存在未压缩的的字母，则返回true，否则返回false。
 *
 *
 * 示例 1：
 *
 * 输入：
 * ["StringIterator", "next", "next", "next", "next", "next", "next", "hasNext", "next", "hasNext"]
 * [["L1e2t1C1o1d1e1"], [], [], [], [], [], [], [], [], []]
 * 输出：
 * [null, "L", "e", "e", "t", "C", "o", true, "d", true]
 *
 * 解释：
 * StringIterator stringIterator = new StringIterator("L1e2t1C1o1d1e1");
 * stringIterator.next(); // 返回 "L"
 * stringIterator.next(); // 返回 "e"
 * stringIterator.next(); // 返回 "e"
 * stringIterator.next(); // 返回 "t"
 * stringIterator.next(); // 返回 "C"
 * stringIterator.next(); // 返回 "o"
 * stringIterator.hasNext(); // 返回 True
 * stringIterator.next(); // 返回 "d"
 * stringIterator.hasNext(); // 返回 True
 *
 *
 * 提示：
 *
 * 1 <= compressedString.length <= 1000
 * compressedString 由小写字母、大写字母和数字组成。
 * 在 compressedString 中，单个字符的重复次数在 [1,10 ^9] 范围内。
 * next 和 hasNext 的操作数最多为 100 。
 * 通过次数2,960
 * 提交次数7,818
 */

type StringIterator struct {
	Bytes  []byte
	Counts []int
	Index  int
	Curr   int
	N      int
}

func Constructor(s string) StringIterator {
	arr := []byte(s)
	n := len(arr)
	var bytes []byte
	var counts []int
	for i := 0; i < n; i++ {
		if ('a' <= arr[i] && arr[i] <= 'z') || ('A' <= arr[i] && arr[i] <= 'Z') {
			bytes = append(bytes, arr[i])
		} else {
			j := i
			v := 0
			for ; j < n && '0' <= arr[j] && arr[j] <= '9'; j++ {
				v = v*10 + int(arr[j]-'0')
			}
			counts = append(counts, v)
			i = j - 1
		}
	}
	// assert len(bytes) == len(count)
	return StringIterator{
		Bytes:  bytes,
		Counts: counts,
		Index:  0,
		Curr:   0,
		N:      len(bytes),
	}
}

func (si *StringIterator) Next() byte {
	if !si.HasNext() {
		return ' '
	}
	si.Curr++
	b := si.Bytes[si.Index]
	if si.Curr == si.Counts[si.Index] {
		si.Curr = 0
		si.Index++
	}
	return b
}

func (si *StringIterator) HasNext() bool {
	return si.Index < si.N
}

/**
 * Your StringIterator object will be instantiated and called as such:
 * obj := Constructor(compressedString);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
