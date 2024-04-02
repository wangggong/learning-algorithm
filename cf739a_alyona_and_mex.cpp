#include <iostream>

using namespace std;
const int maxn = 1e5;

// **构造题**
// Q1: 答案最大 **能** 是多少?
// A: 最大是最短区间长度. `[l, r)` 区间随便赋值的话可以无脑赋 `0, 1, ..., r-l-1`, 此时 `mex(a[l..r]) = r-l`. 如果极大化这个极小指标顶天到最短区间长度.
// Q2: 能取到吗?
// A: 能. 如何让所有长度为 `k` 的区间取 `mex` 都能得到 `k`, 需要让这些区间都有 `0, 1, ..., k-1`. 问题转化为如何保证一个长度 `>= k` 的区间都能取到这些数字----
// 答案是 **直接去构造 `0 ~ k` 的周期序列.** 傻逼了吧.
int main() {
    int n, m, k = maxn;
    cin >> n >> m;
    while (m--) {
        int l, r;
        cin >> l >> r;
        k = min(k, r - l + 1);
    }
    cout << k << endl;
    for (int i = 0; i < n; i++)
        cout << (i % k) << " ";
    cout << endl;
    return 0;
}
