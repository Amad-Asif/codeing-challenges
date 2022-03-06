package examples

import "fmt"

/*
Each time Sunny and Johnny take a trip to the Ice Cream Parlor, they pool their money to buy ice cream. On any given day,
the parlor offers a line of flavors. Each flavor has a cost associated with it.

Given the value of money and the cost of each flavor for t trips to the Ice Cream Parlor, help Sunny and Johnny choose two distinct
flavors such that they spend their entire pool of money during each visit. ID numbers are the 1- based index number associated with a . For each trip to the parlor, print the ID numbers for the two types of ice cream that Sunny and Johnny purchase as two space-separated integers on a new line. You must print the smaller ID first and the larger ID second.

*/

func IceCreamParlor(cost []int32, money int32) {
	costMap := map[int32][]int{}

	// Cost as key and index as value
	for i := len(cost) - 1; i >= 0; i-- {
		// Keep cost of each as a slice to cater for duplicates
		if costMap[cost[i]] == nil {
			costMap[cost[i]] = []int{}
		}
		costMap[cost[i]] = append(costMap[cost[i]], i)
	}

	var ind1, ind2 int
	for i := 0; i < len(cost); i++ {
		ind1 = i
		c1 := cost[i]
		remCost := money - c1
		if indArr, ok := costMap[remCost]; ok {
			for _, v := range indArr {
				// Just to ensure index other than the one already selected
				if v != ind1 {
					ind2 = v
					break
				}
			}
			break
		}
	}

	fmt.Printf("%d %d\n", ind1+1, ind2+1)
}
