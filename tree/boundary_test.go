package tree

import "testing"

func TestPrintBoundary(t *testing.T) {
	r1 := &Tree{Value: 1}
	r2 := &Tree{Value: 2}
	r3 := &Tree{Value: 3}
	r4 := &Tree{Value: 4}
	r5 := &Tree{Value: 5}
	r6 := &Tree{Value: 6}
	r7 := &Tree{Value: 7}
	r8 := &Tree{Value: 8}
	r9 := &Tree{Value: 9}
	r10 := &Tree{Value: 10}
	r11 := &Tree{Value: 11}
	r12 := &Tree{Value: 12}
	r13 := &Tree{Value: 13}
	r14 := &Tree{Value: 14}
	r15 := &Tree{Value: 15}
	r16 := &Tree{Value: 16}

	r1.Left = r2
	r1.Right = r3

	r2.Right = r4

	r3.Left = r5
	r3.Right = r6

	r4.Left = r7
	r4.Right = r8

	r5.Left = r9
	r5.Right = r10

	r8.Right = r11

	r9.Left = r12

	r11.Left = r13
	r11.Right = r14

	r12.Left = r15
	r12.Right = r16

	PrintBoundary(r1)
}
