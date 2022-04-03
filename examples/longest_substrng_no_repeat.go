package examples

// GetLongestSubstringNoRepetitions
func GetLongestSubstringNoRepetitions(s string) int {
	var longest int
	var theChars = map[string]struct{}{}
	initial := 0
	for i := 0; i < len(s); i++ {
		r := string(s[i])
		if len(theChars) == 0 {
			initial = i
		}
		if _, ok := theChars[r]; !ok {
			theChars[r] = struct{}{}
			continue
		}
		if initial < len(s) {
			i = initial
		}
		curEles := len(theChars)
		if curEles > longest {
			longest = curEles
		}
		theChars = map[string]struct{}{}
	}
	finalLen := len(theChars)
	if longest == 0 || finalLen > longest {
		longest = len(theChars)
	}
	return longest
}
