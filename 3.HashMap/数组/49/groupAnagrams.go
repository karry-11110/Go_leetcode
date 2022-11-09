package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	//定义一个map，是字母异位词的对应同一个key
	m := map[string][]string{}
	for _, value := range strs {
		valueByte := []byte(value)
		sort.Slice(valueByte, func(i, j int) bool {
			return valueByte[i] < valueByte[j]
		})
		valueStr := string(valueByte)
		m[valueStr] = append(m[valueStr], value)
	}
	result := [][]string{}
	for _, value := range m {
		s
		result = append(result, value)
	}
	return result
}

// func groupAnagrams(strs []string) [][]string {
// 	mp := map[[26]int][]string{}
// 	for _, value := range strs {
// 		count := [26]int{}
// 		for _, str := range value {
// 			count[str-'a']++
// 		}
// 		mp[count] = append(mp[count], value)
// 	}
// 	result := [][]string{}
// 	for _, v := range mp {
// 		result = append(result, v)
// 	}
// 	return result
// }
