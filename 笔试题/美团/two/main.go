package main
import "fmt"
func main() {
	n := 0
	fmt.Scan(&n)
	one, two, three, dis := [2]int{}, [2]int{}, [2]int{}, [3]int{}
	for i := 0; i < 2; i++ {
		fmt.Scan(&one[i])
	}
	for i := 0; i < 2; i++ {
		fmt.Scan(&two[i])
	}
	for i := 0; i < 2; i++ {
		fmt.Scan(&three[i])
	}
	for i := 0; i < 3; i++ {
		fmt.Scan(&dis[i])
	}
	fmt.Println(one, two, three, dis)
	for x := one[0]-dis[0]; x < one[0]+dis[0]+1; x++ {
		if x >= 1 && x <= n {
			y := one[1]-(dis[0]-abs(x-one[0]))
			if y >= 1{
				if abs(two[0]-x) + abs(two[1]-y) == dis[1] && abs(three[0]-x) + abs(three[1]-y) == dis[2] {
					fmt.Println(x, y)
					break
				}
			}
			y = one[1]+(dis[0]-abs(x-one[0]))
			if y <= n{
				if abs(two[0]-x) + abs(two[1]-y) == dis[1] && abs(three[0]-x) + abs(three[1]-y) == dis[2] {
					fmt.Println(x, y)
					break
				}
			}
		}
	}
}
func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
