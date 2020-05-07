#include <bits/stdc++.h>

using namespace std;
const int MAX  =  1e5+7;
int a[MAX];

int main(){
    int n, k, q;

    cin >> n >> k >> q;
    int cnt = 0;

    for (int i = 1; i <= n; i++)
        a[i] = k;

    for (int i = 1; i <= q; i++) {
        cin>>k;
        a[k]++;
    }

    for (int i = 1; i <= n; i++) {
        if (a[i] <= q)
            puts("No");
        else
            puts("Yes");
    }

    return 0;
}
