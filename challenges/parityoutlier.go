package challenges

import "fmt"

// if return -1, means NO OUTLIER
func ParityOutlier(arr []int) string {
	odds := []int{}
	evens := []int{}
	for _, v := range arr {
		checkOutlier(v, &odds, &evens)
	}

	if len(odds) == 1 {
		return fmt.Sprintf("%d, the only odd number", odds[0])
	} else if len(evens) == 1 {
		return fmt.Sprintf("%d, the only even number", evens[0])
	} else {
		if len(odds) > 0 {
			return "false, all odd number, NO OUTLIER"
		} else if len(evens) > 0 {
			return "false, all even number, NO OUTLIER"
		} else {
			return "NO OUTLIER"
		}
	}
}

func checkOutlier(data int, odds *[]int, evens *[]int) {
	if data%2 == 0 {
		*evens = append(*evens, data)
	} else {
		*odds = append(*odds, data)
	}
}
