func permutation(s string) []string {
	str := []byte(s)
	sort.Slice(str, func(i, j int) bool {
		return str[i] < str[j]
	})
	var result []string
	var path []byte
	used := make([]bool, len(str))
	var backtracking func(str []byte, start int)
	backtracking = func(str []byte, start int) {
		if len(path) == len(str) {
			result = append(result, string(path))
			return
		}
		for i := 0; i < len(str); i++ {
			if i > 0 && str[i] == str[i-1] && used[i-1] == false {
				continue
			}
			if used[i] == true {
				continue
			}
			used[i] = true
			path = append(path, str[i])
			backtracking(str, start+1)
			path = path[:len(path)-1]
			used[i] = false
		}
		return
	}
	backtracking(str, 0)
	return result
}