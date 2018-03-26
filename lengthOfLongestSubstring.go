func lengthOfLongestSubstring(s string) int {
	runes := []rune(s)
	start := 0
	length := 0
	longest := 0
	runeMap := make(map[rune]int)
	for i,rn:=range runes{
		if val,ok:=runeMap[rn];!ok||val<i-start {
			
			length = start+1
		}else{
			length = i - val
		}
	
		if length > longest{
			longest = length
		}
		runeMap[rn] = i
		start = length	
	}
	return longest
}