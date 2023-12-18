
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

#define inf 0x3f3f3f3f

auto printvector = [](vector<int> nums) { for(auto x:nums) {cout << x << " ";} cout << endl;};

struct node {
    int to;
    int val;
};

int f;
int n, m, w;
int s, e, t;

// spfa

bool spfa(vector<vector<node>> g, int n) {
    vector<int> dist(n+1, inf);
    vector<bool> vis(n+1);
    vector<int> cnt(n+1);

    queue<int> q;
    for(int i = 1; i <= n; i++) q.push(i);

    while(q.size()) {
        auto cur = q.front(); q.pop();

        vis[cur] = false;

        for(auto e : g[cur]) {
            int v = e.to;
            int w = e.val;
            if(dist[v] > dist[cur] + w) {
                dist[v] = dist[cur] + w;
                cnt[v] = cnt[cur] + 1;  // 记录边数
                if(cnt[v] >= n) {
                    // cout << "Yes" << endl;
                    return true;
                }
                if(!vis[cur]) {
                    q.push(v);
                    vis[v] = true;
                }
            }
        }
    }
    return false;
}

int main() {
    ios::sync_with_stdio(false); cin.tie(nullptr); cout.tie(nullptr);
    #ifdef ACM_DEBUG
    freopen("acm.txt", "r", stdin);
    #endif  

    cin >> f;

    while(f--) {
        cin >> n >> m >> w;
        vector<vector<node>> g(n+1);

        for(int i = 0; i < m; i++) {
            cin >> s >> e >> t;
            g[s].pb(node{e, t});
            g[e].pb(node{s, t});
        }
        for(int i = 0; i < w; i++) {
            cin >> s >> e >> t;
            g[s].pb(node{e, -t});
        }

        // 虚拟节点
        for(int i = 1; i <= n; i++) {
            g[i].pb(node{i, 0});
        }

        if(spfa(g, n)) {
            cout << "YES" << endl;
        } else {
            cout << "NO" << endl;
        }
    }

    return 0;
}
