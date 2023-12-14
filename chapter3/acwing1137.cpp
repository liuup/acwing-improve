
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

int n, m, s;
int p, q, t;
int w;
int a;  // w个车站

int main() {
    ios::sync_with_stdio(false); cin.tie(nullptr); cout.tie(nullptr);
    #ifdef ACM_DEBUG
    freopen("acm.txt", "r", stdin);
    #endif  

    // 建反向图

    while(cin >> n >> m >> s) {
        vector<vector<edge>> g(n+1);
        vector<int> dist(n+1, inf);
        
        while(m--) {
            cin >> p >> q >> t;
            g[q].push_back(edge{p, t});
        }

        auto cmp = [](edge a, edge b){ return a.val > b.val; };
        pq<edge, vector<edge>, decltype(cmp)> q(cmp);
        
        cin >> w;
        vi ws;  // 求s到ws的最小距离
        while(w--) {
            cin >> a;
            ws.push_back(a);
        } 

        // 从s到各个点的dijkstra
        q.push(edge{s, 0});
        dist[s] = 0;

        while(q.size()) {
            auto cur = q.top(); q.pop();

            if(dist[cur.to] < cur.val) continue;

            for(auto e : g[cur.to]) {
                int d = dist[cur.to] + e.val;
                if(d < dist[e.to]) {
                    dist[e.to] = d;
                    q.push(edge{e.to, d});
                }
            }
        }

        int ans = inf;
        for(auto idx : ws) {
            ans = min(ans, dist[idx]);
        }

        if(ans > inf/2) {
            cout << -1 << endl;
        } else {
            cout << ans << endl;
        }
    }

    return 0;
}
