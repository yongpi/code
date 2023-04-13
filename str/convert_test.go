package str

import (
	"fmt"
	"testing"
)

func TestConvertMinDistance(t *testing.T) {
	fmt.Println(ConvertMinDistance("abc", "cab", []string{"cab", "acc", "cbc", "ccc", "cac", "cbb", "aab", "abb"}))
}
