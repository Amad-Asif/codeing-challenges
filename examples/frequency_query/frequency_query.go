package main

import "fmt"

const (
	_ int32 = iota
	add
	remove
	present
)

/*
Need two data structures here
1 - Insert x in your data structure
2 - Delete one occurrence of y from your data structure, if present
3 - Check if any integer is present whose frequency is exactly z. If yes, print 1 else 0.
*/

var valCountStructure = map[int32]int32{}

func insertVal(val int32) {
	if v, ok := valCountStructure[val]; ok {
		valCountStructure[val] = v + 1
	} else {
		valCountStructure[val] = 1
	}
}

func removeVal(val int32) {
	if v, ok := valCountStructure[val]; ok {
		valCountStructure[val] = v - 1
		if valCountStructure[val] <= 0 {
			delete(valCountStructure, val)
		}
	}
}

// TODO Need to improve the check method, this is linear and very slow
func check(val int32) int32 {
	for _, v := range valCountStructure {
		if v == val {
			return 1
		}
	}
	return 0
}

func freqQuery(queries [][]int32) []int32 {
	var results []int32

	for _, q := range queries {
		if len(q) != 2 {
			continue
		}
		oper := q[0]
		val := q[1]

		switch oper {
		case add:
			insertVal(val)
		case remove:
			removeVal(val)
		case present:
			v := check(val)
			results = append(results, v)
		}
	}

	return results
}

func main() {
	input := [][]int32{{1, 1}, {2, 2}, {3, 2}, {1, 1}, {1, 1}, {2, 1}, {3, 2}}
	results := freqQuery(input)
	fmt.Println(results)
}
