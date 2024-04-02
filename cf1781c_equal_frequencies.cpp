// 参考: https://codeforces.com/contest/1781/submission/240233629
#include <iostream>
#include <cstring>
#include <algorithm>
#include <vector>

using namespace std;
typedef long long ll;
typedef pair<ll, char> plc;

const ll N = 1e5, ALPHA = 26;
ll t, n, k, m, maxSave, c[ALPHA];
char s[N + 10];

void getMaxSave(vector<plc> &vec) {
    k = 0, maxSave = 0;
    for (ll i = 1; i <= ALPHA; i++) {
        if (n % i)
            continue;
        ll save = 0;
        for (ll j = 0; j < i; j++)
            save += min(-vec[j].first, n / i);
        if (save > maxSave)
            maxSave = save, k = i, m = n / k;
    }
}

void setMaxSave(vector<plc> &vec) {
    vector<char> needed;
    for (ll i = 0; i < k; i++)
        if (c[vec[i].second] > m)
            c[vec[i].second] = m;
        else
            for (ll j = c[vec[i].second]; j < m; j++)
                needed.push_back(vec[i].second + 'a');
    for (ll i = k; i < ALPHA; i++)
        c[vec[i].second] = 0;
    for (ll i = 0, j = 0; i < n; i++)
        if (c[s[i] - 'a'])
            c[s[i] - 'a']--;
        else
            s[i] = needed[j++];
}

int main() {
    for (scanf("%lld", &t); t; t--) {
        scanf("%lld%s", &n, s);
        memset(c, 0, sizeof c);
        for (ll i = 0; i < n; i++)
            c[s[i] - 'a']++;
        vector<plc> vec;
        for (ll i = 0; i < ALPHA; i++)
            vec.push_back({-c[i], i});
        sort(vec.begin(), vec.end());
        getMaxSave(vec);
        setMaxSave(vec);
        printf("%lld\n%s\n", n - maxSave, s);
    }
    return 0;
}
