package main

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccuInfo := map[rune]int{}
	startIdx := 0
	ans := 0
	for i, ch := range []rune(s) {
		if lastIdx, ok := lastOccuInfo[ch]; ok && lastIdx >= startIdx {
			startIdx = lastIdx + 1

		}
		if i-startIdx+1 > ans {
			ans = i - startIdx + 1
		}
		lastOccuInfo[ch] = i
	}
	return ans
}
