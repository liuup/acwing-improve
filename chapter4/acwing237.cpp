
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
    // int from;
    int to;
    int val;
};


// 并查集模板
struct Unionset {
    // vector<int> father;
    unordered_map<int, int> father; // 离散数据
    // int count;

    Unionset(int n);
    int Find(int i);
    void Union(int i, int j);
    bool isConnected(int i, int j);
};

Unionset::Unionset(int n) {
    // father = vector<int>(n+1);
    // for(int i = 1; i <= n; i++) father[i] = i;

    father = unordered_map<int, int>(); // 离散数据    
    // count = n;
}

int Unionset::Find(int i) {
    if(this->father[i] == i) {
        return i;
    }
    this->father[i] = this->Find(this->father[i]);  // 路径压缩
    return this->father[i];
}

void Unionset::Union(int i, int j) {
    int i_fa = this->Find(i);
    int j_fa = this->Find(j);

    if(i_fa == j_fa) {
        return;
    }
    this->father[i_fa] = j_fa;
    // this->count--;
}

bool Unionset::isConnected(int i, int j) {
    return this->Find(i) == this->Find(j);
}

int t;
int n;
int x, y, z;    //  表示问题 1相等约束；0不相等约束

auto us = Unionset(1);

int main() {
    ios::sync_with_stdio(false); cin.tie(nullptr); cout.tie(nullptr);
    #ifdef ACM_DEBUG
    freopen("acm.txt", "r", stdin);
    #endif  

    cin >> t;
    while(t--) {
        cin >> n;

        us.father.clear();

        vector<pair<int, int>> eql; // 相等的约束
        vector<pair<int, int>> uneql;   // 不相等的约束
        
        bool fg = false;
        for(int i = 0; i < n; i++) {
            cin >> x >> y >> z;
            if(z == 1) {
                eql.pb(make_pair(x, y));
            } else {
                uneql.pb(make_pair(x, y));
            }
        }

        for(auto p : eql) {
            us.father[p.first] = p.first;
            us.father[p.second] = p.second;
        }

        for(auto p : eql) {
            us.Union(p.first, p.second);
        }

        for(auto p : uneql) {
            if(us.father[p.first] == 0) {
                us.father[p.first] = p.first;
            }
            if(us.father[p.second] == 0) {
                us.father[p.second] = p.second;
            }
        }
        

        // 最后还要检查一遍满不满足不相等的约束
        for(auto p : uneql) {
            if(us.isConnected(p.first, p.second)) {
                fg = true;
                break;
            }
        }

        if(fg) {
            cout << "NO" << endl;
        } else {
            cout << "YES" << endl;
        }
    }


    return 0;
}
