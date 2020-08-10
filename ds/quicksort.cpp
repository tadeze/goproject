#include<iostream>
#include<vector>
using namespace std;

int partition(vector<int> &a, int l, int h){

    int i = l - 1; 
    for(int j=l;j<h;j++){
        if(a[j] < a[h]){
            i++;
            swap(a[i], a[j]);
        }
    }
    swap(a[i+1], a[h]);
    return i+1;

}


void quick_sort(vector<int> &a, int l, int h){
    if (l < h){
        int p = partition(a, l, h);
        quick_sort(a, l, p-1);
        quick_sort(a, p+1, h);
    }
}

int main() {

    vector<int> a = { 3,10, 9,4,2,7 }; 
    quick_sort(a, 0, a.size() - 1);

    for(auto e : a) cout<<e;
        
}
