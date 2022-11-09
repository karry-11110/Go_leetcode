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
		n, _ := strconv.Atoi(strings.Split(input.Text(), " ")[0])
		c, _ := strconv.Atoi(strings.Split(input.Text(), " ")[1])
		nums := make([][2]int, n)
		for j := 0; j < n; j++ {
			input.Scan()
			nums[j][0], _ = strconv.Atoi(strings.Split(input.Text(), " ")[0])
			nums[j][1], _ = strconv.Atoi(strings.Split(input.Text(), " ")[1])
		}
		result := solve(nums, c)
		fmt.Printf("Case #%d: %d", i+1, result)
		fmt.Println()
	}
}
func solve(nums [][2]int, c int) int {
	return 1
}

