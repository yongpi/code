package tree

import "testing"

func TestMorris(t *testing.T) {
	r1 := &Tree{Value: 1}
	r2 := &Tree{Value: 2}
	r3 := &Tree{Value: 3}
	r4 := &Tree{Value: 4}
	r5 := &Tree{Value: 5}
	r6 := &Tree{Value: 6}
	r7 := &Tree{Value: 7}

	r4.Left = r2
	r4.Right = r6

	r2.Left = r1
	r2.Right = r3

	r6.Left = r5
	r6.Right = r7

	Morris(r4)
}

func TestMorrisPre(t *testing.T) {
	r1 := &Tree{Value: 1}
	r2 := &Tree{Value: 2}
	r3 := &Tree{Value: 3}
	r4 := &Tree{Value: 4}
	r5 := &Tree{Value: 5}
	r6 := &Tree{Value: 6}
	r7 := &Tree{Value: 7}

	r4.Left = r2
	r4.Right = r6

	r2.Left = r1
	r2.Right = r3

	r6.Left = r5
	r6.Right = r7

	MorrisPre(r4)

	MorrisPre2(r4)
}

func TestMorrisMid(t *testing.T) {
	r1 := &Tree{Value: 1}
	r2 := &Tree{Value: 2}
	r3 := &Tree{Value: 3}
	r4 := &Tree{Value: 4}
	r5 := &Tree{Value: 5}
	r6 := &Tree{Value: 6}
	r7 := &Tree{Value: 7}

	r4.Left = r2
	r4.Right = r6

	r2.Left = r1
	r2.Right = r3

	r6.Left = r5
	r6.Right = r7

	MorrisMid(r4)
}
