package str

import (
	"fmt"
	"testing"
)

func TestBinaryNullStr(t *testing.T) {
	fmt.Println(BinaryNullStr([]string{"", "a", "", "a", "", "b", ""}, "a"))
	fmt.Println(BinaryNullStr([]string{"", "a", "", "a", "", "b", ""}, "d"))
}
