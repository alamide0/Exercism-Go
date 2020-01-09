package binarysearchtree

type SearchTreeData struct {
	left  *SearchTreeData
	data  int
	right *SearchTreeData
}

func Bst(data int) *SearchTreeData {
	return &SearchTreeData{nil, data, nil}
}

func (s *SearchTreeData) Insert(data int) {

	td := &SearchTreeData{nil, data, nil}
	p := s

	for p != nil {
		if data > p.data {
			if p.right == nil {
				p.right = td
				break
			}
			p = p.right
		} else {
			if p.left == nil {
				p.left = td
				break
			}
			p = p.left
		}

	}
}

func (s *SearchTreeData) MapString(fn func(int) string) []string {
	ints := walkTree(s)
	res := make([]string, len(ints))
	for i, v := range ints {
		res[i] = fn(v)
	}
	return res
}

func (s *SearchTreeData) MapInt(fn func(int) int) []int {
	return walkTree(s)
}

func walkTree(s *SearchTreeData) []int {

	if s == nil {
		return nil
	}

	ints := make([]int, 0)

	ints = append(ints, walkTree(s.left)...)
	ints = append(ints, s.data)
	ints = append(ints, walkTree(s.right)...)

	return ints
}
