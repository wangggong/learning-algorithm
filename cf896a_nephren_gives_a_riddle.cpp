#include <iostream>

using namespace std;
typedef long long ll;

const ll N = 1e5, K = 1e18 + 1;
ll q, n, k;
ll offset[N + 10][6];
string s = "What are you doing at the end of the world? Are you busy? Will you save us?";
string ss[] = {"What are you doing while sending \"", "\"? Are you busy? Will you send \"", "\"?"};

void init() {
    ll size = s.size();
    for (ll i = 1; i <= N; i++)
        offset[i][1] = min((ll) ss[0].size(), K), offset[i][2] = min(offset[i][1] + size, K), offset[i][3] = min(offset[i][2] + (ll) ss[1].size(), K), offset[i][4] = min(offset[i][3] + size, K), offset[i][5] = min(offset[i][4] + (ll) ss[2].size(), K), size = offset[i][5];
    return;
}

char get(ll n, ll k) {
    if (n == 0)
        return k >= s.size() ? '.' : s[k];
    if (offset[n][0] <= k && k < offset[n][1])
        return ss[0][k];
    if (offset[n][2] <= k && k < offset[n][3])
        return ss[1][k - offset[n][2]];
    if (offset[n][4] <= k && k < offset[n][5])
        return ss[2][k - offset[n][4]];
    if (offset[n][1] <= k && k < offset[n][2])
        return get(n - 1, k - offset[n][1]);
    if (offset[n][3] <= k && k < offset[n][4])
        return get(n - 1, k - offset[n][3]);
    return '.';
}

int main() {
    init();
    cin >> q;
    for (ll i = 0; i < q; i++)
        cin >> n >> k, cout << get(n, k - 1);
    return 0;
}
