package link

import (
	"testing"
)

func TestSamePart(t *testing.T) {
	l1 := &Link{
		Value: 0,
		Next: &Link{
			Value: 2,
			Next: &Link{
				Value: 4,
				Next: &Link{
					Value: 5,
					Next:  nil,
				},
			},
		},
	}

	l2 := &Link{
		Value: -1,
		Next: &Link{
			Value: 2,
			Next: &Link{
				Value: 4,
			},
		},
	}

	SamePart(l1, l2)
}
