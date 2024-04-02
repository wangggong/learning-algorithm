#include <iostream>

using namespace std;
typedef long long ll;

string s, ans, c;
ll n, k, last;

int main() {
    cin >> n >> k >> s;
    for (; ans.size() < k; ans = ans + s);
    ans = ans.substr(0, k);
    for (ll i = s.size() - 1; i > 0; i--) {
        for (c = ""; c.size() < k; c = c + s.substr(0, i));
        c = c.substr(0, k);
        if (c < ans)
            ans = c;
    }
    cout << ans << endl;
    return 0;
}
