package examples

import (
	"fmt"
	"strconv"
	"unicode"
)

// "3[a]2[bc]"
func DecodeString(s string) string {
	var getSequence func(index, num int) (string, int)

	extractNumber := func(index int) (int, int) {
		numStr := ""
		for i := index; i < len(s); i++ {
			if unicode.IsNumber(rune(s[i])) {
				numStr += string(s[i])
				index += 1
			} else {
				break
			}
		}
		num, _ := strconv.Atoi(numStr)
		return num, index
	}

	getSequence = func(index, num int) (string, int) {
		seqStr := ""
		for i := index; i < len(s); i++ {
			if unicode.IsNumber(rune(s[i])) {
				num, ind := extractNumber(i)
				i = ind
				if string(s[i]) == "[" {
					i += 1
					res, fin := getSequence(i, num)
					i = fin
					seqStr += res
					index = fin
				}
			}
			if string(s[i]) == "]" {
				break
			}
			seqStr += string(s[i])
			index += 1
		}

		final := ""
		for i := 0; i < num; i++ {
			final += seqStr
		}

		return final, index
	}
	var finSeq string
	for i := 0; i < len(s); i++ {
		if unicode.IsNumber(rune(s[i])) {
			num, ind := extractNumber(i)
			i = ind
			if string(s[i]) == "[" {
				i += 1
				res, fin := getSequence(i, num)
				i = fin
				finSeq += res
			}
		} else if string(s[i]) != "]" {
			finSeq += string(s[i])
		}
	}
	fmt.Println(finSeq)
	return finSeq
}
