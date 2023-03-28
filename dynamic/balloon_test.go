package dynamic

import (
	"fmt"
	"testing"
)

func TestBalloon(t *testing.T) {
	fmt.Println(Balloon([]int{4, 2, 3, 5, 1, 6}))
	fmt.Println(BalloonDynamic([]int{4, 2, 3, 5, 1, 6}))
}
