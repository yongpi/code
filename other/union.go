package other

type UnionSearch struct {
	Data map[int]int
	Rank map[int]int
}

func NewUnionSearch() *UnionSearch {
	return &UnionSearch{
		Data: make(map[int]int),
		Rank: make(map[int]int),
	}
}

func (s *UnionSearch) Union(a, b int) {
	ap := s.FindParent(a)
	bp := s.FindParent(b)

	if ap == bp {
		return
	}

	if s.Rank[ap] > s.Rank[bp] || (s.Rank[ap] == s.Rank[bp] && ap > bp) {
		s.Data[b] = ap
		s.Rank[ap] += s.Rank[bp] + 1
	} else {
		s.Data[a] = bp
		s.Rank[bp] += s.Rank[ap] + 1
	}
}

func (s *UnionSearch) FindParent(a int) int {
	if _, ok := s.Data[a]; !ok {
		s.Data[a] = a
	}

	list := make([]int, 0)
	ap := a
	for {
		list = append(list, ap)
		tmp, ok := s.Data[ap]
		if !ok || tmp == ap {
			break
		}
		ap = tmp
	}

	for _, item := range list {
		s.Data[item] = ap
	}

	return ap
}

func (s *UnionSearch) SameParent(a, b int) bool {
	return s.FindParent(a) == s.FindParent(b)
}
