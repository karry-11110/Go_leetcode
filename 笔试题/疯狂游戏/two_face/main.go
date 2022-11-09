package main
import "fmt"
func main() {
	res := solve([2]int{4,2}, [][2]int{{1, 1},{2,1},{4,2},{6,5},{3,4},{1, 2}})
	fmt.Println(res)
}
func solve(point [2]int, graph [][2]int) bool {
	nCross := 0
	for i := 0; i < len(graph)-1; i++ {
		p1, p2 := graph[i], graph[(i+1)%len(graph)]
		if p1[1] == p2[1] || point[1] < min(p1[1],p2[1]) || point[1] >= max(p1[1],p2[1]) {
			continue
		}
		var x float32 = (float32)(point[1] -p1[1] ) * (float32)(p2[0] - p1[0]) / (float32)(p2[1] - p1[1]) + float32( p1[0])
		fmt.Println(p1, p2)
		fmt.Println(x)
		if x > float32(point[0]) {
			nCross++
		}
	}
	if nCross % 2 == 1{
		return true
	}
	return false

}
func min (i, j int) int {
	if i < j {
		return i
	}
	return j
}
func max (i, j int) int {
	if i > j {
		return i
	}
	return j
}
