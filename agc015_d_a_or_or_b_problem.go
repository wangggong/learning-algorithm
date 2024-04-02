// 参考: https://www.luogu.com.cn/blog/endlesscheng/solution-at-agc015-d

package main

import (
	"fmt"
	"io"
	"math/bits"
	"os"
)

func agc015(in io.Reader, out io.Writer) {
	var a uint
	var b uint
	fmt.Fscan(in, &a, &b)
	if a == b {
		fmt.Fprintln(out, 1)
		return
	}
	ans := b - a + 1
	mask := uint(1)<<(bits.Len(a^b)-1) - 1
	a &= mask
	b &= mask
	nb := bits.Len(b)
	if bits.Len(a) <= nb {
		ans += mask - b
	} else {
		ans += mask - a + 1<<nb - b
	}
	fmt.Fprintln(out, ans)
}

func main() { agc015(os.Stdin, os.Stdout) }
