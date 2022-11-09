package main

func canConstruct(ransomNote string, magazine string) bool {
	mgz := []byte(magazine)
	rs := []byte(ransomNote)
	var a1 [26]int
	for i := 0; i < len(mgz); i++ {
		a1[mgz[i]-'a']++
	}
	for i := 0; i < len(rs); i++ {
		a1[rs[i]-'a']--
		if a1[rs[i]-'a'] < 0 {
			return false
		}
	}
	return true
}

// func canConstruct(ransomNote string, magazine string) bool {

// 	var a1 [26]int
// 	for _, value := range magazine {
// 		a1[value-'a']++
// 	}
// 	for _, value := range ransomNote {
// 		a1[value-'a']--
// 		if a1[value-'a'] < 0 {
// 			return false
// 		}
// 	}
// 	return true
// }
