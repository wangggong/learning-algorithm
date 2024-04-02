#include <iostream>
#include <vector>
#include <map>

using namespace std;

int main() {
    int n;
    cin >> n;
    vector<string> vec;
    map<string, int> m;
    string a, b;
    for (int i = 0; i < 2 * n - 2; i++) {
        string s;
        cin >> s;
        vec.push_back(s);
        if (!m.count(s))
            m[s] = i;
        if (s.size() == n - 1) {
            if (a == "")
                a = s;
            else
                b = s;
        }
    }
    string s;
    if (a.substr(1, n - 2) == b.substr(0, n - 2)) {
        s = a + b.substr(n - 2, 1);
        for (int i = 1; i < n; i++)
            if (!m.count(s.substr(0, i))) {
                s = b + a.substr(n - 2, 1);
                break;
            }
    }
    else
        s = b + a.substr(n - 2, 1);
    string ans(2 * n - 2, 'S');
    for (int i = 1; i < n; i++)
        ans[m[s.substr(0, i)]] = 'P';
    cout << ans << endl;
    return 0;
}

