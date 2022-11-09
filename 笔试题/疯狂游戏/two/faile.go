//package main
//import(
//	"fmt"
//	"bufio"
//	"os"
//	"strings"
//	"strconv"
//	"sort")
//
//func main() {
//	input := bufio.NewScanner(os.Stdin)
//	input.Scan()
//	str := input.Text()[:len(input.Text())-2]
//	strs := strings.Split(str,",")
//	nums := []int{}
//	for i := 0; i < len(strs); i++ {
//		tmp, _ := strconv.Atoi(strs[i])
//		nums = append(nums, tmp)
//	}
//	nb := input.Text()[len(input.Text())-1]
//	number, _ := strconv.Atoi(string(nb))
//	res := solve(nums, number)
//
//	if len(res) != 0 {
//		st(res)
//		var i int
//		var j int
//		for i = 0; i < len(res); i++ {
//			for j = 0; j < len(res[i])-1; j++ {
//				fmt.Printf("%d,", res[i][j])
//			}
//			if i != len(res)-1 {
//				fmt.Printf("%d;", res[i][j])
//			} else {
//				fmt.Printf("%d", res[i][j])
//			}
//
//		}
//	}
//
//}
//func st(res [][]int) {
//	sort.Slice(res, func(i, j int)bool {
//		return len(res[i]) < len(res[j])
//	})
//	sort.Slice(res, func(i, j int)bool {
//		if len(res[i]) == len(res[j]) {
//			if res[i][0] ==  res[j][0] {
//				return res[i][1] < res[j][1]
//			} else {
//				return res[i][0] < res[j][0]
//			}
//		}
//		return false
//	})
//
//}
//func solve(nums []int, target int) [][]int {
//	var track, res = []int{}, [][]int{}
//	history := make(map[int]bool)
//	sort.Ints(nums)
//	back(0, 0, target, nums, track, &res, history)
//	return res
//}
//func back(index, sum, target int, nums, track []int, res *[][]int, history map[int]bool){
//	if sum == target {
//		tmp := make([]int, len(track))
//		copy(tmp, track)
//		*res = append(*res, tmp)
//		return
//	}
//	if sum > target {
//		return
//	}
//	for i := index; i < len(nums); i++ {
//		if i > 0 && nums[i] == nums[i-1]&&history[i-1]==false {
//			continue
//		}
//		track = append(track, nums[i])
//		sum += nums[i]
//		history[i] = true
//		back(i+1, sum, target, nums, track, res, history)
//		track = track[:len(track)-1]
//		sum -= nums[i]
//		history[i] = false
//	}
//}
