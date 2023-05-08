package dynamic

import (
	"fmt"
	"testing"
)

func TestMinCash(t *testing.T) {
	fmt.Println(MinCash([]int{5, 2, 3, 1}, 100))
	fmt.Println(MinCashDynamic([]int{5, 2, 3, 1}, 100))
	fmt.Println(MinCashDynamic2([]int{5, 2, 3, 1}, 100))
}

func TestAllCash(t *testing.T) {
	fmt.Println(AllCash([]int{1, 2, 3}, 6))
	fmt.Println(AllCashDynamic([]int{1, 2, 3}, 6))
}
