// 参考: https://codeforces.com/problemset/submission/26/230000559
#include <iostream>

using namespace std;
typedef long long ll;
string s;
ll ans, stk;

int main() {
    cin >> s;
    ans = s.size();
    for (char c : s)
        if (c == '(')
            stk++;
        else if (stk > 0)
            stk--;
        else
            ans--;
    cout << ans - stk << endl;
    return 0;
}
