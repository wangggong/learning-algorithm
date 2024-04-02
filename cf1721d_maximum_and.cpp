#include <iostream>
#include <vector>

using namespace std;
typedef long long ll;
typedef vector<ll> vll;
typedef pair<vll, vll> pvv;
const ll D = 30;
ll t, n, x;

int main() {
    scanf("%lld", &t);
    while (t--) {
        vll a, b;
        scanf("%lld", &n);
        for (ll i = 0; i < n; i++)
            scanf("%lld", &x), a.push_back(x);
        for (ll i = 0; i < n; i++)
            scanf("%lld", &x), b.push_back(x);
        ll ans = 0;
        vector<pvv> Q;
        Q.push_back({a, b});
        for (ll d = D; d >= 0; d--) {
            bool valid = true;
            vector<pvv> nQ;
            for (auto q : Q) {
                vll na[2];
                vll nb[2];
                for (auto v : q.first)
                    na[(v >> d) & 1].push_back(v);
                for (auto v : q.second)
                    nb[(v >> d) & 1].push_back(v);
                if (na[0].size() != nb[1].size()) {
                    valid = false;
                    break;
                }
                if (na[0].size())
                    nQ.push_back({na[0], nb[1]});
                if (na[1].size())
                    nQ.push_back({na[1], nb[0]});
            }
            if (valid) {
                Q = nQ;
                ans |= 1 << d;
            }
        }
        printf("%lld\n", ans);
    }
    return 0;
}
