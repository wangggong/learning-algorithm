/*
 * @lc app=leetcode.cn id=design-a-text-editor lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6093] 设计一个文本编辑器
 *
 * https://leetcode.cn/problems/design-a-text-editor/description/
 *
 * algorithms
 * Hard (32.85%)
 * Total Accepted:    2.4K
 * Total Submissions: 7.3K
 * Testcase Example:  '["TextEditor","addText","deleteText","addText","cursorRight","cursorLeft","deleteText","cursorLeft","cursorRight"]\n' +
  '[[],["leetcode"],[4],["practice"],[3],[8],[10],[2],[6]]'
 *
 * 请你设计一个带光标的文本编辑器，它可以实现以下功能：
 *
 *
 * 添加：在光标所在处添加文本。
 * 删除：在光标所在处删除文本（模拟键盘的删除键）。
 * 移动：将光标往左或者往右移动。
 *
 *
 * 当删除文本时，只有光标左边的字符会被删除。光标会留在文本内，也就是说任意时候 0 <= cursor.position <=
 * currentText.length 都成立。
 *
 * 请你实现 TextEditor 类：
 *
 *
 * TextEditor() 用空文本初始化对象。
 * void addText(string text) 将 text 添加到光标所在位置。添加完后光标在 text 的右边。
 * int deleteText(int k) 删除光标左边 k 个字符。返回实际删除的字符数目。
 * string cursorLeft(int k) 将光标向左移动 k 次。返回移动后光标左边 min(10, len) 个字符，其中 len
 * 是光标左边的字符数目。
 * string cursorRight(int k) 将光标向右移动 k 次。返回移动后光标左边 min(10, len) 个字符，其中 len
 * 是光标左边的字符数目。
 *
 *
 *
 *
 * 示例 1：
 *
 * 输入：
 * ["TextEditor", "addText", "deleteText", "addText", "cursorRight",
 * "cursorLeft", "deleteText", "cursorLeft", "cursorRight"]
 * [[], ["leetcode"], [4], ["practice"], [3], [8], [10], [2], [6]]
 * 输出：
 * [null, null, 4, null, "etpractice", "leet", 4, "", "practi"]
 *
 * 解释：
 * TextEditor textEditor = new TextEditor(); // 当前 text 为 "|" 。（'|' 字符表示光标）
 * textEditor.addText("leetcode"); // 当前文本为 "leetcode|" 。
 * textEditor.deleteText(4); // 返回 4
 * ⁠                         // 当前文本为 "leet|" 。
 * ⁠                         // 删除了 4 个字符。
 * textEditor.addText("practice"); // 当前文本为 "leetpractice|" 。
 * textEditor.cursorRight(3); // 返回 "etpractice"
 * ⁠                          // 当前文本为 "leetpractice|".
 * ⁠                          // 光标无法移动到文本以外，所以无法移动。
 * ⁠                          // "etpractice" 是光标左边的 10 个字符。
 * textEditor.cursorLeft(8); // 返回 "leet"
 * ⁠                         // 当前文本为 "leet|practice" 。
 * ⁠                         // "leet" 是光标左边的 min(10, 4) = 4 个字符。
 * textEditor.deleteText(10); // 返回 4
 * ⁠                          // 当前文本为 "|practice" 。
 * ⁠                          // 只有 4 个字符被删除了。
 * textEditor.cursorLeft(2); // 返回 ""
 * ⁠                         // 当前文本为 "|practice" 。
 * ⁠                         // 光标无法移动到文本以外，所以无法移动。
 * ⁠                         // "" 是光标左边的 min(10, 0) = 0 个字符。
 * textEditor.cursorRight(6); // 返回 "practi"
 * ⁠                          // 当前文本为 "practi|ce" 。
 * ⁠                          // "practi" 是光标左边的 min(10, 6) = 6 个字符。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= text.length, k <= 40
 * text 只含有小写英文字母。
 * 调用 addText ，deleteText ，cursorLeft 和 cursorRight 的 总 次数不超过 2 * 10^4 次。
 *
 *
*/

/*
 * // []byte 模拟
 * const maxn int = 1e6
 *
 * type TextEditor struct {
 * 	SF, ST []byte
 * 	Index  int
 * 	Len    int
 * 	FT     bool
 * }
 *
 * func Constructor() TextEditor {
 * 	return TextEditor{
 * 		SF: make([]byte, maxn),
 * 		ST: make([]byte, maxn),
 * 	}
 * }
 *
 * func (te *TextEditor) AddText(s string) {
 * 	from, to := te.SF, te.ST
 * 	if te.FT {
 * 		from, to = te.ST, te.SF
 * 	}
 *
 * 	copy(to[:te.Index], from[:te.Index])
 * 	newIndex := te.Index + len(s)
 * 	copy(to[te.Index:newIndex], []byte(s))
 * 	newLen := newIndex + te.Len - te.Index
 * 	copy(to[newIndex:newLen], from[te.Index:te.Len])
 *
 * 	te.Index = newIndex
 * 	te.Len = newLen
 * 	te.FT = !te.FT
 * 	return
 * }
 *
 * func (te *TextEditor) DeleteText(k int) int {
 * 	from, to := te.SF, te.ST
 * 	if te.FT {
 * 		from, to = te.ST, te.SF
 * 	}
 *
 * 	k = min(k, te.Index)
 *
 * 	newIndex := te.Index - k
 * 	copy(to[:newIndex], from[:newIndex])
 * 	newLen := newIndex + te.Len - te.Index
 * 	copy(to[newIndex:newLen], from[te.Index:te.Len])
 *
 * 	te.Index = newIndex
 * 	te.Len = newLen
 * 	te.FT = !te.FT
 * 	return k
 * }
 *
 * func (te *TextEditor) CursorLeft(k int) string {
 * 	from := te.SF
 * 	if te.FT {
 * 		from = te.ST
 * 	}
 *
 * 	k = min(k, te.Index)
 * 	te.Index -= k
 * 	return string(from[te.Index-min(10, te.Index) : te.Index])
 * }
 *
 * func (te *TextEditor) CursorRight(k int) string {
 * 	from := te.SF
 * 	if te.FT {
 * 		from = te.ST
 * 	}
 *
 * 	k = min(k, te.Len-te.Index)
 * 	te.Index += k
 * 	return string(from[te.Index-min(10, te.Index) : te.Index])
 * }
 *
 * func min(x, y int) int {
 * 	if x < y {
 * 		return x
 * 	}
 * 	return y
 * }
 */

type listNode struct {
	Val        byte
	Prev, Next *listNode
}

func (n *listNode) insertPrev(node *listNode) {
	prev := n.Prev
	node.Prev, node.Next, n.Prev = prev, n, node
	if prev != nil {
		prev.Next = node
	}
	return
}

type TextEditor struct {
	Head, Tail, Curr *listNode
}

func Constructor() TextEditor {
	head, tail := &listNode{}, &listNode{}
	head.Next, tail.Prev = tail, head
	return TextEditor{
		Head: head,
		Tail: tail,
		Curr: tail,
	}
}

func (te *TextEditor) AddText(s string) {
	for _, b := range []byte(s) {
		te.Curr.insertPrev(&listNode{Val: b})
	}
	return
}

func (te *TextEditor) DeleteText(k int) int {
	ans := 0
	curr := te.Curr.Prev
	for ; ans < k && curr != te.Head; ans++ {
		curr = curr.Prev
	}
	te.Curr.Prev, curr.Next = curr, te.Curr
	return ans
}

func (te *TextEditor) text() string {
	var ans []byte
	for i, curr := 0, te.Curr.Prev; i < 10 && curr != te.Head; i, curr = i+1, curr.Prev {
		ans = append(ans, curr.Val)
	}
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return string(ans)
}

func (te *TextEditor) CursorLeft(k int) string {
	for i := 0; i < k && te.Curr != te.Head.Next; i++ {
		te.Curr = te.Curr.Prev
	}
	return te.text()
}

func (te *TextEditor) CursorRight(k int) string {
	for i := 0; i < k && te.Curr != te.Tail; i++ {
		te.Curr = te.Curr.Next
	}
	return te.text()
}

/**
 * Your TextEditor object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddText(text);
 * param_2 := obj.DeleteText(k);
 * param_3 := obj.CursorLeft(k);
 * param_4 := obj.CursorRight(k);
 */
