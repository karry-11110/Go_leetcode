package main

import "math"

func minWindow(s string, t string) string {
	mT, mS := map[byte]int{}, map[byte]int{} //定义两个map映射两个字符串
	//处理好mT
	for i := 0; i < len(t); i++ {
		mT[t[i]]++
	}
	//定义判断当前ms包含mt不的函数
	check := func() bool {
		for k, v := range mT {
			if mS[k] < v {
				return false
			}
		}
		return true
	}
	length := math.MaxInt32
	resultL, resultR := -1, -1
	//进行滑动窗口
	for left, right := 0, 0; right < len(s); right++ {
		if mT[s[right]] > 0 { //进右边界数据
			mS[s[right]]++
		}
		for check() && left <= right { //处理左边界
			if right-left+1 < length {
				length = right - left + 1
				resultL, resultR = left, left+length
			}
			if _, ok := mT[s[left]]; ok {
				mS[s[left]]--
			}
			left++
		}
	}
	if resultL == -1 {
		return ""
	}
	return s[resultL:resultR]
}
