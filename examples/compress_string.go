package examples

import "fmt"

/*
Input: chars = ["a","a","b","b","c","c","c"]
Output: Return 6, and the first 6 characters of the input array should be: ["a","2","b","2","c","3"]
Explanation: The groups are "aa", "bb", and "ccc". This compresses to "a2b2c3"
*/

func CompressString(chars []byte) (string, int) {
	getRepetitions := func(a string, index int) (string, int) {
		count := 0
		i := index
		for i < len(chars) {
			if string(chars[i]) == a {
				index += 1
				count += 1
			} else {
				break
			}
			i += 1
		}
		if count == 1 {
			return a, index
		}

		return fmt.Sprintf("%s%d", a, count), index
	}

	finalStr := ""
	i := 0
	for i < len(chars) {
		newStr, index := getRepetitions(string(chars[i]), i)
		i = index
		finalStr += newStr
	}
	fmt.Println(finalStr)
	return finalStr, len(finalStr)
}
