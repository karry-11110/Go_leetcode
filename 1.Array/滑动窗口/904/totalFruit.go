package main

func totalFruit(fruits []int) int {
	//2种水果的最大数量，返回连续最大长度，问题k值化
	var mostK func(fruits []int, k int) int
	mostK = func(fruits []int, k int) int {
		//定义窗口左边界及结果
		left, result := 0, 0
		//存放种类数量的哈希表
		m := map[int]int{}
		//遍历水果树，即移动右边界。
		for right := 0; right < len(fruits); right++ {
			//未出现，加入，种类数-1
			if m[fruits[right]] == 0 {
				k--
			}
			//出现，则相应数量+1
			m[fruits[right]]++
			//种类数超过的话，窗口就要进行移动，也要处理相应的map,k值
			for k < 0 {
				//map数量改变
				m[fruits[left]]--
				//种类数出现变化
				if m[fruits[left]] == 0 {
					k++
				}
				//收缩左边界
				left++
			}
			//每一轮都要判断当前结果是否更新
			if result < right-left+1 {
				result = right - left + 1
			}
		}
		return result
	}
	return mostK(fruits, 2)
}
