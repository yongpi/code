package str

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	root := &Node{}
	Insert(root, "abc")
	Insert(root, "ac")

	fmt.Println(PreCount(root, "a"))
	fmt.Println(PreCount(root, "ac"))
	fmt.Println(PreCount(root, "ab"))

	Delete(root, "ac")
	fmt.Println(PreCount(root, "a"))
	fmt.Println(PreCount(root, "ac"))
	fmt.Println(PreCount(root, "ab"))

}
