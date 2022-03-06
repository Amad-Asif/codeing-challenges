package examples

/*
It is New Year's Day and people are in line for the Wonderland roller-coaster ride. Each person wears a sticker indicating
their initial position in the queue from 1 to N. Any person can bribe the person directly in front of them to swap positions,
but they still wear their original sticker. One person can bribe at most two others.

Determine the minimum number of bribes that took place to get to a given queue order. Print the number of bribes, or,
if anyone has bribed more than two people, print Too chaotic.

minimumBribes has the following parameter(s):

int q[n]: the positions of the people after all bribes

*/

/*
CASES -
2 1 5 3 4
2 5 1 3 4
5 1 2 3 7 8 6 4
1 2 5 3 7 8 6 4
1 2 5 3 4 7 8 6


*/

// MinimumBribes
func MinimumBribes(arr []int32) (int32, bool) {
	max := func(a, b int32) int32 {
		if a > b {
			return a
		}
		return b
	}

	var isChaos bool
	var bribes int32
	for i := int32(0); i < int32(len(arr)); i++ {
		// Get a slice of positions as input, start from first element and check if (index - value_at_index) is greater than 2
		// If the difference is greater than 2 then break
		if arr[i]-(i+1) > 2 {
			isChaos = true
			break
		}

		// Since person can't bribe more than twice so just check last 2 values
		// Get max to avoid Panic
		jMax := max(0, arr[i]-2)
		for j := jMax; j <= i; j++ {
			if arr[j] > arr[i] {
				bribes++
			}
		}
	}

	return bribes, isChaos
}
