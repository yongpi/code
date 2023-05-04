package tree

import (
	"fmt"
	"testing"
)

func TestMaxSumLevel(t *testing.T) {
	r1 := &Tree{Value: -3}
	r2 := &Tree{Value: 3}
	r3 := &Tree{Value: -9}
	r4 := &Tree{Value: 1}
	r5 := &Tree{Value: 0}
	r6 := &Tree{Value: 2}
	r7 := &Tree{Value: 1}

	r8 := &Tree{Value: 1}
	r9 := &Tree{Value: 6}

	r1.Left = r2
	r1.Right = r3

	r2.Left = r4
	r2.Right = r5

	r3.Left = r6
	r3.Right = r7

	r5.Left = r8
	r5.Right = r9

	fmt.Println(MaxSumLevel(r1, 6))
	fmt.Println(MaxSumLevel2(r1, 6))
}
