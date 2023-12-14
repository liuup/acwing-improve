
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

#define inf 0x3f3f3f3f

auto printvector = [](vector<int> nums) { for(auto x:nums) {cout << x << " ";} cout << endl;};

struct edge {
    int to;
    int val;
};

int n, m;
int x, y;

int main() {
    ios::sync_with_stdio(false); cin.tie(nullptr); cout.tie(nullptr);
    #ifdef ACM_DEBUG
    freopen("acm.txt", "r", stdin);
    #endif  

    cin >> n >> m;
    int mod = 100003;

    vvi g(n+1);

    while(m--) {
        cin >> x >> y;
        g[x].push_back(y);
        g[y].push_back(x);
    }
   
    // bfs
    queue<int> q;
    vi dist(n+1, inf);
    vi cnt(n+1, 0); // 从1到达i的最短路径的条数

    dist[1] = 0;
    cnt[1] = 1;
    q.push(1);

    while(q.size()) {
        auto cur = q.front(); q.pop();

        // 记录一下这个点被更新时,更新他的前驱 
        for(auto e : g[cur]) {
            if(dist[cur] + 1 < dist[e] ) {
                dist[e] = dist[cur] +1;
                cnt[e] = cnt[cur];
                q.push(e);
            } else if(dist[cur]+1 == dist[e]) {
                cnt[e] = (cnt[e] + cnt[cur])%mod;
            }
        }
    }

    for(int i = 1; i <= n; i++) {
        cout << cnt[i] << endl;
    }

    return 0;
}
