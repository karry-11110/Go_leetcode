package main

import "sort"

func reconstructQueue(people [][]int) [][]int {
	//先将身高按从大到小！！！（注意点）的顺序排列，确定好相对位置(z注意点)
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	//再按照k值来排序,插入排序
	result := make([][]int, 0)
	for _, value := range people {
		//先拿进来，再将此元素之前的前k个整体后移
		result = append(result, value)
		copy(result[value[1]+1:], result[value[1]:])
		//再放入空挡
		result[value[1]] = value
	}
	return result
}
