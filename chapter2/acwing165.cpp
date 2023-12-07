// comment this line when upload

// #define ACM_DEBUG

#pragma GCC optimize(2)


#include <bits/stdc++.h>
using namespace std;

#define LL long long
#define ULL unsigned long long

#define PII pair<int,int>
#define all(a) a.begin(), a.end()

#define umap unordered_map
#define pq priority_queue

#define vi vector<int>
#define vvi vector<vector<int>>
auto printvector = [](vector<int> nums) { for(auto x:nums) {cout << x << " ";} cout << endl;};

const int N = 2e1;

int n, w;

int cat[N], cab[N];
int ans;

void dfs(int now, int cnt) {
    if(cnt >= ans) {
        return;
    }

    if(now == n+1) {
        ans = min(ans, cnt);
        return;
    }

    // 尝试分配到已经租用的缆车上
    for(int i = 1; i <= cnt; i++) {
        if(cab[i] + cat[now] <= w) {
            cab[i] += cat[now];
            dfs(now+1, cnt);
            cab[i] -= cat[now]; // 还原
        }
    }
    
    // 新开一辆缆车
    cab[cnt+1] = cat[now];
    dfs(now+1, cnt+1);
    cab[cnt+1] = 0;
}


int main() {
    ios::sync_with_stdio(false);
    std::cin.tie(nullptr);
    std::cout.tie(nullptr);
    #ifdef ACM_DEBUG
    freopen("acm.txt", "r", stdin);
    #endif

    cin >> n >> w;

    for(int i = 1; i <= n; i++) cin >> cat[i];

    sort(cat+1, cat+1+n, [](int a, int b){ return a > b; });

    ans = n;

    dfs(1, 0);

    cout << ans << endl;

    return 0;
}
