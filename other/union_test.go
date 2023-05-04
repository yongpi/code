package other

import (
	"fmt"
	"testing"
)

func TestNewUnionSearch(t *testing.T) {
	search := NewUnionSearch()
	search.Union(1, 2)
	search.Union(3, 1)
	search.Union(4, 2)
	search.Union(5, 3)

	fmt.Println(search.Data)
	fmt.Println(search.Rank)
}
