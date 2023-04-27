package list

import (
	"fmt"
	"testing"
)

func TestMaxMoney(t *testing.T) {
	cost := []int{5, 4, 1, 3}
	profit := []int{3, 5, 3, 2}
	fmt.Println(MaxMoney(cost, profit, 3, 2))
}
