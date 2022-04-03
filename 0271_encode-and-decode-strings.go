/*
 * Medium
 * 通过次数3,129
 * 提交次数5,591
 *
 * 请你设计一个算法，可以将一个 字符串列表 编码成为一个 字符串。这个编码后的字符串是可以通过网络进行高效传送的，并且可以在接收端被解码回原来的字符串列表。
 *
 * 1 号机（发送方）有如下函数：
 *
 * string encode(vector<string> strs) {
 *   // ... your code
 *   return encoded_string;
 * }
 * 2 号机（接收方）有如下函数：
 *
 * vector<string> decode(string s) {
 *   //... your code
 *   return strs;
 * }
 * 1 号机（发送方）执行：
 *
 * string encoded_string = encode(strs);
 * 2 号机（接收方）执行：
 *
 * vector<string> strs2 = decode(encoded_string);
 * 此时，2 号机（接收方）的 strs2 需要和 1 号机（发送方）的 strs 相同。
 *
 * 请你来实现这个 encode 和 decode 方法。
 *
 * 注意：
 *
 * 因为字符串可能会包含 256 个合法 ascii 字符中的任何字符，所以您的算法必须要能够处理任何可能会出现的字符。
 * 请勿使用 “类成员”、“全局变量” 或 “静态变量” 来存储这些状态，您的编码和解码算法应该是非状态依赖的。
 * 请不要依赖任何方法库，例如 eval 又或者是 serialize 之类的方法。本题的宗旨是需要您自己实现 “编码” 和 “解码” 算法。
 *
 */

import (
	"fmt"
)

type Codec struct{}

// Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
	ans := ""
	for _, s := range strs {
		ans = ans + fmt.Sprintf("[%v]%s", len(s), s)
	}
	return ans
}

// Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {
	var ans []string
	bytes := []byte(strs)
	p, q := 0, len(bytes)
	for p < q {
		p++
		n := 0
		for p < q && bytes[p] != ']' {
			n *= 10
			n += int(bytes[p] - '0')
			p++
		}
		p++
		ans = append(ans, string(bytes[p:p+n]))
		p += n
	}
	return ans
}

// 注: 官方解给的是直接用非 ascii 编码的字符把字符串列表 join 在一起; 解码时直接 split.

// Your Codec object will be instantiated and called as such:
// var codec Codec
// codec.Decode(codec.Encode(strs));
