package main

import (
	"fmt"
	"io"
	"os"
)

type keyType struct {
	D       int
	Product int64
	IsLimit bool
	IsNum   bool
}

func abc208e(in io.Reader, out io.Writer) {
	var s string
	var k int64
	fmt.Fscan(in, &s, &k)
	m := make(map[keyType]int64)
	var dfs func(int, int64, bool, bool) int64
	dfs = func(d int, product int64, isLimit, isNum bool) int64 {
		key := keyType{d, product, isLimit, isNum}
		if v, ok := m[key]; ok {
			return v
		}
		m[key] = 0
		if d == len(s) {
			if isNum && product <= k {
				m[key] = 1
			}
			return m[key]
		}
		if !isNum {
			m[key] += dfs(d+1, product, false, isNum)
		}
		l, r := int64(0), int64(9)
		if isLimit {
			r = int64(s[d] - '0')
		}
		if !isNum {
			l = 1
		}
		for i := l; i <= r; i++ {
			m[key] += dfs(d+1, product*i, isLimit && i == r, true)
		}
		return m[key]
	}
	fmt.Fprintln(out, dfs(0, 1, true, false))
}

func main() { abc208e(os.Stdin, os.Stdout) }
