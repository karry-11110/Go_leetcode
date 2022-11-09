package main

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	word := make(map[string]bool)
	for _, value := range wordDict {
		word[value] = true
	}
	for i := 1; i < len(s)+1; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && word[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len((s))]
}
