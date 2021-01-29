package main

type set struct {
	elements map[int]bool
}

func newSet(vals ...int) *set {
	els := make(map[int]bool)
	for i := range vals {
		els[vals[i]]=true
	}

	return &set{
		elements: els,
	}
}

func (s *set) size() int {
	return len(s.elements)
}

func (s *set) has(val int) bool {
	_, ok := s.elements[val]
	return ok
}

func (s *set) add(val int) {
	s.elements[val] = true
}