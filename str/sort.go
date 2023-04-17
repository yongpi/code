package str

import (
	"sort"
	"strings"
)

func MaxSumList(list []string) string {
	sort.Slice(list, func(i, j int) bool {
		return list[i] > list[j]
	})

	return strings.Join(list, "")
}
