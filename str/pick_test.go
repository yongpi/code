package str

import (
	"fmt"
	"testing"
)

func TestPickMin(t *testing.T) {
	fmt.Println(PickMin("dbcacbca"))
}

func TestPickMaxSub(t *testing.T) {
	fmt.Println(PickMaxSub("abcdefjj"))
}

func TestPickMinContainSub(t *testing.T) {
	fmt.Println(PickMinContainSub("abcdea", "eac"))
}
