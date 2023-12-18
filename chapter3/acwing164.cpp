
// #define ACM_DEBUG   // comment this line when upload !!!

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
#define pb push_back
auto printvector = [](vector<int> nums) { for(auto x:nums) {cout << x << " ";} cout << endl;};

struct node {
    // int from;
    int to;
    int val;
};

int n, m;
int x, y;

const int N = 3e4+10;
bitset<N> f[N];


int main() {
    ios::sync_with_stdio(false); cin.tie(nullptr); cout.tie(nullptr);
    #ifdef ACM_DEBUG
    freopen("acm.txt", "r", stdin);
    #endif  

    cin >> n >> m;

    vvi g(n+1);
    vi in(n+1);

    for(int i = 0; i < m; i++) {
        cin >> x >> y;
        g[x].pb(y);
        in[y]++;    // 入度统计
    }

    // 拓扑排序
    queue<int> q;
    for(int i = 1; i <= n; i++) {
        if(!in[i]) q.push(i);
    }

    vi path;

    while(q.size()) {
        auto cur = q.front(); q.pop();

        path.pb(cur);

        for(auto e : g[cur]) {
            in[e]--;
            if(in[e] == 0) {
                q.push(e);
            }
        }
    }

    // 从后往前递推
    for(int i = n-1; i >= 0; i--) {
        int j = path[i];
        f[j][j] = 1;    // 这个点可以到达自己
        for(int cur : g[j]) { // 所有能到达的点
            f[j] |= f[cur];
        }
    }

    for(int i = 1; i <= n; i++) {
        cout << f[i].count() << endl;
    }

    return 0;
}