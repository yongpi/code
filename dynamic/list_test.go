package dynamic

import (
	"fmt"
	"testing"
)

func TestMinIncrList(t *testing.T) {
	fmt.Println(MinIncrList([]int{2, 1, 5, 3, 6, 4, 8, 9, 7}))
}

func TestMaxXORSubList(t *testing.T) {
	fmt.Println(MaxXORSubList([]int{3, 2, 1, 9, 0, 7, 0, 2, 1, 3}))
}

func TestFirstOrSecond(t *testing.T) {
	fmt.Println(FirstOrSecond([]int{1, 2, 100, 4}))
}

func TestIncrList(t *testing.T) {
	fmt.Println(IncrList([]int{100, 4, 200, 1, 3, 2}))
}
