#include <iostream>
#include <vector>

using namespace std;
typedef long long ll;

const ll D = 9;
ll n, m, x, ans;
vector<ll> A, B;

int main() {
    cin >> n >> m;
    for (ll i = 0; i < n; i++)
        cin >> x, A.push_back(x);
    vector<vector<ll>> I(n, vector<ll>());
    for (ll i = 0; i < m; i++)
        cin >> x, B.push_back(x);
    for (ll i = 0; i < n; i++)
        for (ll j = 0; j < m; j++)
            I[i].push_back(B[j]);
    for (ll d = D; d >= 0; d--) {
        bool flag = false;
        vector<vector<ll>> ne(n, vector<ll>());
        for (ll i = 0; i < n; i++) {
            if (!(A[i] & (1 << d))) {
                ne[i] = I[i];
                continue;
            }
            for (auto b : I[i])
                if (!(b & (1 << d)))
                    ne[i].push_back(b);
            if (ne[i].empty())
                flag = true, ne[i] = I[i];
        }
        if (flag)
            ans |= 1 << d;
        else
            I = ne;
    }
    cout << ans << endl;
    return 0;
}
