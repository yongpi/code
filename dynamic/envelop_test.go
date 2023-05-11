package dynamic

import (
	"fmt"
	"testing"
)

func TestMaxLevelEnvelop(t *testing.T) {
	list := [][2]int{{3, 4}, {2, 3}, {4, 5}, {1, 3}, {2, 2}, {3, 6}, {1, 2}, {3, 2}, {2, 4}}
	fmt.Println(MaxLevelEnvelop(list))
	fmt.Println(MaxLevelEnvelop2(list))
}
