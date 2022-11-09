package main
import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)
func main(){
	t := 0
	fmt.Scanln(&t)
	input := bufio.NewScanner(os.Stdin)
	for i := 0; i < t; i++ {
		input.Scan()
		r, _ := strconv.Atoi(strings.Split(input.Text(), " ")[0])
		c, _ := strconv.Atoi(strings.Split(input.Text(), " ")[1])
		nums := make([][]int, r)
		for i := 0; i < r; i++ {
			input.Scan()
			nums[i] = make([]int, c)
			for j := 0; j < c; j++ {
				nums[i][j], _ = strconv.Atoi(strings.Split(input.Text(), " ")[j])
			}
		}
		result := solve(nums)
		fmt.Printf("Case #%d: %d", i+1, result)
		fmt.Println()
	}
}
func solve(nums [][]int) int {
	return 1
}

