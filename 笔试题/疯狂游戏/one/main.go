package main
import(
	"fmt")
var res [][]int

func main() {
	sumValue := 0
	fmt.Scan(&sumValue)
	r := solve(sumValue)
	fmt.Println(r)
	//for i := 0; i < len(res)-1; i++ {
	//	fmt.Printf("%d,", res[i])
	//}
	//fmt.Printf("%d", res[len(res)-1])
	//fmt.Println()
}
//func solve(sumValue int) [][]int {
//	path := []int{1,4,5}
//	dfs(sumValue, 0, path)
//	return res
//}
//func dfs(target, sum int, path []int) {
//	if sum == target {
//		tmp := make([]int, len(path))
//		copy(tmp, path)
//		res = append(res, tmp)
//		return
//	}
//	if sum > target {
//		return
//	}
//	if sum+5 <= target {
//		path[2]++
//		dfs(target, sum+5, path)
//		path[2]--
//	}
//	if sum+4 <= target {
//		path[1]++
//		dfs(target, sum+4, path)
//		path[1]--
//	}
//	if sum+1 <= target {
//		path[0]++
//		dfs(target, sum+1, path)
//		path[0]--
//	}
//	return
//}

func solve(sumValue int) []int {
	dp := make([][2]int, sumValue+1)
	dp[0][0] = 0
	for i := 1; i <= sumValue; i++ {
		if i -1 >= 0 {
			dp[i][0] = dp[i-1][0]
			dp[i][1] = i-1
		}
		if i-4 >= 0 && dp[i][0] > dp[i-4][0] {
			dp[i][0] = dp[i-4][0]
			dp[i][1] = i-4
		}
		if i-5 >= 0 && dp[i][0] > dp[i-5][0] {
			dp[i][0] = dp[i-5][0]
			dp[i][1] = i-5
		}
		dp[i][0] += 1
	}
	tail := sumValue
	tmp := [6]int{0,0,0,0,0,0}
	for tail >0 {
		tmp[tail-dp[tail][1]] += 1
		tail = dp[tail][1]
	}
	res := []int{tmp[1], tmp[4], tmp[5]}
	return res
}
