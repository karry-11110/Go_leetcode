package main
import(
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"sort")

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	str := input.Text()[:len(input.Text())-2]
	strs := strings.Split(str,",")
	nums := []int{}
	for i := 0; i < len(strs); i++ {
		tmp, _ := strconv.Atoi(strs[i])
		nums = append(nums, tmp)
	}
	nb := input.Text()[len(input.Text())-1]
	n, _ := strconv.Atoi(string(nb))
	sort.Ints(nums)
	fmt.Println(solve(nums, n))
}
func solve(nums []int, target int) string {
	tmp := make([]int, len(nums))
	return dfs(0, 0, target, "", nums, tmp, 0)
}
func dfs(index, sum, target int, s string, price []int, tmp []int, tail int) string {
	if sum > target  {
		return s
	}
	if sum == target {
		t := ""
		for i := 0; i < tail; i++ {
			t += strconv.Itoa(tmp[i]) + ","
		}

		if s == ""{
			s = t[: len(t)]
		} else {
			s = s + ";" + t[: len(t)]
		}
		return s
	}
	if index >= len(price) {
		return s
	}
	tmp[tail] = price[index]
	tail++
	t1 := dfs(index+1, sum+price[index], target, s, price, tmp, tail)
	tail--
	t2 := dfs(index+1, sum, target, t1, price, tmp, tail)
	return t2
}
