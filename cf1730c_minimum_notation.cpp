#include <iostream>
#include <cstring>
#include <algorithm>

using namespace std;
typedef long long ll;
const ll N = 2e5;
ll t;
char s[N + 10];

int main() {
    for (scanf("%lld", &t); t--; ) {
        scanf("%s", s);
        char c = '9' + 1;
        for (ll i = strlen(s) - 1; i >= 0; i--)
            if (s[i] <= c)
                c = s[i];
            else
                s[i] = min((char)(s[i] + 1), '9');
        sort(s, s + strlen(s));
        puts(s);
    }
    return 0;
}
