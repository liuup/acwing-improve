
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

struct node {
    int to;
    int val;
};

int n;  // 牧场数
int x, y;   // 坐标
string s;   // 邻接矩阵

// 两点间距离
double len(int x1, int y1, int x2, int y2) {
    return sqrt(pow((x1-x2), 2) + pow((y1-y2), 2));
}

int main() {
    ios::sync_with_stdio(false); cin.tie(nullptr); cout.tie(nullptr);
    #ifdef ACM_DEBUG
    freopen("acm.txt", "r", stdin);
    #endif  

    cin >> n;
    vector<PII> locs(n+1);    // 每个牧场所对应的下标
    for(int i = 1; i <= n; i++) {
        cin >> x >> y;
        locs[i] = make_pair(x, y);
    }

    vector<vector<double>> dist(n+1, vector<double>(n+1));

    for(int i = 1; i <= n; i++) {
        cin >> s;
        for(int j = 1; j <= n; j++) {
            if(i != j) {
                if(s[j-1] == '1') {   // 连通则记录距离
                    dist[i][j] = len(locs[i].first, locs[i].second, locs[j].first, locs[j].second);
                } else {
                    dist[i][j] = inf;
                }
            }  
        }
    }
    
    // floyd
    for(int k = 1; k <= n; k++) {
        for(int i = 1; i <= n; i++) {
            for(int j = 1; j <= n; j++) {
                dist[i][j] = min(dist[i][j], dist[i][k] + dist[k][j]);
            }
        }
    }

    double ans1 = 0, ans2 = inf;

    vector<double> maxstep(n+1);

    for(int i = 1; i <= n; i++) {
        for(int j = 1; j <= n; j++) {
            if(dist[i][j] != inf) {
                maxstep[i] = max(dist[i][j], maxstep[i]);
            }
            ans1 = max(ans1, maxstep[i]);
        }
    }

    for(int i = 1; i <= n; i++) {
        for(int j = 1; j <= n; j++) {
            if(dist[i][j] == inf) {
                ans2 = min(maxstep[i] + len(locs[i].first, locs[i].second, locs[j].first, locs[j].second) + maxstep[j], ans2);
            }
        }
    }
    
    printf("%.6lf", max(ans1, ans2));


    return 0;
}
