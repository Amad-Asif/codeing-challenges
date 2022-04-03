package examples

import (
	"fmt"
	"sort"
)

func GetMedian(arr []int) float64 {
	var val float64

	sort.Ints(arr)
	mNum := len(arr) / 2
	if len(arr)%2 == 0 {
		val = (float64(arr[mNum-1]) + float64(arr[mNum])) / 2.0
	} else {
		val = float64(arr[mNum])
	}
	return val
}

func ActivityNotifications(expenditure []int, d int) int {
	var alerts int

	if len(expenditure) < d+1 {
		return 0
	}

	for i := d + 1; i < len(expenditure); i++ {
		mArr := expenditure[i-d-1 : i-1]
		sliceCpy := append([]int(nil), mArr...)
		median := GetMedian(sliceCpy)
		if expenditure[i] >= int(median*2) {
			alerts += 1
		}

	}
	fmt.Println(alerts)
	return alerts
}
