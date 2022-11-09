package main

import (
	"bufio"
	"fmt"
	"os"
)

////3
//func main() {
//	n := 0
//	fmt.Scan(&n)
//	res := make([][]int, n)
//	for i := range res {
//		res[i] = make([]int, n)
//	}
//	m := map[int]bool{}
//	for i := 1; i <= n*n; i++ {
//		m[i] = false
//	}
//	if dfs(res, m) {
//		for i := 0; i < len(res); i++ {
//			for j := 0; j < len(res[0]); j++ {
//				fmt.Print(res[i][j])
//				if j < len(res[0])-1 {
//					fmt.Print(" ")
//				}
//			}
//			fmt.Println()
//		}
//	} else {
//		fmt.Println(-1)
//	}
//
//}
//func isvaild(i, j, num int, res [][]int) bool {
//	if check(i, j-1, len(res)) && check(i-1, j, len(res)) && check(i-1, j-1, len(res)) {
//		return (num+res[i][j-1]+res[i-1][j]+res[i-1][j-1])%2 == 1
//	}
//	return true
//}
//func check(row, col, n int) bool {
//	if row >= 0 && row <= n-1 {
//		if col >= 0 && col <= n-1 {
//			return true
//		}
//	}
//	return false
//}
//func dfs(res [][]int, m map[int]bool) bool {
//	for i := 0; i <= len(res)-1; i++ {
//		for j := 0; j <= len(res)-1; j++ {
//
//			if res[i][j] != 0 {
//				continue
//			}
//			for num := range m {
//				if m[num] == false {
//					if isvaild(i, j, num, res) {
//						res[i][j] = num
//						m[num] = true
//						if dfs(res, m) {
//							return true
//						}
//						m[num] = false
//						res[i][j] = 0
//					}
//				}
//			}
//			return false
//		}
//	}
//	return true
//}

2
func getCount(lens, slide int) int {
	return lens - slide + 1
}
func main() {
	input := bufio.NewScanner(os.Stdin)
	bs := make([]byte, 2000*1024)
	input.Buffer(bs, len(bs))
	input.Scan()
	s := input.Text()
	input.Scan()
	k := input.Text()
	count := 0
	for i := 1; i < len(k); i++ {
		count += getCount(len(s), i)
	}
	tmp := ""
	for i := 0; i < len(s)-len(k)+1; i++ {
		tmp = s[i : i+len(k)]

		if tmp < k {
			count++
		}
	}
	fmt.Println(count)
}

//1
//#include <vector>
//#include <stack>
//#include <iostream>
//using namespace std;
//
//int main() {
//int n,q;
//cin >> n >> q;
//vector<int> nums;
//int tmp;
//for (int i=0; i<n;i++){
//cin >> tmp;
//nums.push_back(tmp);
//}
//int type, index, target;
//for(int i=0; i <q;i++) {
//cin >> type>>index>>target;
//if (type==1) {
//nums[index-1] = target;
//}
//else {
//int res = 0;
//for (int j=0; j < index; j++) {
//if (nums[j]==target) {
//res++;
//}
//}
//cout << res << endl;
//}
//}
//return 0;
//}
