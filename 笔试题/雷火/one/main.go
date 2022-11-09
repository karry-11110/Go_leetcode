package main

import (
"fmt"
)

func main() {
	input := [][4]int{
		{1,1,12,5},
		{10,1,12,9},
		{8,8,12,10},
		{1,1,8,10},
	}
	max_x,max_y := 0, 0
	for i := range input {
	  if input[i][2] > max_x {
		  max_x = input[i][2]
	  }
		if input[i][3] > max_y {
			max_y = input[i][3]
		}
	}
	//max_x 是列的范围，max_y是行的范围
	graph := make([][]int, max_y+1)
	for i := range graph {
		graph[i] = make([]int, max_x+1)
	}
	for i := 0; i < len(input); i++ {
		for x := input[i][1];  x <= input[i][3]; x++ {
			for y := input[i][0]; y <= input[i][2]; y++ {
				if graph[x][y] == 0{
					graph[x][y] = 1
				}
			}
		}
	}

	result := make([][2]int, 0)
	for x := 0; x < len(graph); x++ {
		for y := 0; y < len(graph[0]); y++ {
			if graph[x][y] == 1 {
				if check(graph, [2]int{x, y}) {
					result = append(result, [2]int{x, y})
				}
			}
		}
	}
	fmt.Println(result)
}
func check(graph [][]int, point [2]int) bool {
	point_top_x, point_top_y := point[0] - 1, point[1]
	point_bottom_x, point_bottom_y := point[0] +1 , point[1]
	point_left_x, point_left_y := point[0], point[1] - 1
	point_right_x, point_right_y := point[0], point[1] + 1
	v1, v2 , v3, v4 := 0,0,0,0
	if point_top_x < 0 || point_top_x  > len(graph)-1 || point_top_y < 0 || point_top_y > len(graph[0])-1 {
		v1 = -1
	} else {
		v1 = graph[point_top_x][point_top_y]
	}
	if point_bottom_x < 0 || point_bottom_x  > len(graph)-1 ||point_bottom_y < 0 || point_bottom_y > len(graph[0])-1 {
		v2 = -1
	} else {
		v2 = graph[point_bottom_x][point_bottom_y]
	}
	if point_left_x < 0 || point_left_x  > len(graph)-1 || point_left_y < 0 || point_left_y > len(graph[0])-1 {
		v3 = -1
	} else {
		v3 =graph[point_left_x][point_left_y]
	}
	if point_right_x < 0 || point_right_x  > len(graph)-1 || point_right_y < 0 || point_right_y > len(graph[0])-1 {
		v4 = -1
	} else {
		v4 = graph[point_right_x][point_right_y]
	}
    if v1 == 1 && v2 == 1 && v3 == 1 && v4 == 1 {
		if graph[point[0]-1][point[1]-1]
	}
	return false
}

