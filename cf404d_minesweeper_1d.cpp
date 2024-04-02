// 参考: https://atcoder.jp/contests/abc155/submissions/45503273

#include <iostream>
#include <cstring>

using namespace std;
typedef long long ll;

const ll N = 1e6;
const ll MOD = 1e9 + 7;
char s[N + 10];
ll memo[N + 10][2 + 10];

// `dfs(i, j)` 代表 `i` 位置的状态:
// j == 0: 当前位置没有雷
// j == 1: 当前位置有雷
// j == 1: 下一个位置有雷
ll dfs(int i, int j) {
    if (i < 0)
        return j != 1;
    if (memo[i][j] != -1)
        return memo[i][j];
    switch (j) {
        case 0:
            switch (s[i]) {
                case '0':
                    return memo[i][j] = dfs(i - 1, 0);
                case '1':
                    return memo[i][j] = dfs(i - 1, 1);
                case '?':
                    return memo[i][j] = (dfs(i - 1, 0) + dfs(i - 1, 1)) % MOD;
                default:
                    return memo[i][j] = 0;
            }
        case 1:
            return memo[i][j] = (s[i] == '0' || s[i] == '1' || s[i] == '2'
                ? 0
                : dfs(i - 1, 2));
        default:    // 2
            switch (s[i]) {
                case '0':
                    return memo[i][j] = 0;
                case '1':
                    return memo[i][j] = dfs(i - 1, 0);
                case '2':
                    return memo[i][j] = dfs(i - 1, 1);
                case '*':
                    return memo[i][j] = dfs(i - 1, 2);
                default:
                    return memo[i][j] = (dfs(i - 1, 0) + dfs(i - 1, 1) + dfs(i - 1, 2)) % MOD;
            }
    }
}

int main() {
    memset(memo, -1, sizeof memo);
    scanf("%s", s);
    ll n = (ll)strlen(s);
    printf("%lld\n", (dfs(n - 1, 0) + dfs(n - 1, 1)) % MOD);
    return 0;
}
