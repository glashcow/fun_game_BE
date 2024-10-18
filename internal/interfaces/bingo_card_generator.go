package interfaces

import (
	"fmt"
	"math/rand"
	"time"
)

// generateTwoDistinctNumbers generates x distinctive numbers between 0 and max
func generateXDistinctNumbers(x int, max int) ([]int, error) {
	if x < 0 || max < 0 {
		return nil, fmt.Errorf("x and max must be greater than 0")
	}
	if x > max {
		return nil, fmt.Errorf("x must be lower than the max number")
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	nums := make([]int, 0)
	firstNumber := r.Intn(max + 1)
	nums = append(nums, firstNumber)

	trackNums := make(map[int]bool)
	trackNums[firstNumber] = true
	for len(nums) < x {
		nextNumber := r.Intn(max + 1)
		if !trackNums[nextNumber] {
			nums = append(nums, nextNumber)
			trackNums[nextNumber] = true
		}
	}

	return nums, nil
}
