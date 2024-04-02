#include <iostream>
#include <cstring>
#include <stack>
#define __FAST__ ios::sync_with_stdio(false), cin.tie(nullptr), cout.tie(nullptr)

using namespace std;
typedef long long ll;
typedef pair<ll, ll> pll;

const ll N = 2e5;
ll t, n, a[N + 10], c[N + 10], ans[N + 10];

int main() {
	__FAST__;
	cin >> t;
	while (t--) {
		memset(c, 0, sizeof c);
		memset(ans, -1, sizeof ans);
		cin >> n;
		for (ll i = 0; i < n; i++)
			cin >> a[i];
		for (ll i = 0; i < n; i++)
			c[a[i]]++;
		stack<pll> S;
		ll tot = 0, last = 0;
		for (ll i = 0; i <= n; i++) {
			if (tot < i)
				break;
			tot += c[i];
			if (i >= 1 && c[i - 1] == 0) {
				auto p = S.top(); S.pop();
				last += (i - 1) - p.first;
				if (p.second > 1)
					S.push({p.first, p.second - 1});
			}
			ans[i] = last + c[i];
			if (c[i] > 1)
				S.push({i, c[i] - 1});
		}
		for (ll i = 0; i <= n; i++)
			cout << ans[i] << ' ';
		cout << endl;
	}
	return 0;
}
