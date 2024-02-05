#pragma GCC optimize(2)
// main.cpp for upload

#include <bits/stdc++.h>
using namespace std;

#define __FASTIO__ ios::sync_with_stdio(false);cin.tie(nullptr);cout.tie(nullptr);

#define int long long
#define ll long long
#define ull unsigned long long

#define all(a) a.begin(), a.end()
#define fori(i, a, b) for(ll i=a;i<b;i++)
#define endl "\n"
#define umap unordered_map
#define pq priority_queue
#define pb push_back
#define pii pair<int,int>
#define vi vector<int>
#define vll vector<long long>
#define vvi vector<vector<int>>
#define len(x) (int)((x).size())

#define inf 1e16

template<typename T>void coutvc(vector<T> nums){for(auto x:nums){cout<<x<<" ";}cout<<endl;}
void coutb(int a){cout<<bitset<sizeof(a)*8>(a)<<endl;}  // 打印数字对应的二进制

void _yes(){cout<<"YES"<<endl;}
void _no(){cout<<"NO"<<endl;}
void _1(){cout<<-1<<endl;}

struct node {
    int to;
    int val;
};

void solve();

// 一组还是多组数据

#define SOLVE_SINGLE

// modify

signed main(void) {
    __FASTIO__
    #ifdef ACM_DEBUG
    freopen("acm.in", "r", stdin); 
    #endif

    #ifdef SOLVE_SINGLE
    solve();
    #else
    int t;cin>>t;while(t--){solve();}
    #endif

    #ifdef ACM_DEBUG
    fclose(stdin);
    #endif
    return 0;
}


// 线段树,懒标记
class LazySeg {
public:
    const int todoInit = 0;
    struct seg {
        int l, r;
        int sum;
        int todo;    
    };
    vector<seg> t;

    LazySeg() {};
    LazySeg(const vector<int>& nums);

    void build(const vector<int>& nums, int o, int l, int r);
    void maintain(int o);
    int mergeInfo(int a, int b);

    void spread(int o);
    void do_(int o, int v);

    void update(int o, int l, int r, int v);
    int query(int o, int l, int r);
    int queryAll();
};

LazySeg::LazySeg(const vector<int>& nums) {
    int n = nums.size();
    if(n == 0) {
        cout << "can't be empty!" << endl;
        return ;
    }
    t = vector<seg>(4 * n);   // 图方便开4 * n
    build(nums, 1, 1, n);
}

void LazySeg::build(const vector<int>& nums, int o, int l, int r) {
    t[o].l = l; t[o].r = r;
    t[o].todo = todoInit;
    if(l == r) { t[o].sum = nums[l-1]; return; }
    int m = (l + r) >> 1;
    build(nums, o<<1, l, m);
    build(nums, o<<1|1, m+1, r);
    maintain(o);
}

void LazySeg::maintain(int o) {
    t[o].sum = mergeInfo(t[o<<1].sum, t[o<<1|1].sum);
}

int LazySeg::mergeInfo(int a, int b) {
    // 或者其他的操作
    // return a + b;
    return max(a, b);
}

void LazySeg::spread(int o) {
    int v = t[o].todo;
    if(v != todoInit) {
        do_(o<<1, v);
        do_(o<<1|1, v);
        t[o].todo = todoInit;
    }
}

void LazySeg::do_(int o, int v) {
    t[o].sum += v * (t[o].r - t[o].l + 1);  // 更新v对整个区间的影响
    t[o].todo += v; // 更新v对左右儿子的影响
}

// 区间修改
void LazySeg::update(int o, int l, int r, int v) {
    if(l <= t[o].l && t[o].r <= r) { do_(o, v); return; }
    spread(o);
    int m = (t[o].l + t[o].r) >> 1;
    if(l <= m) update(o<<1, l, r, v);
    if(m < r) update(o<<1|1, l, r, v);
    maintain(o);
}

int LazySeg::query(int o, int l, int r) {
    if(l <= t[o].l && t[o].r <= r) { return t[o].sum; }
    spread(o);
    int m = (t[o].l + t[o].r) >> 1;
    if(r <= m) return query(o<<1, l, r);
    if(m < l) return query(o<<1|1, l, r);
    int vl = query(o<<1, l, r);
    int vr = query(o<<1|1, l, r);
    return mergeInfo(vl, vr);
}

int LazySeg::queryAll() { return t[1].sum; }

void solve() {
    int m, p; cin >> m >> p;
    string op; int num;

    auto sg = LazySeg(vector<int>(2 * 1e5, 0)); // 先随便建一个长度吧

    int qans = 0;   // 询问的答案

    int length = 0;
    fori(i, 0, m) {
        cin >> op >> num;
        if(op == "A") { // 添加
            length++;

            sg.update(1, length, length, (num + qans) % p);
        } else {    // 问询
            int tmp = sg.query(1, length-num+1, length);
            qans = tmp;
            cout << tmp << endl;
        }
    }
}
