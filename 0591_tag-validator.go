/*
 * @lc app=leetcode.cn id=tag-validator lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [591] 标签验证器
 *
 * https://leetcode-cn.com/problems/tag-validator/description/
 *
 * algorithms
 * Hard (34.13%)
 * Total Accepted:    6.5K
 * Total Submissions: 13.4K
 * Testcase Example:  '"<DIV>This is the first line <![CDATA[<div>]]></DIV>"'
 *
 * 给定一个表示代码片段的字符串，你需要实现一个验证器来解析这段代码，并返回它是否合法。合法的代码片段需要遵守以下的所有规则：
 *
 *
 * 代码必须被合法的闭合标签包围。否则，代码是无效的。
 *
 * 闭合标签（不一定合法）要严格符合格式：<TAG_NAME>TAG_CONTENT</TAG_NAME>。其中，<TAG_NAME>是起始标签，</TAG_NAME>是结束标签。起始和结束标签中的
 * TAG_NAME 应当相同。当且仅当 TAG_NAME 和 TAG_CONTENT 都是合法的，闭合标签才是合法的。
 * 合法的 TAG_NAME 仅含有大写字母，长度在范围 [1,9] 之间。否则，该 TAG_NAME 是不合法的。
 * 合法的 TAG_CONTENT 可以包含其他合法的闭合标签，cdata
 * （请参考规则7）和任意字符（注意参考规则1）除了不匹配的<、不匹配的起始和结束标签、不匹配的或带有不合法 TAG_NAME
 * 的闭合标签。否则，TAG_CONTENT 是不合法的。
 * 一个起始标签，如果没有具有相同 TAG_NAME 的结束标签与之匹配，是不合法的。反之亦然。不过，你也需要考虑标签嵌套的问题。
 * 一个<，如果你找不到一个后续的>与之匹配，是不合法的。并且当你找到一个<或</时，所有直到下一个>的前的字符，都应当被解析为
 * TAG_NAME（不一定合法）。
 * cdata 有如下格式：<![CDATA[CDATA_CONTENT]]>。CDATA_CONTENT 的范围被定义成 <![CDATA[
 * 和后续的第一个 ]]>之间的字符。
 * CDATA_CONTENT 可以包含任意字符。cdata
 * 的功能是阻止验证器解析CDATA_CONTENT，所以即使其中有一些字符可以被解析为标签（无论合法还是不合法），也应该将它们视为常规字符。
 *
 *
 * 合法代码的例子:
 *
 *
 * 输入: "<DIV>This is the first line <![CDATA[<div>]]></DIV>"
 *
 * 输出: True
 *
 * 解释:
 *
 * 代码被包含在了闭合的标签内： <DIV> 和 </DIV> 。
 *
 * TAG_NAME 是合法的，TAG_CONTENT 包含了一些字符和 cdata 。
 *
 * 即使 CDATA_CONTENT 含有不匹配的起始标签和不合法的 TAG_NAME，它应该被视为普通的文本，而不是标签。
 *
 * 所以 TAG_CONTENT 是合法的，因此代码是合法的。最终返回True。
 *
 *
 * 输入: "<DIV>>>  ![cdata[]] <![CDATA[<div>]>]]>]]>>]</DIV>"
 *
 * 输出: True
 *
 * 解释:
 *
 * 我们首先将代码分割为： start_tag|tag_content|end_tag 。
 *
 * start_tag -> "<DIV>"
 *
 * end_tag -> "</DIV>"
 *
 * tag_content 也可被分割为： text1|cdata|text2 。
 *
 * text1 -> ">>  ![cdata[]] "
 *
 * cdata -> "<![CDATA[<div>]>]]>" ，其中 CDATA_CONTENT 为 "<div>]>"
 *
 * text2 -> "]]>>]"
 *
 *
 * start_tag 不是 "<DIV>>>" 的原因参照规则 6 。
 * cdata 不是 "<![CDATA[<div>]>]]>]]>" 的原因参照规则 7 。
 *
 *
 * 不合法代码的例子:
 *
 *
 * 输入: "<A>  <B> </A>   </B>"
 * 输出: False
 * 解释: 不合法。如果 "<A>" 是闭合的，那么 "<B>" 一定是不匹配的，反之亦然。
 *
 * 输入: "<DIV>  div tag is not closed  <DIV>"
 * 输出: False
 *
 * 输入: "<DIV>  unmatched <  </DIV>"
 * 输出: False
 *
 * 输入: "<DIV> closed tags with invalid tag name  <b>123</b> </DIV>"
 * 输出: False
 *
 * 输入: "<DIV> unmatched tags with invalid tag name  </1234567890> and
 * <CDATA[[]]>  </DIV>"
 * 输出: False
 *
 * 输入: "<DIV>  unmatched start tag <B>  and unmatched end tag </C>  </DIV>"
 * 输出: False
 *
 *
 * 注意:
 *
 *
 * 为简明起见，你可以假设输入的代码（包括提到的任意字符）只包含数字, 字母, '<','>','/','!','[',']'和' '。
 *
 *
 */
const (
	cdataTag    = "[CDATA["
	cdataEndTag = "]]>"
)

var S []string

func empty(S []string) bool { return len(S) == 0 }

func isUpper(b byte) bool { return 'A' <= b && b <= 'Z' }

func indexOf(text, pattern []byte) int {
	n, m := len(text), len(pattern)
	for i := 0; i < n; i++ {
		ok := true
		for j := 0; j < m; j++ {
			if text[i+j] != pattern[j] {
				ok = false
				break
			}
		}
		if ok {
			return i
		}
	}
	return -1
}

func indexOfByte(text []byte, b byte) int {
	for i, c := range text {
		if c == b {
			return i
		}
	}
	return -1
}

func isValid(code string) bool {
	S = nil
	c := []byte(code)
	n := len(c)
	if n == 0 {
		return false
	}
	// 注意, 整体 check 完需要判定 stack 是否为空.
	return valid(c, 0, n) && empty(S)
}

func valid(arr []byte, l, r int) bool {
	if l >= r {
		return true
	}
	if arr[l] != '<' {
		if empty(S) {
			return false
		}
		ind := indexOfByte(arr[l:r], '<')
		if ind < 0 {
			return false
		}
		return valid(arr, l+ind, r)
	}
	if l+1 >= r {
		return false
	}
	switch arr[l+1] {
	case '/':
		i := indexOfByte(arr[l+2:r], '>')
		if i < 0 {
			return false
		}
		if empty(S) || S[len(S)-1] != string(arr[l+2:l+2+i]) {
			return false
		}
		S = S[:len(S)-1]
		// 要求代码必须被标签 **包围**, 因此 `<tag1>...</tag1>...<tag2>...</tag2>` 是非法的.
		if empty(S) && l+3+i < r {
			return false
		}
		return valid(arr, l+3+i, r)
	case '!':
		// 除了 tag 本身外, 所有其他情形均需要保证 stack 非空.
		if empty(S) {
			return false
		}
		if l+9 > r || string(arr[l+2:l+9]) != cdataTag {
			return false
		}
		i := indexOf(arr[l+9:r], []byte(cdataEndTag))
		if i < 0 {
			return false
		}
		return valid(arr, l+9+i+3, r)
	default:
		if !isUpper(arr[l+1]) {
			return false
		}
		i := indexOfByte(arr[l+1:r], '>')
		// tag 长度要 **小于** 10, 注意边界判断.
		if i < 0 || i > 9 {
			return false
		}
		// tag 要求内部均为大写字母.
		for j := 0; j < i; j++ {
			if !isUpper(arr[l+1+j]) {
				return false
			}
		}
		S = append(S, string(arr[l+1:l+1+i]))
		return valid(arr, l+2+i, r)
	}
	return false
}

/*
 *
 * test cases:
 * - base cases:
 * "<DIV>This is the first line <![CDATA[<div>]]></DIV>"
 * "<DIV>>>  ![cdata[]] <![CDATA[<div>]>]]>]]>>]</DIV>"
 * "<A>  <B> </A>   </B>"
 * "<DIV>  div tag is not closed  <DIV>"
 * "<DIV>  unmatched <  </DIV>"
 * "<DIV> closed tags with invalid tag name  <b>123</b> </DIV>"
 * "<DIV> unmatched tags with invalid tag name  </1234567890> and <CDATA[[]]>  </DIV>"
 * "<DIV>  unmatched start tag <B>  and unmatched end tag </C>  </DIV>"
 *
 * - bad cases:
 * "<DIV>This is the first line <![CDATA[<div>]]><DIV>"
 * "<![CDATA[wahaha]]]><![CDATA[]> wahaha]]>"
 * "<TRUe><![CDATA[123123]]]]><![CDATA[>123123]]></TRUe>"
 * "<A></A><B></B>"
 * "<TRUE><![CDATA[wahaha]]]><![CDATA[]> wahaha]]></TRUE>"
 * "<A><A></A></A>"
 * "<A"
 * "<AAAAAAAAAA></AAAAAAAAAA>"
 *
 */
