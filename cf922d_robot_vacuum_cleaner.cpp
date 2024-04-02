#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

typedef long long ll;

ll cnt(vector<string> &vec) {
    ll ans = 0, cs = 0;
    for (string s : vec)
        for (char c : s)
            if (c == 's')
                cs++;
            else
                ans += cs;
    return ans;
}

int main() {
    int n;
    cin >> n;
    vector<string> vec;
    for (int i = 0; i < n; i++) {
        string s;
        cin >> s;
        vec.push_back(s);
    }
    sort(vec.begin(), vec.end(), [](string &s, string &t) {
        vector<string> st = vector<string>{s + t}, ts = vector<string>{t + s};
        return cnt(st) > cnt(ts);
    });
    cout << cnt(vec) << endl;
    return 0;
}
