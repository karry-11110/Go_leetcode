package main
import(
	"fmt"
	"os"
	"bufio"
)
func main(){
	t := 0
	fmt.Scanln(&t)
	input := bufio.NewScanner(os.Stdin)
	for i := 0; i < t; i++ {
		input.Scan()
		input.Scan()
		str := input.Text()
		fmt.Printf("Case #%d:", i+1)
		solve(str)
		fmt.Println()
	}
}
func solve(s string) {
	if len(s) == 1{
		fmt.Printf(" 1")
		return
	}
	count := 1
	fmt.Printf(" 1")
	for i := 1; i < len(s); i++ {
		if s[i] > s[i-1] {
			count++
		} else {
			count = 1
		}
		fmt.Printf(" %d", count)
	}
	return
}
