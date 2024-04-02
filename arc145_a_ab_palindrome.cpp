#include <iostream>
#include <cstring>

using namespace std;
typedef long long ll;

const ll N = 2e5;
ll n;
char s[N + 10];

bool check() {
    if (!strcmp(s, "BA"))
        return false;
    if (s[0] == 'A' && s[n - 1] == 'B')
        return false;
    return true;
}

int main() {
    scanf("%lld%s", &n, s);
    puts(check() ? "Yes" : "No");
    return 0;
}
