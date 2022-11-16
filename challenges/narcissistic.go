package challenges

import (
	"fmt"
	"math"
	"strconv"
)

func Narcissistic(number int) bool {
	digits := len(fmt.Sprint(number))

	sum := 0
	for _, v := range fmt.Sprint(number) {
		val, _ := strconv.Atoi(string(v))

		sum += int(math.Pow(float64(val), float64(digits)))
	}

	if sum == number {
		return true
	}

	return false
}
